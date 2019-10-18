/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator.controller;

import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import io.fabric8.kubernetes.client.CustomResource;
import io.fabric8.kubernetes.client.KubernetesClient;
import io.fabric8.kubernetes.client.informers.ResourceEventHandler;
import io.fabric8.kubernetes.client.informers.SharedIndexInformer;
import io.fabric8.kubernetes.client.informers.SharedInformerFactory;
import io.fabric8.kubernetes.client.informers.cache.Cache;
import io.fabric8.kubernetes.client.informers.cache.Lister;

/**
 * Kubernetes controller for managing SiteWhere instances.
 */
public abstract class SiteWhereResourceController<T extends CustomResource> {

    /** Static logger instance */
    private static Logger LOGGER = LoggerFactory.getLogger(SiteWhereResourceController.class);

    /** Kubernetes client */
    private KubernetesClient client;

    /** Informer factory */
    private SharedInformerFactory informerFactory;

    /** Informer */
    private SharedIndexInformer<T> informer;

    /** Lister */
    private Lister<T> lister;

    /** Work queue */
    private BlockingQueue<ResourceChange<T>> workQueue;

    public SiteWhereResourceController(KubernetesClient client, SharedInformerFactory informerFactory) {
	this.client = client;
	this.informerFactory = informerFactory;
	this.informer = createInformer();
	this.lister = new Lister<>(getInformer().getIndexer());
	this.workQueue = new ArrayBlockingQueue<>(1024);
	startEventHandling();
    }

    /**
     * Start event handling.
     */
    protected void startEventHandling() {
	LOGGER.info("Starting event handler for resource updates.");
	getInformer().addEventHandler(new ResourceEventHandler<T>() {
	    @Override
	    public void onAdd(T resource) {
		String key = Cache.metaNamespaceKeyFunc(resource);
		if (key != null && !key.isEmpty()) {
		    LOGGER.info(String.format("Add with key %s for %s", key, resource.getMetadata().getName()));
		    getWorkQueue().add(new ResourceChange<T>(ResourceChangeType.CREATE, key, resource));
		}
	    }

	    @Override
	    public void onUpdate(T oldResource, T newResource) {
		if (oldResource.getMetadata().getResourceVersion() == newResource.getMetadata().getResourceVersion()) {
		    return;
		}
		String key = Cache.metaNamespaceKeyFunc(newResource);
		if (key != null && !key.isEmpty()) {
		    LOGGER.info(String.format("Update with key %s for %s", key, newResource.getMetadata().getName()));
		    getWorkQueue().add(new ResourceChange<T>(ResourceChangeType.UPDATE, key, newResource));
		}
	    }

	    @Override
	    public void onDelete(T resource, boolean b) {
		String key = Cache.metaNamespaceKeyFunc(resource);
		if (key != null && !key.isEmpty()) {
		    LOGGER.info(String.format("Delete with key %s for %s", key, resource.getMetadata().getName()));
		    getWorkQueue().add(new ResourceChange<T>(ResourceChangeType.DELETE, key, resource));
		}
	    }
	});
    }

    /**
     * Create a runnable event loop.
     * 
     * @return
     */
    public Runnable createEventLoop() {
	return new EventLoop();
    }

    /**
     * Create informer in subclass.
     * 
     * @return
     */
    public abstract SharedIndexInformer<T> createInformer();

    /**
     * Reconcile a resource change.
     * 
     * @param type
     * @param resource
     */
    public abstract void reconcileResourceChange(ResourceChangeType type, T resource);

    /**
     * Runs event loop in a separate thread.
     */
    public class EventLoop implements Runnable {

	/*
	 * @see java.lang.Runnable#run()
	 */
	public void run() {
	    LOGGER.info("Starting event processing loop.");
	    while (!getInformer().hasSynced()) {
		try {
		    Thread.sleep(200);
		} catch (InterruptedException e) {
		}
	    }

	    while (true) {
		try {
		    LOGGER.info("Waiting on work queue.");
		    if (getWorkQueue().isEmpty()) {
			LOGGER.info("Work Queue is empty");
		    }
		    ResourceChange<T> change = getWorkQueue().take();
		    String key = change.getKey();
		    LOGGER.info("Processing key " + key);
		    if (key == null || key.isEmpty()) {
			LOGGER.warn("Resource key was null or empty: " + key);
			continue;
		    }
		    String[] parts = key.split("/");
		    key = parts.length > 1 ? key : parts[0];

		    // Get the resource's name from key which is in format namespace/name
		    boolean isDelete = change.getType() == ResourceChangeType.DELETE;
		    T resource = isDelete ? change.getReference() : getLister().get(key);
		    if (resource == null) {
			LOGGER.error("Resource " + key + " in work queue no longer exists.");
			return;
		    }
		    reconcileResourceChange(change.getType(), resource);
		} catch (InterruptedException e) {
		    LOGGER.info("Shutting down event loop.");
		    return;
		} catch (Throwable t) {
		    LOGGER.error("Unhandled exception in controller.", t);
		}
	    }
	}
    }

    protected KubernetesClient getClient() {
	return client;
    }

    protected SharedInformerFactory getInformerFactory() {
	return informerFactory;
    }

    protected SharedIndexInformer<T> getInformer() {
	return informer;
    }

    protected Lister<T> getLister() {
	return lister;
    }

    protected BlockingQueue<ResourceChange<T>> getWorkQueue() {
	return workQueue;
    }
}
