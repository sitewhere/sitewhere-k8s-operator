/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator.controller;

import io.fabric8.kubernetes.client.dsl.base.CustomResourceDefinitionContext;

public class ResourceContexts {

    /** Namespace scoped resource */
    public static final String SCOPE_NAMESPACED = "Namespaced";

    /** Cluster scoped resource */
    public static final String SCOPE_CLUSTER = "Cluster";

    /** Context used for accessing instances */
    public static final CustomResourceDefinitionContext INSTANCE_CONTEXT = new CustomResourceDefinitionContext.Builder()
	    .withVersion(ApiConstants.SITEWHERE_API_VERSION).withGroup(ApiConstants.SITEWHERE_API_GROUP)
	    .withPlural(ApiConstants.SITEWHERE_INSTANCE_CRD_PLURAL).withScope(SCOPE_CLUSTER).build();

    /** Context used for accessing microservices */
    public static final CustomResourceDefinitionContext MICROSERVICE_CONTEXT = new CustomResourceDefinitionContext.Builder()
	    .withVersion(ApiConstants.SITEWHERE_API_VERSION).withGroup(ApiConstants.SITEWHERE_API_GROUP)
	    .withPlural(ApiConstants.SITEWHERE_MICROSERVICE_CRD_PLURAL).withScope(SCOPE_NAMESPACED).build();

    /** Context used for accessing tenants */
    public static final CustomResourceDefinitionContext TENANT_CONTEXT = new CustomResourceDefinitionContext.Builder()
	    .withVersion(ApiConstants.SITEWHERE_API_VERSION).withGroup(ApiConstants.SITEWHERE_API_GROUP)
	    .withPlural(ApiConstants.SITEWHERE_TENANT_CRD_PLURAL).withScope(SCOPE_NAMESPACED).build();
}
