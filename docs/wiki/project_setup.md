# Project setup

The project use **Docker** and **Docker Compose** to build and run local and distant images in our workspace.

## Stack

All the images use the **same network**, more informations at [docker-compose.yml](docker-compose.yml)

| CONTAINER | PORT        | IMAGE                                                                    |
| :-------- | :---------- | :----------------------------------------------------------------------- |
| GOLANG    | `3333:3333` | [build/pakage/sample-api/Dockerfile](build/pakage/sample-api/Dockerfile) |
| ADMINER   | `3334:8080` | [build/package/adminer/Dockerfile](build/package/adminer/Dockerfile)     |
| POSTGRES  | `5432:5432` | [postgres:latest](https://hub.docker.com/_/postgres)                     |
| WEBAPP    | `3000:80`   | [build/pakage/web/Dockerfile.prod](build/pakage/web/Dockerfile.prod)     |

> Adminer is a GUI that allows us to **manage your database** by permetting to to **create, edit, delete** the different entities, tables, etc.

## Makefile

### TL;DR <!-- omit in toc -->

```bash
make setup-env start logs
```

### `make help` <!-- omit in toc -->

**Display** informations about other commands.

### `make setup-env` <!-- omit in toc -->

**Copy** the sample environment files.

### `make start` <!-- omit in toc -->

Up the containers with **full cache reset** to avoid cache errors.

### `make stop` <!-- omit in toc -->

**Down** the containers.

### `make logs` <!-- omit in toc -->

**Display and follow** the logs.

### `make init-web` <!-- omit in toc -->

**Install Dependencies and build the app.**

### `make init` <!-- omit in toc -->

**All previous actions, without down.**

### `make lint` <!-- omit in toc -->

**Lint** the Go files using `gofmt`.
