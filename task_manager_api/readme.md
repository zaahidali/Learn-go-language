# Task Manager API

The ***Task Manager API*** is a simple yet powerful API built using ***Go*** and ***Gin***. It provides endpoints to manage tasks with in-memory storage. This project is designed to help you learn Go language and understand the fundamentals of building APIs.

## Features

- Get a list of all tasks.
- Get the details of a specific task.
- Update a specific task.
- Delete a specific task.
- Create a new task.

## Installation

To run the Task Manager API locally, follow these steps:

1. Clone the repository: `git clone https://github.com/zaahidali/Learn-go-language.git`
2. Navigate to the `task_manager_api` directory.
3. Install the required dependencies: `go get github.com/gin-gonic/gin`.
4. Run the API server: `go run task_manager_api.go`.
5. The API server will start running at `http://localhost:8080`.

## API Endpoints

- `GET /tasks`
  - Get a list of all tasks.

- `GET /tasks/:id`
  - Get the details of a specific task.

- `PUT /tasks/:id`
  - Update a specific task.
  - Request Body: JSON object with the updated task details.

- `DELETE /tasks/:id`
  - Delete a specific task.

- `POST /tasks`
  - Create a new task.
  - Request Body: JSON object with the task's title, description, due date, and status.

## Usage

1. Make sure you have Go installed on your system.
2. Clone the repository and navigate to the `task_manager_api` directory.
3. Start the API server by running `go run task_manager_api.go`.
4. Use a tool like Postman or any other API testing tool to send requests to the API endpoints.

### Examples

#### GET /tasks
- Send a GET request to `http://localhost:8080/tasks` to retrieve all tasks.

#### GET /tasks/:id
- Send a GET request to `http://localhost:8080/tasks/:id` to retrieve details of a specific task.

#### PUT /tasks/:id
- Send a PUT request to `http://localhost:8080/tasks/:id` with a JSON body containing the updated task details.

#### DELETE /tasks/:id
- Send a DELETE request to `http://localhost:8080/tasks/:id` to delete a specific task.

#### POST /tasks
- Send a POST request to `http://localhost:8080/tasks` with a JSON body containing the task's details to create a new task.

## Contributions

Contributions to the Task Manager API are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request on the [GitHub repository](https://github.com/zaahidali/Learn-go-language/tree/master/task_manager_api).

## License

This project is licensed under the [MIT License](https://github.com/zaahidali/Learn-go-language/blob/master/LICENSE).

Feel free to explore and learn with the Task Manager API. Happy coding!
