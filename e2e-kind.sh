#!/bin/bash
set -e

 # Usage: ./e2e-kind.sh [--keep]
KEEP=false
if [[ "$1" == "--keep" ]]; then
  KEEP=true
fi
#!/bin/bash
set -e

CLUSTER_NAME="fantasy-pokedex-e2e"
FRONTEND_IMAGE="ghcr.io/jaydee94/fantasy-pokedex-frontend:latest"
BACKEND_IMAGE="ghcr.io/jaydee94/fantasy-pokedex-backend:latest"
KIND_CONFIG="kind-config.yaml"

# Create kind config for extra ports if needed
cat <<EOF > $KIND_CONFIG
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane
    extraPortMappings:
      - containerPort: 3000
        hostPort: 3000
        protocol: TCP
      - containerPort: 3001
        hostPort: 3001
        protocol: TCP
      - containerPort: 8080
        hostPort: 8080
        protocol: TCP
EOF


# Start kind cluster with podman as provider
export KIND_EXPERIMENTAL_PROVIDER=podman
if kind get clusters | grep -q "^$CLUSTER_NAME$"; then
  echo "Kind cluster $CLUSTER_NAME already exists. Skipping cluster creation."
else
  systemd-run --scope --user -p "Delegate=yes" kind create cluster --name "$CLUSTER_NAME" --config "$KIND_CONFIG"
fi

# Build images
cd frontend
podman build -t $FRONTEND_IMAGE .
cd ../fantasy_pokedex
podman build -t $BACKEND_IMAGE .
cd ..

# Load images into kind
kind load docker-image $FRONTEND_IMAGE
kind load docker-image $BACKEND_IMAGE

 # Deploy Helm chart
cd chart
helm repo add bitnami https://charts.bitnami.com/bitnami
helm dependency build
helm upgrade --install fantasy-pokedex . \
  --set frontend.image.repository=$FRONTEND_IMAGE \
  --set backend.image.repository=$BACKEND_IMAGE \
  --set frontend.image.tag=latest \
  --set backend.image.tag=latest \
  --set postgresql.auth.password=postgres \
  --set postgresql.auth.username=postgres \
  --set postgresql.auth.database=fantasy_pokedex \
  --wait
cd ..

# Print status
kubectl get pods -A
kubectl get svc -A
kubectl get ingress -A || true

# End-to-end test: check health endpoints
kubectl wait --for=condition=ready pod -l app=fantasy-pokedex --timeout=120s
kubectl port-forward svc/fantasy-pokedex-frontend 3000:3000 &

sleep 5
PF1=$!

PF2=$!
curl -f http://localhost:3000/ || echo "Frontend not ready"
PF3=$!

# Trap for cleanup on exit or Ctrl+C
cleanup() {
  echo "Cleaning up port-forwards..."
  kill $PF1 $PF2 $PF3 2>/dev/null || true
}
trap cleanup EXIT INT TERM
curl -f http://localhost:3001/healthz || echo "Frontend healthz failed"
curl -f http://localhost:8080/healthz || echo "Backend healthz failed"

echo "End-to-end test completed."
if [ "$KEEP" = true ]; then
  echo "App is running. Press Ctrl+C to stop and clean up port-forwards."
  wait
else
  cleanup
  echo "End-to-end test completed."
fi
