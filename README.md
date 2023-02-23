# jumia-clone

## About
This is a jumia app clone project.
The main goal of this project is to create the backend for the app.

## Service Architecture
The service implements `Clean Architecture` which helps to separate concerns by organizing code into several layers with a very explicit rule which enables us to create a testable and maintainable project.  [The Clean Architecture](https://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html).


### The project will be divided into five layers:
#### 1. Presentation
This represents logic that consume the business logic from the `Usecase Layer`
and renders to the view. Here you can choose to render the view in e.g `rest`

#### 2. Usecases
The code in this layer contains application specific business rules.
This represents the pure business logic of the application.
The rules of the application also shouldn't rely on the UI or the persistence frameworks being used.

#### 3. Interfaces
Clean architecture dictates that dependency should only point inwards therefore the inner layers(the usecase layer) should not have any idea of the implementations of the database, third party interactions. So this is just an interface.
This will ensures that the system is independent of a database and any third party agencies making it easier to switch them without affecting the business logic.

#### 4. Infrastructure
These are the `ports` that allow the system to talk to 'outside things' which
could be a `database` for persistence or a `web server` for the UI. None of
the inner use cases or domain entities should know about the implementation of
these layers and if we choose to change them they should not cause change to any of our business rules.

#### 5. Domain
Here we have `business objects` or `entities` and should represent and encapsulate the fundamental business rules.


## Technologies
 1. Golang(version 1.17)

## How to use.
1. First, install golang. You can refer to this link [Install golang](https://go.dev/doc/install).

2. Clone the code from the repository.

3. The next step is setting up your local environment by creating an `env.sh` and add your envs.

```bash
    # env.sh
    export PORT="8080"

    (other service specific variables)
```
4. Install the dependencies.

5. Finally, run to see the server at http://localhost:8080

```bash
    user@user:~$ go run server.go
```
