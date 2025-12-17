# Tilt + Kubernetes Local Development Guide

## 🎯 What You'll Learn

This guide teaches you:
- How to set up a local Kubernetes cluster
- How to use Tilt for Kubernetes development
- How to develop against Kubernetes (production-like environment)
- Hot-reload in Kubernetes
- Kubernetes basics (Deployments, Services, Pods)

## 📚 Prerequisites

You should know:
- ✅ Docker (images, containers, volumes, networks)
- ✅ Docker Compose (orchestration)

You'll learn:
- 🆕 Kubernetes basics
- 🆕 Tilt for local K8s development
- 🆕 Kubernetes manifests (YAML files)

## 🤔 Why Kubernetes + Tilt?

### The Problem with Docker Compose

Docker Compose is great for local development, but:
- Production uses Kubernetes (not Docker Compose)
- Different environments = different problems
- "Works locally" ≠ "Works in production"

### The Solution: Tilt

Tilt lets you:
- **Develop against real Kubernetes** (same as production)
- **Hot-reload in K8s** (see changes instantly)
- **Visual dashboard** (see all services)
- **Automatic port-forwarding** (access services easily)

**Think of it like**: Docker Compose, but for Kubernetes, with superpowers.

## 🏗️ Kubernetes Basics (Quick Primer)

If you know Docker, here's the Kubernetes equivalent:

| Docker Concept | Kubernetes Equivalent | What It Does |
|---------------|----------------------|--------------|
| Container | **Pod** | Runs one or more containers |
| Image | **Image** (same) | Container image |
| `docker run` | **Deployment** | Manages Pods (replicas, updates) |
| Port mapping | **Service** | Exposes Pods to network |
| `docker-compose.yml` | **Manifests (YAML)** | Defines all resources |

### Key Kubernetes Concepts

**Pod**: The smallest deployable unit. Contains one or more containers.

**Deployment**: Manages Pods. Handles:
- How many replicas (copies)
- Rolling updates
- Rollbacks
- Health checks

**Service**: Network access to Pods. Like a load balancer:
- Routes traffic to Pods
- Provides stable IP/name
- Handles service discovery

**Namespace**: Logical grouping of resources (like folders).

## 📦 Step 1: Install Kubernetes

You need a local Kubernetes cluster. Choose one:

### Option A: Docker Desktop (Easiest)

1. Open Docker Desktop
2. Go to Settings → Kubernetes
3. Check "Enable Kubernetes"
4. Click "Apply & Restart"
5. Wait for Kubernetes to start (green icon)

**Verify:**
```bash
kubectl cluster-info
# Should show: Kubernetes control plane is running at https://...
```

### Option B: minikube

```bash
# Install minikube
brew install minikube  # macOS
# or download from https://minikube.sigs.k8s.io/docs/start/

# Start minikube
minikube start

# Verify
kubectl cluster-info
```

### Option C: kind (Kubernetes in Docker)

```bash
# Install kind
brew install kind  # macOS
# or: go install sigs.k8s.io/kind@v0.20.0

# Create cluster
kind create cluster --name cloudnativeapp

# Verify
kubectl cluster-info
```

**Recommendation**: Start with Docker Desktop (easiest).

## 🔧 Step 2: Install Tilt

### macOS:
```bash
brew install tilt-dev/tap/tilt
```

### Linux:
```bash
curl -fsSL https://raw.githubusercontent.com/tilt-dev/tilt/master/scripts/install.sh | bash
```

### Windows:
```powershell
# Using Chocolatey
choco install tilt

# Or download from: https://github.com/tilt-dev/tilt/releases
```

**Verify:**
```bash
tilt version
# Should show version number
```

## 🎯 Step 3: Understand the Setup

Before we start, let's understand what we have:

### Project Structure:
```
infrastructure/
├── Tiltfile              # Tilt configuration
└── k8s/
    └── base/
        ├── backend/
        │   ├── deployment.yaml   # How to run backend
        │   └── service.yaml     # How to access backend
        ├── frontend/
        │   ├── deployment.yaml
        │   └── service.yaml
        ├── postgres/
        │   ├── deployment.yaml
        │   └── service.yaml
        └── redis/
            ├── deployment.yaml
            └── service.yaml
```

### What Tilt Does:

1. **Reads Tiltfile** → Understands your setup
2. **Builds Docker images** → Creates images for backend/frontend
3. **Applies K8s manifests** → Deploys to your cluster
4. **Watches for changes** → Detects file edits
5. **Hot-reloads** → Updates running containers
6. **Port-forwards** → Makes services accessible

## 🚀 Step 4: Start Tilt

### Navigate to Infrastructure Directory:
```bash
cd infrastructure
```

### Start Tilt:
```bash
tilt up
```

**What happens:**
1. Tilt reads `Tiltfile`
2. Builds Docker images
3. Deploys to Kubernetes
4. Sets up port forwarding
5. Opens browser to Tilt UI (http://localhost:10350)

### First Time Setup:

Tilt will ask about your Kubernetes context. It should detect:
- `docker-desktop` (if using Docker Desktop)
- `minikube` (if using minikube)
- `kind-cloudnativeapp` (if using kind)

**Allow it**: Tilt needs permission to manage resources.

## 🎨 Step 5: Use Tilt UI

Tilt opens a web UI at **http://localhost:10350**

### What You'll See:

**Resource List** (left side):
- `backend` - Your Go API
- `frontend` - Your React app
- `postgres` - Database
- `redis` - Cache

**Status Indicators**:
- 🟢 Green = Running
- 🟡 Yellow = Building/Updating
- 🔴 Red = Error

**Click on a Resource** to see:
- Build logs
- Runtime logs
- Port forwarding info
- Resource usage

### Key Features:

1. **Live Logs**: See logs from all services in real-time
2. **Port Forwarding**: Automatically forwards ports (8080, 3000, etc.)
3. **Build Status**: See when images are building
4. **Resource Usage**: CPU/Memory usage

## 🔥 Step 6: Test Hot-Reload

### Test Backend Hot-Reload:

1. **Edit** `backend/cmd/api/main.go`:
```go
r.GET("/health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "super healthy",  // Changed this
        "service": "cloudnativeapp-api",
    })
})
```

2. **Save the file**

3. **Watch Tilt UI**:
   - Backend resource turns yellow (updating)
   - Logs show rebuild
   - Turns green when done

4. **Test**:
```bash
curl http://localhost:8080/health
# Should show: {"status":"super healthy",...}
```

**No manual restart needed!** Tilt detected the change and updated it.

### Test Frontend Hot-Reload:

1. **Edit** `web/src/App.tsx`:
```tsx
<h1>CloudNativeApp - Hello World - UPDATED!</h1>
```

2. **Save the file**

3. **Watch Tilt UI**:
   - Frontend resource updates
   - Vite hot-reloads

4. **Refresh browser** at http://localhost:3000
   - See your changes!

## 🌐 Step 7: Access Your Services

Tilt automatically sets up port forwarding:

- **Backend API**: http://localhost:8080
  - Health: http://localhost:8080/health
- **Frontend**: http://localhost:3000
- **PostgreSQL**: localhost:5432
- **Redis**: localhost:6379

**No manual `kubectl port-forward` needed!**

## 🔍 Step 8: Understand What's Happening

### Check Kubernetes Resources:

```bash
# See all pods (running containers)
kubectl get pods

# See all services
kubectl get services

# See all deployments
kubectl get deployments

# Describe a pod (detailed info)
kubectl describe pod backend-xxxxx
```

### Example Output:
```
NAME                      READY   STATUS    RESTARTS   AGE
backend-xxxxx            1/1     Running   0          5m
frontend-xxxxx           1/1     Running   0          5m
postgres-xxxxx           1/1     Running   0          5m
redis-xxxxx              1/1     Running   0          5m
```

## 📝 Understanding the Tiltfile

Let's break down what's in `Tiltfile`:

```python
# Load Kubernetes manifests
k8s_yaml([
    './k8s/base/backend/deployment.yaml',
    './k8s/base/backend/service.yaml',
    # ... more manifests
])
```
**What it does**: Tells Tilt to deploy these K8s resources.

```python
# Backend service
docker_build(
    'backend',
    '../backend',
    dockerfile='../backend/Dockerfile.dev',
    live_update=[
        sync('../backend', '/app'),
        run('go mod download', trigger=['../backend/go.mod']),
    ]
)
```
**What it does**:
- Builds Docker image from `../backend`
- Uses `Dockerfile.dev` (development version)
- `live_update`: Syncs code changes to container
- `sync`: Copies files
- `run`: Runs commands when files change

```python
k8s_resource('backend', port_forwards=8080)
```
**What it does**: 
- Manages the `backend` K8s resource
- Forwards port 8080 to localhost

## 🛠️ Common Commands

### Tilt Commands:

```bash
# Start Tilt
tilt up

# Stop Tilt (keeps resources running)
tilt down

# Restart Tilt
tilt down && tilt up

# View logs
tilt logs backend

# Get status
tilt get uiresources
```

### Kubernetes Commands:

```bash
# Get all pods
kubectl get pods

# Get logs from a pod
kubectl logs -f deployment/backend

# Describe a resource
kubectl describe deployment backend

# Delete a resource (Tilt will recreate it)
kubectl delete deployment backend

# Exec into a pod (like docker exec)
kubectl exec -it deployment/backend -- sh
```

## 🐛 Troubleshooting

### Issue: Tilt can't find Kubernetes

**Error**: `No Kubernetes context found`

**Solution**:
```bash
# Check if Kubernetes is running
kubectl cluster-info

# If not, start it:
# Docker Desktop: Enable in Settings
# minikube: minikube start
# kind: kind create cluster

# Check context
kubectl config get-contexts

# Set context
kubectl config use-context docker-desktop
```

### Issue: Images not building

**Error**: Build fails in Tilt UI

**Solution**:
```bash
# Check Tilt logs
tilt logs

# Check if Docker is running
docker ps

# Rebuild manually
tilt down
tilt up
```

### Issue: Port forwarding not working

**Error**: Can't access http://localhost:8080

**Solution**:
```bash
# Check if port is forwarded in Tilt UI
# Click on resource → see port forwarding

# Or manually forward:
kubectl port-forward deployment/backend 8080:8080
```

### Issue: Changes not hot-reloading

**Error**: Code changes don't appear

**Solution**:
1. Check Tilt UI - is resource updating?
2. Check file paths in `live_update` section
3. Ensure you're editing files in watched directories
4. Check container logs: `kubectl logs -f deployment/backend`

### Issue: Pods not starting

**Error**: Pods stuck in `Pending` or `CrashLoopBackOff`

**Solution**:
```bash
# Check pod status
kubectl get pods

# Describe pod for details
kubectl describe pod <pod-name>

# Check logs
kubectl logs <pod-name>

# Common issues:
# - Not enough resources (CPU/Memory)
# - Image pull errors
# - Configuration errors
```

## 🎓 Key Learnings

### What You've Learned:

1. **Kubernetes Basics**:
   - Pods, Deployments, Services
   - How K8s manages containers
   - K8s manifests (YAML)

2. **Tilt Workflow**:
   - Develop against real Kubernetes
   - Hot-reload in K8s
   - Visual dashboard
   - Automatic port-forwarding

3. **Production-Like Development**:
   - Same environment as production
   - Catch issues early
   - Understand K8s before deploying

### Why This Matters:

- **Interview Ready**: Shows K8s experience
- **Production Skills**: Develop like you deploy
- **Faster Debugging**: Issues appear locally
- **Team Alignment**: Everyone uses same setup

## 🚀 Next Steps

Now that you have Tilt + Kubernetes working:

1. **Practice**: Make changes, see hot-reload
2. **Explore**: Check K8s resources with `kubectl`
3. **Learn**: Read the K8s manifests to understand structure
4. **Module 2**: Move to OpenAPI code generation
5. **Module 3**: Learn Terraform (Infrastructure as Code)

## 📚 Additional Resources

- **Kubernetes Docs**: https://kubernetes.io/docs/
- **Tilt Docs**: https://docs.tilt.dev/
- **kubectl Cheat Sheet**: https://kubernetes.io/docs/reference/kubectl/cheatsheet/

## ✅ Success Checklist

You're ready when:
- [ ] Kubernetes cluster is running
- [ ] Tilt is installed
- [ ] `tilt up` works without errors
- [ ] Tilt UI shows all services running
- [ ] Can access backend at http://localhost:8080
- [ ] Can access frontend at http://localhost:3000
- [ ] Hot-reload works (edit code, see changes)
- [ ] Understand basic K8s concepts

---

**Congratulations!** You now have a production-like local development environment with Kubernetes! 🎉

