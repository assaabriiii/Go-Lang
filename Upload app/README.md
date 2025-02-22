Download and install Go from the official website: https://go.dev/dl/.

download `gin` from the official website:

```bash
go get -u github.com/gin-gonic/gin
```

then run the project using:

```bash
go run main.go
```

- POST a file:

```bash
curl -X POST http://localhost:8080/upload \
-F "file=@/path/to/your/file.txt"
```

- GET all file names:

```bash
curl -X GET http://localhost:8080/files
```

- GET a file:

```bash
curl -X GET http://localhost:8080/download/test.txt --output test.txt
```