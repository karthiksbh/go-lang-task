
# ![image](https://user-images.githubusercontent.com/83284855/219548047-04913622-09f4-4098-9db8-c1958f00e007.png)
 Go Lang Task

Shows an understanding of data, data types, and ability to develop features using GoLang on Scale. The task incorporates an understanding of XML and JSO2 type data formats as well as APIs using GoFiber. 




## Installation

1. Install [Go](https://golang.org/dl/)

2. Clone this repository and change into its directory:
```bash
  git clone https://github.com/karthiksbh/go-lang-task.git
  cd go-lang-task
```
3. Install the required packages using `go get`:
```bash
  go get -u github.com/gin-gonic/gin
  go get -u github.com/karthiksbh/go-lang-task/config
  go get -u github.com/karthiksbh/go-lang-task/routes
  go get -u gorm.io/gorm
```

4. Create .env file to store environment variables to connect to PostgreSQL database.

5. Build and run the application
```bash
go build
./go-lang-task
```

This will start the application at `http://localhost:8081`
    
## Structure
```
├── config
│   ├── db.go  DB Configuration
├── controller 
│   ├── AddressController.go //APIs to create Address
|   ├── PropertyVals.go //APIs to create Property values
├── data
|    ├── db  // Contains the database files
├── models
|    ├── MainModels.go //Contains all the models used
├── routes
|    ├── procadd.go //Contains all the routes
├── main.go 
```


## APIs

#### /procadd
* `GET` : Get all addresses with pagination
* `POST` : Create a new address

#### /procadd/:id
* `GET` : Get an address for the given identifier
* `PUT` : Update an address for the given identifier
* `DELETE` : Delete an address for the given identifier

#### /filter/
* `POST` : Filter the address based on typecode, name and city


## Documentation

* `HOSTED CODE`: https://golang-task.onrender.com
* `DATABASE SCHEMA` : https://dbdiagram.io/d/63edc654296d97641d815bbd
* `POSTMAN` : https://documenter.getpostman.com/view/18159368/2s93CGRvS4

## Tasks

### Part 1:
* `DATA` : Downloaded the XML data and converted the data into JSON by creating schema (both the files in the directory)

### Part 2:
* `SCHEMA` : Created a database schema with tables, relations, attributes, primary keys etc. 
* `SERVICES` : Built the basic CRUD operations for the tables created.

### Part 3: 
* `REQUEST VALIDATION`
* `DOCUMENTATED APIs` : Documented the APIs using Postman

### Part 4: 
* `DEPLOY DATABASE` : Deployed the database on ElephantSQL
* `DEPLOY CODE` : Code deployed on render. 





## Result

- [x] Support basic REST APIs.
- [x] Proper Database Schema
- [x] Deployed Code and Database
- [x] Dockerized Code
- [x] Proper Postman Documentation 
