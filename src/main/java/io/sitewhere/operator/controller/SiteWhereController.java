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
import java.util.logging.Level;
import java.util.logging.Logger;

import io.fabric8.kubernetes.client.KubernetesClient;
import io.fabric8.kubernetes.client.informers.ResourceEventHandler;
import io.fabric8.kubernetes.client.informers.SharedIndexInformer;
import io.fabric8.kubernetes.client.informers.cache.Cache;
import io.fabric8.kubernetes.client.informers.cache.Lister;
import io.sitewhere.operator.crd.SiteWhereInstance;

public class SiteWhereController {

    public static Logger LOGGER = Logger.getLogger(SiteWhereController.class.getName());
    public static String APP_LABEL = "app";

    private BlockingQueue<String> workQueue;
    private SharedIndexInformer<SiteWhereInstance> instanceInformer;
    private Lister<SiteWhereInstance> instanceLister;
    private KubernetesClient kubernetesClient;

    public SiteWhereController(KubernetesClient kubernetesClient,
	    SharedIndexInformer<SiteWhereInstance> instanceInformer) {
	this.kubernetesClient = kubernetesClient;
	this.instanceLister = new Lister<>(instanceInformer.getIndexer(), "default");
	this.instanceInformer = instanceInformer;
	this.workQueue = new ArrayBlockingQueue<>(1024);
    }

    public void create() {
	getInstanceInformer().addEventHandler(new ResourceEventHandler<SiteWhereInstance>() {
	    @Override
	    public void onAdd(SiteWhereInstance instance) {
		enqueueInstance(instance);
	    }

	    @Override
	    public void onUpdate(SiteWhereInstance oldInstance, SiteWhereInstance newInstance) {
		if (oldInstance.getMetadata().getResourceVersion() == newInstance.getMetadata().getResourceVersion()) {
		    return;
		}
		enqueueInstance(newInstance);
	    }

	    @Override
	    public void onDelete(SiteWhereInstance pod, boolean b) {
	    }
	});
    }

    private void enqueueInstance(SiteWhereInstance instance) {
	LOGGER.log(Level.INFO, "enqueueInstance(" + instance.getMetadata().getName() + ")");
	String key = Cache.metaNamespaceKeyFunc(instance);
	LOGGER.log(Level.INFO, "Going to enqueue key " + key);
	if (key != null && !key.isEmpty()) {
	    LOGGER.log(Level.INFO, "Adding item to workqueue");
	    getWorkQueue().add(key);
	}
    }

    public void run() {
	LOGGER.log(Level.INFO, "Starting SiteWhere controller");
	while (!getInstanceInformer().hasSynced()) {
	}

	while (true) {
	    try {
		LOGGER.log(Level.INFO, "Waiting on work queue.");
		if (getWorkQueue().isEmpty()) {
		    LOGGER.log(Level.INFO, "Work Queue is empty");
		}
		String key = getWorkQueue().take();
		LOGGER.log(Level.INFO, "Got " + key);
		if (key == null || key.isEmpty() || (!key.contains("/"))) {
		    LOGGER.log(Level.WARNING, "invalid resource key: " + key);
		}

		// Get the resource's name from key which is in format namespace/name
		String name = key.split("/")[1];
		SiteWhereInstance instance = getInstanceLister().get(key.split("/")[1]);
		if (instance == null) {
		    LOGGER.log(Level.SEVERE, "SiteWhereInstance " + name + " in workqueue no longer exists.");
		    return;
		}
		LOGGER.info("Would be reconciling SiteWhereInstance.");
	    } catch (InterruptedException interruptedException) {
		LOGGER.log(Level.SEVERE, "controller interrupted..");
	    }
	}
    }

    protected BlockingQueue<String> getWorkQueue() {
	return workQueue;
    }

    protected SharedIndexInformer<SiteWhereInstance> getInstanceInformer() {
	return instanceInformer;
    }

    protected Lister<SiteWhereInstance> getInstanceLister() {
	return instanceLister;
    }

    protected KubernetesClient getKubernetesClient() {
	return kubernetesClient;
    }
}
