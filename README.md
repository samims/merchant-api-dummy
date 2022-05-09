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
- project will be up on port 3001

## Test the service
To test the service:
- run from root directory 
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

## Technologies & Tools Used
[Docker](https://www.docker.com/)
[Docker Compose](https://docs.docker.com/compose/overview/)
[GoLang](https://golang.org/doc/default_packages/net.html)
[GO Chi chi](https://github.com/go-chi/chi)
[Postgres](https://www.postgresql.org/)
[Beego ORM](https://beego.me/docs/mvc/basics.html)
[Mockery](https://github.com/vektra/mockery)
[Swagger](https://swagger.io/)
[Git](https://git-scm.com/)
[VSCode](https://code.visualstudio.com/)


    

## Run the service

There is a Dockerfile in the `docker/` directory of the project.
To run the service:
- go to the project directory `/docker`
- rename .env.example to .env
- run `docker-compose up`
- project will be up on port 3001

## Test the service
To test the service:
run from root directory `go test ./... -v`
