/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator.controller.tenant;

import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.google.common.base.CaseFormat;

import io.fabric8.kubernetes.client.KubernetesClient;
import io.fabric8.kubernetes.client.informers.SharedIndexInformer;
import io.fabric8.kubernetes.client.informers.SharedInformerFactory;
import io.sitewhere.k8s.crd.ResourceContexts;
import io.sitewhere.k8s.crd.ResourceLabels;
import io.sitewhere.k8s.crd.controller.ResourceChangeType;
import io.sitewhere.k8s.crd.controller.SiteWhereResourceController;
import io.sitewhere.k8s.crd.microservice.SiteWhereMicroservice;
import io.sitewhere.k8s.crd.microservice.SiteWhereMicroserviceList;
import io.sitewhere.k8s.crd.tenant.SiteWhereTenant;
import io.sitewhere.k8s.crd.tenant.SiteWhereTenantList;
import io.sitewhere.k8s.crd.tenant.configuration.TenantConfigurationTemplate;
import io.sitewhere.k8s.crd.tenant.engine.SiteWhereTenantEngine;
import io.sitewhere.k8s.crd.tenant.engine.SiteWhereTenantEngineList;
import io.sitewhere.k8s.crd.tenant.engine.SiteWhereTenantEngineSpec;
import io.sitewhere.k8s.crd.tenant.engine.configuration.TenantEngineConfigurationTemplate;
import io.sitewhere.operator.controller.microservice.SiteWhereMicroserviceController;

/**
 * Resource controller for SiteWhere microservice monitoring.
 */
public class SiteWhereTenantController extends SiteWhereResourceController<SiteWhereTenant> {

    /** Static logger instance */
    private static Logger LOGGER = LoggerFactory.getLogger(SiteWhereTenantController.class);

    /** Resync period in milliseconds */
    private static final int RESYNC_PERIOD_MS = 10 * 60 * 1000;

    /** Workers for handling microservice resource tasks */
    private ExecutorService workers = Executors.newFixedThreadPool(2);

    public SiteWhereTenantController(KubernetesClient client, SharedInformerFactory informerFactory) {
	super(client, informerFactory);
    }

    /**
     * Create informer.
     */
    public SharedIndexInformer<SiteWhereTenant> createInformer() {
	return getInformerFactory().sharedIndexInformerForCustomResource(ResourceContexts.TENANT_CONTEXT,
		SiteWhereTenant.class, SiteWhereTenantList.class, RESYNC_PERIOD_MS);
    }

    /**
     * Get list of all microservices in the same namespace as the tenant.
     * 
     * @param tenant
     * @return
     */
    protected SiteWhereMicroserviceList getAllMicroservices(SiteWhereTenant tenant) {
	return getSitewhereClient().getMicroservices().inNamespace(tenant.getMetadata().getNamespace()).list();
    }

    /**
     * Get map of tenant engines for a tenant (indexed by microservice name).
     * 
     * @param tenant
     * @return
     */
    protected Map<String, SiteWhereTenantEngine> getTenantEnginesForTenantByMicroservice(SiteWhereTenant tenant) {
	SiteWhereTenantEngineList list = getSitewhereClient().getTenantEngines()
		.inNamespace(tenant.getMetadata().getNamespace())
		.withLabel(ResourceLabels.LABEL_SITEWHERE_TENANT, tenant.getMetadata().getName()).list();
	Map<String, SiteWhereTenantEngine> byMicroservice = new HashMap<>();
	for (SiteWhereTenantEngine engine : list.getItems()) {
	    String microservice = engine.getMetadata().getLabels().get(ResourceLabels.LABEL_SITEWHERE_MICROSERVICE);
	    if (microservice != null) {
		byMicroservice.put(microservice, engine);
	    }
	}
	return byMicroservice;
    }

    /**
     * Locate tenant engine configuration template based on tenant configuration
     * template associated with tenant.
     * 
     * @param tenant
     * @param microservice
     * @return
     */
    protected TenantEngineConfigurationTemplate getTenantEngineConfigurationTemplate(SiteWhereTenant tenant,
	    SiteWhereMicroservice microservice) {
	TenantConfigurationTemplate tenantTemplate = getSitewhereClient().getTenantConfigurationTemplates()
		.withName(tenant.getSpec().getConfigurationTemplate()).get();
	if (tenantTemplate == null) {
	    String message = String.format("Tenant references non-existent configuration template '%s'.",
		    tenant.getSpec().getConfigurationTemplate());
	    LOGGER.warn(message);
	    throw new RuntimeException(message);
	}
	String target = CaseFormat.LOWER_HYPHEN.to(CaseFormat.LOWER_CAMEL,
		SiteWhereMicroserviceController.getFunctionalArea(microservice));
	String tecTemplateName = tenantTemplate.getSpec().getTenantEngineTemplates().get(target);
	if (tecTemplateName == null) {
	    String message = String.format("Missing tenant engine template mapping for '%s'.",
		    SiteWhereMicroserviceController.getFunctionalArea(microservice));
	    LOGGER.warn(message);
	    return null;
	}

	return getSitewhereClient().getTenantEngineConfigurationTemplates().withName(tecTemplateName).get();
    }

    /**
     * Create a new tenant engine for a tenant/microservice combination.
     * 
     * @param tenant
     * @param microservice
     */
    protected void createNewTenantEngine(SiteWhereTenant tenant, SiteWhereMicroservice microservice) {
	SiteWhereTenantEngine engine = new SiteWhereTenantEngine();
	String tenantEngineName = String.format("%s-%s-%s", tenant.getMetadata().getName(),
		microservice.getMetadata().getName(), String.valueOf(System.currentTimeMillis()));
	engine.getMetadata().setName(tenantEngineName);
	engine.getMetadata().setNamespace(tenant.getMetadata().getNamespace());

	Map<String, String> labels = new HashMap<>();
	labels.put(ResourceLabels.LABEL_SITEWHERE_TENANT, tenant.getMetadata().getName());
	labels.put(ResourceLabels.LABEL_SITEWHERE_MICROSERVICE, microservice.getMetadata().getName());
	labels.put(ResourceLabels.LABEL_SITEWHERE_FUNCTIONAL_AREA,
		SiteWhereMicroserviceController.getFunctionalArea(microservice));
	engine.getMetadata().setLabels(labels);

	// Look up tenant configuration template for tenant/microservice combination.
	TenantEngineConfigurationTemplate tecTemplate = getTenantEngineConfigurationTemplate(tenant, microservice);
	if (tecTemplate == null) {
	    LOGGER.warn("Unable to resolve default tenant engine configuration. Skipping engine creation.");
	    return;
	}

	// Copy template configuration into spec.
	SiteWhereTenantEngineSpec spec = new SiteWhereTenantEngineSpec();
	spec.setConfiguration(tecTemplate.getSpec().getConfiguration());
	engine.setSpec(spec);

	getSitewhereClient().getTenantEngines().withName(tenantEngineName).createOrReplace(engine);
	LOGGER.info(String.format("Created new tenant engine for tenant `%s` microservice `%s`",
		tenant.getMetadata().getName(), microservice.getMetadata().getName()));
    }

    /**
     * For a given tenant, verify that tenant engines exist for each microservice in
     * the tenant namespace.
     * 
     * @param tenant
     */
    protected void validateTenantEngines(SiteWhereTenant tenant) {
	// Index existing tenant engines by microservice.
	Map<String, SiteWhereTenantEngine> enginesByMicroservice = getTenantEnginesForTenantByMicroservice(tenant);

	// List all microservices and check whether engines exist for each.
	SiteWhereMicroserviceList allMicroservices = getAllMicroservices(tenant);
	for (SiteWhereMicroservice microservice : allMicroservices.getItems()) {
	    boolean supportsMultitenant = microservice.getSpec().isMultitenant();

	    // Create engine if not found.
	    if (supportsMultitenant && enginesByMicroservice.get(microservice.getMetadata().getName()) == null) {
		createNewTenantEngine(tenant, microservice);
	    }
	}
    }

    /**
     * Deletes any tenant engines associated with the tenant.
     * 
     * @param tenant
     * @return
     */
    protected boolean deleteTenantEngines(SiteWhereTenant tenant) {
	SiteWhereTenantEngineList list = getSitewhereClient().getTenantEngines()
		.inNamespace(tenant.getMetadata().getNamespace())
		.withLabel(ResourceLabels.LABEL_SITEWHERE_TENANT, tenant.getMetadata().getName()).list();
	LOGGER.info(String.format("Deleting %s tenant engines for tenant '%s'", String.valueOf(list.getItems().size()),
		tenant.getMetadata().getName()));
	return getSitewhereClient().getTenantEngines().inNamespace(tenant.getMetadata().getNamespace())
		.withLabel(ResourceLabels.LABEL_SITEWHERE_TENANT, tenant.getMetadata().getName()).delete();
    }

    /*
     * @see io.sitewhere.operator.controller.SiteWhereResourceController#
     * reconcileResourceChange(io.sitewhere.operator.controller.ResourceChangeType,
     * io.fabric8.kubernetes.client.CustomResource)
     */
    @Override
    public void reconcileResourceChange(ResourceChangeType type, SiteWhereTenant tenant) {
	LOGGER.info(String.format("Detected %s resource change in tenant '%s'.", type.name(),
		tenant.getMetadata().getName()));
	if (type == ResourceChangeType.CREATE) {
	    getWorkers().execute(new TenantCreationWorker(type, tenant));
	} else if (type == ResourceChangeType.UPDATE) {
	    getWorkers().execute(new TenantUpdateWorker(type, tenant));
	} else if (type == ResourceChangeType.DELETE) {
	    getWorkers().execute(new TenantDeleteWorker(type, tenant));
	}
    }

    /**
     * Creates k8s resources associated with new SiteWhere tenant.
     */
    protected class TenantCreationWorker extends TenantWorkerRunnable {

	public TenantCreationWorker(ResourceChangeType type, SiteWhereTenant tenant) {
	    super(type, tenant);
	}

	@Override
	public void run() {
	    LOGGER.info("Handling tenant creation.");
	    validateTenantEngines(getTenant());
	}
    }

    /**
     * Updates k8s resources associated with new SiteWhere tenant.
     */
    protected class TenantUpdateWorker extends TenantWorkerRunnable {

	public TenantUpdateWorker(ResourceChangeType type, SiteWhereTenant tenant) {
	    super(type, tenant);
	}

	@Override
	public void run() {
	    LOGGER.info("Handling tenant update.");
	    validateTenantEngines(getTenant());
	}
    }

    /**
     * Deletes k8s resources associated with new SiteWhere tenant.
     */
    protected class TenantDeleteWorker extends TenantWorkerRunnable {

	public TenantDeleteWorker(ResourceChangeType type, SiteWhereTenant tenant) {
	    super(type, tenant);
	}

	@Override
	public void run() {
	    LOGGER.info("Handling tenant deletion.");
	    deleteTenantEngines(getTenant());
	}
    }

    protected ExecutorService getWorkers() {
	return workers;
    }
}
