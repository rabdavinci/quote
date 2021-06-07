# Quotes REST API
This repo contains solutions of "Quotes" REST API on Golang 
## How 2RUN
1. The first need Go installed , then use the below Go command to install Gorilla mux
```
go get -v -u github.com/gorilla/mux
```
2. Clone project 
```
git clone git@github.com:rabdavinci/quote.git .
```
3. Run main.go and visit 0.0.0.0:10000/ (for windows "localhost:10000/") on browser
```
$ go run main.go quote.go worker.go
```
4. Use REST API Methods

4.1. Get all quotes
```
GET http://localhost:10000/
```
4.2. Create quote
```
POST http://localhost:10000/quote
BODY {"Author":"new author","Quote":"new quote","Category":"new category"}
```
4.3. Update quote
```
PUT http://localhost:10000/quote/id
BODY {"Author":"updated author","Quote":"updated quote","Category":"updated category"}
```
4.4. Delete quote
```
DELETE http://localhost:10000/quote/id
```
4.5. Get all quotes by category

```
GET http://localhost:10000/category/id
```
4.6. Get quote by id
```
GET http://localhost:10000/quote/id
```
4.7. Get random quote
```
GET http://localhost:10000/random-quote
```

## GarbageWorker
Added a worker that wakes up every 5 minutes and deletes quotes that were created more than 1 hour ago.

## HOW 2TEST
```
go test
```

## TODO
1. Use RestAPI design with  http status codes for NotFound, Duplicate, InternalError, etc.,
2. Add errors handle.
