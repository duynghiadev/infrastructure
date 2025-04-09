# "ToDo API" Microservice Example

## Introduction

Welcome! 👋

This is an educational repository that includes a microservice written in Go. It is used as the principal example of my video series: [Building Microservices in Go](https://www.youtube.com/playlist?list=PL7yAAGMOat_Fn8sAXIk0WyBfK_sT1pohu).

This repository **is not** a template **nor** a framework, it's a **collection of patterns and guidelines** I've successfully used to deliver enterprise microservices when using Go, and just like with everything in Software Development some trade-offs were made.

My end goal with this project is to help you learn another way to structure your Go project with 3 final goals:

1. It is _enterprise_, meant to last for years,
2. It allows a team to collaborate efficiently with little friction, and
3. It is as idiomatic as possible.

Join the fun at [https://youtube.com/MarioCarrion](https://www.youtube.com/c/MarioCarrion).

## Domain Driven Design

This project uses a lot of the ideas introduced by Eric Evans in his book [Domain Driven Design](https://www.domainlanguage.com/), I do encourage reading that book but before I think reading [Domain-Driven Design Distilled](https://smile.amazon.com/Domain-Driven-Design-Distilled-Vaughn-Vernon/dp/0134434420/) makes more sense, also there's a free to download [DDD Reference](https://www.domainlanguage.com/ddd/reference/) available as well.

On YouTube I created [a playlist](https://www.youtube.com/playlist?list=PL7yAAGMOat_GJqfTdM9PBdTRSH7jXs6mI) that includes some of my favorite talks and webinars, feel free to explore that as well.

## Project Structure

Talking specifically about microservices **only**, the structure I like to recommend is the following, everything using `<` and `>` depends on the domain being implemented and the bounded context being defined.

- [ ] `build/`: defines the code used for creating infrastructure as well as docker containers.
  - [ ] `<cloud-providers>/`: define concrete cloud provider.
  - [ ] `<executableN>/`: contains a Dockerfile used for building the binary.
- [ ] `cmd/`
  - [ ] `<primary-server>/`: uses primary database.
  - [ ] `<replica-server>/`: uses readonly databases.
  - [ ] `<binaryN>/`
- [x] `db/`
  - [x] `migrations/`: contains database migrations.
  - [ ] `seeds/`: contains file meant to populate basic database values.
- [ ] `internal/`: defines the _core domain_.
  - [ ] `<datastoreN>/`: a concrete _repository_ used by the domain, for example `postgresql`
  - [ ] `http/`: defines HTTP Handlers.
  - [ ] `service/`: orchestrates use cases and manages transactions.
- [x] `pkg/` public API meant to be imported by other Go package.

There are cases where requiring a new bounded context is needed, in those cases the recommendation would be to
define a package like `internal/<bounded-context>` that then should follow the same structure, for example:

- `internal/<bounded-context>/`
  - `internal/<bounded-context>/<datastoreN>`
  - `internal/<bounded-context>/http`
  - `internal/<bounded-context>/service`

## Tools

```
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1
go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.6.0
go install github.com/maxbrunsfeld/counterfeiter/v6@v6.3.0
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.5.1
go install goa.design/model/cmd/mdl@v1.7.6
go install goa.design/model/cmd/stz@v1.7.6
```

## Features

Icons meaning:

- ![YouTube icon](https://img.shields.io/badge/-YouTube-black?style=flat&logo=youtube&logoColor=red) means a link to YouTube video.
- ![Blog icon](https://img.shields.io/badge/-Blog-black?style=flat&logo=linktree&logoColor=white) means a link to Blog post.

In no particular order:

- [x] Project Layout
- [x] Dependency Injection
- [x] [Secure Configuration](docs/SECURE_CONFIGURATION.md)
  - [x] Using [Hashicorp Vault](https://www.hashicorp.com/products/vault)
  - [x] Using [AWS SSM](https://aws.amazon.com/systems-manager/features/#Parameter_Store)
- [ ] Infrastructure as code
- [x] [Metrics, Traces and Logging using OpenTelemetry](docs/METRICS_TRACES_LOGGING.md)
- [ ] Caching
  - [x] Memcached
  - [ ] Redis
- [x] Persistent Storage
  - [x] Repository Pattern
  - [x] Database migrations
  - [ ] MySQL
  - [x] [PostgreSQL](docs/PERSISTENT_STORAGE.md)
    - [`jmoiron/sqlx`](https://github.com/jmoiron/sqlx), [`jackc/pgx`](https://github.com/jackc/pgx) and [`database/sql`](https://pkg.go.dev/database/sql)
    - [`go-gorm/gorm`](https://github.com/go-gorm/gorm) and [`volatiletech/sqlboiler`](https://github.com/volatiletech/sqlboiler)
    - [`Masterminds/squirrel`](https://github.com/Masterminds/squirrel) and [`kyleconroy/sqlc`](https://github.com/kyleconroy/sqlc)
- [ ] REST APIs
  - [x] HTTP Handlers
  - [x] Custom JSON Types
  - [ ] Versioning
  - [x] Error Handling
  - [x] [OpenAPI 3 and Swagger-UI](docs/OPENAPI3_SWAGGER.md)
  - [ ] Authorization
- [ ] Events and Messaging
  - [ ] [Apache Kafka](https://kafka.apache.org/)
  - [ ] [RabbitMQ](https://www.rabbitmq.com/)
  - [ ] [Redis](https://redis.io/)
- [ ] Testing
  - [x] Type-safe mocks with [`maxbrunsfeld/counterfeiter`](https://github.com/maxbrunsfeld/counterfeiter)
  - [x] Equality with [`google/go-cmp`](https://github.com/google/go-cmp)
  - [x] Integration tests for Datastores with [`ory/dockertest`](https://github.com/ory/dockertest)
  - [x] REST APIs
- [ ] Containerization using Docker
- [ ] Graceful Shutdown
- [ ] Search Engine using [ElasticSearch](https://www.elastic.co/elasticsearch/)
- [ ] Whatever else I forgot to include

## More ideas

- [2016: Peter Bourgon&#39;s: Repository structure](https://peter.bourgon.org/go-best-practices-2016/#repository-structure)
- [2016: Ben Johnson&#39;s: Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
- [2017: William Kennedy&#39;s: Design Philosophy On Packaging](https://www.ardanlabs.com/blog/2017/02/design-philosophy-on-packaging.html)
- [2017: Jaana Dogan&#39;s: Style guideline for Go packages](https://rakyll.org/style-packages/)
- [2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

## Docker Containers

Please notice in order to run this project locally you need to run a few programs in advance, if you use Docker please refer to the concrete instructions in [`docs/`](docs/) for more details.

There's also a [docker-compose.yml](docker-compose.yml), covered in [Building Microservices In Go: Containerization with Docker](https://youtu.be/u_ayzie9pAQ), however like I mentioned in the video you have to execute `docker-compose` in multiple steps.

Notice that because of the way RabbitMQ and Kafka are being used they are sort of competing with each other, so at the moment we either have to enable Kafka and disable RabbitMQ or the other way around in both the code and the `docker-compose.yml` file, in either case there are Dockerfiles and services defined that cover building and running them.

- Run `docker-compose up`, here both _rest-server_ and _elasticsearch-indexer_ services will fail because the `postgres`, `rabbitmq`, `elasticsearch` and `kafka` services take too long to start.
  - If you're planning to use RabbitMQ, run `docker-compose up rest-server elasticsearch-indexer-rabbitmq`.
  - If you're planning to use Kafka, run `docker-compose up rest-server elasticsearch-indexer-kafka`.
  - If you're planning to use Redis, run `docker-compose up rest-server elasticsearch-indexer-redis`.
- For building the service images you can use:
  - `rest-server` image: `docker-compose build rest-server`.
  - `elasticsearch-indexer-rabbitmq` image: `docker-compose build elasticsearch-indexer-rabbitmq`.
  - `elasticsearch-indexer-kafka` image: `docker-compose build elasticsearch-indexer-kafka`.
  - `elasticsearch-indexer-redis` image: `docker-compose build elasticsearch-indexer-redis`.
- Run `docker-compose run rest-server migrate -path /api/migrations/ -database postgres://user:password@postgres:5432/dbname?sslmode=disable up` to finally have everything working correctly.

## Diagrams

To start a local HTTP server that serves a graphical editor:

```
mdl serve github.com/MarioCarrion/todo-api/internal/doc -dir docs/diagrams/
```

To generate JSON artifact for uploading to [structurizr](https://structurizr.com/):

```
stz gen github.com/MarioCarrion/todo-api/internal/doc
```
