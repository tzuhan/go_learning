# RESTful json API implement practices
A simple restful api with GET, PUT, UPDATE, DELETE methods to query users info
```
    type User struct {
        ID      string `json:"ID"`
        Account string `json:"Title"`
        FirstName   string    `json:"FirstName"`
        LastName    string    `json:"LastName"`
        Birthday    time.Time `json:"BirthDay"`
        Description string    `json:"Description"`
        Gender      int       `json:"Gender"`
    }
```
---
## APIs

**GET**
1. `/users`
- get all users
2. `/users/{id}`
- get particular user's data by id

**POST**
1. `/regist/{id}` + json of new user's data
- add new user by id, show error if user already exists.

**PUT**
1. `/users/{id}` + json of user's data 
- update existing user, show error if user not exists.

**DELETE**
1. `/users/{id}`
- delete user by id, show error if user not exists.

---

## Different implementions
0. raw.go
- using only standard library
1. gm.go
- using gorilla Mux package
2. hr.go
- using httprouter package to speed up routing process


