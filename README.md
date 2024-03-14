
# Go Docker Hub

A simple API and automatic Docker Hub image refresh with Go

## Installation

Pull this repo and follow the steps below

```bash
  git clone https://github.com/WeLedger/wl-auth.git
  cd wl-auth
```

```bash
  make up
```

## API Reference

#### Authentication

>Login

```http
  POST /api/v1/auth/login
```

```bash
curl -X POST http://localhost:8080/api/v1/auth/login
   -H "Content-Type: application/json"
   -d '{"username": "John", "password": "12345678"}'  
```

---
>Register

```http
  POST /api/v1/auth/register
```

```bash
curl -X POST http://localhost:8080/api/v1/auth/register
   -H "Content-Type: application/json"
   -d '{"email": "john@doe.com", "username": "John", "password": "12345678"}'  
```