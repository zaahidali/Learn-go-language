# Book Management API with **gqlgen**

This repository contains the source code for the **Book Management API** implemented using **gqlgen**, a Go library for building **GraphQL** servers. The API allows users to perform **CRUD** (Create, Read, Update, Delete) operations on a collection of **books** stored in a **PostgreSQL** database.

This codebase is associated with the Medium blog post titled "GraphQL Go: Powerful API Development with gqlgen & Go Bun". You can find the blog post [here](https://zaahidali.medium.com/graphql-go-powerful-api-development-with-gqlgen-go-bun-57131c35c089). The blog post provides a detailed explanation of the code and guides you through the process of building the Book Management API using gqlgen and PostgreSQL.

## Features

- Create a new **book** with details such as **title**, **author**, **publication year**, and **genre**.
- Update the details of an existing **book**.
- Delete a **book** based on its **ID**.
- Retrieve a single **book** by its **ID**.
- Retrieve a list of all **books**.

## Getting Started

To set up the project, follow these steps:

1. **Clone** the repository: `git clone https://github.com/zaahidali/Learn-go-language.git`
2. Navigate to the project directory: `cd Learn-go-language/book_management_api_with_gqlgengen`
3. **Install** the dependencies: `go mod download`
4. **Set up the database**: Make sure you have **PostgreSQL** installed and create a new database named `book_management_api`. You can create the database by running the following command in your terminal:
```
psql -U postgres -h localhost -p 5432 -c 'CREATE DATABASE book_management_api;'
```
5. **Configure the environment variables**: Create a `.env` file in the project directory and provide the necessary database connection details. You can refer to the `.env.example` file for the required variables.
6. **Run the migration**: Execute the migration command to set up the necessary database tables: `make migrate-up`
7. **Start the server**: Launch the API server by running the following command: `go run .`

Once the server is running, you can access the **GraphQL playground** at `http://localhost:8080/` to interact with the API.

## API Documentation

The Book Management API supports the following **GraphQL queries** and **mutations**:

- `createBook`: Creates a new **book**.
- `updateBook`: Updates the details of an existing **book**.
- `deleteBook`: Deletes a **book** based on its **ID**.
- `book`: Retrieves a single **book** by its **ID**.
- `books`: Retrieves a list of all **books**.

For detailed information on the API schema and available operations, refer to the **GraphQL playground documentation**.

## Examples

Here are some example **queries** and **mutations** that you can perform:

1. Create a **book**:
   ```
   mutation {
     createBook(title: "Example Book", author: "John Doe", publicationYear: 2023, genre: "Fiction") {
       id
       title
       author
       publicationYear
       genre
     }
   }
   ```

2. Update a **book**:
   ```
   mutation {
     updateBook(id: "123", title: "Updated Book", author: "Jane Smith") {
       id
       title
       author
       publicationYear
       genre
     }
   }
   ```

3. Delete a **book**:
   ```
   mutation {
     deleteBook(id: "123")
   }
   ```

4. Retrieve a **book**:
   ```
   query {
     book(id: "123") {
       id
       title
       author
       publicationYear
       genre
     }
   }
   ```

5. Retrieve all **books**:
   ```
   query {
     books {
       id
       title
       author
       publicationYear
       genre
     }
   }
   ```

## Contributing

Contributions to the Book Management API project are **welcome**! If you find any issues or have suggestions for improvements, please create a new issue or submit a pull request.

## License

This project is licensed under the **MIT License**.

Feel free to use this code as a reference or modify it for your own projects.

## Acknowledgements

This project was built using the following technologies and libraries:

- **gqlgen**: A Go library for building GraphQL servers
- **PostgreSQL**: An open-source relational database management system
- **Go**: An open-source programming language
- **GitHub**: A web-based hosting service for version control repositories

