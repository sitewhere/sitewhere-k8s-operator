/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator.controller.tenant;

import io.sitewhere.k8s.crd.tenant.SiteWhereTenant;
import io.sitewhere.operator.controller.ResourceChangeType;

/**
 * Base class for workers that repond to tenant resource updates.
 */
public abstract class TenantWorkerRunnable implements Runnable {

    /** Resource change type */
    private ResourceChangeType type;

    /** Tenant */
    private SiteWhereTenant tenant;

    public TenantWorkerRunnable(ResourceChangeType type, SiteWhereTenant tenant) {
	this.type = type;
	this.tenant = tenant;
    }

    public ResourceChangeType getType() {
	return type;
    }

    public void setType(ResourceChangeType type) {
	this.type = type;
    }

    public SiteWhereTenant getTenant() {
	return tenant;
    }

    public void setTenant(SiteWhereTenant tenant) {
	this.tenant = tenant;
    }
}
