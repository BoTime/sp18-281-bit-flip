# Inventory Backend

StarBcuks Inventory backend is responsible for keeping track of StarBcuks locations and their corresponding inventory to ensure that a drink order can be fulfilled.

## API Reference
### List Stores

List StarBcuks store locations.

#### GET /inventory/v1/stores
##### Request Headers

None Required.

##### Request Body

Do not supply a request body for this method.

##### Response Body

```json
{
  "stores": [
    {
      "id": "f0b043e7-3388-4a56-b6b0-1adbf471d51f",
      "name": "San Jose"
    },
  ]
}
```

### List Store Inventory

List StarBcuks store inventory.

#### GET /inventory/v1/stores/{store_id}/inventory
##### URL Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| store_id | uuid | Store Identifier |

##### Request Headers

None Required.

##### Request Body

Do not supply a request body for this method.

##### Response Body

```json
{
  "products": [
    {
      "id": "ef2d6b8a-58f0-44b0-970c-8ae77c77eee4",
      "item": "Caramel Macchiato",
      "qty": "500",
      "size": "large"
    },
  ]
}
```

### Allocate Store Inventory

Creates an allocation of store inventory for fulfillment of a drink order. Allocations expire one minute after creation unless confirmed during order processing.

#### POST /inventory/v1/stores/{store_id}/allocations
##### URL Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| store_id | uuid | Store Identifier |

##### Request Headers

| Header | Description |
|--------|-------------|
| Authorization | User Credential for Authorization Verification |

##### Request Body

```json
{
  "products": [
    {
      "id": "ef2d6b8a-58f0-44b0-970c-8ae77c77eee4",
      "item": "Caramel Macchiato",
      "qty": "2",
      "size": "large"
    }
  ]
}
```

##### Response Body

```json
{
  "id": "b5312315-d3fd-43bf-9bc6-e6de1af1724a",
  "store_id": "f0b043e7-3388-4a56-b6b0-1adbf471d51f",
  "status": "Unconfirmed",
  "expires": "2018-05-04T02:48:52.958Z",
  "products": [
    {
      "id": "ef2d6b8a-58f0-44b0-970c-8ae77c77eee4",
      "item": "Caramel Macchiato",
      "qty": "2",
      "size": "large"
    }
  ]
}
```

### Confirm Store Inventory Allocation

Confirm store inventory allocation for order fulfillment.

#### POST /inventory/v1/stores/{store_id}/allocations/{allocation_id}
##### URL Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| store_id | uuid | Store Identifier |
| allocation_id | uuid | Allocation Identifier |

##### Request Headers

| Header | Description |
|--------|-------------|
| Authorization | User Credential for Authorization Verification |

##### Request Body

Do not send a request body for this method.

##### Response Body

```json
{
  "id": "b5312315-d3fd-43bf-9bc6-e6de1af1724a",
  "store_id": "f0b043e7-3388-4a56-b6b0-1adbf471d51f",
  "status": "Confirmed",
  "expires": "2018-05-04T02:48:52.958Z",
  "products": [
    {
      "id": "ef2d6b8a-58f0-44b0-970c-8ae77c77eee4",
      "item": "Caramel Macchiato",
      "qty": "2",
      "size": "large"
    }
  ]
}
```

### Cancel Store Inventory Allocation (Optional)

Delete store inventory allocation for order fulfillment. This is optional as allocations will expire automatically after one minute.

#### DELETE /inventory/v1/stores/{store_id}/allocations/{allocation_id}
##### URL Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| store_id | uuid | Store Identifier |
| allocation_id | uuid | Allocation Identifier |

##### Request Headers

| Header | Description |
|--------|-------------|
| Authorization | User Credential for Authorization Verification |

##### Request Body

Do not send a request body for this method.

##### Response Body

```json
{
  "id": "b5312315-d3fd-43bf-9bc6-e6de1af1724a",
  "store_id": "f0b043e7-3388-4a56-b6b0-1adbf471d51f",
  "status": "Expired",
  "expires": "2018-05-04T02:48:52.958Z",
  "products": [
    {
      "id": "ef2d6b8a-58f0-44b0-970c-8ae77c77eee4",
      "item": "Caramel Macchiato",
      "qty": "2",
      "size": "large"
    }
  ]
}
```
