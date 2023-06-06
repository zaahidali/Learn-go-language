# User Post Comment API

This repository contains code for a **RESTful API built in Go**, designed for **user, post, and comment management**. It provides APIs to **create, retrieve, update, and delete** users, posts, and comments.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Makefile](#makefile)
- [API Endpoints](#api-endpoints)
  - [Users](#users)
  - [Posts](#posts)
  - [Comments](#comments)
- [Development](#development)
- [Support](#support)

## Installation

1. **Clone the Repository**

```bash
git clone https://github.com/zaahidali/Learn-go-language.git
```

2. **Navigate to the directory**

```bash
cd blog_post_orm_system
```

3. **Install dependencies**

Before you can run this project, you must have the following software installed on your system:

- Go (version 1.15 or later)
- PostgreSQL

Ensure you've installed the required dependencies for this project:

```bash
go mod tidy
```

4. **Set up database**

Make sure that a PostgreSQL server is running and accessible through the given connection string. The database named "blog_posts_orm" should exist. If it doesn't, you can create one:

```bash
psql -U postgres -c 'create database blog_posts_orm;'
```

5. **Install go-migrate**

```bash
# Install the migrate library
go get -u github.com/golang-migrate/migrate/v4

# Install the migrate command line tool
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Make sure the bin directory of your GOPATH is in your system's PATH
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc

# Source the bashrc file to load the new PATH into your current shell session
source ~/.bashrc
```

6. **Run migrations**


With the `migrate` tool installed, you can now run migrations to set up your database schema.

```bash
make migrate-up
```

## Usage

To run the server, use the following command:

```bash
make run
```

## Makefile

The `Makefile` includes several commands:

- `make run`: runs the application
- `make migrate-up`: applies all up migrations
- `make migrate-down`: rolls back all migrations
- `make create-migration <name>`: creates a new migration with the specified name
- `make remove-migration <migration>`: removes a specified migration file.


## API Endpoints

### Users

- `POST /users`: Create a new user. Accept a JSON body with the user's name and email.
- `GET /users`: Get a list of all users.
- `GET /users/:id`: Get the details of a specific user.
- `PUT /users/:id`: Update a specific user. Accept a JSON body with the new details of the user.
- `DELETE /users/:id`: Delete a specific user.

### Posts

- `POST /posts`: Create a new post. Accept a JSON body with the post's title, content, and user_id.
- `GET /posts`: Get a list of all posts.
- `GET /posts/:id`: Get the details of a specific post.
- `PUT /posts/:id`: Update a specific post. Accept a JSON body with the new details of the post.
- `DELETE /posts/:id`: Delete a specific post.

### Comments

- `POST /comments`: Create a new comment. Accept a JSON body with the comment's content, user_id, and post_id.
- `GET /comments`: Get a list of all comments.
- `GET /comments/:id`: Get the details of a specific comment.
- `PUT /comments/:id`: Update a specific comment. Accept a JSON body with the new details of the comment.
- `DELETE /comments/:id`: Delete a specific comment.

## Development

When developing, you might want to use some helpful commands. Please refer to the Makefile section for further information.

## Support

If you encounter any issues while setting up or running this project, please file an issue on this repository.
