# ACTE (Axxxxx Code Test Engineering)

![flow-design](internal/loans/architecture-design.png)


# To-Do :
- Proposed Code Problems
    - [x] Billing Engine (System Design and abstraction) [Billings Domain](internal/billings)
    - [x] reconciliation service (Algorithmic and scaling) [Reconciliations Domain](internal/reconciliations)
    - [x] Loan Service (system design and abstraction) [System Design Documentation](internal/loans/README.md)
    - [x] Postman with Example Response

## Technology Overview
- **Golang**: 1.21.5
- **Echo**: Web framework for building RESTful APIs
- **Goose**: Database migration tool
- **Docker**: Containerization platform
- **Viper**: Configuration management tool

## How To Run

### Prerequisites

Before running the application, ensure you have the following installed:

- **Docker**: Used for containerization. [Install Docker](https://docs.docker.com/get-docker/)
- **Go**: Programming language used for the project. [Install Go](https://golang.org/doc/install)
- **Goose**: Database migration tool used for managing database schema. Install it using Go:

  ```sh
  go install github.com/pressly/goose/v3@latest
  ```

### Running the Application

**1. Clone the Repository**

```sh
git clone https://github.com/martinyonatann/acte.git
cd acte
```

**2. Build and Run Database on Docker**
```sh
    make setup
```

**3. Run Database Migrations**

Use Goose to apply database migrations. You may need to configure Goose to point to your database.

```sh
    make migrate-up
```
**4. Start the Application**

 use the following Go command to start the application:
 ```sh
    go run main.go
```

**5. import Postman collection file into Postman**

import file `acte.postman_collection.json` into your postman