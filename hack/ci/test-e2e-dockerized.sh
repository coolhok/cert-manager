#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail
set -o xtrace

# Runs the e2e tests, producing JUnit-style XML test
# reports in ${WORKSPACE}/artifacts. This script is intended to be run from
# cert-manager-test container with a cert-manager repo mapped in. See
# github.com/jetstack/test-infra/scenarios/cert-manager_e2e.py

while true; do if kubectl get nodes; then break; fi; echo "Waiting 5s for kubernetes to be ready..."; sleep 5; done

echo "Installing helm with cluster-admin privileges..."
# Create a service account for tiller
kubectl create serviceaccount -n kube-system tiller
# Bind the tiller service account to the cluster-admin role
kubectl create clusterrolebinding tiller-binding --clusterrole=cluster-admin --serviceaccount kube-system:tiller
# Deploy tiller
helm init --service-account=tiller

echo "Exposing nginx-ingress service with a stable IP (10.80.0.123)"
# Setup service for nginx ingress controller. A DNS entry for *.certmanager.kubernetes.network has been setup to point to 10.80.0.123 for e2e tests
while true; do if kubectl get rc nginx-ingress-controller -n kube-system; then break; fi; echo "Waiting 5s for nginx-ingress-controller rc to be installed..."; sleep 5; done
kubectl expose -n kube-system --port 80 --target-port 80 --type ClusterIP rc nginx-ingress-controller --cluster-ip 10.80.0.123

cd /go/src/github.com/jetstack/cert-manager
echo "Running e2e tests"
make e2e_test E2E_NGINX_CERTIFICATE_DOMAIN=certmanager.kubernetes.network
