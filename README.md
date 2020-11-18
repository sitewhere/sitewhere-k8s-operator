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
