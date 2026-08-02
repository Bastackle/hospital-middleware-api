# 🏥 Hospital Middleware System API

> ⚠️ **DISCLAIMER**
>
> **TH:** ระบบนี้ถูกพัฒนาขึ้นเพื่อวัตถุประสงค์ในการทดสอบและประเมินผลทักษะการพัฒนาระบบ (Technical Assessment) เท่านั้น ข้อมูลทั้งหมด เช่น ชื่อผู้ป่วย เลขบัตรประชาชน และข้อมูลทางการแพทย์ เป็น **Mock Data** ทั้งหมด และไม่สามารถนำไปใช้ในงานจริงได้
>
> **EN:** This project is developed **for technical assessment purposes only**. All patient records and personal information are **mock data** created for demonstration and testing.

---

## 📖 Table of Contents

- [Project Overview](#-project-overview)
- [Key Features](#-key-features)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Getting Started](#-getting-started)
- [Running Unit Tests](#-running-unit-tests)
- [Additional Documentation](#-additional-documentation)

---

# 📌 Project Overview

Hospital Middleware API is developed using **Go (Golang)** and the **Gin Framework**. The system acts as a middleware service for searching and retrieving patient information while enforcing **Hospital-based Multi-tenancy Isolation**, ensuring that each hospital can only access its own data.

The project also simulates integration with an external Hospital Information System (**Hospital A External HIS**) and exposes a RESTful API documented with Swagger.

---

# ✨ Key Features

### 🔒 Hospital Isolation (Multi-tenancy)

- Enforces hospital-level data isolation using **JWT Authentication** and authorization middleware.
- Each staff member can only access patient records belonging to their own hospital.

### 🔍 Dynamic Patient Search

- Supports flexible searching using optional criteria.
- Search by:
  - `national_id`
  - `passport_id`
  - `first_name`
  - `last_name`
  - `phone_number`
- Additional filters can be easily extended.

### 🔗 Hospital A External HIS Integration (Mock)

- Simulates integration with an external Hospital Information System.
- Provides a mock endpoint compatible with Hospital A's API format:
  - `GET /patient/search/{id}`

### 🧪 Comprehensive Unit Testing

- Uses **SQLite In-Memory** for fast, isolated, and repeatable testing.
- Covers:
  - Positive Scenarios
  - Negative Scenarios
  - Hospital Data Isolation
  - Authentication & Authorization

---

# 🛠 Tech Stack

| Category | Technology |
|----------|------------|
| Language | Go 1.20+ |
| Framework | Gin |
| ORM | GORM |
| Database | PostgreSQL (Docker) / SQLite (In-Memory Testing) |
| Authentication | JWT + Bcrypt |
| Reverse Proxy | Nginx |
| Containerization | Docker & Docker Compose |
| API Documentation | Swagger (OpenAPI 3.0) |

---

# 📁 Project Structure

```text
.
├── cmd/
│   └── main.go                  # Application entry point
├── docs/                        # Generated Swagger documentation
├── internal/
│   ├── delivery/
│   │   └── http/                # HTTP Handlers, Middleware & Hospital A Mock API
│   ├── model/                   # Database Models & DTOs
│   └── service/                 # Business Logic
│       ├── patient_service.go
│       ├── patient_service_test.go
│       ├── staff_service.go
│       └── staff_service_test.go
├── pkg/
│   ├── auth/                    # JWT & Password Hashing
│   └── database/                # Database Connection, Migration & Seed
├── docker-compose.yml
├── Dockerfile
├── nginx.conf
├── go.mod
├── go.sum
└── README.md
```

---

# 🚀 Getting Started

## 1. Clone the Repository

```bash
git clone https://github.com/Bastackle/hospital-middleware-api.git
cd hospital-middleware-system
```

---

## 2. Run with Docker Compose (Recommended)

Start the entire stack (**Nginx + Go API + PostgreSQL**) with a single command.

```bash
docker-compose up --build -d
```

### Available Endpoints

| Service | URL |
|---------|-----|
| Swagger UI | http://localhost/swagger/index.html |

Stop the application:

```bash
docker-compose down
```

---

## 3. Run Locally (Development Mode)

Install project dependencies:

```bash
go mod tidy
```

Run the application:

```bash
go run cmd/main.go
```

> **Note**
>
> If PostgreSQL environment variables are not configured, the application automatically falls back to an **SQLite In-Memory** database.

### Available Endpoints

| Service | URL |
|---------|-----|
| Swagger UI | http://localhost:8080/swagger/index.html |

---

# 🧪 Running Unit Tests

Run all unit tests:

```bash
go clean -testcache && go test -v ./...
```

The test suite includes:

- ✅ Positive Scenarios
- ✅ Negative Scenarios
- ✅ Hospital Data Isolation (Multi-tenancy)
- ✅ Authentication & Authorization

---

# 📚 Additional Documentation

A detailed Google Document is included with the project submission, covering:

- ER Diagram
- Database Design
- API Specification
- API Usage Examples
- System Design Explanation
- Assumptions & Design Decisions
