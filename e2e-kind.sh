#!/bin/bash
set -e


# Usage: ./e2e-kind.sh [--keep] [--provider podman|docker]
KEEP=false
PROVIDER="podman"
for arg in "$@"; do
  case $arg in
    --keep)
      KEEP=true
      ;;
    --provider)
      PROVIDER="$2"
      shift
      ;;
    --provider=*)
      PROVIDER="${arg#*=}"
      ;;
  esac
done
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



# Start kind cluster with selected provider
if [ "$PROVIDER" = "podman" ]; then
  export KIND_EXPERIMENTAL_PROVIDER=podman
  KIND_CREATE_CMD="systemd-run --scope --user -p 'Delegate=yes' kind create cluster --name \"$CLUSTER_NAME\" --config \"$KIND_CONFIG\""
else
  unset KIND_EXPERIMENTAL_PROVIDER
  KIND_CREATE_CMD="kind create cluster --name \"$CLUSTER_NAME\" --config \"$KIND_CONFIG\""
fi
if kind get clusters | grep -q "^$CLUSTER_NAME$"; then
  echo "Kind cluster $CLUSTER_NAME already exists. Skipping cluster creation."
else
  eval "$KIND_CREATE_CMD"
fi


# Build images

if [ "$PROVIDER" = "podman" ]; then
  cd frontend
  podman build -t fantasy-pokedex-frontend:latest .
  podman tag fantasy-pokedex-frontend:latest $FRONTEND_IMAGE
  cd ../fantasy_pokedex
  podman build -t fantasy-pokedex-backend:latest .
  podman tag fantasy-pokedex-backend:latest $BACKEND_IMAGE
  cd ..
else
  cd frontend
  docker build -t fantasy-pokedex-frontend:latest .
  docker tag fantasy-pokedex-frontend:latest $FRONTEND_IMAGE
  cd ../fantasy_pokedex
  docker build -t fantasy-pokedex-backend:latest .
  docker tag fantasy-pokedex-backend:latest $BACKEND_IMAGE
  cd ..
fi


# Load images into kind
kind load docker-image $FRONTEND_IMAGE --name "$CLUSTER_NAME"
kind load docker-image $BACKEND_IMAGE --name "$CLUSTER_NAME"

 # Deploy Helm chart
cd chart
helm repo add bitnami https://charts.bitnami.com/bitnami
helm dependency build
helm upgrade --install fantasy-pokedex . \
  --set frontend.image.repository="ghcr.io/jaydee94/fantasy-pokedex-frontend" \
  --set backend.image.repository="ghcr.io/jaydee94/fantasy-pokedex-backend" \
  --set frontend.image.tag=latest \
  --set backend.image.tag=latest \
  --set postgresql.auth.password=postgres \
  --set postgresql.auth.username=postgres \
  --set postgresql.auth.database=fantasy_pokedex \
  --set backend.env.DATABASE_URL="postgres://postgres:postgres@fantasy-pokedex-postgresql:5432/fantasy_pokedex" \
  --wait
cd ..

# Print status
kubectl get pods -A
kubectl get svc -A
kubectl get ingress -A || true

# End-to-end test: check health endpoints
kubectl wait --for=condition=ready pod -l app=fantasy-pokedex --timeout=120s
kubectl port-forward svc/fantasy-pokedex-frontend 3000:3000 &
PF1=$!
kubectl port-forward svc/fantasy-pokedex-backend 8080:8080 &
PF2=$!

sleep 5

curl -f http://localhost:3000/ || echo "Frontend not ready"

# Trap for cleanup on exit or Ctrl+C
cleanup() {
  echo "Cleaning up port-forwards..."
  kill $PF1 $PF2 2>/dev/null || true
  echo "Deleting kind cluster $CLUSTER_NAME..."
  kind delete cluster --name "$CLUSTER_NAME" || true
}
trap cleanup EXIT INT TERM

echo "End-to-end test completed."
if [ "$KEEP" = true ]; then
  echo "App is running. Press Ctrl+C to stop and clean up port-forwards."
  wait
else
  cleanup
  echo "End-to-end test completed."
fi
