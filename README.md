# Go CRUD API

A RESTful CRUD API built with Go, Gin framework, and PostgreSQL. This project demonstrates basic CRUD operations for a blog post management system.

## Features

- Create, Read, Update, and Delete blog posts
- RESTful API endpoints
- PostgreSQL database integration
- Environment variable configuration
- Clean project structure

## Prerequisites

- Go 1.16 or higher
- PostgreSQL
- Git

## Modules and Frameworks

### Core Framework
- [Gin](https://github.com/gin-gonic/gin) - Web framework for Go
- [GORM](https://gorm.io/) - ORM library for Go

### Key Dependencies
- [godotenv](https://github.com/joho/godotenv) - Environment variable management
- [lib/pq](https://github.com/lib/pq) - PostgreSQL driver for Go
- [gin-gonic/gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [gorm.io/gorm](https://gorm.io/) - ORM library
- [gorm.io/driver/postgres](https://gorm.io/docs/connecting_to_the_database.html) - PostgreSQL driver for GORM

### Development Tools
- [go.mod](https://go.dev/doc/modules/gomod-ref) - Go module management
- [go.sum](https://go.dev/ref/mod#go-sum-files) - Go module checksum file

## Project Structure

```
go-crud/
├── controllers/     # Request handlers
├── initializers/    # Database and environment setup
├── migrate/         # Database migrations
├── models/          # Data models
├── main.go         # Application entry point
├── go.mod          # Go module file
└── go.sum          # Go module checksum
```

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/shamsher4499/go-crud.git
   cd go-crud
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables:
   Create a `.env` file in the root directory with the following variables:
   ```
   DB_HOST=localhost
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database_name
   DB_PORT=5432
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST   | /posts   | Create a new post |
| GET    | /posts   | Get all posts |
| GET    | /posts/:id | Get a specific post |
| PUT    | /posts/:id | Update a post |
| DELETE | /posts/:id | Delete a post |

## Example Usage

### Create a Post
```bash
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"title":"My First Post","content":"This is the content of my first post"}'
```

### Get All Posts
```bash
curl http://localhost:8080/posts
```

### Get a Specific Post
```bash
curl http://localhost:8080/posts/1
```

### Update a Post
```bash
curl -X PUT http://localhost:8080/posts/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Title","content":"Updated content"}'
```

### Delete a Post
```bash
curl -X DELETE http://localhost:8080/posts/1
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

- Shamsher Singh
- GitHub: [@shamsher4499](https://github.com/shamsher4499)
