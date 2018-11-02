# Subscription Cntroller Service
A lab RESTful service to find out all notification-service's subscriptions which refer to invalid accounts. 

## Getting started
If you want just make it run:
```
dep ensure
go run main.go
```

And check the usage from the command:
```
go run main.go server run
```

If `dep: command not found` is making you sad:
```
brew install go dep
```
(And try again)

## Status
Implemented features:
* ✅ Create GET/PUT/PATCH/POST/DELETE endpoints
* ✅ Configure and enstablish db connection to Mongo
* ✅ Configure external services 
 