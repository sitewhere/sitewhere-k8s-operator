/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator.controller;

public class ApiConstants {

    /** SiteWhere application name */
    public static final String SITEWHERE_APP_NAME = "sitewhere";

    /** SiteWhere CRD API version */
    public static final String SITEWHERE_API_VERSION = "v1alpha3";

    /** SiteWhere CRD group */
    public static final String SITEWHERE_API_GROUP = "sitewhere.io";

    /** SiteWhereInstance CRD plural */
    public static final String SITEWHERE_INSTANCE_CRD_PLURAL = "sitewhereinstances";

    /** InstanceConfigurationTemplate CRD plural */
    public static final String SITEWHERE_ICT_CRD_PLURAL = "instanceconfigurationtemplates";

    /** SiteWhereMicroservice CRD plural */
    public static final String SITEWHERE_MICROSERVICE_CRD_PLURAL = "sitewheremicroservices";

    /** SiteWhereInstance CRD name */
    public static final String SITEWHERE_INSTANCE_CRD_NAME = SITEWHERE_INSTANCE_CRD_PLURAL + "." + SITEWHERE_API_GROUP;

    /** InstanceConfigurationTemplate CRD name */
    public static final String SITEWHERE_ICT_CRD_NAME = SITEWHERE_ICT_CRD_PLURAL + "." + SITEWHERE_API_GROUP;

    /** SiteWhereMicroservice CRD name */
    public static final String SITEWHERE_MICROSERVICE_CRD_NAME = SITEWHERE_MICROSERVICE_CRD_PLURAL + "."
	    + SITEWHERE_API_GROUP;
}
