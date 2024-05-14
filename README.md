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

## http/server

### 1. Each time the URL is visited, a counter is incremented and the value is returned

`go run http/server/1-handle.go`

`curl localhost:8080/count && curl localhost:8080/count && curl localhost:8080/count`

*Counter: 1*

*Counter: 2*

*Counter: 3*

### 2. Create a simple HTTP server listening on port 8080

`go run http/server/2-handlefunc.go`

`curl localhost:8080`

*Hello, there*

### 3. Send the http.StatusOK for the /status path

`go run http/server/3-status-code.go`

`curl -I localhost:8080/status`

*HTTP/1.1 200 OK*

*Date: Tue, 23 Apr 2024 06:00:35 GMT*

### 4. We have three endpoints, for anything other than these three endpoints, we return a 404 error message

`go run http/server/4-not-found-handler.go`

`curl localhost:8080/about && curl localhost:8080 && curl localhost:8080/contact`

*about page*

*home page*

*404 - not found*

### 5. Get the User-Agent header from the Header field and return it back to the caller

`go run http/server/5-get-header-server.go`

`go run http/server/5-get-header-client.go`

*User agent: Go-http-client/1.1*

### 6. Get the URL path value with r.URL.Path[1:] and build a message with the data

`go run http/server/6-path-param.go`

`curl localhost:8080/John`

### 7. Accept a name parameter and get the parameter with r.URL.Query()["name"]

`go run http/server/7-query-param.go`

`curl localhost:8080/?name=Peter`

*Hello Peter!*

### 8. A file server and a simple hello handler

`go run http/server/8-file-server.go`

`curl localhost:8080`

*Home page. This text in "http/server/8-file-server/index.html" without HTML-tags.*

`curl localhost:8080/about.html`

*About page. This text in "http/server/8-file-server/about.html" without HTML-tags.*

### 9. Servers processes the GET and the POST requests from a client (no CSRF protection!)

`go run http/server/9-get-post.go`

See the HTML form located at "http/server/9-get-post/form.html" at the following link: [localhost:8080](http://localhost:8080)

`curl -d 'name=John' -d 'occupation=Teacher' localhost:8080`

*John is a Teacher*

### 10. Create a simple web server that sends an image to the client

`go run http/server/10-serve-image.go`

`curl -I localhost:8080/image`

*HTTP/1.1 200 OK*

*Content-Type: image/png*

*Date: Tue, 23 Apr 2024 06:35:11 GMT*

### 11. Return an HTML page with a table of users for the /users URL path

`go run http/server/11-template.go`

See the HTML page with a table of users at the following link: [localhost:8080/users](http://localhost:8080/users)

## http/client

### 1. Create a GET request to a small website and get the status code of the request

`go run http/client/1-status-code.go`

*200 OK*

*200*

### 2. Create a simple GET request and print the received data to the console

`go run http/client/2-get-request.go`

### 3. Issue a HEAD request with http.Head and prints all the data of the response header

`go run http/client/3-head-request.go`

### 4. Set a User-Agent header for its GET request

`go run http/client/4-user-agent-header.go`

### 5. Send a POST request to an online testing website

`go run http/client/5-postform.go`

## json

### 1. Transform a Go struct and a slice of structs into JSON format

`go run json/1-encoding.go`

### 2. Decode a JSON string into a Go struct and a slice of structs

`go run json/2-decoding.go`

### 3. Prettify output with MarshalIndent

`go run json/3-pretty_print.go`

### 4. Read JSON data from a file and decode it to a slice of user structs

`go run json/4-read_file.go`

### 5. Read data from a web service using the built-in Go HTTP client

`go run json/5-astros.go`

## csv

### 1. Read values from "csv/csv-files/numbers.csv" file

`go run csv/1-read.go`

### 2. ReadAll values from "csv/csv-files/users.csv" file

`go run csv/2-readall.go`

### 3. Read "csv/csv-files/users-delimiter.csv" where fields are separated with a semicolon

`go run csv/3-delimiter.go`

### 4. Write records to "csv/csv-files/users-writed.csv" file

`go run csv/4-write.go`

### 5. Append a new record to "csv/csv-files/users-writed.csv" file

`go run csv/5-append.go`

### 6. Another example to read values from "csv/csv-files/users.csv" file

`go run csv/6-read.go`

### 7. CSV-to-JSON: read "csv/csv-files/users.csv" file and convert its data to JSON

`go run csv/7-csv-json.go`

### Extra 1. Write records to "csv/csv-files/persons-writed.csv" file using "github.com/gocarina/gocsv"

`go run csv/extra/1-write-gocsv.go`

### Extra 2. Read "csv/csv-files/persons-writed.csv" using "github.com/gocarina/gocsv"

`go run csv/extra/2-read-gocsv.go`

### Extra 3. Read "csv/extra/json2csv_input.json" and convert it to "csv/csv-files/json2csv_output.csv" CSV-file using "github.com/yukithm/json2csv"

`go run csv/extra/3-json2csv.go`

## basic

### for-loop

`go run basic/for-loop.go`

## mutexes

### Example 1 of the race condition

`go run mutexes/1-race-condition.go && go run mutexes/1-race-condition.go && go run mutexes/1-race-condition.go && go run mutexes/1-race-condition.go && go run mutexes/1-race-condition.go`

`go run mutexes/1-mutex.go && go run mutexes/1-mutex.go && go run mutexes/1-mutex.go && go run mutexes/1-mutex.go && go run mutexes/1-mutex.go`

### Example 2 of the race condition

`go run mutexes/2-race-condition.go && go run mutexes/2-race-condition.go && go run mutexes/2-race-condition.go && go run mutexes/2-race-condition.go && go run mutexes/2-race-condition.go`

`go run mutexes/2-mutex.go && go run mutexes/2-mutex.go && go run mutexes/2-mutex.go && go run mutexes/2-mutex.go && go run mutexes/2-mutex.go`

### Concurrent map iteration and map write

`go run mutexes/3-concurrent-iter-write.go`

`go run mutexes/3-mutex.go`

`go run mutexes/3-rwmutex.go`

## basic

`go run basic/0-basic-program.go`

`go run basic/1-for-loop.go`

`go run basic/2-multiple-values-from-func.go`

`go run basic/3-strings-builder.go`

`go run basic/4-pointer.go`

`go run basic/5-type-conversion.go`

`go run basic/6-copy-map.go`

`go run basic/7-reverse-int-slice.go`

`go run basic/8-oop.go`

`go run basic/9-check-var-type.go`

`go run basic/10-concatenate-str.go`

`go run basic/11-func-closures.go`

`go run basic/12-implement-stack-lifo.go`

`go run basic/13-slice-chars-permutate.go`

`go run basic/14-vars-values-swap.go`

`go run basic/15-min-max.go`

`go run basic/16-check-slice-is-empty.go`

`go run basic/17-sum-of-squares.go`

`go run basic/18-sum-even-from-1-to-100.go`

`go run basic/19-largest-smallest-in-array.go`

`go run basic/20-prime-num.go`

`go run basic/21-factorial-of-num-recursion.go`

`go run basic/22-num-of-words.go`

`go run basic/23-bubble-sort.go`

`go run basic/24-permutations.go`

`go run basic/25-boyer-moore-search.go`

`go run basic/26-longest-subsequence-in-array.go`

`go run basic/27-binary-tree.go`

`go run basic/28-stack-array.go`

`go run basic/29-split-strings.go`

## atomic

### Example 1

`go run atomic/1-wrong-output.go && go run atomic/1-wrong-output.go && go run atomic/1-wrong-output.go && go run atomic/1-wrong-output.go && go run atomic/1-wrong-output.go`

`go run atomic/1-correct-output.go && go run atomic/1-correct-output.go && go run atomic/1-correct-output.go && go run atomic/1-correct-output.go && go run atomic/1-correct-output.go`

## channels

`go run channels/1-hello.go`

`go run channels/2-sum-of-squares.go`

`go run channels/3-merge-sort.go`

## problem-list - easy

### 1. Two sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target. You may assume that each input would have exactly one solution, and you may not use the same element twice. You can return the answer in any order.

`go run problem-list/easy/1-two-sum.go`

### 2. Palindrome number without converting the integer to a string

Given an integer x, return true if x is a palindrome, and false otherwise.

`go run problem-list/easy/2-palindrome-no-int2str.go`

### 3. Roman to integer

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

`go run problem-list/easy/3-roman-to-int.go`
