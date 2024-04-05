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
  - [Run Locally](#running-run-locally)
- [Usage](#eyes-usage)
  - [Access Swagger UI](#access-swagger-ui)
  - [Build Docker Image](#build-docker-image)
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
  <img src="https://github.com/DuckyMomo20012/go-todo/assets/64480713/efa3b6cc-b6b1-4437-bb8e-30d65ef5cc1b" alt="swagger_ui" />
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

<!-- Env Variables -->

### :key: Environment Variables

> [!NOTE]
> All the environment variables file are required to run this project.

To run this project, you will need to add the following environment variables file:

- `internal/tasks/configs/cfg.env`:

  - `HOST`: The host of the server. Default is `localhost`.
  - `PORT`: The port of the server. Default is `8080`.

  - `DB_HOST`: The host of the Postgres database. Default is `localhost`.
  - `DB_PORT`: The port of the database. Default is `5432`.
  - `DB_USER`: The user of the database. Default is `postgres`.
  - `DB_PASSWORD`: The password of the database. Default is `postgres`.
  - `DB_NAME`: The name of the database. Default is `task`.
  - `CORS_ALLOW_ORIGIN`: The origin that is allowed to access the server.
    Default is `*`.


E.g:

```
# internal/tasks/configs/cfg.env
HOST=localhost
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=task
CORS_ALLOW_ORIGIN=*
```

You can also check out the file `internal/tasks/configs/cfg.env.example` to see
all required environment variables.

<!-- Getting Started -->

## :toolbox: Getting Started

<!-- Prerequisites -->

### :bangbang: Prerequisites

- Go: `1.22.1`.
- golangci-lint: `1.57.2`.

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

Start the Postgres database:

```bash
docker-compose up -d
```

Start the server:

```bash
go run ./cmd/tasks/main.go
```

<!-- Usage -->

## :eyes: Usage

### Access Swagger UI

Open your browser and go to `http://localhost:8081`.

### Build Docker Image

> [!NOTE]
> The file `internal/tasks/configs/cfg.env` is required to build the Docker image.

```bash
docker build -t go-todo -f ./docker/tasks/Dockerfile .
```

### Linting

```bash
make lint
```

<!-- Roadmap -->

## :compass: Roadmap

- [ ] gRPC support.

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

- Question 1

  - Answer 1.

- Question 2

  - Answer 2.

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
