# My Microservice Project

This is a small blogging app that I decided to build with a microservice architecture. Normally, an app like this would take me 2 to 3 hours to build using a monolithic approach, but it took me far longer (3 weeks) with microservices.

## Tech Stack

### Frontend
- React
- TypeScript

### Backend
- Node
- Express
- TypeScript

### Databases
- MongoDB

### Infrastructure
- Docker
- Kubernetes
- GCP
- NATS Streaming Server (Event Bus)

## App Architecture

So far, this application is comprised of four services all running on seperate kubernetes **pods**. Each service contains its own **MongoDB database** which are also running on their own pod. Communication between pods are done via a **Cluster IP Service** which is not shown in the diagram.

Outside communcation is done through an **Ingress Ngnix** service that forwards requests to the appropriate **Cluster IP** service in the kubernetes cluster.

Communication between services are done **asynchronously** through the use of an **event bus** known as **NATS Streaming Server**. Apache Kafka was my original option, but it was too complex to understand and implement into this introductory project. I do intend to learn it in the future, however. 

<img width="1084" alt="Screen Shot 2021-04-04 at 8 36 50 AM" src="https://user-images.githubusercontent.com/35265876/113508971-1c86fc80-9521-11eb-98c5-bc45151ac148.png">

## Authentication

Authentication is a very tricky thing to accomplish in a microservice architecture. This is because all of the authentication logic is housed in the auth service, seperate from the other services. However, knowing whether a user is authenticated or not is very important for any service that contains protected routes.

One approach to communicate authentication information between services is for a service to **synchronously** communucate with the auth service.

<img width="843" alt="Screen Shot 2021-04-04 at 9 30 42 AM" src="https://user-images.githubusercontent.com/35265876/113510346-72ab6e00-9528-11eb-9975-35c3245ed4d2.png">

However, this introduces dependencies between multiple services and the auth service. This is not ideal because if the auth service goes down, multiple routes within other services will not be functional. 

A better approach is to teach each service how to determine if a user is authenticated or not. In our application, an authenticated user will return a **json web token** in each request header. Each service should have the ability to get the JWT, verify it and extract the payload.

However, the generation of the JWT, and other functionality such as signing in, signing up, and logging out can only be done in the auth service.

## Error Handling

One of the perks of microservices is that they allow you to use multiple different languages and frameworks across multiple services. This however can come with its own drawbacks. 

For instance, different libraries or frameworks can return different error structures. This is problematic, as we don't want our React App to have to figure out how to parse the structure of each error it gets thrown its way.

To fix this, I created a common library that will be used by each service and will ensure a consistent error structure.

## Service Communication

There are two man ways communication between services can occur; synchronously or asynchronously. Synchronously communication occurs when one service communicates directly with another via an HTTP request. This is a simple solution, but it may lead to many issues. When services communicate with one another, it can lead to dependencies between services. Thus if one service goes down, certain routes in other services might become comprised.

Asynchronous communication solves this issue. Communication between services is down via events that are emitted by the publisher and consumed by the subscribers. This can be done with the use of an event bus like Kafka, RabbitMQ, or NATS. This ensures that there is no dependencies across services.
