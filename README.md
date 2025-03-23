# JSON Web Token Authentication in Go

> Disclaimer: This repository is not about implementing the JWT itself.
> Rather it was about applying JWT to get access to a protected backend service

## Features

- Login & Register
- CRUD user operation for authenticated user

## Limitations

- Update user won't update it's password

## Perquisites

1. Acquire self-signed certificates

Generate private key

```bash
openssl genpkey -algorithm ed25519 -out certs/priv.key
```

Generate certificates

```bash
openssl req -x509 -key certs/priv.key -out certs/localhost.pem -subj "/CN=localhost"
```

2. Run database migration

```bash
go run cmd/database/migration.go
```

## Usage

Make sure you have your `.env` file ready and loaded, it only read from os.GetEnv() not from `.env` file. You could do something like: `docker run --env-file=.env`

```bash
go run cmd/server/main.go
```

## API Reference

Base URL: `https://localhost:8080`

### Authentication

#### Register

```http
  POST /register
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Name of user |
| `email`      | `string` | **Required**. Email of user |
| `password`      | `string` | **Required**. Password of user |

#### Login

```http
  POST /login
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email`      | `string` | **Required**. Email of user |
| `password`      | `string` | **Required**. Password of user |

### User

#### Get all user

```http
  GET /users

  Authorization: Bearer <token>
```

#### Get user

```http
  GET /users/${id}

  Authorization: Bearer <token>
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of user |

#### Create user

```http
  POST /users

  Authorization: Bearer <token>
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Name of user |
| `email`      | `string` | **Required**. Email of user |
| `password`      | `string` | **Required**. Password of user |

#### Update user

```http
  PUT /users/${id}

  Authorization: Bearer <token>
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of user |

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Name of user |
| `email`      | `string` | **Required**. Email of user |

#### Delete user

```http
  DELETE /users/${id}

  Authorization: Bearer <token>
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of user |
