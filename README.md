cmpe 281 Team Project

Architecture Diagram :

Please use below link for editing (Please sign into your SJSU gmail id)
https://docs.google.com/drawings/d/1IqZc8vxy2CkHh_zAqYUndz0EAhEl5wDZS-HAGB9p8Pg/edit?usp=sharing

Raceleg Challenge:

https://docs.google.com/document/d/172zN_JmlNBy1MiGxuYDfQZS04yMvDICOmocGRcR0Vzw/edit?usp=sharing

Load Balancer URLs:
- User

    url: `http://cmpe281-team-project-user-api-995132055.us-west-1.elb.amazonaws.com/`

    port: `80`

- Payment


- Inventory


- Order

### Tools

- Live reload Go web application

    gin: https://github.com/codegangsta/gin
    ```
    go get github.com/codegangsta/gin
    gin -h
    gin run main.go
    ```
