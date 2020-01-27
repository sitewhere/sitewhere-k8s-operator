/*
 * Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com
 *
 * The software in this package is published under the terms of the CPAL v1.0
 * license, a copy of which has been included with this distribution in the
 * LICENSE.txt file.
 */
package io.sitewhere.operator.controller;

import java.util.HashMap;
import java.util.Map;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.google.common.base.CaseFormat;

import io.sitewhere.k8s.crd.ResourceLabels;
import io.sitewhere.k8s.crd.SiteWhereKubernetesClient;
import io.sitewhere.k8s.crd.microservice.SiteWhereMicroservice;
import io.sitewhere.k8s.crd.tenant.SiteWhereTenant;
import io.sitewhere.k8s.crd.tenant.SiteWhereTenantList;
import io.sitewhere.k8s.crd.tenant.configuration.TenantConfigurationTemplate;
import io.sitewhere.k8s.crd.tenant.engine.SiteWhereTenantEngine;
import io.sitewhere.k8s.crd.tenant.engine.SiteWhereTenantEngineList;
import io.sitewhere.k8s.crd.tenant.engine.SiteWhereTenantEngineSpec;
import io.sitewhere.k8s.crd.tenant.engine.configuration.TenantEngineConfigurationTemplate;
import io.sitewhere.operator.controller.microservice.SiteWhereMicroserviceController;

/**
 * Utility methods for common operations used by operator.
 */
public class OperatorUtils {

    /** Static logger instance */
    private static Logger LOGGER = LoggerFactory.getLogger(OperatorUtils.class);

    /**
     * Get all tenants.
     * 
     * @param client
     * @return
     */
    public static SiteWhereTenantList getAllTenants(SiteWhereKubernetesClient client) {
	return client.getTenants().list();
    }

    /**
     * Get map of tenant engines for a tenant (indexed by microservice name).
     * 
     * @param client
     * @param tenant
     * @return
     */
    public static Map<String, SiteWhereTenantEngine> getTenantEnginesForTenantByMicroservice(
	    SiteWhereKubernetesClient client, SiteWhereTenant tenant) {
	SiteWhereTenantEngineList list = client.getTenantEngines().inNamespace(tenant.getMetadata().getNamespace())
		.withLabel(ResourceLabels.LABEL_SITEWHERE_TENANT, tenant.getMetadata().getName()).list();
	Map<String, SiteWhereTenantEngine> byMicroservice = new HashMap<>();
	for (SiteWhereTenantEngine engine : list.getItems()) {
	    String microservice = engine.getMetadata().getLabels().get(ResourceLabels.LABEL_SITEWHERE_MICROSERVICE);
	    if (microservice != null) {
		byMicroservice.put(microservice, engine);
	    }
	}
	return byMicroservice;
    }

    /**
     * Get map of tenant engines for a microservice (indexed by tenant id).
     * 
     * @param client
     * @param microservice
     * @return
     */
    public static Map<String, SiteWhereTenantEngine> getTenantEnginesForMicroserviceByTenant(
	    SiteWhereKubernetesClient client, SiteWhereMicroservice microservice) {
	SiteWhereTenantEngineList list = client.getTenantEngines()
		.inNamespace(microservice.getMetadata().getNamespace())
		.withLabel(ResourceLabels.LABEL_SITEWHERE_MICROSERVICE, microservice.getMetadata().getName()).list();
	Map<String, SiteWhereTenantEngine> byMicroservice = new HashMap<>();
	for (SiteWhereTenantEngine engine : list.getItems()) {
	    String tenantId = engine.getMetadata().getLabels().get(ResourceLabels.LABEL_SITEWHERE_TENANT);
	    if (tenantId != null) {
		byMicroservice.put(tenantId, engine);
	    }
	}
	return byMicroservice;
    }

    /**
     * Locate tenant engine configuration template based on tenant configuration
     * template associated with tenant.
     * 
     * @param client
     * @param tenant
     * @param microservice
     * @return
     */
    public static TenantEngineConfigurationTemplate getTenantEngineConfigurationTemplate(
	    SiteWhereKubernetesClient client, SiteWhereTenant tenant, SiteWhereMicroservice microservice) {
	TenantConfigurationTemplate tenantTemplate = client.getTenantConfigurationTemplates()
		.withName(tenant.getSpec().getConfigurationTemplate()).get();
	if (tenantTemplate == null) {
	    String message = String.format("Tenant references non-existent configuration template '%s'.",
		    tenant.getSpec().getConfigurationTemplate());
	    LOGGER.warn(message);
	    throw new RuntimeException(message);
	}
	String target = CaseFormat.LOWER_HYPHEN.to(CaseFormat.LOWER_CAMEL,
		SiteWhereMicroserviceController.getFunctionalArea(microservice));
	String tecTemplateName = tenantTemplate.getSpec().getTenantEngineTemplates().get(target);
	if (tecTemplateName == null) {
	    String message = String.format("Missing tenant engine template mapping for '%s'.",
		    SiteWhereMicroserviceController.getFunctionalArea(microservice));
	    LOGGER.warn(message);
	    return null;
	}

	return client.getTenantEngineConfigurationTemplates().withName(tecTemplateName).get();
    }

    /**
     * Create a new tenant engine for a tenant/microservice combination.
     * 
     * @param client
     * @param tenant
     * @param microservice
     */
    public static void createNewTenantEngine(SiteWhereKubernetesClient client, SiteWhereTenant tenant,
	    SiteWhereMicroservice microservice) {
	SiteWhereTenantEngine engine = new SiteWhereTenantEngine();
	String tenantEngineName = String.format("%s-%s-%s", tenant.getMetadata().getName(),
		microservice.getMetadata().getName(), String.valueOf(System.currentTimeMillis()));
	engine.getMetadata().setName(tenantEngineName);
	engine.getMetadata().setNamespace(tenant.getMetadata().getNamespace());

	Map<String, String> labels = new HashMap<>();
	labels.put(ResourceLabels.LABEL_SITEWHERE_TENANT, tenant.getMetadata().getName());
	labels.put(ResourceLabels.LABEL_SITEWHERE_MICROSERVICE, microservice.getMetadata().getName());
	labels.put(ResourceLabels.LABEL_SITEWHERE_FUNCTIONAL_AREA,
		SiteWhereMicroserviceController.getFunctionalArea(microservice));
	engine.getMetadata().setLabels(labels);

	// Look up tenant configuration template for tenant/microservice combination.
	TenantEngineConfigurationTemplate tecTemplate = getTenantEngineConfigurationTemplate(client, tenant,
		microservice);
	if (tecTemplate == null) {
	    LOGGER.warn("Unable to resolve default tenant engine configuration. Skipping engine creation.");
	    return;
	}

	// Copy template configuration into spec.
	SiteWhereTenantEngineSpec spec = new SiteWhereTenantEngineSpec();
	spec.setConfiguration(tecTemplate.getSpec().getConfiguration());
	engine.setSpec(spec);

	client.getTenantEngines().withName(tenantEngineName).createOrReplace(engine);
	LOGGER.info(String.format("Created new tenant engine for tenant `%s` microservice `%s`",
		tenant.getMetadata().getName(), microservice.getMetadata().getName()));
    }
}
