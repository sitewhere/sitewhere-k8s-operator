[![Build Status](https://travis-ci.org/sitewhere/sitewhere-k8s-operator.svg?branch=master)](https://travis-ci.org/sitewhere/sitewhere-k8s-operator) [![Go Report Card](https://goreportcard.com/badge/github.com/sitewhere/sitewhere-k8s-operator)](https://goreportcard.com/report/github.com/sitewhere/sitewhere-k8s-operator) [![GoDoc](https://godoc.org/github.com/sitewhere/sitewhere-k8s-operator?status.svg)](https://godoc.org/github.com/sitewhere/sitewhere-k8s-operator) 

# SiteWhere Kubernetes Operator

Manages SiteWhere lifecycle based on Kubernetes Custom Resource Definitions (CRDs)
being added, updated, and deleted.

## Install SiteWhere Custom Resource Definitions

```console
kubectl apply -f deploy/cert-manager/cert-manager.yaml
make deploy
```

## Cleanup

```console
make undeploy
kubectl delete -f deploy/cert-manager/cert-manager.yaml
```

## Operator SDK

### sitewhere.io

```console
operator-sdk create api --group sitewhere.io --version v1alpha4 --kind SiteWhereInstance --controller --namespaced=false --resource
operator-sdk create webhook --defaulting --group sitewhere.io --version v1alpha4 --kind SiteWhereInstance --programmatic-validation
operator-sdk create api --group sitewhere.io --version v1alpha4 --kind SiteWhereMicroservice --controller --namespaced=true --resource
operator-sdk create api --group sitewhere.io --version v1alpha4 --kind SiteWhereTenant --controller --namespaced=true --resource
operator-sdk create api --group sitewhere.io --version v1alpha4 --kind SiteWhereTenantEngine --controller --namespaced=true --resource
```

### scripting.sitewhere.io

```console
operator-sdk create api --group scripting.sitewhere.io --version v1alpha4 --kind SiteWhereScript --controller=false --namespaced=true --resource
operator-sdk create api --group scripting.sitewhere.io --version v1alpha4 --kind SiteWhereScriptCategory --controller=false --namespaced=false --resource
operator-sdk create api --group scripting.sitewhere.io --version v1alpha4 --kind SiteWhereScriptTemplate --controller=false --namespaced=false --resource
operator-sdk create api --group scripting.sitewhere.io --version v1alpha4 --kind SiteWhereScriptVersion --controller=false --namespaced=true --resource
```

### templates.sitewhere.io

```console
operator-sdk create api --group templates.sitewhere.io --version v1alpha4 --kind InstanceConfigurationTemplate --controller=false --namespaced=false --resource
operator-sdk create api --group templates.sitewhere.io --version v1alpha4 --kind InstanceDatasetTemplate --controller=false --namespaced=false --resource
operator-sdk create api --group templates.sitewhere.io --version v1alpha4 --kind TenantConfigurationTemplate --controller=false --namespaced=false --resource
operator-sdk create api --group templates.sitewhere.io --version v1alpha4 --kind TenantDatasetTemplate --controller=false --namespaced=false --resource
operator-sdk create api --group templates.sitewhere.io --version v1alpha4 --kind TenantEngineConfigurationTemplate --controller=false --namespaced=false --resource
operator-sdk create api --group templates.sitewhere.io --version v1alpha4 --kind TenantEngineDatasetTemplate --controller=false --namespaced=false --resource
```
