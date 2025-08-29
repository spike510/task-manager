# Task Manager API

Task management API built with Go, Gin, PostgreSQL, JWT, Swagger, Docker, and CI/CD.

## 🚀 Features
- User registration and login with JWT authentication  
- Task CRUD (create, read, update, delete)  
- Swagger API documentation  
- Unit and integration tests  
- Docker + Docker Compose setup  
- CI/CD pipeline: lint, tests, build, Docker image push  
- Optional Kubernetes deployment  

## 📦 Installation / Local Setup
1. Clone the repository:
   ```bash
   git clone <repo-url>
   cd task-manager
   ```
2. Create a `.env` file (see `.env.example` for reference).  
3. Start the app with Docker Compose:
   ```bash
   docker-compose up --build
   ```
4. Access the API:
   - API: [http://localhost:8085](http://localhost:8085)  
   - Swagger docs: [http://localhost:8085/swagger/index.html](http://localhost:8085/swagger/index.html)  

## 🧪 Tests
Run all tests:
```bash
go test ./... -v
```

## 📈 CI/CD
GitHub Actions runs linting, tests, builds the Docker image, and pushes it to the container registry.  

## 📝 Tech Stack
- **Backend:** Go, Gin  
- **Database:** PostgreSQL  
- **Auth:** JWT  
- **Docs:** Swagger  
- **DevOps:** Docker, Docker Compose, Kubernetes (optional), GitHub Actions  

## 💻 API Endpoints
- `POST /users/register` – User registration  
- `POST /users/login` – User login  
- `GET /tasks` – List tasks  
- `POST /tasks` – Create a task  
- `PUT /tasks/:id` – Update a task  
- `DELETE /tasks/:id` – Delete a task  