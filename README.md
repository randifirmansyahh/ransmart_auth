# Randi Firmansyah

## Getting started

Welcome!

## To run this project you can type text below on the terminal / bash :

```cmd
go run .
```

# E-Commerce Service for

- User

## LOGIN

- POST http://localhost:1000/login

```json
{
  "username": "randi",
  "password": "randi"
}
```

## REGISTER

- POST http://localhost:1000/register

```json
{
  "firstname": "randi",
  "lastname": "firmansyah",
  "username": "randi04",
  "password": "123321",
  "email": "randy@gmail.com",
  "no_hp": "+628515123123",
  "image": "https://1757140519.rsc.cdn77.org/blog/wp-content/uploads/2018/05/1-google-logo.png"
}
```

## USER (Login Required)

- GET http://localhost:1000/user
- GET http://localhost:1000/user/{id}
- DELETE http://localhost:1000/user/{id}
- PUT http://localhost:1000/user/{id}

```json
{
  "nama": "Randi Firmansyah",
  "username": "randifirmansyahh",
  "password": "randi123!",
  "email": "randykelvin26@gmail.com",
  "no_hp": "0854545454",
  "image": "https://1757140519.rsc.cdn77.org/blog/wp-content/uploads/2018/05/1-google-logo.png"
}
```

- POST http://localhost:1000/user

```json
{
  "nama": "Randi Firmansyah",
  "username": "randifirmansyahh",
  "password": "randi123!",
  "email": "randykelvin26@gmail.com",
  "no_hp": "0854545454",
  "image": "https://1757140519.rsc.cdn77.org/blog/wp-content/uploads/2018/05/1-google-logo.png"
}
```
