# Starbucks Online Ordering Service Components

## Web Frontend

The Web Frontend serves the user interface which present services available to the user. Web Frontend inspects user session to present only services relevant to the current user's role.

Roles include:
* Anonymous User
* Starbucks Customer
* Starbucks Employee

### Anonymous User UI

The Starbucks Anonymous User UI presents Company Branding and encourages the customer to register for an account with which they can begin making online Starbucks drink orders and save their favorite drinks for future online orders.

### Customer UI

The Starbucks Customer UI supplements the interface present to an Anonymous User to expose additional customer-centric services such as placing an online Starbucks drink order and saving the customer's drink order for future orders.

### Employee UI

The Starbucks Employee UI provides an employee at a Starbucks location access to view and fulfill online Starbucks drink orders.

## Order Backend

The Order Backend provides an API which allows Starbucks customers to place orders and Starbucks employees to view and process online orders placed by customers.
Below three functional modularization enable functional splitting.

### SIGN-IN/UP/OUT

The Three user authentication functionalities can be modularized using ../SIGN/ route for identification of the incoming request. The functionalities will include some mechanism to identify logged-in/out users on distributed servers.

### Order Placement

The Signed-In users will be able to post orders(POST),get orders(GET),modify order(PUT) and delete order(DELETE). The functionalities related to order placemets can be modularized using ../ORDER/ route for identification of the incoming requests.

### Order Processing

A user on sucessfull payement, will be able to place orders. The processing of orders will be done in the module routed by ../ORDERS.
TBD (auto or using EMPLOYEE role)

## Load Balancing

Each backend module is replicated and load balanced by a load balancer to evenly distribute load. Each load balancer is access points for Frontends.

