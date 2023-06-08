# Education Management System API

The **Education Management System API** provides endpoints to manage **students**, **courses**, **enrollments**, and **teachers**. It allows you to perform CRUD operations for each entity and retrieve related information.

## Table of Contents
- [Education Management System API](#education-management-system-api)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Makefile](#makefile)
  - [Endpoints](#endpoints)
    - [Students](#students)
    - [Courses](#courses)
    - [Enrollments](#enrollments)
    - [Relationships](#relationships)
    - [Teachers](#teachers)
  - [Usage](#usage)
  - [Contributions](#contributions)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/zaahidali/Learn-go-language.git
cd education-management-system
```

2. Install the dependencies:
```bash
go mod download
```

3. Create the PostgreSQL database named **education_management_system**:
```bash
sudo -u postgres createdb education_management_system
```

4. Apply the database migrations:
```bash
make migrate-up
```

## Makefile

The project includes a Makefile that provides convenient commands for running the server and applying database migrations.

- `make run`: Run the server.
- `make migrate-up`: Apply the database migrations.
- `make migrate-down`: Roll back the applied migrations.
- `make create-migration <name>`: Create a new migration file.
- `make remove-migration <migration>`: Remove a migration file.

Make sure to update the `DATABASE_URL` variable in the Makefile with your PostgreSQL connection details.


## Endpoints

### Students
- `GET /students`: List all students
- `GET /students/{id}`: Get a single student by ID
- `POST /students`: Create a new student
- `PUT /students/{id}`: Update a student by ID
- `DELETE /students/{id}`: Delete a student by ID

### Courses
- `GET /courses`: List all courses
- `GET /courses/{id}`: Get a single course by ID
- `POST /courses`: Create a new course
- `PUT /courses/{id}`: Update a course by ID
- `DELETE /courses/{id}`: Delete a course by ID

### Enrollments
- `GET /enrollments`: List all enrollments
- `GET /enrollments/{id}`: Get a single enrollment by ID
- `POST /enrollments`: Create a new enrollment
- `PUT /enrollments/{id}`: Update an enrollment by ID
- `DELETE /enrollments/{id}`: Delete an enrollment by ID

### Relationships
- `GET /students/{id}/courses`: Get all courses for a student
- `GET /courses/{id}/students`: Get all students for a course
- `GET /courses/:id/teacher`: Fetch the teacher of a specific course

### Teachers
- `GET /teachers`: Fetch all teachers
- `GET /teachers/{id}`: Fetch a specific teacher by ID
- `POST /teachers`: Create a new teacher
- `PUT /teachers/{id}`: Update a teacher by ID
- `DELETE /teachers/{id}`: Delete a teacher by ID

## Usage

To run the server, use the following command:
```bash
make run
```

To get started with the Education Management System API, follow these steps:
1. Install the required dependencies (Go, Gin, Bun, PostgreSQL, etc.).
2. Set up your database and configure the connection details in the application.
3. Run the API server using the appropriate command.
4. Use an API testing tool like Postman to interact with the endpoints.

Make sure to replace `{id}` with the actual ID value in the endpoint URLs.

## Contributions

Contributions to the Education Management System API are welcome! If you find any issues or have suggestions for improvement, feel free to open an issue or submit a pull request.
