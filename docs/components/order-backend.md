# Order Backend

Starbcuks Order backend is responsible for accepting Starbcuks drink orders as placed by the customer. An order consists of both the items being ordered as well as payment details which are forwarded to the Payment service for processing for approval or decline.

## API Reference
### Create Order

Create a Starbcuks drink order.

#### POST /orders/v1/orders
##### Request Headers

| Header | Description |
|--------|-------------|
| Authorization | User Credential for Authorization Verification |

##### Request Body

```json
{
"store": "store1",
  "product":
   [ { "item": "Cappuccino", "qty": "1", "size": "small" },
     { "item": "Expresso Machiato", "qty": "2", "size": "medium" },
     { "item": "Lemon Tea", "qty": "1", "size": "large" } ],
"first": "aa",
"second": "vv",
"card": "1111111111111111",
"exp_month": "9",
"exp_yr": "2018",
"add1": "hh",
"add2": "hhh",
"city": "hh",
"state": "hh",
"pin": "11111",
"amount": "27"
}
```

| Property Name | Type | Description |
|---------------|------|-------------|
| `items[]` | [Drink\[\]](#Drink-Resource) | List of drink items for the order. |
| `payment` | [Payment](#Payment-Resource) | Payment details for the order. |

### List Orders

List all Starbcuks drink orders.

#### GET /orders/v1/orders
##### Request Headers

| Header | Description |
|--------|-------------|
| Authorization | User Credential for Authorization Verification |

##### Request Body

Do not supply a request body for this method.

##### Response Body

```json
{
  "orders": [
    order Resource
  ]
}
```

### Get Order
#### GET /orders/v1/orders/{orderid}
##### Request Headers

| Header | Description |
|--------|-------------|
| Authorization | User Credential for Authorization Verification |

Retrieve details for a Starbcuks drink order.

##### Request Body

Do not supply a request body for this method.

##### Response Body

Returns an [Order Resource](#Order-Resource)

### Update Order
#### PATCH /orders/v1/orders/{orderid}

Not Supported

### Delete Order
#### DELETE /orders/v1/orders/{orderid}

Not Supported

#### Resources
##### Order Resource
```json
{
  "items": [
    {
      "product": "string",
      "size": "string",
      "decaf": "boolean"
    }
  ],
  "payment": {
    "card": {
      "number": "string",
      "ccv": "string"
    },
    "billing": {
      "firstname": "string",
      "lastname": "string",
      "line1": "string",
      "line2": "string",
      "city": "string",
      "state": "string",
      "zipcode": "string"
    }
  }
}

```

| Property Name | Type | Description |
|---------------|------|-------------|
| `items[].product` | string | Name of the drink being ordered. |
| `items[].size` | string | Size of drink. Choose from: `short`, `tall`, `grande`, `venti`. |
| `items[].decaf` | string | Decaffinated drink. |
| `payment.card.number` | string | Credit or Debit Card Number. |
| `payment.card.ccv` | string | Credit or Debit Security Code. |
| `payment.billing.firstname` | string | Customer Billing First Name. |
| `payment.billing.lastname` | string | Customer Billing Last Name. |
| `payment.billing.line1` | string | Customer Billing Line 1. |
| `payment.billing.line2` | string | Customer Billing Line 2. |
| `payment.billing.city` | string | Customer Billing City. |
| `payment.billing.state` | string | Customer Billing State. |
| `payment.billing.zipcode` | string | Customer Billing Zip Code. |

##### Payment Resource
TODO(bbamsch): Move this to the Payment Service Design Document
```json
{
  "card": {
    "number": "string",
    "ccv": "string"
  },
  "billing": {
    "firstname": "string",
    "lastname": "string",
    "line1": "string",
    "line2": "string",
    "city": "string",
    "state": "string",
    "zipcode": "string"
  }
}
```

| Property Name | Type | Description |
|---------------|------|-------------|
| `card.number` | string | Credit or Debit Card Number. |
| `card.ccv` | string | Credit or Debit Security Code. |
| `billing.firstname` | string | Customer Billing First Name. |
| `billing.lastname` | string | Customer Billing Last Name. |
| `billing.line1` | string | Customer Billing Line 1. |
| `billing.line2` | string | Customer Billing Line 2. |
| `billing.city` | string | Customer Billing City. |
| `billing.state` | string | Customer Billing State. |
| `billing.zipcode` | string | Customer Billing Zip Code. |
