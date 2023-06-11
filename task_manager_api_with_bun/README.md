# Go Bun Task Manager API

This project showcases the implementation of a **Task Manager REST API** using Go Bun. It provides a complete CRUD (Create, Read, Update, Delete) functionality for managing tasks, with data persistence in a PostgreSQL database.

## Features

- Create, read, update, and delete tasks
- Store tasks in a PostgreSQL database using Go Bun ORM
- API endpoints for interacting with tasks
- Error handling and validation of input data
- Integration with Gin web framework for HTTP routing

## Requirements

- Go 1.16 or higher
- PostgreSQL database

## Getting Started

1. Clone the repository:

   ```shell
   git clone https://github.com/zaahidali/Learn-go-language
   ```

2. Install the dependencies:

   ```shell
   go mod tidy
   ```

3. Configure the PostgreSQL database connection in the `main.go` file:

   ```go
   // Replace the connection string with your own PostgreSQL database credentials
   dsn := "postgres://your-username:your-password@localhost:5432/your-database?sslmode=disable"
   ```

4. Run the application:

   ```shell
   go run main.go
   ```

5. Access the API at `http://localhost:8080` and start managing your tasks!

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please create a new issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
