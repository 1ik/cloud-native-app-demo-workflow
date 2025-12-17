# CloudNativeApp - Cloud-Native Learning Project

## 🎯 Project Vision

This is a **learning-focused project** to master cloud-native development workflows. The app itself is intentionally minimal—a simple API + frontend used as a sandbox—while we focus on infrastructure, delivery, and operational patterns.

### What We're Building (End Goal)

A small reference app that:
- **Stays minimal** - Hello-world style API and React UI
- **Supports learning** - Used to practice cloud-native workflows end-to-end
- **Is technology-focused** - No domain/product features yet

### Current Focus: Phase 1 - Learning Cloud-Native Workflows

**We are NOT building the full product yet.**

**Learning Objectives:**
1. ✅ Local development with Docker Compose and Tilt (Kubernetes)
2. ✅ OpenAPI code generation (type-safe APIs)
3. ✅ Infrastructure as Code (Terraform)
4. ✅ CI/CD pipelines (GitHub Actions)
5. ✅ Cloud deployment (AWS ECS Fargate → EKS)
6. ✅ Multi-environment strategies (dev → staging → prod)

**Development Workflow:**
```
Local Development → Git → CI/CD → Cloud Deployment
     ↓                ↓      ↓          ↓
  Docker/Tilt    GitHub   Actions   AWS (ECS/EKS)
```

## 📁 Project Structure

```
cloudnativeapp/
├── backend/              # Go API (minimal hello world)
│   ├── cmd/api/         # Application entry point
│   ├── Dockerfile       # Production container
│   └── go.mod
├── web/                  # React frontend (minimal hello world)
│   ├── src/            # React source code
│   ├── Dockerfile      # Production container
│   └── package.json
├── infrastructure/       # Infrastructure as Code
│   ├── docker-compose.yml  # Local development (Docker Compose)
│   ├── Tiltfile         # Local Kubernetes development (Tilt)
│   └── terraform/       # AWS infrastructure (coming later)
├── api/                 # OpenAPI specification (coming later)
├── docs/                # Learning documentation
│   ├── project-vision.md
│   ├── learning-path.md
│   ├── module-1-local-dev.md
│   └── ...
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- **Docker** and **Docker Compose** (for local development)
- **Kubernetes** (Docker Desktop K8s, minikube, or kind) + **Tilt** (for K8s dev)
- **Go** 1.21+
- **Node.js** 18+
- **AWS CLI** (for later modules)

### Quick Start

#### Option 1: Docker Compose (Simpler)

```bash
# Start all services
docker-compose -f infrastructure/docker-compose.yml up

# Backend API: http://localhost:8080
# Frontend: http://localhost:3000
```

#### Option 2: Tilt (Kubernetes - Advanced)

```bash
# Start Tilt (requires local Kubernetes cluster)
tilt up

# Access Tilt UI: http://localhost:10350
# Services automatically deployed to local K8s
```

## 📚 Learning Path

See [docs/](docs/) for detailed guides:

1. **Module 1**: Local Development Setup (Docker Compose + Tilt)
2. **Module 2**: OpenAPI Code Generation
3. **Module 3**: Infrastructure as Code (Terraform)
4. **Module 4**: CI Pipeline (GitHub Actions)
5. **Module 5**: Container Registry (ECR)
6. **Module 6**: Cloud Deployment (ECS Fargate)
7. **Module 7**: CD Pipeline (Automated Deployment)
8. **Module 8**: Multi-Environment Strategy
9. **Module 9**: Database and Managed Services
10. **Module 10**: Production Deployment

## 🎓 Development Philosophy

**Learn First, Build Later:**
- Start with minimal "Hello World" apps
- Master workflows before adding features
- Understand "why" behind each tool
- Build interview-ready cloud-native skills

**Workflow Focus:**
1. ✅ Local development (Docker Compose + Tilt)
2. ✅ Git workflow (feature branches, PRs)
3. ✅ CI/CD pipelines (automated testing & deployment)
4. ✅ Cloud deployment (AWS)
5. ✅ Then: Product development

## 📖 Documentation

- [Project Vision](docs/project-vision.md) - What we're building and why
- [Learning Path](docs/learning-path.md) - Complete learning roadmap
- [Module 1: Local Development](docs/module-1-local-dev.md) - Docker Compose & Tilt setup

## 🏗️ Current Status

✅ **Phase 1 - Learning Workflows** (Foundation Complete)
- [x] Project structure and documentation
- [x] Minimal Go API (Hello World)
- [x] Minimal React frontend (Hello World)
- [x] Docker Compose setup with hot-reload
- [x] Tilt setup for Kubernetes development
- [ ] CI/CD pipeline (Next: Module 4)
- [ ] Cloud deployment (Next: Module 6)

**Current Focus:** Testing local development setup, then moving to OpenAPI code generation (Module 2)

**Next Phase:** CI/CD and cloud deployment workflows, then product development

## 🤝 Contributing

This is a personal learning project. Focus is on understanding cloud-native principles and building interview-ready skills.

## 📝 License

MIT
