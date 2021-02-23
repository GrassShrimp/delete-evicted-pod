# delete-evicted-pod

This is a tool for delete evicted pod on kubernetes cluster

## Prerequisites

- [docker](https://www.docker.com/products/docker-desktop)
- [skaffold](https://skaffold.dev/docs/install/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Usage

switch to kubernetes context which you want to clean evicted pod

```bash
kubectl config use-context <context> 
```

build image and deploy to kubernetes cluster

```bash
skaffold dev
```
