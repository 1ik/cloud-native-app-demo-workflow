# Kubernetes Basics for Docker Users

## 🎯 Quick Reference: Docker → Kubernetes

If you know Docker, here's the Kubernetes equivalent:

| Docker | Kubernetes | Purpose |
|--------|-----------|---------|
| `docker run` | `kubectl run` | Run a container |
| Container | **Pod** | Smallest deployable unit |
| Image | Image (same) | Container image |
| `docker-compose.yml` | **Manifests (YAML)** | Define resources |
| `docker-compose up` | `kubectl apply` | Deploy resources |
| Port mapping | **Service** | Expose pods |
| Volume | **Volume** (same concept) | Persistent storage |

## 🏗️ Core Kubernetes Concepts

### 1. Pod

**What it is**: The smallest deployable unit in Kubernetes.

**Think of it as**: A wrapper around one or more containers.

**Key points**:
- Pods share network and storage
- Usually one container per pod
- Pods are ephemeral (can be recreated)

**Example**:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: backend-pod
spec:
  containers:
  - name: backend
    image: backend:latest
    ports:
    - containerPort: 8080
```

### 2. Deployment

**What it is**: Manages Pods. Handles replicas, updates, rollbacks.

**Think of it as**: Docker Compose service, but smarter.

**Key points**:
- Manages multiple Pod replicas
- Handles rolling updates (zero downtime)
- Can rollback on failure
- Self-healing (recreates failed pods)

**Example**:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend
spec:
  replicas: 3  # Run 3 copies
  selector:
    matchLabels:
      app: backend
  template:
    metadata:
      labels:
        app: backend
    spec:
      containers:
      - name: backend
        image: backend:latest
```

**Why use Deployments?**
- High availability (multiple replicas)
- Rolling updates (no downtime)
- Automatic rollback
- Self-healing

### 3. Service

**What it is**: Network access to Pods. Provides stable IP/name.

**Think of it as**: Load balancer + DNS name.

**Key points**:
- Routes traffic to Pods
- Provides stable IP address
- Service discovery (pods can find each other)
- Load balancing

**Example**:
```yaml
apiVersion: v1
kind: Service
metadata:
  name: backend
spec:
  selector:
    app: backend  # Routes to pods with this label
  ports:
  - port: 8080
    targetPort: 8080
  type: ClusterIP  # Internal only
```

**Service Types**:
- `ClusterIP`: Internal only (default)
- `NodePort`: Expose on node IP
- `LoadBalancer`: Cloud load balancer
- `ExternalName`: External service

### 4. Namespace

**What it is**: Logical grouping of resources.

**Think of it as**: Folders for organizing resources.

**Key points**:
- Isolate resources
- Resource quotas
- Access control

**Common namespaces**:
- `default`: Default namespace
- `kube-system`: System resources
- `kube-public`: Public resources

## 🔄 How It All Works Together

### Simple Flow:

```
User Request
    ↓
Service (backend)  ← Provides stable IP
    ↓
Deployment (backend)  ← Manages pods
    ↓
Pod (backend-xxxxx)  ← Runs container
    ↓
Container (backend)  ← Your app
```

### Example: Our Backend

1. **Deployment** creates Pods:
   ```bash
   kubectl get pods
   # NAME                    READY   STATUS
   # backend-abc123          1/1     Running
   ```

2. **Service** routes traffic:
   ```bash
   kubectl get services
   # NAME      TYPE        CLUSTER-IP      PORT(S)
   # backend   ClusterIP   10.96.1.2      8080/TCP
   ```

3. **Access**:
   - From outside: `kubectl port-forward service/backend 8080:8080`
   - From inside cluster: `http://backend:8080`

## 📝 Understanding Our Manifests

Let's look at our backend deployment:

```yaml
apiVersion: apps/v1        # K8s API version
kind: Deployment           # Type of resource
metadata:
  name: backend            # Name of resource
spec:
  replicas: 1             # How many copies
  selector:
    matchLabels:
      app: backend         # Which pods to manage
  template:               # Pod template
    spec:
      containers:
      - name: backend
        image: backend:latest
        ports:
        - containerPort: 8080
```

**Breaking it down**:
- `replicas: 1` = Run 1 copy (can scale to 3, 5, etc.)
- `selector` = Which pods this deployment manages
- `template` = What each pod should look like
- `containers` = What containers to run

## 🎮 Common kubectl Commands

### View Resources:

```bash
# Get all pods
kubectl get pods

# Get all services
kubectl get services

# Get all deployments
kubectl get deployments

# Get everything
kubectl get all
```

### Describe Resources:

```bash
# Detailed info about a pod
kubectl describe pod <pod-name>

# Detailed info about a deployment
kubectl describe deployment backend

# Detailed info about a service
kubectl describe service backend
```

### Logs:

```bash
# Logs from a pod
kubectl logs <pod-name>

# Logs from a deployment
kubectl logs -f deployment/backend

# Logs from all pods with label
kubectl logs -l app=backend
```

### Exec (like docker exec):

```bash
# Execute command in pod
kubectl exec -it <pod-name> -- sh

# Execute in deployment
kubectl exec -it deployment/backend -- sh
```

### Port Forwarding:

```bash
# Forward service port
kubectl port-forward service/backend 8080:8080

# Forward pod port
kubectl port-forward <pod-name> 8080:8080
```

### Apply/Delete:

```bash
# Apply manifest
kubectl apply -f k8s/base/backend/deployment.yaml

# Delete resource
kubectl delete deployment backend

# Delete all resources
kubectl delete all --all
```

## 🔍 Debugging Tips

### Pod Not Starting:

```bash
# Check pod status
kubectl get pods

# Describe pod (see events)
kubectl describe pod <pod-name>

# Check logs
kubectl logs <pod-name>

# Common issues:
# - ImagePullBackOff: Can't pull image
# - CrashLoopBackOff: Container keeps crashing
# - Pending: Not enough resources
```

### Service Not Working:

```bash
# Check service
kubectl get service backend

# Check endpoints (which pods service routes to)
kubectl get endpoints backend

# Test from inside cluster
kubectl run -it --rm debug --image=busybox -- sh
# Then: wget -O- http://backend:8080/health
```

### Deployment Issues:

```bash
# Check deployment status
kubectl get deployment backend

# Check rollout status
kubectl rollout status deployment/backend

# View rollout history
kubectl rollout history deployment/backend

# Rollback
kubectl rollout undo deployment/backend
```

## 🎓 Key Differences from Docker

### 1. Pods vs Containers

- **Docker**: You run containers directly
- **Kubernetes**: You run Pods (which contain containers)

### 2. Networking

- **Docker**: Containers on same network can talk
- **Kubernetes**: Services provide DNS names (e.g., `backend:8080`)

### 3. Scaling

- **Docker Compose**: Manual scaling
- **Kubernetes**: Automatic scaling, self-healing

### 4. Updates

- **Docker**: Stop, rebuild, start
- **Kubernetes**: Rolling updates (zero downtime)

### 5. Configuration

- **Docker Compose**: One file
- **Kubernetes**: Multiple YAML files (more flexible)

## ✅ Quick Checklist

You understand Kubernetes when you can:

- [ ] Explain what a Pod is
- [ ] Explain what a Deployment does
- [ ] Explain what a Service does
- [ ] Use `kubectl get pods` to see running pods
- [ ] Use `kubectl logs` to view logs
- [ ] Use `kubectl describe` to debug issues
- [ ] Understand how Services route to Pods
- [ ] Know how to scale a Deployment

## 🚀 Next Steps

1. **Practice**: Use `kubectl` commands to explore
2. **Read Manifests**: Understand our YAML files
3. **Experiment**: Change replicas, see what happens
4. **Learn More**: Kubernetes documentation

---

**Remember**: Kubernetes is Docker, but with orchestration superpowers! 🚀

