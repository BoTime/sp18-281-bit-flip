# Order Backend

Starbcuks Order backend is responsible for accepting Starbcuks drink orders as placed by the customer. An order consists of both the items being ordered as well as payment details which are forwarded to the Payment service for processing for approval or decline.

## API Reference

### Ping without Auth token

#### GET /orders/v1/
##### Request Headers

No header

##### Request Body

No body needed

##### Response Body

"API ORDER ALIVE!"  -Success

### Ping with Auth token

#### GET /orders/v1/ping
##### Request Headers

| Header | Description |
|--------|-------------|
| Authorization | User Credential for Authorization Verification |

##### Request Body

No body needed

##### Response Body
"Authenticated"  -StatusOK- Success for valid Token<br>
"Bad Authentication" -StatusUnauthorized - Failure for invalid Token

### Create Order

Create a Starbcuks drink order.

#### POST /orders/v1/order
##### Request Headers

| Header | Description |
|--------|-------------|
| Authorization | User Credential for Authorization Verification |

##### Request Body

```json
{ "store": "San Jose",
  "product":
   [ 
     { "item": "Cafe Americano", "qty": "2", "size": "medium" },
     { "item": "English Spice Tea", "qty": "3", "size": "large" } ],
   "payment":
    { "billing_details":
	   { "first_name": "vimmi",
		 "last_name": "swami",
		 "line1": "sanjose",
		 "line2": "vbnm",
		 "city": "vb n",
		 "state": "vb nm",
		 "zip_code": "52366" },
	  "card_details":
	   { "number": "1111111111111111",
		 "exp_month": "08",
		 "exp_year": "2018",
		 "cvv": "111" },
	  "amount": "49" 
	}
} 
```
##### Response Body
"created", StatusCreated on sucsess<br>
"Something went bad with payload", internal server error for bad payload<br>
"Bad Dependency-Inventory"/"Bad Dependency-Payment", internal server error for dependencies<br>
"Bad Card Details", internal server error if Payment declines<br>
"Error creating session", internal server error while unsuccessful session creation<br>
"Failed to create order", internal server error while order creation<br>
"Inventory Lacking", internal server error incase Inventory doesnt confirm the order request<br>


| Property Name | Type | Description |
|---------------|------|-------------|
| `store` | ("San Jose", "Mountain View") as string | The store identifier for location of purchase |
| `product` | [Order](https://github.com/nguyensjsu/team281-bit-flip/edit/master/docs/components/order-backend.md#L164) | List of drink items, quantity and size for the order. |
| `payment` | [Payment](https://github.com/nguyensjsu/team281-bit-flip/blob/master/docs/components/payment-backend.md) | Payment details for the order. |

### List Orders

List all Starbcuks drink orders.

#### GET /orders/v1/order
##### Request Headers

| Header | Description |
|--------|-------------|
| Authorization | User Credential for Authorization Verification |

##### Request Body

Do not supply a request body for this method.

##### Response Body on sucess -StatusOK

```jsonString
"[
	{
		\"pay_id\":\"06da992b-485f-11e8-93bb-204747ddadd5\",
		\"status\":\"processed\",
		\"store\":\"San Jose\",
		\"product\":[
						{\"item\":\"Cappuccino\",\"qty\":\"1\",\"size\":\"small\"},
						{\"item\":\"Cappuccino\",\"qty\":\"5\",\"size\":\"small\"}
					]
		},
	{
		\"pay_id\":\"3b0b9639-485f-11e8-93bc-204747ddadd5\",
		\"status\":\"processed\",
		\"store\":\"Mountain View\",
		\"product\":[
						{\"item\":\"Cafe Americano\",\"qty\":\"1\",\"size\":\"medium\"},
						{\"item\":\"Caffe Late\",\"qty\":\"1\",\"size\":\"small\"}
					]
	}
]"
```
err/msg data in case of internalserver error.<br>
"Unable to Authenticate", StatusUnauthorized for invalid token request.<br>
"Bad Authentication", StatusUnauthorized for  unauthorized request.<br>
"Error creating session", internal server error while unsuccessful session creation<br>

### Update Order
#### PATCH /orders/v1/order

Not Supported, internally handled by Go routine

### Delete Order
#### DELETE /orders/v1/order

##### Request Headers

| Header | Description |
|--------|-------------|
| Authorization | User Credential for Authorization Verification |

##### Request Body

{ "pid": "e5440db4-468d-11e8-84b1-204747ddadd5"}

As the userId will be decoded from Auth token.

##### Response<br>
"Deleted", StatusAccepted on Success<br>
"Something went bad with payload", internal server error for bad pid data/payload<br>
err/msg data in case of internalserver error<br>
"Unable to Authenticate",StatusUnauthorized for invalid token request.<br>
"Error creating session", internal server error while unsuccessful session creation<br>
"Bad Authentication", StatusUnauthorized for  unauthorized request.
<br>
#### Resources
##### Order Resource
```json
{ "store": "San Jose",
  "product":
   [ 
     { "item": "Cafe Americano", "qty": "2", "size": "medium" },
     { "item": "English Spice Tea", "qty": "3", "size": "large" } ],
   "payment":
    { "billing_details":
	   { "first_name": "vimmi",
		 "last_name": "swami",
		 "line1": "sanjose",
		 "line2": "vbnm",
		 "city": "vb n",
		 "state": "vb nm",
		 "zip_code": "52366" },
	  "card_details":
	   { "number": "1111111111111111",
		 "exp_month": "08",
		 "exp_year": "2018",
		 "cvv": "111" },
	  "amount": "49" 
	}
} 

```
