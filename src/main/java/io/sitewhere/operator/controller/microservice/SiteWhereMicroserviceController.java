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
import java.util.Map;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

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
import io.sitewhere.k8s.crd.instance.SiteWhereInstance;
import io.sitewhere.k8s.crd.microservice.SiteWhereMicroservice;
import io.sitewhere.k8s.crd.microservice.SiteWhereMicroserviceList;
import io.sitewhere.operator.controller.ApiConstants;
import io.sitewhere.operator.controller.ResourceAnnotations;
import io.sitewhere.operator.controller.ResourceChangeType;
import io.sitewhere.operator.controller.ResourceContexts;
import io.sitewhere.operator.controller.ResourceLabels;
import io.sitewhere.operator.controller.SiteWhereComponentRoles;
import io.sitewhere.operator.controller.SiteWhereResourceController;

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
	    getWorkers().execute(new MicroserviceCreationValidator(type, microservice));
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
     * Get name for a microservice.
     * 
     * @param microservice
     * @return
     */
    protected String getMicroserviceName(SiteWhereMicroservice microservice) {
	return String.format("%s-%s", ApiConstants.SITEWHERE_APP_NAME, microservice.getSpec().getFunctionalArea());
    }

    /**
     * Get deployment name for a microservice.
     * 
     * @param microservice
     * @return
     */
    protected String getDeploymentName(SiteWhereMicroservice microservice) {
	return String.format("%s-%s", microservice.getSpec().getHelm().getReleaseName(),
		microservice.getSpec().getFunctionalArea());
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
     * Get debug service name for a microservice.
     * 
     * @param microservice
     * @return
     */
    protected String getDebugServiceName(SiteWhereMicroservice microservice) {
	return String.format("%s-debug-svc", getDeploymentName(microservice));
    }

    /**
     * Labels for a microservice resource deployment.
     * 
     * @return
     */
    protected Map<String, String> deploymentLabels(SiteWhereMicroservice microservice) {
	Map<String, String> labels = new HashMap<>();
	labels.put(ResourceLabels.LABEL_SITEWHERE_NAME, microservice.getSpec().getFunctionalArea());
	labels.put(ResourceLabels.LABEL_SITEWHERE_ROLE, SiteWhereComponentRoles.ROLE_MICROSERVICE);
	labels.put(ResourceLabels.LABEL_SITEWHERE_INSTANCE, microservice.getSpec().getInstanceName());
	labels.put(ResourceLabels.LABEL_K8S_NAME, ApiConstants.SITEWHERE_APP_NAME);
	labels.put(ResourceLabels.LABEL_K8S_INSTANCE, microservice.getSpec().getInstanceName());
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
	labels.put(ResourceLabels.LABEL_K8S_INSTANCE, microservice.getSpec().getInstanceName());
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
	labels.put(ResourceLabels.LABEL_K8S_INSTANCE, microservice.getSpec().getInstanceName());
	labels.put(ResourceLabels.LABEL_SITEWHERE_NAME, microservice.getSpec().getFunctionalArea());
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
	labels.put(ResourceLabels.LABEL_SITEWHERE_NAME, microservice.getSpec().getFunctionalArea());
	labels.put(ResourceLabels.LABEL_SITEWHERE_ROLE, SiteWhereComponentRoles.ROLE_MICROSERVICE);
	labels.put(ResourceLabels.LABEL_SITEWHERE_INSTANCE, microservice.getSpec().getInstanceName());
	labels.put(ResourceLabels.LABEL_K8S_NAME, ApiConstants.SITEWHERE_APP_NAME);
	labels.put(ResourceLabels.LABEL_K8S_INSTANCE, microservice.getSpec().getInstanceName());
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
		    microservice.getSpec().getPodSpec().getImageRepository(),
		    microservice.getSpec().getFunctionalArea(), microservice.getSpec().getPodSpec().getImageTag());
	} else {
	    return String.format("%s/%s/service-%s:%s", microservice.getSpec().getPodSpec().getImageRegistry(),
		    microservice.getSpec().getPodSpec().getImageRepository(),
		    microservice.getSpec().getFunctionalArea(), microservice.getSpec().getPodSpec().getImageTag());
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

	// Create pod spec.
	builder.withNewSpec().addNewContainer().withName(getMicroserviceName(microservice))
		.withImage(buildPodImageName(microservice))
		.withImagePullPolicy(microservice.getSpec().getPodSpec().getImagePullPolicy())
		.withPorts(microservice.getSpec().getPodSpec().getPorts())
		.withEnv(microservice.getSpec().getPodSpec().getEnv())
		.withResources(microservice.getSpec().getPodSpec().getResources())
		.withReadinessProbe(microservice.getSpec().getPodSpec().getReadinessProbe())
		.withLivenessProbe(microservice.getSpec().getPodSpec().getLivenessProbe()).endContainer().endSpec();

	return builder.build();
    }

    /**
     * Verifies an existing deployment or creates a new one for the microservice.
     * 
     * @param microservice
     * @return
     */
    protected Deployment verifyOrCreateDeployment(SiteWhereMicroservice microservice) {
	String deployName = getDeploymentName(microservice);

	// Check for existing deployment.
	Deployment deployment = getClient().apps().deployments().inNamespace(microservice.getMetadata().getNamespace())
		.withName(deployName).get();
	if (deployment != null) {
	    LOGGER.info(String.format("Found existing deployment at '%s'", deployName));
	    return deployment;
	}

	// Build new deployment.
	DeploymentBuilder dbuilder = new DeploymentBuilder();

	// Build deployment metadata.
	dbuilder.withNewMetadata().withName(deployName).withNamespace(microservice.getMetadata().getNamespace())
		.addToLabels(deploymentLabels(microservice)).endMetadata();

	// Build deployment spec.
	dbuilder.withNewSpec().withReplicas(microservice.getSpec().getReplicas()).withNewSelector()
		.withMatchLabels(deploymentMatchLabels(microservice)).endSelector()
		.withTemplate(buildPodTemplate(microservice)).endSpec();

	// Create deployment.
	deployment = getClient().apps().deployments().create(dbuilder.build());
	return deployment;
    }

    /**
     * Verifies an existing service for making debug ports available or creates a
     * new one.
     * 
     * @param microservice
     * @return
     */
    protected Service verifyOrCreateDebugService(SiteWhereMicroservice microservice) {
	String debugSvcName = getDebugServiceName(microservice);
	Service service = getClient().services().inNamespace(microservice.getMetadata().getNamespace())
		.withName(debugSvcName).get();
	if (service != null) {
	    LOGGER.info(String.format("Found existing debug service at '%s'", debugSvcName));
	    return service;
	}

	ServiceBuilder builder = new ServiceBuilder();

	// Create debug service metadata.
	builder.withNewMetadata().withName(debugSvcName).withNamespace(microservice.getMetadata().getNamespace())
		.addToLabels(serviceLabels(microservice)).endMetadata();

	// Create debug service spec.
	builder.withNewSpec().withType("LoadBalancer").addNewPort().withName("tcp-jdwp")
		.withPort(microservice.getSpec().getDebug().getJdwpPort())
		.withTargetPort(
			new IntOrStringBuilder().withIntVal(microservice.getSpec().getDebug().getJdwpPort()).build())
		.withProtocol("TCP").endPort().addNewPort().withName("tcp-jmx")
		.withPort(microservice.getSpec().getDebug().getJmxPort())
		.withTargetPort(
			new IntOrStringBuilder().withIntVal(microservice.getSpec().getDebug().getJmxPort()).build())
		.withProtocol("TCP").endPort()
		.addToSelector(ResourceLabels.LABEL_K8S_NAME, getMicroserviceName(microservice))
		.addToSelector(ResourceLabels.LABEL_K8S_INSTANCE, microservice.getSpec().getInstanceName()).endSpec();

	// Create debug service.
	service = getClient().services().create(builder.build());
	return service;
    }

    /**
     * Validates k8s resources associated with new SiteWhere microservice.
     */
    protected class MicroserviceCreationValidator extends MicroserviceWorkerRunnable {

	public MicroserviceCreationValidator(ResourceChangeType type, SiteWhereMicroservice microservice) {
	    super(type, microservice);
	}

	@Override
	public void run() {
	    LOGGER.info("Validating created SiteWhereMicroservice.");

	    // Validate that Helm metadata is provided.
	    if (getMicroservice().getSpec() == null || getMicroservice().getSpec().getHelm() == null) {
		LOGGER.error("Missing spec or Helm metadata. Can not add microservice resources.");
		return;
	    }

	    Deployment deployment = verifyOrCreateDeployment(getMicroservice());
	    LOGGER.info(String.format("Created deployment:\n\n%s", deployment.toString()));

	    // Create debug service if debug enabled.
	    if (!getMicroservice().getSpec().getDebug().isEnabled()) {
		verifyOrCreateDebugService(getMicroservice());
	    }
	}
    }

    protected ExecutorService getWorkers() {
	return workers;
    }
}
