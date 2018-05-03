## CMPE 281 Team Project

### Project Submisison Report edit link:
https://docs.google.com/document/d/1GHot9Dl56YPM3TW3IDoWyk40Agwt-fhS8CGDteBda-I/edit?usp=sharing

### Deploy Frontend

Frontend Node.js app is deployed to Heroku (Bo's account) using Git subtree command.

App URL: https://infinite-atoll-21952.herokuapp.com/


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

### Backend Auto scaling
1. Create an ec2 instance with the go app build executable file<br>
2. Modify /etc/rc.d/rc.local file in sudo mode to append below sample code(please modify based on your structure)<br>

echo "Executing user data script........."<br>

echo "Exporting Path"<br>
export PATH=$PATH:/usr/local/go/bin<br>
export GOPATH='/home/ec2-user/cmpe281-vimmis/goapi_orders'<br>
echo "Running app..."<br>
./home/ec2-user/cmpe281-vimmis/goapi_orders/src/starbcuks/starbcuks<br>

3. Stop and start the ec2 instance to check if system log shows logs base don above<br>
4. stop the ec2 instance, create an image of it.<br>
5. Create a template from above image (In netwrk interface: subnet and Security Group (make sure your app port are exposed)should match)<br>
6. Create Auto scale group based on above(version should be latest) and link it with the load balancer.<br>
7. Check load balancer, instances should registered which are made through ASG.<br>

Incase you wish to do midfications to the backend, create a new AMI of above with modifications, and create new template version to the existing. Autoscale group should be able to use the latest.

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
