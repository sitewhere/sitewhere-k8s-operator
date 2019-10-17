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
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import io.fabric8.kubernetes.api.model.Event;
import io.fabric8.kubernetes.api.model.EventBuilder;
import io.fabric8.kubernetes.client.KubernetesClient;
import io.fabric8.kubernetes.client.dsl.base.CustomResourceDefinitionContext;
import io.fabric8.kubernetes.client.informers.SharedIndexInformer;
import io.fabric8.kubernetes.client.informers.SharedInformerFactory;
import io.sitewhere.k8s.crd.instance.SiteWhereInstance;
import io.sitewhere.k8s.crd.microservice.SiteWhereMicroservice;
import io.sitewhere.k8s.crd.microservice.SiteWhereMicroserviceList;
import io.sitewhere.operator.controller.ApiConstants;
import io.sitewhere.operator.controller.ResourceChangeType;
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

    /** Context used for accessing instances */
    private static CustomResourceDefinitionContext CONTEXT = new CustomResourceDefinitionContext.Builder()
	    .withVersion(ApiConstants.SITEWHERE_API_VERSION).withGroup(ApiConstants.SITEWHERE_API_GROUP)
	    .withPlural(ApiConstants.SITEWHERE_MICROSERVICE_CRD_PLURAL).build();

    public SiteWhereMicroserviceController(KubernetesClient client, SharedInformerFactory informerFactory) {
	super(client, informerFactory);
    }

    /**
     * Create informer.
     */
    public SharedIndexInformer<SiteWhereMicroservice> createInformer() {
	return getInformerFactory().sharedIndexInformerForCustomResource(CONTEXT, SiteWhereMicroservice.class,
		SiteWhereMicroserviceList.class, RESYNC_PERIOD_MS);
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
     * Validates k8s resources associated with new SiteWhere microservice.
     */
    protected class MicroserviceCreationValidator extends MicroserviceWorkerRunnable {

	public MicroserviceCreationValidator(ResourceChangeType type, SiteWhereMicroservice instance) {
	    super(type, instance);
	}

	@Override
	public void run() {
	    LOGGER.info("Validating created SiteWhereInstance.");
	}
    }

    protected ExecutorService getWorkers() {
	return workers;
    }
}
