# Payment Backend

Starbucks Payment backend is responsible for accepting and processing payment for Starbucks drink orders placed by customers. Payment requests made to the Payments API typically originate from other backend services (e.g. Order API) which requires that a customer pay for the products specified in the order.

## API Reference
### Create Payment

Process a payment request.

#### POST /starbucks/v1/payments
##### Request Headers

| Header | Description |
|--------|-------------|
| Authorization | User Credential for Authorization Verification |

##### Request Body

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
    "expiration": "04/18"
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
| `card_details.expiration` | string | Customer Payment Card Expiration. |

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
    "expiration": "04/18"
  },
  "payment_id": "52a1972c-4451-11e8-842f-0ed5f89f718b",
  "status": "Approved",
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
| `card_details.expiration` | string | Customer Payment Card Expiration. |
| `payment_id` | string | Payment Identifier. |
| `status` | string | Payment Status. |

### Get Payment

Retrieve a payment by ID.

#### GET /starbucks/v1/payments/{payment_id}
##### Request Headers

| Header | Description |
|--------|-------------|
| Authorization | User Credential for Authorization Verification |

##### Request Body

Do not submit a request body.

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
    "expiration": "04/18"
  },
  "payment_id": "52a1972c-4451-11e8-842f-0ed5f89f718b",
  "status": "Approved",
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
| `card_details.expiration` | string | Customer Payment Card Expiration. |
| `payment_id` | string | Payment Identifier. |
| `status` | string | Payment Status. |

### List Payments

List all payments for current user.

#### GET /starbucks/v1/payments
##### Request Headers

| Header | Description |
|--------|-------------|
| Authorization | User Credential for Authorization Verification |

##### Request Query Parameters

| Query Parameter | Description |
|-----------------|-------------|
| `limit` | Limit number of payments returned. |
| `page_token` | Lists results after the provided page token. |

##### Request Body

Do not submit a request body.

##### Response Body

```json
{
  "payments": [
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
        "expiration": "04/18"
      },
      "payment_id": "52a1972c-4451-11e8-842f-0ed5f89f718b",
      "status": "Approved",
    },
  ]
}
```

| Property Name | Type | Description |
|---------------|------|-------------|
| `payments[].billing_details.first_name` | string | Customer Billing First Name. |
| `payments[].billing_details.last_name` | string | Customer Billing Last Name. |
| `payments[].billing_details.line1` | string | Customer Billing Address Line 1. |
| `payments[].billing_details.line2` | string | Customer Billing Address Line 2. |
| `payments[].billing_details.city` | string | Customer Billing City. |
| `payments[].payments[].billing_details.state` | string | Customer Billing State. |
| `payments[].billing_details.zip_code` | string | Customer Billing Zip Code. |
| `payments[].card_details.number` | string | Customer Payment Card Number. |
| `payments[].card_details.expiration` | string | Customer Payment Card Expiration. |
| `payments[].payment_id` | string | Payment Identifier. |
| `payments[].status` | string | Payment Status. |
