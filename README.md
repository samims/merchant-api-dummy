# Merchant CRUD API service

Merchant CRUD API service is a RESTful API service that provides the following functionality:

## Authentication

    - Signup
    - Signin by email and password(JWT)

## User APIs

    - List of all users
    - Update a user

## Merchant APIs

    - List of all merchants
    - Create a new merchant
    - Read a merchant
    - Update a merchant if associated with the user
    - Delete a merchant if associated with the user
    - Add team member to a merchant if associated with the user
    - Remove team member from a merchant if associated with the user


## Run the service

There is a Dockerfile in the `docker/` directory of the project.
To run the service:
- go to the project directory `/docker`
- rename .env.example to .env
- run `docker-compose up`

