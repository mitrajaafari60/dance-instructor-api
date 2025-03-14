## Copyright

© 2025 [Mitra Jafari].

# Dance Instructor API

A RESTful API built with Go and Gin to manage a directory of dance instructors for a dance studio. This project uses JWT (JSON Web Token) authentication with RSA signing for security and an in-memory store for data persistence.


## Project Structure

dance-instructor-api/
├── cmd/                
│   └── server.go
├── docs/               
│   └── docs.go
│   └── swagger.json
│   └── swagger.yaml
├── internal/           
│   ├── auth/          
│   │   └── jwt.go
│   ├── config/       
│   │   └── config.go
│   ├── handlers/      
│   │   └── instructor.go
│   ├── middleware/    
│   │   └── jwt.go
│   ├── models/        
│   │   └── instructor.go
│   └── storage/      
│       └── memory.go
├── tests/             
│   ├── instructor_test.go
│   ├── jwt_test.go
│   └── ...          
├── private.pem        
├── public.pem         
├── go.mod             
├── main.go            
└── README.md         


## Features

- **CRUD operations** for dance instructors (Create, Read, Update, Delete)
- **Secure JWT authentication** with RSA public/private key signing
- **In-memory storage** for simplicity (scalable to database in future)
- **Unit and integration tests** with `testify`
- **Configurable port** via environment variables


## Prerequisites

- Go 1.21 or higher
- Git
- OpenSSL (for generating RSA keys)
- A terminal or command-line interface

---
## Installation
1. **Clone the Repository:**

```
git clone https://github.com/mitrajaafari60/dance-instructor-api
cd dance-instructor-api
```

## Install Dependencies:

```   
go mod download
```
##  Generate RSA Key Pair: 
Generate private and public keys for JWT authentication:

```
openssl genrsa -out private.pem 2048
openssl rsa -in private.pem -pubout -out public.pem
```

## Running the API

```
go run main.go
```    
    
The API will start on http://localhost:8081 by default.
Run in Release Mode (Production): To disable debug logging:

```
export GIN_MODE=release
go run main.go
```

Custom Port (Optional): Set a different port via environment variable:

```
export PORT=:8080
go run main.go
```

## API Endpoints

All endpoints are under the /api/v1 prefix and require JWT authentication.

Method	Endpoint	        Description	                    Request Body (if applicable)
GET	    /instructors	    List all instructors	        -
GET	    /instructors/:id	Get instructor by ID	        -
POST	/instructors	    Create a new instructor         {"name": "", "bio": "", "specialty": "","availability": ""}
PUT	    /instructors/:id	Update an existing instructor   {"name": "", "bio": "", "specialty": "", "availability": ""}
DELETE	/instructors/:id	Delete an instructor by ID	    -

## API Documentation

You can view the Swagger documentation below:
    http://localhost:8081/swagger/index.html

# Example Requests

Create Instructor:

    curl -X POST http://localhost:8081/api/v1/instructors \
        -H "Authorization: Bearer <jwt-token>" \
        -H "Content-Type: application/json" \
        -d '{"name": "Jane Doe", "bio": "Expert dancer", "specialty": "Ballet", "availability": "Mon-Fri"}'

Get All Instructors:
    curl http://localhost:8081/api/v1/instructors \
     -H "Authorization: Bearer <jwt-token>"


## Authentication

The API uses JWT with RSA signing for authentication. Tokens must be included in the Authorization header as Bearer <token>.

Generating a JWT Token
    
```
export PRINT_TEST_TOKEN=true
go run main.go
```

## Testing

Run the unit and integration tests:

```    
cd tests
go test -v
```
