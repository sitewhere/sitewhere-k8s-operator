# SiteWhere Kubernetes Operator

Manages SiteWhere lifecycle based on Kubernetes Custom Resource Definitions (CRDs)
being added, updated, and deleted.

## Install SiteWhere Custom Resource Definitions

Before the operator may be used, the SiteWhere CRDs must be installed. Install the
CRDs via Helm using the following command:

```
helm install --name sitewhere-crds installer/crds/.
```

## Install Default Instance and Tenant Templates

In order to bootstrap an instance, default instance and tenant templates must
be installed. These templates determine the default system configuration and
may be customized after installation. The command below will install the 
default templates:

```
helm install --name sitewhere-templates installer/templates/.
```

## Install SiteWhere Operator

The SiteWhere operator is the orchrestrator which uses the CRDs and templates
in order to realize SiteWhere instances at runtime. Install the operator
via Helm as shown below:

```
helm install --name sitewhere-operator installer/operator/.
```
