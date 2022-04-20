# Url Shortener Api

Using Go

## Endpoints

- Get and redirect to long url

  `GET: "/:id"`

- Create new shorten url
  `POST: "/urls"`

  Request Body:

  ```json
  {
    "long_url": "https://facebook.com",
    "title": "Facebook"
  }
  ```

  Response Body:

  ```json
  {
    "id": "DVqU6gCRvCnsVWkPcCDT8E",
    "short_url": "http://localhost:8080/DVqU6gCRvCnsVWkPcCDT8E",
    "long_url": "https://facebook.com",
    "title": "Facebook"
  }
  ```

- Get all shorten url
  `GET: "/urls"`

  Response Body:

  ```json
  [
    {
      "id": "hQhMt48Ws4vPhVsT4DBnkf",
      "short_url": "http://localhost:8080/hQhMt48Ws4vPhVsT4DBnkf",
      "long_url": "https://manjaro.org",
      "title": "Manjaro"
    },
    {
      "id": "FsbFyFZXWaFodL4Qxkangd",
      "short_url": "http://localhost:8080/FsbFyFZXWaFodL4Qxkangd",
      "long_url": "https://go.dev/tour/concurrency/1",
      "title": "Golang Concurrency"
    },
    {
      "id": "CxQtx3mKz7QrDCYtnw3Uxk",
      "short_url": "http://localhost:8080/CxQtx3mKz7QrDCYtnw3Uxk",
      "long_url": "https://google.com",
      "title": ""
    }
  ]
  ```

- Get a shorten url
  `GET: "/urls/:id"`

  Response Body:

  ```json
  {
    "id": "CxQtx3mKz7QrDCYtnw3Uxk",
    "short_url": "http://localhost:8080/CxQtx3mKz7QrDCYtnw3Uxk",
    "long_url": "https://google.com",
    "title": ""
  }
  ```

- Delete a shorten url
  `DELETE: "/urls/:id"`

## Tech

| \*           | Tech                              |
| ------------ | --------------------------------- |
| Web          | github.com/gin-gonic/gin          |
| Database     | Sqlite                            |
| ORM          | gorm.io/gorm                      |
| Id Generator | github.com/lithammer/shortuuid/v3 |
| Testing      | github.com/stretchr/testify       |

## Architecture

Using layered architecture, introduced by Uncle Bob

`controller -> service -> repository`
