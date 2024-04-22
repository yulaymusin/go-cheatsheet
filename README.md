# go-cheatsheet

## postgres

### 1. Connect

`go run postgres/1-connect.go`

*Hello, world!*

## gorilla/mux

### 1. Return a small text message for the /hello path

`go run gorillamux/1-newrouter.go`

`curl localhost:8080/hello`

*Hello there!*

### 2. Methods function adds a matcher for HTTP methods

Send http.StatusOK for the HEAD method, other methods are not allowed.

`go run gorillamux/2-methods.go`

`curl -I localhost:8080/hello`

*HTTP/1.1 200 OK*

*Date: Mon, 22 Apr 2024 13:56:03 GMT*

### 3. Send a name as a query string

`go run gorillamux/3-query-params.go`

`curl localhost:8080/hello?name=John%20Doe`

*Hello John Doe!*

### 4. Process a name variable from the URL path

`go run gorillamux/4-path-vars.go`

`curl localhost:8080/hello/John%20Doe`

*Hello John Doe!*

### 5. Determine the current datetime and send it as JSON response

`go run gorillamux/5-json-response.go`

`curl localhost:8080/now`

*{"now":"Mon Apr 22 14:01:01 2024"}*

### 6. Create two subroutes to organize handlers into logical groups

`go run gorillamux/6-subroutes.go`

`curl localhost:8080/path1/`

*Subroute 1*

`curl localhost:8080/path2/`

*Subroute 2*

### 7. Serve static files from a subdirectory for the app URL path

`go run gorillamux/7-static-files.go`

`curl localhost:8080/app/`

*This text in "gorillamux/7-static-files/index.html" without HTML-tags.*
