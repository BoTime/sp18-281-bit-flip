# Starbcuks Online Ordering Service Design Doc

* Status: Final
* Last Updated: 2018/05/04

## Goals

* Provide Starbcuks customers with convenient way to place Starbcuks drink orders online before arriving at a Starbcuks location to increase overall customer satisfaction.
* Allow Starbcuks employees on location to access orders placed by customers online ahead of time so that a customer's drink order will be prepared and ready for pickup up before customer arrives in store.

## Non-Goals

* Online ordering is not intended to replace in store Starbcuks experience via delivery or other means. Online ordering is intended to enhance the customer experience by reducing overall customer wait time when making a purchase from a Starbcuks location.

## Overview

The Starbcuks online ordering system is composed of 4 microservices which interact to process an order on behalf of a customer.

The components are as follows:
 - Web Frontend
 - User API
 - Orders API
 - Payment API
 - Inventory API

## Detailed Design

### Web Frontend

Handles serving the user interface that the user sees in their web browser and handles some additional checks on inputs and token validity. Performs redirects for unauthenticated users and form data validation before service backends receive a request as a convenience for the user. This allows the system to actively provide feedback to the customer in cases where invalid data is entered without paying the penalty of an extra network call.

Note: Client side validations are only supplemental to those validations enforced on the server side -- they are not a replacement for server side validation.

### User API

Handles Registration, Login, Logout and JWT token generation on behalf of the entire online ordering system. When a user first interacts with Starbcuks web interface, they first have to create an account so that they can place orders. Registering with the system assigns the user a unique UUID v4 identifier which is used to key data in all service backends.

### Orders API

Handles Placement and Processing of online orders for StarBcuks stores. The Orders API backend interfaces with both the Inventory API and Payment API to orchestrate an order for processing. When an order is executed, inventory must first be allocated to the order and, if the inventory allocates successfully, payment must be taken for the order -- otherwise the company will go bankrupt.

### Payment API

Handles processing of Payments on behalf of the user for a drink order placed with StarBcuks. Provides a uniform interface that allows the Order System to accept a variety of payment forms without having to deal with the routing of payments to their respective processors. Performs card number and expiration validation on behalf of the Orders API.

### Inventory API

Handles management of Store details such as the products offered and inventory available at a given location. Inventory API is a high throughput API as it will be required during all order processing. As such, the database backing Inventory API is sharded to better handle an immense workload. Sharding is performed by Store ID which allows a store to remain consistent in its own inventory totals while not hindering overall ability to scale to increased workload. When an order is executed, an inventory allocation is made which removes inventory from the total store inventory and holds the products in a separate queue for further processing.

## Security Considerations

We take security seriously, as such we enforce authentication at the lowest level in the API stack -- at the service backends themselves. Requests authenticate via JWT tokens which allows for passing of critical identifiers such as the identifier of the user without additional network activity (this helps performance).

JWT tokens include a signature which allows the service backends to verify that the user identifier has not been modified by a malicious actor. Without access to the signing key, spoofing the identity of a user becomes a near impossibility without intercepting the token itself.

## Scalability Considerations

To scale effectively to an increasing customer base, the storage systems backing the Starbcuks online ordering system are configured in AP Mode such that the system is able to function during split network connectivity.

## Testing Plan

Individual API components are implemented against API Reference documentation to ensure agreement in the interfaces between the systems. Testing is performed at the API level as well as system level tests during Quality Assurance.

Our team uses a combination Postman collections, CURL requests, and manual testing to ensure the quality of the system both in its components and as a whole.

## Document History

| Date | Author | Description |
|:----:|:------:|-------------|
| 2018/04/04 | bbamsch | System Design Doc Template & System Goals + Non-Goals |
| 2018/05/04 | bbamsch | Finalize System Design Doc |
