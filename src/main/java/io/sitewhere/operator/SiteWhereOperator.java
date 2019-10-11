/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator;

import io.fabric8.kubernetes.client.DefaultKubernetesClient;
import io.fabric8.kubernetes.client.KubernetesClient;
import io.fabric8.kubernetes.client.dsl.base.CustomResourceDefinitionContext;
import io.fabric8.kubernetes.client.informers.SharedIndexInformer;
import io.fabric8.kubernetes.client.informers.SharedInformerFactory;
import io.sitewhere.operator.controller.SiteWhereController;
import io.sitewhere.operator.crd.SiteWhereInstance;
import io.sitewhere.operator.crd.SiteWhereInstanceList;

/**
 * Main class for operator.
 */
public class SiteWhereOperator {

    public static void main(String[] args) {
	try (KubernetesClient client = new DefaultKubernetesClient()) {

	    CustomResourceDefinitionContext instanceDefinitionContext = new CustomResourceDefinitionContext.Builder()
		    .withVersion("v1alpha1").withScope("Namespaced").withGroup("sitewhere.io")
		    .withPlural("sitewhereinstances").build();

	    SharedInformerFactory informerFactory = client.informers();

	    SharedIndexInformer<SiteWhereInstance> instanceSharedIndexInformer = informerFactory
		    .sharedIndexInformerForCustomResource(instanceDefinitionContext, SiteWhereInstance.class,
			    SiteWhereInstanceList.class, 10 * 60 * 1000);
	    SiteWhereController sitewhereController = new SiteWhereController(client, instanceSharedIndexInformer);

	    sitewhereController.create();
	    informerFactory.startAllRegisteredInformers();

	    sitewhereController.run();
	}
    }
}
