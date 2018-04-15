# Starbucks Online Ordering Service Components

## Web Frontend

The Web Frontend serves the user interface which present services available to the user. Web Frontend inspects user session to present only services relevant to the current user's role.

Roles include:
* Anonymous User
* Starbucks Customer

### Anonymous User UI

The Starbucks Anonymous User UI presents Company Branding and encourages the customer to register for an account with which they can begin making online Starbucks orders.

### Customer UI

The Starbucks Customer UI supplements the interface present to an Anonymous User to expose additional customer-centric services such as placing an online Starbucks drink order and viewing the past order history.

## App Backend Modules

###  User - Sing In/Up/Out
The user authentication functionalities can be combined as a module to distiguish between signed-in/out user . The functionalities will include some mechanism to identify logged-in/out users on distributed servers and corrsponding APIs.

### Order 

The Order placement and processing functionalities can be combined in a module. The module is responsible for authenticating requests via User module and processing orders based on inventory and payment details.

### Inventory

Inventory of amount of products available in the two stores are managed by the module. It aids in order processing based on its inventory checks. It maintains two databases of each store.

### Payment

Allow payment processing of orders placed in a module. It provides validation/invalidation of payment details provided during order processing. 

