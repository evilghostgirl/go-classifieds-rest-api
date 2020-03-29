# Simple CRUD server written in Go!

This is little project what serves data (classifieds, localizations, users info) for my android app - Locally. 
## GET Endpoints:
* localhost:3000/classifieds/
* localhost:3000/classifieds/{item}
* localhost:3000/classifieds?title={title}
* localhost:3000/classifieds?localizationid={localizationid}
* localhost:3000/classifieds/{id}
* localhost:3000/categories
* localhost:3000/categories/{id}
* localhost:3000/users
* localhost:3000/users/{id}
* localhost:3000/users?localizationid={localizationid}

## Setup
``` sh 
make postgres
make adminer
make migrate
go run main.go
```
## Database
http://localhost:8080/
* Credentials:
user: postgres, 
password: secret

## In the future:
### Adding GET endpoints
* localhost:3000/localizations
* localhost:3000/localizations?province={province}

### Adding another CRUD operations on api

### Restructuring project
* Better structure of directories
* Refactoring to separate files

### Adding auth
* Authorized users could add new data
* Owner can edit his data, but nobody else
* New roles - basic user, admin

### Tests
