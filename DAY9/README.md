# Day 9: HTTP Server & REST API in Go (Optional)

## Learn:
- Build a simple HTTP server.
- Handle GET and POST requests.
- JSON API using `net/http`.

## Resources:
- [Writing an HTTP Server](https://gobyexample.com/http-servers)
- [Building a REST API in Go](https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-28g5)

## Exercises:
1. **Create an HTTP server that listens on port `8080`.**
2. **Implement a simple API endpoint that returns a JSON response.**
3. **Handle POST requests and parse incoming JSON data.**

## Notes:
- `http.HandleFunc()` maps URLs to handler functions.
- `json.NewEncoder()` encodes responses in JSON.
- `json.NewDecoder()` decodes JSON request bodies.
