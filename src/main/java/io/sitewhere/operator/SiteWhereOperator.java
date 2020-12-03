/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator;

import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.ThreadFactory;
import java.util.concurrent.atomic.AtomicInteger;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import io.fabric8.kubernetes.client.Config;
import io.fabric8.kubernetes.client.ConfigBuilder;
import io.fabric8.kubernetes.client.DefaultKubernetesClient;
import io.fabric8.kubernetes.client.NamespacedKubernetesClient;
import io.fabric8.kubernetes.client.informers.SharedInformerFactory;
import io.sitewhere.operator.controller.instance.SiteWhereInstanceController;
import io.sitewhere.operator.controller.microservice.SiteWhereMicroserviceController;
import io.sitewhere.operator.controller.tenant.SiteWhereTenantController;

/**
 * Main class for operator.
 */
public class SiteWhereOperator {

    /** Static logger instance */
    private static Logger LOGGER = LoggerFactory.getLogger(SiteWhereOperator.class);

    /** Tracks shutdown */
    private static CountDownLatch SHUTDOWN = new CountDownLatch(1);

    public static void main(String[] args) {
	// Thread pool used for controller loops.
	ExecutorService loopsPool = Executors.newFixedThreadPool(3, new EventLoopThreadFactory());

	// Catch shutdown signal.
	addShutdownHook();

	LOGGER.info("\n\nStarting SiteWhere Kubernetes Operator\n");

	NamespacedKubernetesClient client = null;
	SharedInformerFactory informerFactory = null;
	try {
	    Config config = new ConfigBuilder().withNamespace(null).build();
	    client = new DefaultKubernetesClient(config);
	    informerFactory = client.informers();

	    // Create controllers.
	    SiteWhereInstanceController instanceController = new SiteWhereInstanceController(client, informerFactory);
	    SiteWhereMicroserviceController microserviceController = new SiteWhereMicroserviceController(client,
		    informerFactory);
	    SiteWhereTenantController tenantController = new SiteWhereTenantController(client, informerFactory);

	    // Start informers.
	    informerFactory.startAllRegisteredInformers();

	    // Start event loops.
	    loopsPool.execute(instanceController.createEventLoop());
	    loopsPool.execute(microserviceController.createEventLoop());
	    loopsPool.execute(tenantController.createEventLoop());

	    SHUTDOWN.await();
	} catch (InterruptedException e) {
	    LOGGER.warn("Operator shutting down.");
	} finally {
	    if (informerFactory != null) {
		LOGGER.info("Shutting down informers.");
		informerFactory.stopAllRegisteredInformers();
	    }
	    if (loopsPool != null) {
		LOGGER.info("Shutting down event loops.");
		loopsPool.shutdownNow();
	    }
	    if (client != null) {
		LOGGER.info("Closing client.");
		client.close();
	    }
	}
    }

    /**
     * Add hook to close everything at shutdown.
     */
    protected static void addShutdownHook() {
	Runtime.getRuntime().addShutdownHook(new Thread() {
	    public void run() {
		SHUTDOWN.countDown();
	    }
	});
    }

    /** Used for naming controller event loop threads */
    private static class EventLoopThreadFactory implements ThreadFactory {

	/** Counts threads */
	private AtomicInteger counter = new AtomicInteger();

	public Thread newThread(Runnable r) {
	    return new Thread(r, "Event Loop " + counter.incrementAndGet());
	}
    }
}
