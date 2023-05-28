# Simple Blog API

The **Simple Blog API** is a **RESTful API** built with **Go** and **Gin framework**. It allows you to **create, retrieve, update, and delete** blog posts, as well as **manage comments** for each post.

## Features

- **Create** a new blog post with **title, content, and author**
- **Retrieve** a list of **all blog posts**
- **Retrieve** details of a specific blog post **by ID**
- **Update** a specific blog post with **new details**
- **Delete** a specific blog post
- **Add** a new comment to a specific blog post
- **Retrieve** all comments for a specific blog post
- **Delete** a specific comment from a blog post

## API Endpoints

### Blog Posts

- `POST /posts` - Create a new blog post
- `GET /posts` - Retrieve a list of all blog posts
- `GET /posts/:id` - Retrieve details of a specific blog post
- `PUT /posts/:id` - Update a specific blog post
- `DELETE /posts/:id` - Delete a specific blog post

### Comments

- `POST /posts/:id/comments` - Add a new comment to a blog post
- `GET /posts/:id/comments` - Retrieve all comments for a blog post
- `DELETE /posts/:id/comments/:commentId` - Delete a specific comment from a blog post
