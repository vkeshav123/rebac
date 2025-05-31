# Relationship-Based Access Control (ReBAC) API

This project implements a Relationship-Based Access Control (ReBAC) system with a PostgreSQL backend and a RESTful API using Go, Gin, and GORM.

## Project Structure
```
rebac/
├── db/           # Database connection and migration
├── handlers/     # HTTP handler functions for CRUD APIs
├── models/       # GORM models matching the PostgreSQL schema
├── main.go       # Entry point, Gin router setup
├── go.mod        # Go module and dependencies
└── README.md     # This file
```

## Prerequisites
- Go 1.19+
- PostgreSQL

## Database Setup
1. Create a PostgreSQL database (default is `rebac`).
2. Ensure your tables match the schema (auto-migration will create them if missing).
3. Set the `DATABASE_URL` environment variable, e.g.:
   ```sh
   export DATABASE_URL="host=localhost user=postgres password=postgres dbname=rebac port=5432 sslmode=disable"
   ```
   Or edit the fallback DSN in `db/db.go`.

## Install Dependencies
```
go mod tidy
```

## Running the Server
```
go run main.go
```
The server runs on [http://localhost:8080](http://localhost:8080) by default.

## API Endpoints

### Models CRUD
- `POST   /models`      - Create a new model
- `GET    /models`      - List all models
- `GET    /models/:id`  - Get a model by ID
- `PUT    /models/:id`  - Update a model by ID
- `DELETE /models/:id`  - Delete a model by ID

> More endpoints for Types, Relations, DirectRelations, ImpliedRelations, and Policies can be added following the same pattern.

## Extending the API
- Add new CRUD handlers in `handlers/handlers.go`.
- Register routes in `main.go`.
- Update models in `models/models.go` as needed.

## Notes
- On startup, all models are auto-migrated to the database.
- Use tools like Postman or curl to test the endpoints.
- For production, configure environment variables and review security settings.

---

Feel free to extend the API for other entities as needed!