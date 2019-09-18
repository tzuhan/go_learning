# RESTful json API implement practices
A simple restful api with GET, PUT, UPDATE, DELETE methods to query users info
```
    type User struct {
        ID string
        Name string
        Birthday string
        Description string
        Gender int
    }
```
---
## APIs

**GET**
1. `/ids`
- get all users ids
2. `/id/:id`
- get particular user's info by id

**PUT**
1. `/regist` + json of new user's data
- add new user by id, show error if user already exists.

**UPDATE**
1. `/update` + json of user's data 
- update existing user, show error if user not exists.

**DELETE**
1. `/delete/:id`
- delete user by id, show error if user not exists.

---

## Different implementions
0. raw.go
- using only standard library
1. gm.go
- using gorilla Mux package
2. hr.go
- using httprouter package to speed up routing process


