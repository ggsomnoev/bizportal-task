# BizPortal Task

A simple [Go](https://golang.org/) solution to the provided [Rope Bridge](https://docs.google.com/document/d/1aK8g7BQ6IahHblh1QBZWEmXX_PjUlFsyFRGT_aXCivQ/edit?usp=sharing) problem.
I decided to use [Clean Architecture](https://manakuro.medium.com/clean-architecture-with-go-bce409427d31) for all the benefits which this architecture provides.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Project](#running-the-project)
- [Usage](#usage)
  - [Running Unit Tests](#running-unit-tests)
  - [Viewing Documentation](#viewing-documentation)

## Getting Started

### Prerequisites

- I was using latest [Go](https://golang.org/) version (1.21.4)
- Docker (optional, for running with Docker)

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/ggsomnoev/task-rope-bridge-georgi-somnoev.git
   ```

2. Change into the project directory:

   ```shell
   cd bizportal-task
   ```

3. Install project dependencies:

   ```shell
   go mod tidy
   ```

## Usage

### Running the Project

To run the project with docker:

```shell
docker-compose up -d unique-tail-moves-api
docker exec -it unique-tail-moves-api bash
```

Inside the container use the following command:

```shell
go run main.go
```
If you want to add a custom input please change the content in the "head_movements.txt":


### Running Unit Tests

To run unit tests, execute the following command:

```shell
go test ./...
```

This will run all unit tests in the project. You can use additional flags or specify specific test files or directories as needed.

### Viewing Documentation

To view the documentation locally, use the following steps:

1. To generate the documentation:
```shell
godoc -http=:6060
```
2. To view it directly in the terminal:
```shell
go doc package/path
```