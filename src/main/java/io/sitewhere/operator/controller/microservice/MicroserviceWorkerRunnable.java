/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator.controller.microservice;

import io.sitewhere.k8s.crd.controller.ResourceChangeType;
import io.sitewhere.k8s.crd.microservice.SiteWhereMicroservice;

/**
 * Base class for workers that repond to microservice resource updates.
 */
public abstract class MicroserviceWorkerRunnable implements Runnable {

    /** Resource change type */
    private ResourceChangeType type;

    /** Microservice */
    private SiteWhereMicroservice microservice;

    public MicroserviceWorkerRunnable(ResourceChangeType type, SiteWhereMicroservice microservice) {
	this.type = type;
	this.microservice = microservice;
    }

    public ResourceChangeType getType() {
	return type;
    }

    public void setType(ResourceChangeType type) {
	this.type = type;
    }

    public SiteWhereMicroservice getMicroservice() {
	return microservice;
    }

    public void setMicroservice(SiteWhereMicroservice microservice) {
	this.microservice = microservice;
    }
}
