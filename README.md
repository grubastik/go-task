# Create HTTP Rest API.

It should includes the following steps:
1. Use echo for web handler
2. Implement login endpoint with JWT token, and simple middleware that checks header for 'Authorization: Bearer %jwt_token%' in each request. Otherwise return 403 and json struct with error
3. Implement endpoint that will use OAuth2 authorization for FB, to login and issue access_token
4. Log each request including status code (logrus or https://github.com/uber-go/zap)
5. Implement persistence with MySQL and Gorm (https://github.com/jinzhu/gorm or sqlx)
6. Use tool of your choice for DB migrations
7. Implement save endpoint for Task object
8. Implement update endpoint for Task object
9. Implement get endpoint for Task object
10. Implement delete endpoint for Task object (just update IsDeleted field)
11. Use CORS (reply with header Access-Control-Allow-Origin: *)
12. Add support for OPTION HTTP method for each endpoints
13. Configure daemon over simple YAML config. Specify path as process flag for daemon. Required params: ListenAddress, DatabaseUri.
14. Implement 3-rd party libs vendoring with tool of your choice. (glide)

Put in comments below description of taken architectural decisions

Here is a schema of the document:
```
    type Task struct {
        Id          int64
        Title       string
        Description string
        Priority    int
        CreatedAt   *time.Time
        UpdatedAt   *time.Time
        CompletedAt *time.Time
        IsDeleted   bool
        IsCompleted bool
    }
```
