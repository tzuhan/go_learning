1. Swagger api doc
2. dockerfile
3. deploy to gcp server
4. non serverless
5. CI
6. Version control
7. go framework => gin
8. API allow CORS origin: http://localhost:3002
9. return api role list
```
// API endpoint: /api/v1/role
// HTTP method: GET
[
{
"id": 1,
"name": "Admin"
},
{
"id": 2,
"name": "Moderator"
},
{
"id": 3,
"name": "Viewer"
}
]
// Note
// 1, Admin (最高權限者)
// 2, Moderator (審核者)
// 3, Viewer (檢視者)
```
10. Use kubernetes
11. CD
12. HTTPS
13. Load balancing