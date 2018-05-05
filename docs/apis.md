# Starbucks Online Ordering Service API Doc

* Status: Draft
* Last Updated: 2018/04/07

### User Information / Authentication
| URL        | Method | Description | Response |
|:-----------|:------ | :---------- | :----- |
| /signin    | POST   | Verify username and password | {} |
| /signup    | POST   | Verify username and password | {} |
| /signout   | POST   | Log out user | {} |
| /user/:user_id/{?} | GET   | Get user information | {} |

### Menu / Product Information
| URL        | Method | Description | Response |
|:-----------|:------ | :---------- | :----- |
| /menu      | GET   | Get menu | [] |
| /product/:id| GET   | Get product information by id | {} |
| /product | POST  | Add a new product | {} |
| /product/:id| PUT| Modify a product by id | {} |
| /product/:id| DELETE| Delete a product by id | {} |

### Order Management
| URL        | Method | Description | Response |
|:-----------|:------ | :---------- | :----- |
| /order     | POST   | Place an order | "" |
| /order     | GET   | Get all orders for a user | "" |
| /order     | DELETE   | Delete order for a user | "" |
