# Quick Start Guide

## Prerequisites

Before you begin, ensure you have:

- **Docker** and **Docker Compose** installed
- **Kubernetes** cluster (optional, for Tilt):
  - Docker Desktop with Kubernetes enabled, OR
  - minikube, OR
  - kind (Kubernetes in Docker)
- **Tilt** (optional, for K8s dev): `brew install tilt-dev/tap/tilt` or see [Tilt docs](https://docs.tilt.dev/install.html)
- **Go** 1.21+ (for local Go development)
- **Node.js** 18+ (for local React development)

## Option 1: Docker Compose (Recommended for Beginners)

### Step 1: Navigate to Infrastructure Directory

```bash
cd infrastructure
```

### Step 2: Start All Services

```bash
docker-compose up --build
```

This will:
- Build the Go backend API
- Build the React frontend
- Start PostgreSQL database
- Start Redis cache
- Set up networking between services

### Step 3: Access the Application

- **Backend API**: http://localhost:8080
  - Health check: http://localhost:8080/health
- **Frontend**: http://localhost:3000
- **PostgreSQL**: localhost:5432
- **Redis**: localhost:6379

### Step 4: Test Hot-Reload

1. Edit `backend/cmd/api/main.go` - change the health message
2. Watch the terminal - the container should automatically rebuild
3. Refresh http://localhost:8080/health to see changes

1. Edit `web/src/App.tsx` - change the text
2. Watch the terminal - Vite should hot-reload
3. Refresh http://localhost:3000 to see changes

### Step 5: Stop Services

```bash
# Press Ctrl+C to stop, or in another terminal:
docker-compose down
```

## Option 2: Tilt (Kubernetes - Advanced)

### Step 1: Ensure Kubernetes is Running

```bash
# Check if Kubernetes is running
kubectl cluster-info

# If using Docker Desktop, enable Kubernetes in Settings
# If using minikube:
minikube start

# If using kind:
kind create cluster
```

### Step 2: Navigate to Infrastructure Directory

```bash
cd infrastructure
```

### Step 3: Start Tilt

```bash
tilt up
```

This will:
- Build Docker images
- Deploy to your local Kubernetes cluster
- Set up port forwarding
- Enable hot-reload

### Step 4: Access Tilt UI

Open http://localhost:10350 in your browser to see:
- All services and their status
- Build logs
- Resource usage
- Port forwarding information

### Step 5: Access the Application

Tilt automatically sets up port forwarding:
- **Backend API**: http://localhost:8080
- **Frontend**: http://localhost:3000

### Step 6: Test Hot-Reload

1. Edit any file in `backend/` or `web/`
2. Watch Tilt UI - it should detect changes and rebuild
3. Changes appear automatically (no manual refresh needed)

### Step 7: Stop Tilt

```bash
# Press Ctrl+C, or in another terminal:
tilt down
```

## Troubleshooting

### Docker Compose Issues

**Port already in use:**
```bash
# Check what's using the port
lsof -i :8080
lsof -i :3000

# Kill the process or change ports in docker-compose.yml
```

**Build fails:**
```bash
# Rebuild without cache
docker-compose build --no-cache

# Check logs
docker-compose logs -f backend
docker-compose logs -f frontend
```

**Services can't connect:**
```bash
# Ensure all services are on the same network
docker network ls
docker network inspect cloudnativeapp_app-network
```

### Tilt Issues

**Kubernetes not running:**
```bash
# Check cluster status
kubectl cluster-info

# If using Docker Desktop, enable Kubernetes in Settings
# If using minikube:
minikube status
minikube start
```

**Tilt can't find Kubernetes context:**
```bash
# List available contexts
kubectl config get-contexts

# Set the context
kubectl config use-context docker-desktop
# or
kubectl config use-context minikube
```

**Images not building:**
```bash
# Check Tilt logs
tilt logs

# Restart Tilt
tilt down
tilt up
```

**Port forwarding not working:**
```bash
# Check port forwarding in Tilt UI
# Or manually forward ports:
kubectl port-forward deployment/backend 8080:8080
kubectl port-forward deployment/frontend 3000:80
```

## Next Steps

Once local development is working:

1. **Module 2**: Set up OpenAPI code generation
2. **Module 3**: Learn Terraform for infrastructure
3. **Module 4**: Set up CI/CD with GitHub Actions

See [Learning Path](learning-path.md) for the complete roadmap.

## Development Tips

### Backend Development

- Use `Dockerfile.dev` for hot-reload with Air
- Go code changes trigger automatic rebuilds
- Check logs: `docker-compose logs -f backend`

### Frontend Development

- Use `Dockerfile.dev` for hot-reload with Vite
- React code changes appear instantly
- Check logs: `docker-compose logs -f frontend`

### Database Access

```bash
# Connect to PostgreSQL
docker-compose exec postgres psql -U postgres -d cloudnativeapp

# Connect to Redis
docker-compose exec redis redis-cli
```

### Viewing Logs

```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f backend
docker-compose logs -f frontend
```

## Getting Help

- Check the [Module 1 Guide](module-1-local-dev.md) for detailed explanations
- Review [Project Vision](project-vision.md) to understand the goals
- See [Learning Path](learning-path.md) for the complete roadmap

