# Go CRUD API

This project is a simple CRUD (Create, Read, Update, Delete) API built with Go. It allows you to manage a collection of users.

## Routes

### Create a User
- **URL:** `/users`
- **Method:** `POST`
- **Description:** Adds a new user to the collection.
- **Request Body:**
    ```json
    {
        "name": "John Doe",
        "email": "john.doe@example.com"
    }
    ```

### Get All Users
- **URL:** `/users`
- **Method:** `GET`
- **Description:** Retrieves a list of all users.

### Get a User by ID
- **URL:** `/users/{id}`
- **Method:** `GET`
- **Description:** Retrieves a user by their ID.

### Update a User
- **URL:** `/users/{id}`
- **Method:** `PUT`
- **Description:** Updates the details of a user.
- **Request Body:**
    ```json
    {
        "name": "John Doe",
        "email": "john.doe@newexample.com"
    }
    ```

### Delete a User
- **URL:** `/users/{id}`
- **Method:** `DELETE`
- **Description:** Deletes a user by their ID.

## Example Usage with cURL

### Create a User
```sh
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{
    "name": "Dante",
    "email": "dante@example.com"
}'
```

### Get All Users
```sh
curl http://localhost:8080/users
```

### Get a User by ID
```sh
curl http://localhost:8080/users/1
```

### Update a User
```sh
curl -X PUT http://localhost:8080/users/1 -H "Content-Type: application/json" -d '{
    "name": "Dante Alighieri",
    "email": "dante.alighieri@example.com"
}'
```

### Delete a User
```sh
curl -X DELETE http://localhost:8080/users/1
```

## Running with Docker and Docker Compose

To run this project using Docker and Docker Compose, follow these steps:

1. Ensure you have Docker and Docker Compose installed on your machine.

2. Create a `.env` file in the root of your project with the following content:
    ```env
    DATABASE_URL=host=localhost user=youruser password=yourpass dbname=yourdb sslmode=disable
    ```

3. Build and start the containers using Docker Compose:
    ```sh
    docker-compose up --build
    ```

4. The API will be available at `http://localhost:8000`.

## Acknowledgements

This project was created based on the code from the [go-crud-live](https://github.com/FrancescoXX/go-crud-live) repository. Special thanks to [FrancescoXX](https://github.com/FrancescoXX/) for the sharing his knowledge.