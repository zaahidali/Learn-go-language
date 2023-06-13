# My Library API

This is a simple API for managing books in a library. It provides various routes to perform CRUD operations on books.

## Prerequisites

- Go version 1.16 or higher installed on your machine.
- The Gin web framework package (`github.com/gin-gonic/gin`) installed. You can install it by running:
  ```
  go get -u github.com/gin-gonic/gin
  ```

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/zaahidali/Learn-go-language
   ```

2. Change to the project directory:
   ```
   cd library_management_api
   ```

3. Start the API server:
   ```
   go run library_management_api.go
   ```

## Getting Started

Once you have the API up and running locally, you can make HTTP requests to interact with it. The server will start at `http://localhost:8000`.


## Routes

- **`GET /`** - Welcome page that greets users with a message.

- **`GET /books`** - Retrieves a list of all books in the library.

- **`GET /books/:id`** - Retrieves details of a specific book based on the provided `id` parameter.

- **`PUT /books/:id`** - Updates an existing book with new information based on the provided `id` parameter.

- **`DELETE /books/:id`** - Deletes a specific book from the library based on the provided `id` parameter.

- **`POST /books`** - Adds a new book to the library.

## Usage

To use the API, you can make HTTP requests to the corresponding routes using a tool like cURL or Postman. Here are some examples:

- **Get all books:**
  ```
  GET /books
  ```

- **Get details of a specific book with ID "1":**
  ```
  GET /books/1
  ```

- **Update the information of a book with ID "2":**
  ```
  PUT /books/2
  Request Body: {"id": "2", "title": "Updated Title", "author": "Updated Author"}
  ```

- **Delete a book with ID "1":**
  ```
  DELETE /books/1
  ```

- **Add a new book to the library:**
  ```
  POST /books
  Request Body: {"id": "3", "title": "New Book", "author": "New Author"}
  ```

## Book Structure

Each book in the library has the following structure:

```json
{
  "id": "string",
  "title": "string",
  "author": "string"
}
```

## Example Data

Here are some example books that are initially present in the library:

```json
[
  {
    "id": "1",
    "title": "12 Rules of Life",
    "author": "Jordan Peterson"
  },
  {
    "id": "2",
    "title": "Harry Potter",
    "author": "J.K. Rowling"
  }
]
```

## Contributing

**Contributions are always welcome!** If you'd like to contribute to the My Library API project, follow these steps:

1. Fork the repository and clone it to your local machine:

```bash
git clone https://github.com/zaahidali/Learn-go-language
```

2. Create a new branch for your feature or bug fix:

```bash
git checkout -b feature/my-feature
```

3. Make the necessary changes and ensure that the code is properly formatted.

4. Commit your changes and push them to your forked repository:

```bash
git add .
git commit -m "Add my feature"
git push origin feature/my-feature
```

5. Submit a pull request with a clear description of your changes and the problem they solve.

We appreciate your contributions and will review your pull request as soon as possible. Thank you for helping to improve the ***My Library API.***
