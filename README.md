<div align="center">

  <h1>Go Todo</h1>

  <p>
    Simple Todo api using Go
  </p>

<!-- Badges -->
<p>
  <a href="https://github.com/DuckyMomo20012/go-todo/graphs/contributors">
    <img src="https://img.shields.io/github/contributors/DuckyMomo20012/go-todo" alt="contributors" />
  </a>
  <a href="">
    <img src="https://img.shields.io/github/last-commit/DuckyMomo20012/go-todo" alt="last update" />
  </a>
  <a href="https://github.com/DuckyMomo20012/go-todo/network/members">
    <img src="https://img.shields.io/github/forks/DuckyMomo20012/go-todo" alt="forks" />
  </a>
  <a href="https://github.com/DuckyMomo20012/go-todo/stargazers">
    <img src="https://img.shields.io/github/stars/DuckyMomo20012/go-todo" alt="stars" />
  </a>
  <a href="https://github.com/DuckyMomo20012/go-todo/issues/">
    <img src="https://img.shields.io/github/issues/DuckyMomo20012/go-todo" alt="open issues" />
  </a>
  <a href="https://github.com/DuckyMomo20012/go-todo/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/DuckyMomo20012/go-todo.svg" alt="license" />
  </a>
</p>

<h4>
    <a href="https://github.com/DuckyMomo20012/go-todo/">View Demo</a>
  <span> · </span>
    <a href="https://github.com/DuckyMomo20012/go-todo">Documentation</a>
  <span> · </span>
    <a href="https://github.com/DuckyMomo20012/go-todo/issues/">Report Bug</a>
  <span> · </span>
    <a href="https://github.com/DuckyMomo20012/go-todo/issues/">Request Feature</a>
  </h4>
</div>

<br />

<!-- Table of Contents -->

# :notebook_with_decorative_cover: Table of Contents

- [About the Project](#star2-about-the-project)
  - [Screenshots](#camera-screenshots)
  - [Tech Stack](#space_invader-tech-stack)
  - [Features](#dart-features)
  - [Environment Variables](#key-environment-variables)
- [Getting Started](#toolbox-getting-started)
  - [Prerequisites](#bangbang-prerequisites)
  - [Run with Docker Compose](#whale-run-with-docker-compose)
  - [Run Locally](#running-run-locally)
  - [Running Tests](#test_tube-running-tests)
- [Usage](#eyes-usage)
  - [Access Swagger UI](#access-swagger-ui)
  - [Build Docker Image](#build-docker-image)
  - [Makefile](#makefile)
- [Roadmap](#compass-roadmap)
- [Contributing](#wave-contributing)
  - [Code of Conduct](#scroll-code-of-conduct)
- [FAQ](#grey_question-faq)
- [License](#warning-license)
- [Contact](#handshake-contact)
- [Acknowledgements](#gem-acknowledgements)

<!-- About the Project -->

## :star2: About the Project

<!-- Screenshots -->

### :camera: Screenshots

<div align="center">
  <img src="https://github.com/DuckyMomo20012/go-todo/assets/64480713/911e6a3e-28bb-42b1-8653-4dd99e1e31c2" alt="swagger_ui" />
  <i>Swagger UI</i>
</div>

<!-- TechStack -->

### :space_invader: Tech Stack

<details>
  <summary>Server</summary>
  <ul>
    <li><a href="https://go.dev"> Golang</a></li>
  </ul>
</details>

<details>
<summary>Database</summary>
  <ul>
    <li><a href="https://www.postgresql.org/">PostgreSQL</a></li>
  </ul>
</details>

<details>
<summary>DevOps</summary>
  <ul>
    <li><a href="https://www.docker.com/">Docker</a></li>
  </ul>
</details>

<!-- Features -->

### :dart: Features

- Basic CRUD operations.
- Swagger UI for API documentation.
- Simple CLI for running the server.
- gRPC Gateway server.

<!-- Env Variables -->

### :key: Environment Variables

> [!NOTE]
> All the environment variables file are required to run this project.

To run this project, you will need to add the following environment variables
file:

- `internal/gateway/configs/.env`: Gateway service environment variables.

  - `HOST`: The host of the server. Default is `0.0.0.0`.
  - `PORT`: The port of the server. Default is `8081`.
  - `APP_ENV`: The environment of the application. Default is `development`. It
    can be `production` or `development`.
  - `LOG_LEVEL`: The log level of the application. Default is `0`. See available
    level in
    [zerolog](https://github.com/rs/zerolog?tab=readme-ov-file#leveled-logging).
  - `LOG_SAMPLE_RATE`: The sample rate of the log. Default is `5`.

  - `TASK_SERVER_ADDRESS`: The address of the task service. Example: `localhost:8080`.

  E.g:

  ```
  # internal/gateway/configs/.env
  HOST="0.0.0.0"
  PORT=8081
  APP_ENV="development"
  LOG_LEVEL=0
  LOG_SAMPLE_RATE=5

  TASK_SERVER_ADDRESS="localhost:9000"
  ```

  You can also check out the file `internal/gateway/configs/.env.example` to see
  all required environment variables.


- `internal/task/configs/.env`: Task service environment variables.

  - `PORT`: The port of the server. Default is `8080`.
  - `APP_ENV`: The environment of the application. Default is `development`. It
    can be `production` or `development`.
  - `LOG_LEVEL`: The log level of the application. Default is `0`. See available
    level in
    [zerolog](https://github.com/rs/zerolog?tab=readme-ov-file#leveled-logging).
  - `LOG_SAMPLE_RATE`: The sample rate of the log. Default is `5`.

  - `DB_URL`: The URL of the database. Example:
    `postgres://postgres:postgres@localhost:5432/task?sslmode=disable`.

  E.g:

  ```
  # internal/task/configs/.env
  PORT=9000
  APP_ENV="development"
  LOG_LEVEL=0
  LOG_SAMPLE_RATE=5

  DB_URL="postgresql://postgres:postgres@localhost:5432/task?sslmode=disable"

  ```

  You can also check out the file `internal/task/configs/.env.example` to see
  all required environment variables.

<!-- Getting Started -->

## :toolbox: Getting Started

<!-- Prerequisites -->

### :bangbang: Prerequisites

- Go: `1.22.1`.

- Docker: `26.1.2`.

- Brew tools:

  All required `brew` tools is placed in `internal/tools/Brewfile`:

  ```bash
  make download-deps
  ```

- Go tools:

  All required Go tools is placed in `internal/tools/tools.go`:

  ```bash
  make download-deps
  ```

> [!NOTE]
> These dependencies are not included during build.

<!-- Run with Docker Compose -->

### :whale: Run with Docker Compose

Update the environment variables files:

Please check the [Environment Variables](#key-environment-variables) section to
see all required environment variables.

Run the server:

```bash
docker-compose up -d
```

Access the Swagger UI at `http://localhost/docs`.

Access the API at `http://localhost/api`.

<!-- Run Locally -->

### :running: Run Locally

Clone the project:

```bash
git clone https://github.com/DuckyMomo20012/go-todo.git
```

Go to the project directory:

```bash
cd go-todo
```

Install dependencies:

```bash
go mod download
```

Update the environment variables files:

Please check the [Environment Variables](#key-environment-variables) section to
see all required environment variables.

Start the Postgres database:

```bash
docker-compose up -d
```

Start the server:

```bash
go run ./main.go gateway start

go run ./main.go task start
```

Or with [wgo](https://github.com/bokwoon95/wgo) for live reload:

```bash
wgo run ./main.go gateway start

wgo run ./main.go task start
```

> [!NOTE]
> The gateway API is mount to `localhost:<PORT>/api` not `localhost:<PORT>`.

<!-- Running Tests -->

### :test_tube: Running Tests

#### Test API with Postman

- REST API:

  You can test the REST API with Postman by importing files from
  [`/api/openapi/`](/api/openapi/) to your Postman.

> [!NOTE]
> You may have to add the `Authorization` header to your requests. You can get a
> token by logging in with the `login` API.

- gRPC API:

  You can test the gRPC API with Postman by using Postman reflection feature.

<!-- Usage -->

## :eyes: Usage

### Access Swagger UI

Open your browser and go to `http://localhost/docs`.

### Build Docker Image

> [!NOTE]
> Environment variables files are required to build the Docker image. Check the
> [Environment Variables](#key-environment-variables) section for more
> information.

```bash
docker build -t go-todo -f ./docker/tasks/Dockerfile .
```

Or build ghcr image:

```bash
make docker-build
```

### Makefile

`Makefile` provides some useful targets to help you work with this project:

- `init`: Download tool dependencies and setup `GOPRIVATE` environment variable.

  ```bash
  make init
  ```

> [!NOTE]
> Setup `GOPRIVATE` env for vscode not automatically going to
> [pkg.go.dev](https://pkg.go.dev/) for private modules.

- `download-deps`: Download all tool dependencies.

  ```bash
  make download-deps
  ```

- `gen-proto`: Generate gRPC and gRPC gateway from proto files.

  ```bash
  make gen-proto
  ```

- `lint`: Run lint with `golangci-lint`.

  ```bash
  make lint
  ```

- `docker-build`: Build Docker image for ghcr.io registry.

  ```bash
  make docker-build
  ```

- `docker-push`: Push Docker image to ghcr.io registry.

  ```bash
  make docker-push
  ```

<!-- Roadmap -->

## :compass: Roadmap

- [x] gRPC support.
- [x] gRPC Gateway support.
- [x] CLI support.

<!-- Contributing -->

## :wave: Contributing

<a href="https://github.com/DuckyMomo20012/go-todo/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=DuckyMomo20012/go-todo" />
</a>

Contributions are always welcome!

<!-- Code of Conduct -->

### :scroll: Code of Conduct

Please read the [Code of Conduct](https://github.com/DuckyMomo20012/go-todo/blob/main/CODE_OF_CONDUCT.md).

<!-- FAQ -->

## :grey_question: FAQ

- I can't query requests using the Swagger UI.

  - Currently with the Buf plugin
    [`openapiv2`](https://buf.build/grpc-ecosystem/openapiv2) can only generate
    the OpenAPI v2 spec. Therefore, I can't set hostname using variable which
    only available in OpenAPI v3.

  - The `swagger` service defined in file `docker-compose.yaml` is running with
    port `8082` and sending requests using `localhost:8082`. However, the gRPC
    gateway service is running with port `8081`.

- Why do you migrate to
  [`protoc-gen-openapiv2`](https://pkg.go.dev/github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
  )?

  - The original port for REST API is removed in
    [PR#8](https://github.com/DuckyMomo20012/go-todo/pull/8).

  - Because I decided to switch to
    [`grpc-gateway`](https://github.com/grpc-ecosystem/grpc-gateway) so I can
    automatically generate REST API from proto files.

  - For the OpenAPI spec, I used to use `protoc-gen-openapiv2` (Using buf
    plugin: [`openapiv2`](https://buf.build/grpc-ecosystem/openapiv2)) to
    generate the OpenAPI v2 spec. However, it can only generate the OpenAPI v2
    so I may have to switch to another plugin that can generate OpenAPI v3 spec
    later.

- Why do you rename all vars and files from `tasks` to `task`?

  - Bad naming convention.

- How can test the gRPC server with Postman?

  - Since the gRPC server enabled
    [`reflection`](https://pkg.go.dev/google.golang.org/grpc/reflection) service
    in [PR#9](https://pkg.go.dev/google.golang.org/grpc/reflection), you can use
    option `Using server reflection` from tab `Service definition` in your gRPC
    requests in Postman.


<!-- License -->

## :warning: License

Distributed under MIT license. See
[LICENSE](https://github.com/DuckyMomo20012/go-todo/blob/main/LICENSE)
for more information.

<!-- Contact -->

## :handshake: Contact

Duong Vinh - [@duckymomo20012](https://twitter.com/duckymomo20012) -
tienvinh.duong4@gmail.com

Project Link: [https://github.com/DuckyMomo20012/go-todo](https://github.com/DuckyMomo20012/go-todo).

<!-- Acknowledgments -->

## :gem: Acknowledgements

Here are useful resources and libraries that we have used in our projects:

- [Clean Architecture](https://threedots.tech/post/introducing-clean-architecture/):
  A detailed explanation of Clean Architecture by Three Dots Labs.
- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway): gRPC to JSON
  proxy generator following the gRPC HTTP spec.
- [Buf CLI](https://buf.build/product/cli): A new way to work with Protocol
  Buffers.
