<!-- GOLANG-GRAPHQL-CRUD OPERATION -->
# Project Structure

```bash

├── database
│   ├── database.go
├── graph
│   ├── model
│   │   └── model_gen.go
│   ├── generated.go
│   ├── resolver.go
│   └── schema.graphqls
|   └── schema.resolvers.go
├── go.mod
├── go.sum
├── gqlgen.yml
├── readme.md
├── server.go

# Golang GraphQL CRUD Operation

This is a simple example of how to implement a CRUD operation using GraphQL and Golang.

## How to run

1. Install the dependencies

```bash
go get -u github.com/graphql-go/graphql
go get -u github.com/graphql-go/handler
```

2. To run the server

```bash
go run server.go
```

3. Open the browser and go to the following URL

```bash
http://localhost:8080/
```



# GraphQL Queries

1. Create Employee

```graphql
mutation AddEmployee {
  createEmployee(input: { Name: "Arunprasath", IsTeamLead:true}){
    _id
    Name
  	IsTeamLead
  }
}
```

2. Get Employee

```graphql
query GetEmployee {
  employee(_id: "5e9b9b9b9b9b9b9b9b9b9b9b"){
    _id
    Name
  	IsTeamLead
  }
}
```

3. Get All Employees

```graphql
query GetAllEmployees {
  employees{
    _id
    Name
  	IsTeamLead
  }
}
```

4. Update Employee

```graphql
mutation UpdateEmployee {
  updateEmployee(_id: "5e9b9b9b9b9b9b9b9b9b9b9b", input: { Name: "Arunprasath", IsTeamLead:false}){
    _id
    Name
  	IsTeamLead
  }
}
```

5. Delete Employee

```graphql
mutation DeleteEmployee {
  deleteEmployee(_id: "5e9b9b9b9b9b9b9b9b9b9b9b"){
    _id
    Name
  	IsTeamLead
  }
}
```
