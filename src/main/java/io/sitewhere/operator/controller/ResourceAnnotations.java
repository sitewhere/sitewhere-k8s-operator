/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator.controller;

public class ResourceAnnotations {

    /** Annotation for Prometheus scrape enabled */
    public static final String ANN_PROMETHEUS_SCRAPE = "prometheus.io/scrape";

    /** Annotation for Prometheus scheme */
    public static final String ANN_PROMETHEUS_SCHEME = "prometheus.io/scheme";

    /** Annotation for Prometheus port */
    public static final String ANN_PROMETHEUS_PORT = "prometheus.io/port";
}
