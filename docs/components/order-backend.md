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
"Authenticated"  -Success for valid Token<br>
"Bad Authentication" -Failure for invalid Token

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
"Bad Dependency-Inventory"/"Bad Dependency-Payment" for internal server error for dependencies<br>
"Bad Card Details",StatusBadRequest if Payment declines<br>
"Failed to create order" for internal server error while order creation<br>
"Inventory Lacking", StatusBadRequest incase Inventory doesnt confirm the order request<br>


| Property Name | Type | Description |
|---------------|------|-------------|
| `store` | ("San Jose", "Mountain View") as string | The store identifier for location of purchase |
| `product` | [Order](#Order Resource) | List of drink items, quantity and size for the order. |
| `payment` | [Payment](#Payment Resource) | Payment details for the order. |

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
"Unable to Authenticate" for invalid token request.<br>
"Bad Authentication" for  unauthorized request.<br>

### Update Order
#### PATCH /orders/v1/order

Not Supported, internally handled by Go routine

### Delete Order
#### DELETE /orders/v1/order

##### Request Body

{ "pid": "e5440db4-468d-11e8-84b1-204747ddadd5"}

As the userId will be decoded from Auth token.

##### Response<br>
"Deleted", StatusAccepted on Success<br>
"Something went bad with payload", StatusBadRequest for bad pid data<br>
err/msg data in case of internalserver error<br>
"Unable to Authenticate" for invalid token request.<br>
"Bad Authentication" for  unauthorized request.
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

| Property Name | Type | Description |
|---------------|------|-------------|
| `store` | string | Identifier of the store location of purchase. |
| `product[].item` | string | Name of the drink being ordered. |
| `product[].size` | string | Size of drink. Choose from: `short`, `tall`, `grande`, `venti`. |
| `product[].qty` | string | Amount ordered. |
| `payment.card.number` | string | Credit or Debit Card Number 11digit. |
| `payment.card.ccv` | string | Credit or Debit Security Code 3 digit. |
| `payment.card.exp_month` | string | Credit or Debit expiry month 1-12. |
| `payment.card.exp_year` | string | Credit or Debit expiry year max 20+ yr. |
| `payment.billing.firstname` | string | Customer Billing First Name. |
| `payment.billing.lastname` | string | Customer Billing Last Name. |
| `payment.billing.line1` | string | Customer Billing Line 1. |
| `payment.billing.line2` | string | Customer Billing Line 2. |
| `payment.billing.city` | string | Customer Billing City. |
| `payment.billing.state` | string | Customer Billing State. |
| `payment.billing.zipcode` | string | Customer Billing Zip Code. |
| `payment.amount` | string | Customer Billing Amount. |

##### Payment Resource
##### Request Body

```json
{
  "amount": "10",
  "billing_details": {
    "first_name": "John",
    "last_name": "Doe",
    "line1": "One Washington Square",
    "line2": "",
    "city": "San Jose",
    "state": "CA",
    "zip_code": "95192"
  },
  "card_details": {
    "number": "4111111111111111",
    "cvv": "111",
    "exp_month": "04",
    "exp_year": "2018"
  }
}
```

| Property Name | Type | Description |
|---------------|------|-------------|
| `amount` | float | Payment Amount. |
| `billing_details.first_name` | string | Customer Billing First Name. |
| `billing_details.last_name` | string | Customer Billing Last Name. |
| `billing_details.line1` | string | Customer Billing Address Line 1. |
| `billing_details.line2` | string | Customer Billing Address Line 2. |
| `billing_details.city` | string | Customer Billing City. |
| `billing_details.state` | string | Customer Billing State. |
| `billing_details.zip_code` | string | Customer Billing Zip Code. |
| `card_details.number` | string | Customer Payment Card Number. |
| `card_details.cvv` | string | Customer Payment Card CVV -- Not Stored. |
| `card_details.exp_month` | string | Customer Payment Card Expiration Month. |
| `card_details.exp_year` | string | Customer Payment Card Expiration Year. |

##### Response Body

```json
{
  "amount": 10.00,
  "billing_details": {
    "first_name": "John",
    "last_name": "Doe",
    "line1": "One Washington Square",
    "line2": "",
    "city": "San Jose",
    "state": "CA",
    "zip_code": "95192"
  },
  "card_details": {
    "number": "4111111111111111",
    "cvv": "111",
    "exp_month": "04",
    "exp_year": "2018"
  },
  "payment_id": "52a1972c-4451-11e8-842f-0ed5f89f718b",
  "status": "Approved",
}
```

| Property Name | Type | Description |
|---------------|------|-------------|
| `amount` | string | Payment Amount. |
| `billing_details.first_name` | string | Customer Billing First Name. |
| `billing_details.last_name` | string | Customer Billing Last Name. |
| `billing_details.line1` | string | Customer Billing Address Line 1. |
| `billing_details.line2` | string | Customer Billing Address Line 2. |
| `billing_details.city` | string | Customer Billing City. |
| `billing_details.state` | string | Customer Billing State. |
| `billing_details.zip_code` | string | Customer Billing Zip Code. |
| `card_details.number` | string | Customer Payment Card Number. |
| `card_details.cvv` | string | Customer Payment Card CVV -- Not Stored. |
| `card_details.exp_month` | string | Customer Payment Card Expiration Month. |
| `card_details.exp_year` | string | Customer Payment Card Expiration Year. |
| `payment_id` | string | Payment Identifier. |
| `status` | string | Payment Status. |
