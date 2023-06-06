# Budget Management System

This repository contains code for a Budget Management System written in Go. It provides APIs to manage budgets and transactions.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Makefile](#makefile)
- [Development](#development)
- [API Endpoints](#api-endpoints)
- [Support](#support)

## Installation

1. **Clone the Repository**

```bash
git clone https://github.com/zaahidali/Learn-go-language.git
```

2. **Navigate to the directory**

```bash
cd budget_management_system
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

Make sure that a PostgreSQL server is running and accessible through the given connection string. The database named "budget_management_system" should exist. If it doesn't, you can create one:

```bash
psql -U postgres -c 'create database budget_management_system;'
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

## Development

When making changes, you might want to apply migrations to your database:

```bash
make migrate-up
```

If you wish to roll back your database to its previous state, you can run:

```bash
make migrate-down
```

To create a new migration for a table, you can use:

```bash
make create-migration <table_name>
```

If you want to remove a migration, use:

```bash
make remove-migration <migration_name>
```

## API Endpoints

This application exposes the following API endpoints:

### Budget Endpoints:

- `POST /budgets`: Create a new budget with name and amount.
- `GET /budgets`: Retrieve a list of all budgets.
- `GET /budgets/:id`: Retrieve details of a specific budget.
- `PUT /budgets/:id`: Update a specific budget with new details.
- `DELETE /budgets/:id`: Delete a specific budget.

### Transaction Endpoints:

- `POST /transactions`: Record a new transaction with details including amount, type (income or expense), and associated budget id.
- `GET /transactions`: Retrieve a list of all transactions.
- `GET /transactions/:id`: Retrieve details of a specific transaction.
- `PUT /transactions/:id`: Update a specific transaction with new details.
- `DELETE /transactions/:id`: Delete a specific transaction.


## Support

If you encounter any issues while setting up or running this project, please file an issue on this repository.

