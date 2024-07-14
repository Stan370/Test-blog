# Test-blog
This repository contains a solution to the test task for a backend developer position. The task is to create a small API for a blog platform with CRUD operations.

# Getting Started

To run the application, follow these steps:

Clone the repository to your local machine.
Install the required dependencies using pip install -r requirements.txt (assuming you're using Python).
Set up the database structure using the provided database schema (e.g., MySQL, PostgreSQL, MongoDB).
Run the application using python app.py (assuming you're using Python).
API Endpoints

## ⌨️ Local Development

You can use GitHub Codespaces for online development:

Or clone it for local development:

```fish
$ git clone https://github.com/Stan370/Test-blog.git
$ cd Test-blog
```

### how to start locally

Ask the owner of this repo to get a copy of local config file (config.json), place it in /config directory.
When you have problem get packets from dependency, please check your GOPROXY.

#### Start a MySQL container in Docker

- `docker pull mysql:5.7`

    - For Apple Silicon chips: `docker pull --platform linux/amd64 mysql:5.7`

- `docker run --name Blog -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=public -e MYSQL_DATABASE=blog mysql:5.7`
```
#### Start 'blog' service

- `./run.sh`         in Linux env
**The API has the following endpoints:
**
GET /posts: Retrieves a list of all blog posts.
POST /posts: Creates a new blog post.
GET /posts/{id}: Retrieves a blog post with given ID.
**Database Structure
**
The database structure is designed to store blog post data with the following columns:

id (primary key): Unique identifier for each blog post.
title: Title of the blog post.
content: Content of the blog post.
created_date: Date and time when the blog post was created.

The API is documented using OpenAPI (formerly known as Swagger). The API documentation includes details about the endpoints, request and response formats, and any required headers or query parameters.
A Makefile to run and build the application.
API tests to ensure the API is working as expected.
An additional endpoint /healthcheck to check the state of the application and related services (currently, the database).
Technology Stack

[Database of choice] (e.g., MySQL, PostgreSQL, MongoDB)
OpenAPI (for API documentation)


Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request.

