# Property App

This project is a backend service for managing property listings. It provides functionalities to add, update, delete, and retrieve property information. The service is built using Go

## Features

- Add new properties
- Update existing properties
- Delete properties by ID
- Retrieve all properties
- Retrieve properties by ID


## Setup

1. **Clone the repository**
   ```sh
   git clone https://github.com/ysnkbk/golang-property-app.git
   cd property-app
    ```
   
2. **Install dependencies**
   ```sh
   go mod tidy
   ```
3. **Set up the database**

   - Ensure you have a PostgreSQL database running.
   - Update the database connection string in your configuration.
   
3. **Run the application**
   ```sh
   bash test/scripts/test_db.sh
   go run main.go
   ```

3. **Run the tests**
   ```sh
    go test ./...
    ```

