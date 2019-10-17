/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator.controller;

public class ResourceChange {

    /** Type of change */
    private ResourceChangeType type;

    /** Key */
    private String key;

    public ResourceChange(ResourceChangeType type, String key) {
	this.type = type;
	this.key = key;
    }

    public ResourceChangeType getType() {
	return type;
    }

    public void setType(ResourceChangeType type) {
	this.type = type;
    }

    public String getKey() {
	return key;
    }

    public void setKey(String key) {
	this.key = key;
    }
}
