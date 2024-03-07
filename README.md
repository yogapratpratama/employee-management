
# Project EMPLOYEE MANAGEMENT API (With JWT Authentication)

The Project Test API is contains some API for employee management. This Project use stacks such HTTP Rest API and JWT Authentication also for access main API.


## Tech Stack

**Server:** Golang, JWT Authentication, gin gonic


## Run Locally

#### 1. Local environment
- Make sure postgresql has installed, else you can download https://www.postgresql.org/download/
- Create new connection on db, you can use pgAdmin or access bin folder postgresql. more on https://www.enterprisedb.com/postgres-tutorials/connecting-postgresql-using-psql-and-pgadmin
- Create new folder for save source repository
- Clone the project
```bash
  git clone https://github.com/yogapratpratama/employee-management.git
```
- Go to the project directory
```bash
  cd EmployeeManagementApp
```
- Edit location config development on file .env to source folder
- Run server, as the default environment will be "local"
```bash
  go run ./app/main.go local
```
- Server will be running on http://localhost:9078

## API Reference

#### 1. Register User (User)

```http
  POST /agit/oauth/register
```
##### Request Body JSON (Example)
```json
{
  "username": "samuel15",
  "password": "Samuel123@"
}
```
##### Response Body JSON (Example)
```json
{
  "request_id": "d26648dc-ff21-41a3-8813-df5dabb5ebc8",
  "status": true,
  "message": "Success Register!",
  "data": null
}
```

#### 2. Login User
```http
  POST /agit/oauth/login
```
##### Request Body JSON (Example)
```json
{
  "username": "samuel15",
  "password": "Samuel123@"
}
```
##### Response Body JSON (Example)
```json
{
  "request_id": "44a3d445-be0b-40a0-a5dd-6bbc00aca33c",
  "status": true,
  "message": "Token has generated!",
  "data": null
}
```

#### 3. Logout User
```http
  POST /agit/oauth/logout
```
##### Response Body JSON (Example)
```json
{
  "request_id": "405eb1b0-fb20-4055-9a92-fbbf042b320d",
  "status": true,
  "message": "Logout Success!!",
  "data": null
}
```

#### 4. Add Employee

```http
  POST /agit/employee
```
##### Request Body JSON (Example)
```json
{
  "name": "komeng",
  "nip": "123456766",
  "birthplace": "jakarta",
  "birthdate": "1987-11-24T00:00:00Z",
  "age": 43,
  "address": "jakarta timur",
  "religion": "islam",
  "gender": "laki-laki",
  "phone_number": "082134567898",
  "email": "komeng@mail.com"
}
```
##### Response Body JSON (Example)
```json
{
  "request_id": "37211ed5-7778-4565-b40c-43f8671f9188",
  "status": true,
  "message": "Success Store Data!",
  "data": null
}
```

#### 5. Get ID Employee

```http
  GET /agit/employee/3
```
##### Response Body JSON (Example)
```json
{
  "request_id": "f409d72f-0a9d-46bb-b106-e44bdeebdc34",
  "status": true,
  "message": "Success Get Detail!",
  "data": {
    "id": 3,
    "name": "komeng",
    "nip": "123456766",
    "email": "komeng@mail.com",
    "phone_number": "082134567898",
    "birthplace": "jakarta",
    "birthdate": "1987-11-24T00:00:00Z",
    "age": 43,
    "address": "jakarta timur",
    "religion": "islam",
    "gender": "laki-laki",
    "created_at": "2024-03-08T00:24:53.323714Z",
    "updated_at": "2024-03-08T00:24:53.323714Z"
  }
}
```
# ---Next you can review the postman collection