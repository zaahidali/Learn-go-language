# Task Tracker API

The Task Tracker API is a **GraphQL-based API** for managing tasks and users. It allows users to create, update, and delete tasks, as well as create, update, and delete users. This project is implemented using **Go**, **gqlgen**, and **Bun** for database interaction.

## Features

- **Create a new user** with a unique identifier
- **Update user details** such as name and email
- **Delete a user** by ID
- **Create a task** associated with a user
- **Update task details** such as title and description
- **Delete a task** by ID
- **Retrieve user information** based on ID
- **Retrieve task information** based on ID
- **Retrieve all users**
- **Retrieve all tasks**


 Task Tracker API
```
├── Create User
│   ├── Input: name, email
│   └── Output: created user information (id, name, email)
├── Update User
│   ├── Input: user ID, updated name, updated email
│   └── Output: updated user information (id, name, email)
├── Delete User
│   ├── Input: user ID
│   └── Output: deleted user ID
├── Create Task
│   ├── Input: user ID, task details (title, description)
│   └── Output: created task information (id, title, description, userId)
├── Update Task
│   ├── Input: task ID, updated details (title, description)
│   └── Output: updated task information (id, title, description, userId)
├── Delete Task
│   ├── Input: task ID
│   └── Output: deleted task ID
├── Get User
│   ├── Input: user ID
│   └── Output: user information (id, name, email)
├── Get Task
│   ├── Input: task ID
│   └── Output: task information (id, title, description, userId)
├── Get Users
│   └── Output: list of all users
└── Get Tasks
    └── Output: list of all tasks
```

## Technologies Used

- **Go**: The programming language used for building the API.
- **gqlgen**: A Go library for generating GraphQL servers based on a schema and resolver functions.
- **Bun**: A lightweight and fast Go database toolkit for interacting with the database.

## Getting Started

To get started with the Task Tracker API, follow these steps:

1. **Clone** the repository:

   ```shell
   git clone https://github.com/zaahidali/Learn-go-language.git
   ```

2. **Navigate** to the project directory:

   ```shell
   cd task-tracker-api
   ```

3. **Install** the required dependencies:

   ```shell
   go mod download
   ```

4. **Configure** the database connection:

   Open the `config.yaml` file and update the database connection details according to your setup.

5. **Run** the API:

   ```shell
   go run server.go
   ```

   The API server will be running on `http://localhost:8080`.

## Schema

The GraphQL schema defines the available queries and mutations for interacting with the ***Task Tracker API***. Here is an overview of the schema:

```graphql
type User {
  id: ID!
  name: String!
  email: String!
}

type Task {
  id: ID!
  userId: String!
  title: String!
  description: String!
  status: TaskStatus!
  dueDate: String!
  createdAt: String!
  updatedAt: String!
}

enum TaskStatus {
  COMPLETED
  INPROGRESS
  PENDING
}

type Query {
  getUser(id: ID!): User!
  getUsers: [User!]!
  getTask(id: ID!): Task!
  getTasks: [Task!]!
}

type Mutation {
  createUser(name: String, email: String) : User!
  createTask(userId: String, title: String, description: String, status: TaskStatus, dueDate: String): Task!
  updateUser(id: ID, name: String, email: String): User!
  deleteUser(id: ID): ID!
  updateTask(id: ID, title: String, description: String, status: TaskStatus, dueDate: String): Task!
}
```

## Installing go-migrate

To install the go-migrate library and command line tool, follow these steps:

1. **Install the migrate library**:

   Run the following command to install the migrate library:

   ```shell
   go get -u github.com/golang-migrate/migrate/v4
   ```

2. **Install the migrate command line tool**:

   Run the following command to install the migrate command line tool:

   ```shell
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

3. **Add the bin directory to your PATH**:

   Make sure the bin directory of your GOPATH is added to your system's PATH. Run the following command to add it:

   ```shell
   echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
   ```

4. **Source the bashrc file**:

   Run the following command to load the new PATH into your current shell session:

   ```shell
   source ~/.bashrc
   ```

After completing these steps, you should have go-migrate installed on your system and the migrate command line tool available for use.

## Makefile Usage

This project uses a Makefile for managing database migrations and running the application. Make sure you have `make` and `PostgreSQL` installed on your system to use these commands.

### Setup

1. First, you need to create your PostgreSQL database. Open a `psql` session from your terminal:

    ```bash
    psql -U postgres
    ```

    Then run the following command at the `psql` prompt to create the database:

    ```sql
    CREATE DATABASE task_tracker_db_with_gql_bun;
    ```

    Press `CTRL+D` to exit the `psql` session.

2. Run `make migrate-up` to generate tables using the migrations in the `migrations` directory.

### Commands

- `make run`: Runs the Go application.
- `make migrate-up`: Applies all available database migrations, updating your database schema to its latest version.
- `make migrate-down`: Undoes all applied database migrations, rolling your database schema back to its initial state.
- `make create-migration <name>`: Creates a new migration file in the `migrations` directory. Replace `<name>` with the desired name for the migration. 
- `make remove-migration <migration>`: Removes a specific migration file. Replace `<migration>` with the name of the migration you want to remove. Use this command with caution, and only for migrations which have not been applied.

Please replace `<name>` and `<migration>` with appropriate values when using `make create-migration` and `make remove-migration`, respectively.

These instructions assume that you're using PostgreSQL with a default `postgres` user. If your PostgreSQL installation uses different settings, you may need to adjust these instructions accordingly.

# Planning to start a new Project with gqlgen? 

To start a new project with gqlgen, follow these steps:

1. **Create a new project directory**:

   ```shell
   mkdir gqlgen-demo
   cd gqlgen-demo
   ```

2. **Initialize the Go module**:

   ```shell
   go mod init github.com/[your_user_name]/gqlgen-demo
   ```

3. **Create the tools.go file**:

   Run the following command to create the `tools.go` file with the necessary imports:

   ```shell
   printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
   ```

4. **Tidy the Go module**:

   Run the following command to tidy the Go module and download the required dependencies:

   ```shell
   go mod tidy
   ```

5. **Initialize gqlgen**:

   Run the following command to initialize gqlgen in your project:

   ```shell
   go run github.com/99designs/gqlgen init
   ```

   This will generate the initial project structure and the `gqlgen.yml` configuration file.

6. **Tidy the Go module again**:

   Run the following command to tidy the Go module after the gqlgen initialization:

   ```shell
   go mod tidy
   ```

7. **Run the server**:

   Start the API server by running the following command:

   ```shell
   go run server.go
   ```

   The API server will be running on `http://localhost:8080`.

8. **Generating code for resolvers**:

   Whenever you make changes to the schema or resolver functions, you need to generate the corresponding code. Run the following command to generate the code:

   ```shell
   go run github.com/99designs/gqlgen generate
   ```

   This command will generate the necessary resolvers and resolver methods based on your schema.

Remember to modify the project structure and files according to your specific requirements.


I apologize for the misunderstanding. Here's an updated version with the code and the description added to it:


## Playground Code

The following Playground code showcases how to interact with the Task Tracker API using GraphQL queries and mutations. You can use this code to test and explore the functionality of the API.

### Create User

```graphql
mutation {
  createUser(name: "John Doe", email: "john@example.com") {
    id
    name
    email
  }
}
```

### Update User

```graphql
mutation {
  updateUser(id: "123", name: "John Doe", email: "john@example.com") {
    id
    name
    email
  }
}
```

### Delete User

```graphql
mutation {
  deleteUser(id: "123")
}
```


### Retrieve User

```graphql
query {
  getUser(id: "123") {
    id
    name
    email
  }
}
```

### Retrieve Users

```graphql
query {
  getUsers {
    id
    name
    email
  }
}
```

### Create Task

```graphql
mutation {
  createTask(userId:"6993", title: "Test Title", description: "This is a test", status: COMPLETED, dueDate: "2023-06-05") {
    id
    title
    description
    dueDate
  }
}
```

### Update Task

```graphql
mutation {
  updateTask(id: "4938", title: "Hello Peter", description: "I am here", status: INPROGRESS, dueDate: "2024-06-05") {
    title
    description
    status
    dueDate
  }
}
```

### Delete Task

```graphql
mutation {
  deleteTask(id: "789")
}
```

### Retrieve Task

```graphql
query {
  getTask(id: "789") {
    id
    title
    description
  }
}
```

### Retrieve Tasks

```graphql
query {
  getTasks {
    id
    title
    description
    dueDate
  }
}
```
## Contributing

Your contributions are most welcome! If you'd like to improve this project, please:

1. **Fork** the repository.
2. **Clone** your fork and **create a new branch**.
3. Make and **commit** your changes.
4. **Push** the changes to your fork.
5. Submit a **pull request**.

After you've submitted the pull request, I will review your changes. Thank you for your contribution!

## License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for more information.
