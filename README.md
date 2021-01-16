[![Build Status](https://travis-ci.org/sitewhere/sitewhere-k8s-operator.svg?branch=master)](https://travis-ci.org/sitewhere/sitewhere-k8s-operator) [![Go Report Card](https://goreportcard.com/badge/github.com/sitewhere/sitewhere-k8s-operator)](https://goreportcard.com/report/github.com/sitewhere/sitewhere-k8s-operator) [![GoDoc](https://godoc.org/github.com/sitewhere/sitewhere-k8s-operator?status.svg)](https://godoc.org/github.com/sitewhere/sitewhere-k8s-operator) 

![SiteWhere](https://s3.amazonaws.com/sitewhere-branding/SiteWhereLogo.svg)

# SiteWhere Kubernetes Operator

Manages SiteWhere lifecycle based on Kubernetes Custom Resource Definitions (CRDs)
being added, updated, and deleted.

## Install SiteWhere Custom Resource Definitions

```console
make deploy
```

## Cleanup

```console
make undeploy
```
