# Angular_Go_Mongo_CRUD
Angular_Go_Mongo_CRUD

## Summary
The idea is to provide a Web Front end and back-end API for users. The app must be able to populate it's users table from a random name generation service. 

This was a time-constrained programming challenge, so I leverage bits of code that I already had sitting around. 

## Approach

Angular/TypeScript -> APIs (Go) -> MongoDB


### Front-end in Angular
Angular / Typescript

### APIs in Go
Go Lang 
Go currently amuses me, so the backend is implemented in Go.

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

## Todo 
The "learning-session" component in Angular needs to be renamed to relate to users, instead of the cloned code.
