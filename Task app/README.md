Download and install Go from the official website: https://go.dev/dl/.

download `gin` from the official website:

```bash
go get -u github.com/gin-gonic/gin
```

then run the project using:

```bash
go run main.go
```

- GET all tasks:

```bash
curl -X GET http://localhost:8080/tasks
```

- GET task specified to a ID:

```bash
curl -X GET http://localhost:8080/tasks/1
```

- POST new task:

```bash
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{
  "title": "Deploy the API",
  "description": "Deploy the task manager API to a cloud platform"
}'
```

- PUT and update a task:

```bash
curl -X PUT http://localhost:8080/tasks/1 \
-H "Content-Type: application/json" \
-d '{
  "title": "Learn Go (Updated)",
  "description": "Complete the Go tutorial and build a project"
}'
```

- DELETE a task:

```bash
curl -X DELETE http://localhost:8080/tasks/1
```