# Product notes

The service is a tiny order-ingestion API used by an internal checkout flow.

## Users and jobs

- Checkout services submit orders after payment authorization.
- Operations can list recently accepted orders for diagnostics.

## Current contract

`POST /orders`

```json
{
  "customer": "acme",
  "amount": 12500
}
```

Successful requests return `201 Created` with an order containing `id`, `customer`, and `amount`.

`GET /orders` returns accepted orders for the current process lifetime.

## Known operational problem

Checkout retries requests when a network timeout makes the outcome ambiguous. The current endpoint creates a duplicate order on every retry. Duplicate orders are expensive downstream because they can trigger duplicate fulfillment.

Do not solve this here. Product changes are tracked as GitHub issues and must preserve the existing response fields for existing consumers.
