# SiteWhere Kubernetes Operator

Manages SiteWhere lifecycle based on Kubernetes Custom Resource Definitions (CRDs)
being added, updated, and deleted.

## Install SiteWhere Custom Resource Definitions

Before the operator may be used, the SiteWhere CRDs must be installed. From
the `installer/crds` folder run the following:

```
helm install --name sitewhere-crds .
```

