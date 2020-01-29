/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator.controller.tenant;

import java.util.Map;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

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
import io.sitewhere.k8s.crd.tenant.engine.SiteWhereTenantEngine;
import io.sitewhere.k8s.crd.tenant.engine.SiteWhereTenantEngineList;
import io.sitewhere.operator.controller.OperatorUtils;

/**
 * Resource controller for SiteWhere microservice monitoring.
 */
public class SiteWhereTenantController extends SiteWhereResourceController<SiteWhereTenant> {

    /** Static logger instance */
    private static Logger LOGGER = LoggerFactory.getLogger(SiteWhereTenantController.class);

    /** Resync period in milliseconds */
    private static final int RESYNC_PERIOD_MS = 10 * 60 * 1000;

    /** Workers for handling microservice resource tasks */
    private ExecutorService workers = Executors.newFixedThreadPool(5);

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
     * For a given tenant, verify that tenant engines exist for each microservice in
     * the tenant namespace.
     * 
     * @param tenant
     */
    protected void validateTenantEngines(SiteWhereTenant tenant) {
	// Index existing tenant engines by microservice.
	Map<String, SiteWhereTenantEngine> enginesByMicroservice = OperatorUtils
		.getTenantEnginesForTenantByMicroservice(getSitewhereClient(), tenant);

	// List all microservices and check whether engines exist for each.
	SiteWhereMicroserviceList allMicroservices = getAllMicroservices(tenant);
	for (SiteWhereMicroservice microservice : allMicroservices.getItems()) {
	    boolean supportsMultitenant = microservice.getSpec().isMultitenant();

	    // Create engine if not found.
	    if (supportsMultitenant && enginesByMicroservice.get(microservice.getMetadata().getName()) == null) {
		OperatorUtils.createNewTenantEngine(getSitewhereClient(), tenant, microservice);
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
