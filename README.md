# Angular_Go_Mongo_CRUD
Angular_Go_Mongo_CRUD


## Front-End
Angular / Typescript

## Backend
Go Lang 

Go is not my language of most expertise, but I've been working in it recently

## Approach
Data persistence mongo

APIs in Go

Front-end in Angular

Example create user via CURL/POST
```bash
curl -H "Content-Type: application/json" -H "Accept: application/json" -X POST http://localhost:3001/users --data '{"name":"Zaphod Beeblebrox"}'

```

Example get users 
```bash
curl -H "Content-Type: application/json" -H "Accept: application/json" -X GET http://localhost:3001/users 
```

Trigger retrieval of random names and storing them in the data store
```bash
curl -H "Content-Type: application/json" -H "Accept: application/json" -X GET http://localhost:3001/users/populate_data 
```
