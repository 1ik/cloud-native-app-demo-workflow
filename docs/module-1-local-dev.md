# Module 1: Local Development Setup

## Learning Objectives

By the end of this module, you will:
- Understand containerization basics
- Set up local development with Docker Compose
- Set up local development with Tilt (Kubernetes)
- Learn hot-reload workflows
- Understand service orchestration

## Why It Matters (Cloud-Native)

- **Consistency**: Same environment for all developers
- **Isolation**: Services don't conflict with local machine
- **Portability**: "Works on my machine" → "Works everywhere"
- **Kubernetes Practice**: Tilt lets you develop against K8s locally (production-like)
- **Foundation**: Containers are the building blocks of cloud-native

## What We'll Build

A minimal setup with:
- Simple Go API (one endpoint: `GET /health`)
- Simple React app (one page: "Hello World")
- PostgreSQL database (empty, just running)
- Redis cache (empty, just running)

**Not building**: Product features yet

## Prerequisites

- Docker and Docker Compose installed
- Kubernetes cluster (Docker Desktop K8s, minikube, or kind)
- Tilt installed (`brew install tilt-dev/tap/tilt` or see [Tilt docs](https://docs.tilt.dev/install.html))
- Go 1.21+
- Node.js 18+

## Part 1: Docker Compose Setup (Simpler)

### Step 1: Create Minimal Go API

Create `backend/cmd/api/main.go`:
```go
package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"service": "cloudnativeapp-api",
		})
	})
	
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
```

### Step 2: Create Dockerfile for Backend

Create `backend/Dockerfile`:
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o api ./cmd/api

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["./api"]
```

### Step 3: Create Minimal React App

Create `web/src/App.tsx`:
```tsx
import { useEffect, useState } from 'react';

function App() {
  const [health, setHealth] = useState<string>('checking...');

  useEffect(() => {
    fetch('http://localhost:8080/health')
      .then(res => res.json())
      .then(data => setHealth(JSON.stringify(data, null, 2)))
      .catch(err => setHealth(`Error: ${err.message}`));
  }, []);

  return (
    <div style={{ padding: '2rem' }}>
      <h1>CloudNativeApp - Hello World</h1>
      <p>Backend Health Check:</p>
      <pre>{health}</pre>
    </div>
  );
}

export default App;
```

### Step 4: Create Dockerfile for Frontend

Create `web/Dockerfile`:
```dockerfile
FROM node:18-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

### Step 5: Create docker-compose.yml

Create `infrastructure/docker-compose.yml`:
```yaml
version: '3.8'

services:
  backend:
    build:
      context: ../backend
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - GIN_MODE=debug
    volumes:
      - ../backend:/app
    depends_on:
      - postgres
      - redis

  frontend:
    build:
      context: ../web
      dockerfile: Dockerfile
    ports:
      - "3000:80"
    depends_on:
      - backend

  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: cloudnativeapp
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

volumes:
  postgres_data:
  redis_data:
```

### Step 6: Run with Docker Compose

```bash
cd infrastructure
docker-compose up --build
```

Access:
- Backend: http://localhost:8080/health
- Frontend: http://localhost:3000

## Part 2: Tilt Setup (Kubernetes - Advanced)

### Why Tilt?

Tilt provides:
- **Hot-reload in Kubernetes**: See changes instantly without rebuilding images
- **Production-like environment**: Develop against real K8s
- **Visual dashboard**: See all services and their status
- **Automatic port-forwarding**: Access services easily

### Step 1: Create Kubernetes Manifests

Create `infrastructure/k8s/base/backend/deployment.yaml`:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend
spec:
  replicas: 1
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
        ports:
        - containerPort: 8080
        env:
        - name: GIN_MODE
          value: "debug"
```

Create `infrastructure/k8s/base/backend/service.yaml`:
```yaml
apiVersion: v1
kind: Service
metadata:
  name: backend
spec:
  selector:
    app: backend
  ports:
  - port: 8080
    targetPort: 8080
  type: ClusterIP
```

### Step 2: Create Tiltfile

Create `infrastructure/Tiltfile`:
```python
# Backend service
docker_build('backend', './backend')
k8s_yaml('./k8s/base/backend')
k8s_resource('backend', port_forwards=8080)

# Frontend service
docker_build('frontend', './web')
k8s_yaml('./k8s/base/frontend')
k8s_resource('frontend', port_forwards=3000)

# Enable live updates
local_resource(
    'sync-backend',
    'rsync -avz ./backend/ /tmp/backend-sync',
    deps=['./backend'],
    resource_deps=['backend']
)
```

### Step 3: Run with Tilt

```bash
cd infrastructure
tilt up
```

Access Tilt UI: http://localhost:10350

## Key Concepts

### Dockerfile
- Multi-stage builds (smaller images)
- Layer caching (faster builds)
- Minimal base images (security)

### docker-compose.yml
- Service definitions
- Networking (services can communicate)
- Volumes (persistent data)
- Environment variables

### Tiltfile
- Kubernetes resource definitions
- Hot-reload configuration
- Port forwarding
- Resource dependencies

## Practice Tasks

1. ✅ Create minimal Go API with health endpoint
2. ✅ Create minimal React app
3. ✅ Write Dockerfiles for both
4. ✅ Create docker-compose.yml
5. ✅ Run everything locally with Docker Compose
6. ✅ Test hot-reload (change code, see it update)
7. ✅ Set up Tilt for Kubernetes development
8. ✅ Understand service networking

## Key Takeaways

- "I can containerize applications and run them locally"
- "I understand service orchestration with Docker Compose"
- "I can develop against Kubernetes locally with Tilt"
- "I know how to set up hot-reload for development"

## Next Steps

Once you're comfortable with local development:
- Module 2: OpenAPI Code Generation
- Module 3: Infrastructure as Code (Terraform)

## Troubleshooting

### Docker Compose Issues
- Check ports aren't already in use: `lsof -i :8080`
- Rebuild images: `docker-compose build --no-cache`
- View logs: `docker-compose logs -f`

### Tilt Issues
- Ensure Kubernetes is running: `kubectl cluster-info`
- Check Tilt logs: `tilt logs`
- Restart Tilt: `tilt down && tilt up`

