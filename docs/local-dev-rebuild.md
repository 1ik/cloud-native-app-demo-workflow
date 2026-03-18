# Local Dev Rebuild (from scratch)

This repo intentionally removed all local “cloud-native dev” infrastructure so you can rebuild it **piece by piece** and develop real muscle memory.

## What still exists

- `backend/`: Go API
- `web/`: React frontend
- `Taskfile.yml`: task runner for commands

## What was removed (on purpose)

- `infrastructure/docker-compose.yml`
- all Kubernetes manifests under `infrastructure/k8s/**`
- `infrastructure/Tiltfile`
- local-dev docs that referenced the above

## Rebuild order (recommended)

1. **Run apps without containers**
   - Goal: verify the app works at all.
2. **Containerize one service (backend)**
   - Goal: build image, run container, understand ports/env.
3. **Kubernetes without Tilt**
   - Goal: write a `Deployment` + `Service` for backend.
4. **Add Tilt for backend**
   - Goal: learn `k8s_yaml`, `docker_build`, `k8s_resource`, port-forward.
5. **Add frontend**
6. **Add Postgres + Redis**
7. **Add GitOps (Argo CD) locally**
8. **Move the same GitOps pattern to EKS**

## Next action

Start at step 1 and tell the agent what Kubernetes you’re using:

- Docker Desktop Kubernetes, kind, or minikube

