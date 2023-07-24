### Golang RESTful API Using Go Community Standard Project Layout

## System Requirement:
1. Golang version 1.16+
2. Docker Compose
3. Mockgen (to generate gomock, in ubuntu install using `sudo apt-get install gomock`)
4. Cloudinary (Free Cloudstorage)

## Structure Directory
Source of reference [Golang](https://github.com/golang-standards/project-layout)
```
├── cmd                 # Application directory (main.go)
├── internal            # Private application and library code
│   ├── app
|       └── controller  # Handle incoming request
|       └── mocks       # Gomock generated files
|       └── model       # Database properties
|       └── repository  # Database interaction
|       └── schema      # API Request And Response Body
|       └── service     # API usecase
│   └── pkg
|       └── config      # Config struct
|       └── db          # Database connection
|       └── handler     # Handler incoming request body and response API
|       └── middleware  # Golang middleware
|       └── reason      # Error message
|       └── validator   # API Request Validator

```

### Syntax
1. `make environment` - To run database (PostgreSQL) inside docker container
2. `make server` - To run the application
3. `make test` - To run test case
4. `make test-cover` - To display coverage.out test coverage in browser 

### Step to test API
1. Import postman collection
2. The application will using `localhost:8123/`
3. Create New Movies API `localhost:8123/movies`
   ![This is an alt text.](/post_body.png)

4. Get Movies Lists  `localhost:8123/movies?limit=10&offset=0&sort_by=id&asc=1&title=`
5. Get Movie By ID `localhost:8123/movies/2`
6. Update Movie By ID `localhost:8123/movies/2`
7. Delete Movie By ID `localhost:8123/movies/2`