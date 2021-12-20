# Url Shortener Api

Using Go

## Endpoints

- Get and redirect to long url

    `GET: "/:id"`

- Create new shorten url
    
    `POST: "/"`
  
    Request body:
    ```json
    {
        "long_url": "https://facebook.com"
    }
    ```
    Response:
    ```json
    {
        "short_url": "http://localhost:8080/hRPHui5eqQsqTxwb6ZLQ9M",
        "status": "success"
    }
    ```

## Tech

*|Tech
-----|-----
Web|github.com/gin-gonic/gin
Database|Sqlite
ORM|gorm.io/gorm
Id Generator|github.com/lithammer/shortuuid/v3


## Architecture

Using layered architecture, introduced by Uncle Bob

`controller -> service -> repository`
