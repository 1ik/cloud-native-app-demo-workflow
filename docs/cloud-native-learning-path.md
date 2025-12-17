# Learning Path - Cloud-Native Development

## Overview

This document outlines the complete learning path for mastering cloud-native development workflows. Each module builds on the previous one, with clear learning objectives and hands-on practice.

## Module Structure

Each module includes:
1. **Learning Objectives** - What you'll understand
2. **Why It Matters** - Cloud-native principles
3. **Hands-On Practice** - Minimal example to try
4. **Key Takeaways** - Interview talking points

## Module 1: Local Development Setup

### Learning Objectives
- Understand containerization basics
- Set up local development with Docker Compose
- Set up local development with Tilt (Kubernetes)
- Learn hot-reload workflows
- Understand service orchestration

### Why It Matters
- **Consistency**: Same environment for all developers
- **Isolation**: Services don't conflict with local machine
- **Portability**: "Works on my machine" → "Works everywhere"
- **Kubernetes Practice**: Tilt lets you develop against K8s locally

### What We'll Build
- Minimal Go API (health endpoint)
- Minimal React app (Hello World)
- Docker Compose setup
- Tilt setup for Kubernetes
- PostgreSQL and Redis (for later)

### Key Concepts
- Dockerfile
- docker-compose.yml
- Tiltfile
- Volume mounts (hot-reload)
- Service networking
- Environment variables

## Module 2: OpenAPI Code Generation

### Learning Objectives
- Understand API-first development
- Learn OpenAPI specification
- Set up code generation workflow
- Generate type-safe clients

### Why It Matters
- **Contract-First**: API contract is source of truth
- **Type Safety**: Catch errors at compile time
- **Multi-Client**: Same API for web, mobile
- **Documentation**: Auto-generated, always up-to-date

### What We'll Build
- Minimal OpenAPI spec
- Generate Go server code
- Generate TypeScript client
- Generate Dart client (structure)

## Module 3: Infrastructure as Code (Terraform)

### Learning Objectives
- Understand Infrastructure as Code
- Learn Terraform basics
- Provision AWS resources programmatically
- Manage infrastructure state

### Why It Matters
- **Reproducibility**: Infrastructure is version-controlled
- **Automation**: No manual clicking in AWS console
- **Reviewability**: Infrastructure changes go through PRs
- **Disaster Recovery**: Recreate entire infrastructure from code

### What We'll Build
- Terraform configuration for AWS
- VPC, Security Groups
- ECR (container registry)
- S3 (for Terraform state)

## Module 4: CI Pipeline (GitHub Actions)

### Learning Objectives
- Understand Continuous Integration
- Learn GitHub Actions workflows
- Automate testing and validation
- Build Docker images in CI

### Why It Matters
- **Quality Gates**: Catch bugs before merge
- **Automation**: No manual testing
- **Consistency**: Same checks for every PR
- **Speed**: Fast feedback loop

### What We'll Build
- GitHub Actions workflow
- Lint Go and TypeScript
- Run tests
- Build Docker images
- Validate OpenAPI spec

## Module 5: Container Registry (ECR)

### Learning Objectives
- Understand container registries
- Learn AWS ECR
- Push and pull Docker images
- Tag images for different environments

### Why It Matters
- **Image Storage**: Where containers live
- **Versioning**: Tag images for dev/staging/prod
- **Security**: Private registry
- **Deployment**: Images are deployed, not code

### What We'll Build
- ECR repository setup
- Build and push images
- Image tagging strategy
- Pull images locally

## Module 6: Deploy to Dev (ECS Fargate)

### Learning Objectives
- Understand container orchestration
- Learn AWS ECS
- Deploy containers to cloud
- Understand Fargate

### Why It Matters
- **Orchestration**: Running containers at scale
- **Serverless**: No EC2 management
- **Scalability**: Auto-scale based on demand
- **Production-Ready**: How real apps run

### What We'll Build
- ECS cluster
- ECS task definition
- ECS service
- Application Load Balancer
- Deploy minimal app

## Module 7: CD Pipeline (Automated Deployment)

### Learning Objectives
- Understand Continuous Deployment
- Automate deployment with GitHub Actions
- Deploy to dev automatically
- Understand deployment strategies

### Why It Matters
- **Automation**: Deploy without manual steps
- **Speed**: Fast iteration cycles
- **Reliability**: Consistent deployment process
- **GitOps**: Infrastructure changes via git

### What We'll Build
- GitHub Actions deployment workflow
- Trigger on merge to main
- Build and push images
- Run Terraform
- Update ECS service

## Module 8: Multi-Environment Strategy

### Learning Objectives
- Understand environment separation
- Learn Terraform environments
- Set up dev and staging
- Understand promotion workflow

### Why It Matters
- **Isolation**: Test changes safely
- **Promotion**: Dev → Staging → Prod
- **Risk Reduction**: Catch issues before production
- **Best Practice**: Industry standard

### What We'll Build
- Separate Terraform configs
- Dev environment
- Staging environment
- Deployment workflows

## Module 9: Database and Managed Services

### Learning Objectives
- Understand managed database services
- Learn AWS RDS
- Set up PostgreSQL in cloud
- Connect application to database

### Why It Matters
- **Managed Services**: No database administration
- **High Availability**: Multi-AZ deployments
- **Backups**: Automated backups
- **Scalability**: Easy to scale up/down

### What We'll Build
- RDS PostgreSQL instance
- Security groups
- Database connection
- Migrations setup

## Module 10: Production Deployment

### Learning Objectives
- Understand production requirements
- Learn production best practices
- Set up monitoring and logging
- Implement deployment strategies

### Why It Matters
- **High Availability**: Multi-AZ deployment
- **Monitoring**: Know when things break
- **Scalability**: Handle traffic spikes
- **Reliability**: Zero-downtime deployments

### What We'll Build
- Production Terraform config
- Multi-AZ setup
- CloudWatch monitoring
- Auto-scaling
- Health checks
- Rollback strategy

## Learning Path Summary

**Phase 1: Local Development** (Modules 1-2)
- Docker Compose + Tilt setup
- OpenAPI code generation
- **Goal**: Develop locally with containers/K8s

**Phase 2: Infrastructure** (Modules 3, 9)
- Terraform basics
- AWS resource provisioning
- Managed services (RDS)
- **Goal**: Understand Infrastructure as Code

**Phase 3: CI/CD** (Modules 4, 5, 7)
- GitHub Actions CI
- Container registry (ECR)
- Automated deployment
- **Goal**: Automate testing and deployment

**Phase 4: Cloud Deployment** (Modules 6, 8, 10)
- ECS Fargate deployment
- Multi-environment strategy
- Production deployment
- **Goal**: Deploy to cloud, manage environments

## Key Principles

1. **Start Simple**: Minimal examples, not full features
2. **Build Gradually**: Each module adds one concept
3. **Understand Why**: Cloud-native principles explained
4. **Practice Hands-On**: Try each concept yourself
