# Fantasy Pokedex Helm Chart

This Helm chart deploys the Fantasy Pokedex frontend and backend in a single pod, along with a Bitnami PostgreSQL database.

## Features
- Frontend and backend run together in one pod
- Bitnami PostgreSQL as a subchart

## Installation

```bash
helm dependency build
helm install fantasy-pokedex .
```


## Configurable Values

| Value | Description | Default |
|-------|-------------|---------|
| `frontend.image.repository` | Frontend container image repository | `ghcr.io/jaydee94/fantasy-pokedex-frontend` |
| `frontend.image.tag` | Frontend container image tag | `latest` |
| `frontend.image.pullPolicy` | Frontend image pull policy | `IfNotPresent` |
| `frontend.service.type` | Frontend service type | `ClusterIP` |
| `frontend.service.port` | Frontend service port | `3000` |
| `frontend.env` | Frontend environment variables (map) | `{}` |
| `backend.image.repository` | Backend container image repository | `ghcr.io/jaydee94/fantasy-pokedex-backend` |
| `backend.image.tag` | Backend container image tag | `latest` |
| `backend.image.pullPolicy` | Backend image pull policy | `IfNotPresent` |
| `backend.service.type` | Backend service type | `ClusterIP` |
| `backend.service.port` | Backend service port | `8080` |
| `backend.env.DB_HOST` | Database hostname | `fantasy-pokedex-postgresql` |
| `backend.env.DB_PORT` | Database port | `5432` |
| `backend.env.DB_USER` | Database user | `postgres` |
| `backend.env.DB_PASS` | Database password (can be sourced from secret) | `postgres` |
| `backend.env.DB_NAME` | Database name | `fantasy_pokedex` |
| `backend.env.ADMIN_PASSWORD` | Initial admin password (can be sourced from secret) | `""` |
| `backend.secret.dbPassSecret` | Name of secret for DB_PASS (optional) | `""` |
| `backend.secret.adminPasswordSecret` | Name of secret for ADMIN_PASSWORD (optional) | `""` |
| `postgresql.enabled` | Enable/disable PostgreSQL deployment | `true` |
| `postgresql.auth.username` | PostgreSQL user | `postgres` |
| `postgresql.auth.password` | PostgreSQL password (can be sourced from secret) | `postgres` |
| `postgresql.auth.database` | PostgreSQL database name | `fantasy_pokedex` |
| `postgresql.secret.passwordSecret` | Name of secret for PostgreSQL password (optional) | `""` |
| `postgresql.primary.persistence.enabled` | Enable persistent storage | `true` |
| `postgresql.primary.persistence.size` | Persistent volume size | `8Gi` |
| `postgresql.service.port` | PostgreSQL service port | `5432` |
| `resources` | Pod resource requests/limits | `{}` |
| `nodeSelector` | Node selector for pod scheduling | `{}` |
| `tolerations` | Tolerations for pod scheduling | `[]` |
| `affinity` | Affinity rules for pod scheduling | `{}` |
| `ingress.enabled` | Enable ingress for frontend | `false` |
| `ingress.host` | Ingress host | `fantasy-pokedex.local` |
| `ingress.path` | Ingress path | `/` |
| `ingress.pathType` | Ingress path type | `Prefix` |
| `ingress.className` | Ingress class name | `""` |
| `ingress.annotations` | Ingress annotations | `{}` |
| `ingress.tls` | Enable TLS for ingress | `false` |
| `ingress.tlsSecretName` | TLS secret name for ingress | `""` |


## Ingress

To expose the frontend via Ingress, set `ingress.enabled: true` and configure the host, path, and other options as needed. TLS can be enabled by setting `ingress.tls: true` and providing a secret name in `ingress.tlsSecretName`.

## Secret Integration

To use Kubernetes secrets for sensitive values (like passwords), create a secret containing the relevant key(s) and set the corresponding secret name in `values.yaml`. Example:

```yaml
backend:
  secret:
    dbPassSecret: my-db-secret
    adminPasswordSecret: my-admin-secret
postgresql:
  secret:
    passwordSecret: my-db-secret
```

The secret must contain keys named `DB_PASS` and/or `ADMIN_PASSWORD` for the backend, and `postgres` for PostgreSQL (see Bitnami chart docs).

