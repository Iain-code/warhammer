# AverageDice API (Warhammer Backend)

> RESTful API service powering [AverageDice.com](https://averagedice.com) — a Warhammer roster and model management platform.  
> Written in Go, deployed as a containerized AWS Lambda function, backed by PostgreSQL (AWS RDS).

---

## Features

- **RESTful API** built with [chi](https://github.com/go-chi/chi) router  
- **Serverless runtime** — AWS Lambda container image + Amazon ECR  
- **Persistent storage** — PostgreSQL on AWS RDS  
- **Authentication** — JWT-based auth with refresh tokens  
- **CORS support** for frontend integration (AverageDice.com)  
- **Structured, minimal logging** (to CloudWatch in Lambda)  

---

## Tech Stack

- **Language**: Go 1.23  
- **Framework**: chi router + middleware  
- **Serverless**: AWS Lambda (via [aws-lambda-go](https://github.com/aws/aws-lambda-go) + [httpadapter](https://github.com/awslabs/aws-lambda-go-api-proxy))  
- **Database**: PostgreSQL (`lib/pq`)  
- **Container**: Docker, pushed to AWS ECR  

---

## Deployment

### Prerequisites
- AWS CLI configured with Lambda + ECR permissions  
- Docker installed  
- PostgreSQL instance (e.g. AWS RDS)  

### Environment Variables (`.env`)
```env
dbUrl=postgres://user:password@host:5432/dbname?sslmode=disable
TOKEN_SECRET=your-secret-key
CORS_ALLOW_ORIGINS=https://averagedice.com
