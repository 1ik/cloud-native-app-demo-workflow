# CloudNativeApp - Cloud-Native Learning Project

## 🎯 Project Vision

This is a **learning-focused project** to master cloud-native development workflows. The app itself is intentionally minimal—a simple API + frontend used as a sandbox—while we focus on infrastructure, delivery, and operational patterns.

## Important
Read `.personal/AGENTS.md`.

### What We're Building (End Goal)

A small reference app that:
- **Stays minimal** - Hello-world style API and React UI
- **Supports learning** - Used to practice cloud-native workflows end-to-end
- **Is technology-focused** - No domain/product features yet

### Current Focus: Phase 1 - Learning Cloud-Native Workflows

**We are NOT building the full product yet.**

**Learning Objectives:**
1. 🧱 Local cloud-native dev workflow (being rebuilt piece-by-piece)
2. ✅ OpenAPI code generation (type-safe APIs)
3. ✅ Infrastructure as Code (Terraform)
4. ✅ CI/CD pipelines (GitHub Actions)
5. ✅ Cloud deployment (AWS ECS Fargate → EKS)
6. ✅ Multi-environment strategies (dev → staging → prod)

**Development Workflow:**
```
Local Development → Git → CI/CD → Cloud Deployment
     ↓                ↓      ↓          ↓
  K8s/Tilt/GitOps GitHub  Actions   AWS (EKS/ECS)
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
├── infrastructure/       # Infrastructure as Code (rebuilding from scratch)
├── api/                 # OpenAPI specification (coming later)
├── docs/                # Learning documentation
│   ├── project-vision.md
│   ├── learning-path.md
│   └── ...
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- **Docker**
- **Kubernetes** (Docker Desktop K8s, minikube, or kind)
- **Tilt** (when we reintroduce it)
- **Go** 1.21+
- **Node.js** 18+
- **AWS CLI** (for later modules)

### Quick Start

The local cloud-native dev environment is intentionally being rebuilt from zero (to build muscle memory). For now, run services directly:

```bash
task server:run
task frontend:install
task frontend:start
```

## 📚 Learning Path

See [docs/](docs/) for detailed guides:

1. **Module 1**: Local cloud-native dev workflow (rebuild from scratch)
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
1. 🧱 Local cloud-native dev workflow (rebuild from scratch)
2. ✅ Git workflow (feature branches, PRs)
3. ✅ CI/CD pipelines (automated testing & deployment)
4. ✅ Cloud deployment (AWS)
5. ✅ Then: Product development

## 📖 Documentation

- [Project Vision](docs/project-vision.md) - What we're building and why
- [Learning Path](docs/learning-path.md) - Complete learning roadmap

## 🏗️ Current Status

✅ **Phase 1 - Learning Workflows** (Foundation Complete)
- [x] Project structure and documentation
- [x] Minimal Go API (Hello World)
- [x] Minimal React frontend (Hello World)
- [ ] Local cloud-native dev setup (rebuilding from scratch)
- [ ] CI/CD pipeline (Next: Module 4)
- [ ] Cloud deployment (Next: Module 6)

**Current Focus:** Rebuilding local cloud-native development setup from scratch, then moving to OpenAPI code generation (Module 2)

**Next Phase:** CI/CD and cloud deployment workflows, then product development

## 🤝 Contributing

This is a personal learning project. Focus is on understanding cloud-native principles and building interview-ready skills.

## 📝 License

MIT
