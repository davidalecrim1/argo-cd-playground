# Argo CD Playground: GitOps with Kind, Kubernetes & Go

## Overview
This project is a hands-on playground to learn and demonstrate GitOps using Argo CD, Kubernetes, and Kind. It features a simple Go web application, containerized and deployed to a local Kubernetes cluster managed by Kind, with manifests and automation for continuous delivery via Argo CD.

### What Was Done
- Developed a minimal Go web app with health and ping endpoints.
- Dockerized the app for ARM64 and built/published with Makefile.
- Created Kubernetes manifests (Deployment, Service, Argo CD Application) for automated deployment.
- Documented step-by-step setup for Kind, Argo CD, and app deployment.

## Technologies Used
- Kind (Kubernetes in Docker) - v0.27.0
- Argo CD (GitOps continuous delivery) - v2.14.11
- Go (sample app) - v1.24
- Docker - v27.4.0
- Helm - v3.17.3
- kubectl - v1.33.0

## Quick Start

### 1. Install and Setup Kind
---

For MacOS:
```bash
brew install kind
```

Check installation:
```bash
kind version
```

Create a cluster:
```bash
kind create cluster
```

To interact with your kind cluster, ensure that kubectl is installed. You can install it via Homebrew:
```bash
brew install kubectl
```

After installing kubectl, you can verify the cluster's status:
```bash
kubectl cluster-info --context kind-kind
kubectl get nodes
```

The config from the kind context is in:
```bash
code ~/.kube/config
```


### 2. Install Argo CD
---


Install Helm in MacOS:
```bash
brew install helm
```

Follow instructions below using Helm:
```bash
helm repo add argo https://argoproj.github.io/argo-helm
helm repo update

helm install argocd argo/argo-cd --namespace argocd --create-namespace
helm install argocd argo/argo-cd --namespace argocd --create-namespace -f values.yaml
```

Port forward ArgoCD to the localhost:
```bash
kubectl port-forward service/argocd-server -n argocd 8080:443
```

Get the default password from the installation:
```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```


### 3. Build & Deploy the Sample Go App
---

Use the created app in `sample-go-app` folder.

Build the image locally using the Makefile:
```bash
make build
```

Load the image to kind:
```bash
kind load docker-image sample-go-app:v1.0.1
```

Apply the Argo CD Application manifest:
```bash
cd manifest && kubectl apply -f application.yaml
```

Forward the Go app port to localhost:
```bash
kubectl port-forward service/sample-go-app-service 8081:8081
```