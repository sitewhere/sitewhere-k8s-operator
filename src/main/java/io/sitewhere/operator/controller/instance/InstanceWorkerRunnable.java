/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator.controller.instance;

import io.sitewhere.k8s.crd.instance.SiteWhereInstance;
import io.sitewhere.operator.controller.ResourceChangeType;

/**
 * Base class for workers that repond to instance resource updates.
 */
public abstract class InstanceWorkerRunnable implements Runnable {

    /** Resource change type */
    private ResourceChangeType type;

    /** Instance */
    private SiteWhereInstance instance;

    public InstanceWorkerRunnable(ResourceChangeType type, SiteWhereInstance instance) {
	this.type = type;
	this.instance = instance;
    }

    public ResourceChangeType getType() {
	return type;
    }

    public void setType(ResourceChangeType type) {
	this.type = type;
    }

    public SiteWhereInstance getInstance() {
	return instance;
    }

    public void setInstance(SiteWhereInstance instance) {
	this.instance = instance;
    }
}
