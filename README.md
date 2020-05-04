# Angular_Go_Mongo_CRUD
Angular_Go_Mongo_CRUD



## Approach
Data persistence mongo

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
