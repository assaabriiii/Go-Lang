Download and install Go from the official website: https://go.dev/dl/.

download `gin` from the official website:

```bash
go get -u github.com/gin-gonic/gin
```

then run the project using:

```bash
go run main.go
```


- GET all contacts:

```bash
curl -X GET http://localhost:8080/contacts
```


- GET a contact by ID:

```bash
curl -X GET http://localhost:8080/contacts/1
```


- POST a contact:

```bash
curl -X POST http://localhost:8080/contacts \
-H "Content-Type: application/json" \
-d '{
  "name": "Alice Johnson",
  "email": "alice@example.com",
  "phone": "555-555-5555"
}'
```


- PUT or update a contact:

```bash
curl -X PUT http://localhost:8080/contacts/1 \
-H "Content-Type: application/json" \
-d '{
  "name": "John Doe (Updated)",
  "email": "john.doe@example.com",
  "phone": "111-222-3333"
}'
```


- DELETE a contact:

```bash
curl -X DELETE http://localhost:8080/contacts/1
```