## CMPE 281 Team Project

### Deploy Frontend

Frontend Node.js app is deployed to Heroku (Bo's account) using Git subtree command.

App URL: https://rocky-island-94191.herokuapp.com/


```shell
# Push changes to Heroku remote
# NOTE: YES, it is StarBcuks, not StarBucks
git subtree push --prefix StarBcuks heroku master

```

---
### How to do jwt authentication?
[authentication.md](https://github.com/nguyensjsu/team281-bit-flip/blob/master/apis/src/cmpe281/user/authentication.md)

### Architecture Diagram
![Architecture](images/stack-architecture-diagram.png?raw=true "Architecture Diagram")

### Links

Please use below link for editing (Please sign into your SJSU gmail id)

[Architecture Diagram](https://docs.google.com/drawings/d/1IqZc8vxy2CkHh_zAqYUndz0EAhEl5wDZS-HAGB9p8Pg/edit?usp=sharing)


[Raceleg 2 Challenge](https://docs.google.com/document/d/172zN_JmlNBy1MiGxuYDfQZS04yMvDICOmocGRcR0Vzw/edit?usp=sharing)


Load Balancer URLs:
- [Kong LB URL](http://kong-lb-133222058.us-west-1.elb.amazonaws.com/)
  - Port: `80`
  - APIs:
    - [/users/v1/](http://kong-lb-133222058.us-west-1.elb.amazonaws.com/users/v1/)
    - [/payments/v1/](http://kong-lb-133222058.us-west-1.elb.amazonaws.com/payments/v1/)
    - [/inventory/v1/](http://kong-lb-133222058.us-west-1.elb.amazonaws.com/inventory/v1/)
    - [/orders/v1/](http://kong-lb-133222058.us-west-1.elb.amazonaws.com/orders/v1/)
- [User LB URL](http://cmpe281-team-project-user-api-995132055.us-west-1.elb.amazonaws.com/)
    - Port: `80`
- [Payment LB URL](http://payments-lb-853644621.us-west-1.elb.amazonaws.com/)
    - Port: `80`
- [Inventory LB URL](
http://inventory-lb-1305987865.us-west-1.elb.amazonaws.com/)
    - port: `80`
- [Order LB URL](http://orderLB-2141712569.us-west-1.elb.amazonaws.com/)
    - port: `80`

---

### Tools

- Live reload Go web application

    gin: https://github.com/codegangsta/gin
    ```
    go get github.com/codegangsta/gin
    gin -h
    gin run main.go
    ```
