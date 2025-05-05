# Argo CD Playground with Kind

## Technologies
- Kind 

## Step by Step

### Install and Setup Kind

For MacOS:
```bash
brew install kind
```

Check installation:
```bash
kind version
```

Creater a cluster:
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

The config form the kind context is in:
```bash
code ~/.kube/config
```

### Install Argo CD

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

Port foward ArgoCD to the localhost:
```bash
kubectl port-forward service/argocd-server -n argocd 8080:443
```

Get the default password from the installation:
```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

### Running an app with ArgoCD
Use the created app in `sample-go-app` folder.

Build the image locally using the Makefile:
```bash
make build
```

Load the image to kind:
```bash
kind load docker-image sample-go-app:latest
```

Update the Application manifest as you see fit, then:
```bash
cd manifest && k apply -f application.yaml
```



