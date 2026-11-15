# School Management REST API

A school management backend application developed using Go. The project provides a RESTful API structure for handling school-related data and demonstrates backend application organization, HTTP API development and business logic implementation.

## Overview

The School Management REST API is designed to represent common data-management requirements within an educational institution.

The application is organized into separate packages for application entry points, internal functionality and reusable utilities.

The project provides an opportunity to explore how RESTful APIs can be used to manage relationships between different entities within a school environment.

## Technologies

| Technology | Purpose |
|---|---|
| Go | Backend programming language |
| REST | API communication architecture |
| JSON | Data representation |
| Git | Version control |

Add the database and framework names here once confirmed against the project's `go.mod` and configuration files.

## Key Components

The project is organized around school-related entities, including:

- Students.
- Teachers.
- Executives and administrative personnel.

The codebase also explores teacher-related responsibilities and relationships between teachers and students.

## Project Structure

```text
school_management_RESTapi/
├── cmd/
│   └── api/
├── internal/
├── pkg/
│   └── utils/
├── [project data files]
├── go.mod
└── go.sum
```

### Directory Responsibilities

- `cmd/api`: Application entry point.
- `internal`: Internal application logic and implementation.
- `pkg/utils`: Reusable utility functionality.
- `go.mod`: Go module definition and dependency management.
- `go.sum`: Dependency checksum information.

## Getting Started

### Prerequisites

- Go installed on your machine.
- Git.
- Any additional services required by the project's configuration.

### Clone the Repository

```bash
git clone https://github.com/revernelly/school_management_RESTapi.git
cd school_management_RESTapi
```

### Download Dependencies

```bash
go mod download
```

### Run the Application

If the application entry point is `cmd/api`, run:

```bash
go run ./cmd/api
```

### Run Tests

```bash
go test ./...
```

## API Usage

The application exposes functionality through HTTP endpoints.

Use an API client such as Postman to interact with the available routes.

The exact endpoints, request formats and response examples should be documented here based on the application's implemented route definitions.

## Project Objectives

This project demonstrates practical experience with:

- Go backend application development.
- RESTful API design.
- Organizing Go applications into packages.
- Modeling school-related entities.
- Implementing business logic involving students and teachers.
- Managing reusable utility functionality.
- Using Git for source code management.

## Author

**Carl Anyanwu**

GitHub: https://github.com/revernelly
