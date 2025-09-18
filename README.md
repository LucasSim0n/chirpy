# Chirpy

Minimal REST API written in Go for a Twitter‑like application.

---

## Description

Chirpy is a simple RESTful API built in Go, intended to serve as the backend for a micro‑blogging / “Twitter‑like” application. It aims to be minimal, clean, and easy to extend.

---

## Features

- Clean, minimalistic REST endpoints  
- Written in Go for performance and simplicity  
- SQL schema support (via SQL files)  
- Basic internal organization (handlers, models, etc.)

---

## Tech Stack

- Go  
- SQLC (Generates Go code for sql queries)
- Goose (for database migrations)
- SQL (PostgreSQL or similar)

---

## Project Structure

Here is a high‑level view of the directories / files:

```
/.
├── assets/         Static files (CSS, JS, images, etc.)
├── internal/       Internal Go packages (handlers, models, etc.)
├── sql/            SQL schema / migration files
├── main.go         Entry point for the server
├── go.mod / go.sum Dependences
└── sqlc.yaml       Configuration for sqlc (if used)
```

---

## Getting Started

### Prerequisites

- Go (version ≥ 1.18 or whatever the project requires)  
- Database (e.g. PostgreSQL)  
- `sqlc` tool, for new queries  

### Installation

1. Clone the repo:  
   ```bash
   git clone https://github.com/LucasSim0n/chirpy.git
   cd chirpy
   ```

2. Install dependencies:  
   ```bash
   go mod download
   ```

3. Setup the database:  
   - Create your database  
   - Run migrations / SQL schema in `sql/`  

4. (If applicable) Generate code via sqlc:  
   ```bash
   sqlc generate
   ```

### Running

```bash
go run main.go
```

You should see output indicating the server is running (e.g. listening on a port). Navigate to `http://localhost:PORT` to view the index page or connect via API endpoints.

---

## Configuration

Environment variables or configuration file used for:

- Database connection (host, port, user, password, db name)  
- Server port  
- Any other settings  

Document what variables are expected, sample `.env` if needed.

---

## Database

- Schema files in `sql/`  
- Migrations if any  
- How tables are structured (users, posts, likes, etc.)  

---

## Contributing

If you want to contribute:

- Fork the project  
- Create a feature branch  
- Write tests  
- Open a pull request  

Please follow any coding / style guidelines in the repo.
