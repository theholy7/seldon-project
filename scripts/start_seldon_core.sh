#! /bin/bash
export KUBECONFIG="$HOME/.kube/config"

# This script deploys Seldon Core to Minikube.
echo "Minikube info"
minikube version
minikube status
minikube ip

echo "kubectl info"
kubectl config current-context

echo "using helm to deploy seldon core"
# work around helm v3 not doing namespaces correctly
kubectl config set-context minikube --namespace seldon-system
kubectl create namespace seldon-system
sleep 2
helm --kube-context minikube install seldon-core seldon-core-operator --repo https://storage.googleapis.com/seldon-charts --set usageMetrics.enabled=true 
