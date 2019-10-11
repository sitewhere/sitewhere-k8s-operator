/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator.crd;

import io.fabric8.kubernetes.api.builder.Function;
import io.fabric8.kubernetes.client.CustomResourceDoneable;

public class DoneableSiteWhereInstance extends CustomResourceDoneable<SiteWhereInstance> {

    public DoneableSiteWhereInstance(SiteWhereInstance resource,
	    Function<SiteWhereInstance, SiteWhereInstance> function) {
	super(resource, function);
    }
}
