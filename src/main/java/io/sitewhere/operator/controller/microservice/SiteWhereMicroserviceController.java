/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator.controller.microservice;

import java.time.Clock;
import java.time.Instant;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import io.fabric8.kubernetes.api.model.ContainerPort;
import io.fabric8.kubernetes.api.model.ContainerPortBuilder;
import io.fabric8.kubernetes.api.model.Event;
import io.fabric8.kubernetes.api.model.EventBuilder;
import io.fabric8.kubernetes.api.model.IntOrStringBuilder;
import io.fabric8.kubernetes.api.model.PodTemplateSpec;
import io.fabric8.kubernetes.api.model.PodTemplateSpecBuilder;
import io.fabric8.kubernetes.api.model.Service;
import io.fabric8.kubernetes.api.model.ServiceBuilder;
import io.fabric8.kubernetes.api.model.apps.Deployment;
import io.fabric8.kubernetes.api.model.apps.DeploymentBuilder;
import io.fabric8.kubernetes.client.KubernetesClient;
import io.fabric8.kubernetes.client.informers.SharedIndexInformer;
import io.fabric8.kubernetes.client.informers.SharedInformerFactory;
import io.sitewhere.k8s.crd.ApiConstants;
import io.sitewhere.k8s.crd.ResourceContexts;
import io.sitewhere.k8s.crd.ResourceLabels;
import io.sitewhere.k8s.crd.controller.ResourceChangeType;
import io.sitewhere.k8s.crd.controller.SiteWhereResourceController;
import io.sitewhere.k8s.crd.instance.SiteWhereInstance;
import io.sitewhere.k8s.crd.microservice.SiteWhereMicroservice;
import io.sitewhere.k8s.crd.microservice.SiteWhereMicroserviceList;
import io.sitewhere.k8s.crd.tenant.SiteWhereTenant;
import io.sitewhere.k8s.crd.tenant.SiteWhereTenantList;
import io.sitewhere.k8s.crd.tenant.engine.SiteWhereTenantEngine;
import io.sitewhere.operator.controller.OperatorUtils;
import io.sitewhere.operator.controller.ResourceAnnotations;
import io.sitewhere.operator.controller.SiteWhereComponentRoles;

/**
 * Resource controller for SiteWhere microservice monitoring.
 */
public class SiteWhereMicroserviceController extends SiteWhereResourceController<SiteWhereMicroservice> {

    /** Static logger instance */
    private static Logger LOGGER = LoggerFactory.getLogger(SiteWhereMicroserviceController.class);

    /** Resync period in milliseconds */
    private static final int RESYNC_PERIOD_MS = 10 * 60 * 1000;

    /** Workers for handling microrservice resource tasks */
    private ExecutorService workers = Executors.newFixedThreadPool(2);

    public SiteWhereMicroserviceController(KubernetesClient client, SharedInformerFactory informerFactory) {
	super(client, informerFactory);
    }

    /**
     * Create informer.
     */
    public SharedIndexInformer<SiteWhereMicroservice> createInformer() {
	return getInformerFactory().sharedIndexInformerForCustomResource(ResourceContexts.MICROSERVICE_CONTEXT,
		SiteWhereMicroservice.class, SiteWhereMicroserviceList.class, RESYNC_PERIOD_MS);
    }

    /*
     * @see io.sitewhere.operator.controller.SiteWhereResourceController#
     * reconcileResourceChange(io.sitewhere.operator.controller.ResourceChangeType,
     * io.fabric8.kubernetes.client.CustomResource)
     */
    @Override
    public void reconcileResourceChange(ResourceChangeType type, SiteWhereMicroservice microservice) {
	LOGGER.info(String.format("Detected %s resource change in microservice %s.", type.name(),
		microservice.getMetadata().getName()));
	if (type == ResourceChangeType.CREATE) {
	    getWorkers().execute(new MicroserviceCreationWorker(type, microservice));
	} else if (type == ResourceChangeType.UPDATE) {
	    getWorkers().execute(new MicroserviceUpdateWorker(type, microservice));
	} else if (type == ResourceChangeType.DELETE) {
	    getWorkers().execute(new MicroserviceDeleteWorker(type, microservice));
	}
    }

    /**
     * Create event for a microservice.
     * 
     * @param microservice
     * @param reason
     * @param type
     * @param message
     */
    protected void createEventForMicroservice(SiteWhereMicroservice microservice, String reason, String type,
	    String message) {
	String name = microservice.getMetadata().getName() + "-event-" + System.currentTimeMillis();
	String timestamp = Instant.now(Clock.systemDefaultZone()).toString();
	Event event = new EventBuilder().withNewMetadata().withName(name)
		.withNamespace(microservice.getMetadata().getNamespace()).endMetadata().withCount(1).withReason(reason)
		.withMessage(message).withType(type).withNewInvolvedObject()
		.withKind(SiteWhereInstance.class.getSimpleName())
		.withNamespace(microservice.getMetadata().getNamespace()).withName(microservice.getMetadata().getName())
		.endInvolvedObject().withFirstTimestamp(timestamp).withLastTimestamp(timestamp).withNewSource()
		.withComponent("sitewhere-operator").endSource().build();
	getClient().events().create(event);
    }

    /**
     * Indicates whether microservice debugging is enabled.
     * 
     * @param microservice
     * @return
     */
    protected boolean isDebugEnabled(SiteWhereMicroservice microservice) {
	if (microservice.getSpec().getDebug() != null && microservice.getSpec().getDebug().isEnabled()) {
	    return true;
	}
	return false;
    }

    /**
     * Get instance name for a microservice.
     * 
     * @param microservice
     * @return
     */
    public static String getInstanceName(SiteWhereMicroservice microservice) {
	String instanceName = microservice.getMetadata().getLabels().get(ResourceLabels.LABEL_SITEWHERE_INSTANCE);
	if (instanceName == null) {
	    throw new RuntimeException(String.format("Microservice '%s' does not have an instance name label.",
		    microservice.getMetadata().getName()));
	}
	return instanceName;
    }

    /**
     * Get functional area for a microservice.
     * 
     * @param microservice
     * @return
     */
    public static String getFunctionalArea(SiteWhereMicroservice microservice) {
	String functionalArea = microservice.getMetadata().getLabels()
		.get(ResourceLabels.LABEL_SITEWHERE_FUNCTIONAL_AREA);
	if (functionalArea == null) {
	    throw new RuntimeException(String.format("Microservice '%s' does not have a functional area label.",
		    microservice.getMetadata().getName()));
	}
	return functionalArea;
    }

    /**
     * Get name for a microservice.
     * 
     * @param microservice
     * @return
     */
    protected String getMicroserviceName(SiteWhereMicroservice microservice) {
	return String.format("%s-%s", ApiConstants.SITEWHERE_APP_NAME, getFunctionalArea(microservice));
    }

    /**
     * Get deployment name for a microservice.
     * 
     * @param microservice
     * @return
     */
    protected String getDeploymentName(SiteWhereMicroservice microservice) {
	return String.format("%s-%s", ApiConstants.SITEWHERE_APP_NAME, getFunctionalArea(microservice));
    }

    /**
     * Labels for a microservice resource deployment.
     * 
     * @return
     */
    protected Map<String, String> deploymentLabels(SiteWhereMicroservice microservice) {
	Map<String, String> labels = new HashMap<>();
	labels.put(ResourceLabels.LABEL_SITEWHERE_NAME, getFunctionalArea(microservice));
	labels.put(ResourceLabels.LABEL_SITEWHERE_ROLE, SiteWhereComponentRoles.ROLE_MICROSERVICE);
	labels.put(ResourceLabels.LABEL_SITEWHERE_INSTANCE, getInstanceName(microservice));
	labels.put(ResourceLabels.LABEL_K8S_NAME, ApiConstants.SITEWHERE_APP_NAME);
	labels.put(ResourceLabels.LABEL_K8S_INSTANCE, getInstanceName(microservice));
	labels.put(ResourceLabels.LABEL_K8S_MANAGED_BY, microservice.getSpec().getHelm().getReleaseService());
	labels.put(ResourceLabels.LABEL_HELM_CHART, microservice.getSpec().getHelm().getChartName());
	return labels;
    }

    /**
     * Match labels for locating pods for deployment.
     * 
     * @return
     */
    protected Map<String, String> deploymentMatchLabels(SiteWhereMicroservice microservice) {
	Map<String, String> labels = new HashMap<>();
	labels.put(ResourceLabels.LABEL_K8S_INSTANCE, getInstanceName(microservice));
	labels.put(ResourceLabels.LABEL_K8S_NAME, getMicroserviceName(microservice));
	return labels;
    }

    /**
     * Labels for a microservice resource pod.
     * 
     * @return
     */
    protected Map<String, String> podLabels(SiteWhereMicroservice microservice) {
	Map<String, String> labels = new HashMap<>();
	labels.put(ResourceLabels.LABEL_K8S_NAME, getMicroserviceName(microservice));
	labels.put(ResourceLabels.LABEL_K8S_INSTANCE, getInstanceName(microservice));
	labels.put(ResourceLabels.LABEL_SITEWHERE_NAME, getFunctionalArea(microservice));
	labels.put(ResourceLabels.LABEL_SITEWHERE_ROLE, SiteWhereComponentRoles.ROLE_MICROSERVICE);
	return labels;
    }

    /**
     * Annotations for a microservice resource pod.
     * 
     * @return
     */
    protected Map<String, String> podAnnotations(SiteWhereMicroservice microservice) {
	Map<String, String> labels = new HashMap<>();
	labels.put(ResourceAnnotations.ANN_PROMETHEUS_SCRAPE, "true");
	labels.put(ResourceAnnotations.ANN_PROMETHEUS_SCHEME, "http");
	labels.put(ResourceAnnotations.ANN_PROMETHEUS_PORT, "9090");
	return labels;
    }

    /**
     * Labels for a microservice resource deployment.
     * 
     * @return
     */
    protected Map<String, String> serviceLabels(SiteWhereMicroservice microservice) {
	Map<String, String> labels = new HashMap<>();
	labels.put(ResourceLabels.LABEL_SITEWHERE_NAME, getFunctionalArea(microservice));
	labels.put(ResourceLabels.LABEL_SITEWHERE_ROLE, SiteWhereComponentRoles.ROLE_MICROSERVICE);
	labels.put(ResourceLabels.LABEL_SITEWHERE_INSTANCE, getInstanceName(microservice));
	labels.put(ResourceLabels.LABEL_K8S_NAME, ApiConstants.SITEWHERE_APP_NAME);
	labels.put(ResourceLabels.LABEL_K8S_INSTANCE, getInstanceName(microservice));
	labels.put(ResourceLabels.LABEL_K8S_MANAGED_BY, microservice.getSpec().getHelm().getReleaseService());
	labels.put(ResourceLabels.LABEL_HELM_CHART, microservice.getSpec().getHelm().getChartName());
	return labels;
    }

    /**
     * Build pod image name based on whether debug is enabled.
     * 
     * @param microservice
     * @return
     */
    protected String buildPodImageName(SiteWhereMicroservice microservice) {
	if (microservice.getSpec().getDebug() != null && microservice.getSpec().getDebug().isEnabled()) {
	    return String.format("%s/%s/service-%s:debug-%s", microservice.getSpec().getPodSpec().getImageRegistry(),
		    microservice.getSpec().getPodSpec().getImageRepository(), getFunctionalArea(microservice),
		    microservice.getSpec().getPodSpec().getImageTag());
	} else {
	    return String.format("%s/%s/service-%s:%s", microservice.getSpec().getPodSpec().getImageRegistry(),
		    microservice.getSpec().getPodSpec().getImageRepository(), getFunctionalArea(microservice),
		    microservice.getSpec().getPodSpec().getImageTag());
	}
    }

    /**
     * Build pod template for deployment.
     * 
     * @param microservice
     * @return
     */
    protected PodTemplateSpec buildPodTemplate(SiteWhereMicroservice microservice) {
	PodTemplateSpecBuilder builder = new PodTemplateSpecBuilder();

	// Create pod metadata.
	builder.withNewMetadata().addToLabels(podLabels(microservice)).addToAnnotations(podAnnotations(microservice))
		.endMetadata();

	// Add debug ports if debug is enabled.
	List<ContainerPort> ports = microservice.getSpec().getPodSpec().getPorts();
	if (isDebugEnabled(microservice)) {
	    ContainerPort jwdp = new ContainerPortBuilder()
		    .withContainerPort(microservice.getSpec().getDebug().getJdwpPort()).build();
	    ports.add(jwdp);
	    ContainerPort jmx = new ContainerPortBuilder()
		    .withContainerPort(microservice.getSpec().getDebug().getJmxPort()).build();
	    ports.add(jmx);
	}

	// Create pod spec.
	builder.withNewSpec().addNewContainer().withName(getMicroserviceName(microservice))
		.withImage(buildPodImageName(microservice))
		.withImagePullPolicy(microservice.getSpec().getPodSpec().getImagePullPolicy()).withPorts(ports)
		.withEnv(microservice.getSpec().getPodSpec().getEnv())
		.withResources(microservice.getSpec().getPodSpec().getResources())
		.withReadinessProbe(microservice.getSpec().getPodSpec().getReadinessProbe())
		.withLivenessProbe(microservice.getSpec().getPodSpec().getLivenessProbe()).endContainer().endSpec();

	return builder.build();
    }

    /**
     * Get deployment responsible for microservice pods.
     * 
     * @param microservice
     * @return
     */
    protected Deployment getDeployment(SiteWhereMicroservice microservice) {
	String deployName = getDeploymentName(microservice);
	return getClient().apps().deployments().inNamespace(microservice.getMetadata().getNamespace())
		.withName(deployName).get();
    }

    /**
     * Creates or updates a deployment based on microservice resource configuration.
     * 
     * @param microservice
     * @return
     */
    protected Deployment createOrUpdateDeployment(SiteWhereMicroservice microservice) {
	String deployName = getDeploymentName(microservice);

	// Build deployment metadata.
	DeploymentBuilder dbuilder = new DeploymentBuilder();
	dbuilder.withNewMetadata().withName(deployName).withNamespace(microservice.getMetadata().getNamespace())
		.addToLabels(deploymentLabels(microservice)).endMetadata();

	// Build deployment spec.
	dbuilder.withNewSpec().withReplicas(microservice.getSpec().getReplicas()).withNewSelector()
		.withMatchLabels(deploymentMatchLabels(microservice)).endSelector()
		.withTemplate(buildPodTemplate(microservice)).endSpec();

	// Create deployment.
	Deployment deployment = getClient().apps().deployments().inNamespace(microservice.getMetadata().getNamespace())
		.withName(deployName).createOrReplace(dbuilder.build());
	return deployment;
    }

    /**
     * Delete deployment responsible for microservice pods.
     * 
     * @param microservice
     * @return
     */
    protected Boolean deleteDeployment(SiteWhereMicroservice microservice) {
	Deployment deployment = getDeployment(microservice);
	if (deployment != null) {
	    return getClient().apps().deployments().delete(deployment);
	}
	return true;
    }

    /**
     * Get service name for a microservice.
     * 
     * @param microservice
     * @return
     */
    protected String getServiceName(SiteWhereMicroservice microservice) {
	return String.format("%s-svc", getDeploymentName(microservice));
    }

    /**
     * Get service which exposes microservice pods.
     * 
     * @param microservice
     * @return
     */
    protected Service getService(SiteWhereMicroservice microservice) {
	String svcName = getServiceName(microservice);
	return getClient().services().inNamespace(microservice.getMetadata().getNamespace()).withName(svcName).get();
    }

    /**
     * Creates a service exposing pods (or updates if service exists).
     * 
     * @param microservice
     * @return
     */
    protected Service createOrUpdateService(SiteWhereMicroservice microservice) {
	String svcName = getServiceName(microservice);

	// Create service metadata.
	ServiceBuilder builder = new ServiceBuilder();
	builder.withNewMetadata().withName(svcName).withNamespace(microservice.getMetadata().getNamespace())
		.addToLabels(serviceLabels(microservice)).endMetadata();

	// Create service spec.
	builder.withNewSpec().withType(microservice.getSpec().getServiceSpec().getType())
		.withPorts(microservice.getSpec().getServiceSpec().getPorts())
		.addToSelector(ResourceLabels.LABEL_K8S_NAME, getMicroserviceName(microservice))
		.addToSelector(ResourceLabels.LABEL_K8S_INSTANCE, getInstanceName(microservice)).endSpec();

	// Create debug service.
	Service service = getClient().services().inNamespace(microservice.getMetadata().getNamespace())
		.withName(svcName).createOrReplace(builder.build());
	return service;
    }

    /**
     * Delete service responsible for exposing pods.
     * 
     * @param microservice
     * @return
     */
    protected Boolean deleteService(SiteWhereMicroservice microservice) {
	Service service = getService(microservice);
	if (service != null) {
	    return getClient().services().delete(service);
	}
	return true;
    }

    /**
     * Get debug service name for a microservice.
     * 
     * @param microservice
     * @return
     */
    protected String getDebugServiceName(SiteWhereMicroservice microservice) {
	return String.format("%s-debug-svc", getDeploymentName(microservice));
    }

    /**
     * Get service which exposes debugger ports.
     * 
     * @param microservice
     * @return
     */
    protected Service getDebugService(SiteWhereMicroservice microservice) {
	String debugSvcName = getDebugServiceName(microservice);
	return getClient().services().inNamespace(microservice.getMetadata().getNamespace()).withName(debugSvcName)
		.get();
    }

    /**
     * Creates a service exposing debug ports (or updates if service exists).
     * 
     * @param microservice
     * @return
     */
    protected Service createOrUpdateDebugService(SiteWhereMicroservice microservice) {
	String debugSvcName = getDebugServiceName(microservice);

	// Create debug service metadata.
	ServiceBuilder builder = new ServiceBuilder();
	builder.withNewMetadata().withName(debugSvcName).withNamespace(microservice.getMetadata().getNamespace())
		.addToLabels(serviceLabels(microservice)).endMetadata();

	// Create debug service spec.
	builder.withNewSpec().withType(microservice.getSpec().getServiceSpec().getType()).addNewPort()
		.withName("tcp-jdwp").withPort(microservice.getSpec().getDebug().getJdwpPort())
		.withTargetPort(
			new IntOrStringBuilder().withIntVal(microservice.getSpec().getDebug().getJdwpPort()).build())
		.withProtocol("TCP").endPort().addNewPort().withName("tcp-jmx")
		.withPort(microservice.getSpec().getDebug().getJmxPort())
		.withTargetPort(
			new IntOrStringBuilder().withIntVal(microservice.getSpec().getDebug().getJmxPort()).build())
		.withProtocol("TCP").endPort()
		.addToSelector(ResourceLabels.LABEL_K8S_NAME, getMicroserviceName(microservice))
		.addToSelector(ResourceLabels.LABEL_K8S_INSTANCE, getInstanceName(microservice)).endSpec();

	// Create debug service.
	Service service = getClient().services().inNamespace(microservice.getMetadata().getNamespace())
		.withName(debugSvcName).createOrReplace(builder.build());
	return service;
    }

    /**
     * Delete service responsible for exposing debug ports.
     * 
     * @param microservice
     * @return
     */
    protected Boolean deleteDebugServiceIfExists(SiteWhereMicroservice microservice) {
	Service service = getDebugService(microservice);
	if (service != null) {
	    return getClient().services().delete(service);
	}
	return true;
    }

    /**
     * Creates tenant engines for each tenant that does not already have one for a
     * newly added microservice.
     * 
     * @param microservice
     */
    protected void createTenantEnginesForExistingTenants(SiteWhereMicroservice microservice) {
	SiteWhereTenantList all = OperatorUtils.getAllTenants(getSitewhereClient());
	Map<String, SiteWhereTenantEngine> existing = OperatorUtils
		.getTenantEnginesForMicroserviceByTenant(getSitewhereClient(), microservice);
	for (SiteWhereTenant tenant : all.getItems()) {
	    String tenantId = tenant.getMetadata().getName();
	    if (existing.get(tenantId) == null) {
		LOGGER.info(String.format(
			"No tenant engine found for tenant '%s' on microservice '%s' creation. Adding tenant engine.",
			tenantId, microservice.getMetadata().getName()));
		OperatorUtils.createNewTenantEngine(getSitewhereClient(), tenant, microservice);
	    }
	}
    }

    /**
     * Creates k8s resources associated with new SiteWhere microservice.
     */
    protected class MicroserviceCreationWorker extends MicroserviceWorkerRunnable {

	public MicroserviceCreationWorker(ResourceChangeType type, SiteWhereMicroservice microservice) {
	    super(type, microservice);
	}

	@Override
	public void run() {
	    LOGGER.info("Handling microservice creation.");
	    String name = getMicroserviceName(getMicroservice());

	    // Validate that Helm metadata is provided.
	    if (getMicroservice().getSpec() == null || getMicroservice().getSpec().getHelm() == null) {
		LOGGER.error("Missing spec or Helm metadata. Can not add microservice resources.");
		return;
	    }

	    Deployment deployment = getDeployment(getMicroservice());
	    if (deployment == null) {
		deployment = createOrUpdateDeployment(getMicroservice());
		LOGGER.info(String.format("Created deployment for microservice %s", name));
		if (LOGGER.isDebugEnabled()) {
		    LOGGER.debug(String.format("Created deployment:\n\n%s", deployment.toString()));
		}
	    }

	    Service service = getService(getMicroservice());
	    if (service == null) {
		service = createOrUpdateService(getMicroservice());
		LOGGER.info(String.format("Created service for microservice %s", name));
		if (LOGGER.isDebugEnabled()) {
		    LOGGER.debug(String.format("Created service:\n\n%s", service.toString()));
		}
	    }

	    // Create debug service if debug enabled.
	    if (getMicroservice().getSpec().getDebug().isEnabled()) {
		Service debugSvc = getDebugService(getMicroservice());
		if (debugSvc == null) {
		    debugSvc = createOrUpdateDebugService(getMicroservice());
		    LOGGER.info(String.format("Created debug service for microservice %s", name));
		    if (LOGGER.isDebugEnabled()) {
			LOGGER.debug(String.format("Created debug service:\n\n%s", debugSvc.toString()));
		    }
		}
	    } else {
		deleteDebugServiceIfExists(getMicroservice());
		LOGGER.info(String.format("Deleted debug service for microservice %s", name));
	    }

	    // Create tenant engines for existing tenants if not already present.
	    createTenantEnginesForExistingTenants(getMicroservice());
	}
    }

    /**
     * Updates k8s resources associated with new SiteWhere microservice.
     */
    protected class MicroserviceUpdateWorker extends MicroserviceWorkerRunnable {

	public MicroserviceUpdateWorker(ResourceChangeType type, SiteWhereMicroservice microservice) {
	    super(type, microservice);
	}

	@Override
	public void run() {
	    LOGGER.info("Handling microservice update.");
	    String name = getMicroserviceName(getMicroservice());

	    // Validate that Helm metadata is provided.
	    if (getMicroservice().getSpec() == null || getMicroservice().getSpec().getHelm() == null) {
		LOGGER.error("Missing spec or Helm metadata. Can not add microservice resources.");
		return;
	    }

	    // Update an existing deployment (or create if necessary).
	    Deployment deployment = createOrUpdateDeployment(getMicroservice());
	    LOGGER.info(String.format("Updated deployment for microservice %s", name));
	    if (LOGGER.isDebugEnabled()) {
		LOGGER.debug(String.format("Updated deployment:\n\n%s", deployment.toString()));
	    }

	    // Update an existing service (or create if necessary)
	    Service service = createOrUpdateService(getMicroservice());
	    LOGGER.info(String.format("Updated service for microservice %s", name));
	    if (LOGGER.isDebugEnabled()) {
		LOGGER.debug(String.format("Updated service:\n\n%s", service.toString()));
	    }

	    // Create debug service if debug enabled.
	    if (getMicroservice().getSpec().getDebug().isEnabled()) {
		Service debugSvc = createOrUpdateDebugService(getMicroservice());
		LOGGER.info(String.format("Updated debug service for microservice %s", name));
		if (LOGGER.isDebugEnabled()) {
		    LOGGER.debug(String.format("Updated debug service:\n\n%s", debugSvc.toString()));
		}
	    } else {
		deleteDebugServiceIfExists(getMicroservice());
		LOGGER.info(String.format("Deleted debug service for microservice %s", name));
	    }
	}
    }

    /**
     * Deletes k8s resources associated with new SiteWhere microservice.
     */
    protected class MicroserviceDeleteWorker extends MicroserviceWorkerRunnable {

	public MicroserviceDeleteWorker(ResourceChangeType type, SiteWhereMicroservice microservice) {
	    super(type, microservice);
	}

	@Override
	public void run() {
	    LOGGER.info("Handling microservice deletion.");
	    String name = getMicroserviceName(getMicroservice());

	    // Validate that Helm metadata is provided.
	    if (getMicroservice().getSpec() == null || getMicroservice().getSpec().getHelm() == null) {
		LOGGER.error("Missing spec or Helm metadata. Can not add microservice resources.");
		return;
	    }

	    // Delete deployment.
	    deleteDeployment(getMicroservice());
	    LOGGER.info(String.format("Deleted deployment for microservice %s", name));

	    // Delete service.
	    deleteService(getMicroservice());
	    LOGGER.info(String.format("Deleted service for microservice %s", name));

	    // Delete debug service.
	    deleteDebugServiceIfExists(getMicroservice());
	    LOGGER.info(String.format("Deleted debug service for microservice %s", name));
	}
    }

    protected ExecutorService getWorkers() {
	return workers;
    }
}
