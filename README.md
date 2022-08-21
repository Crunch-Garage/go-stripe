<h1 align="center">
  <img alt="cgapp logo" src="https://raw.githubusercontent.com/CossyCossy/food-delivery/master/html/img/cgapp_logo%402x.png" width="224px"/><br/>
  Go Stripe Payment
</h1>
<p align="center">Create an implementation of a payment gateway with <b>Stripe Api</b> and <b>Golang</b>.</p>

<p align="center"><a href="https://github.com/Crunch-Garage/go-stripe" 
target="_blank"><img src="https://img.shields.io/badge/Go-1.19+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;
<a href="https://github.com/Crunch-Garage/go-stripe" 
target="_blank"><img src="https://img.shields.io/badge/Stripe%20-Stripe Go-red?style=for-the-badge&logo=stripe&logoColor=green" alt="stripe-go version" />
 </p>

 <br />

## 🧐 Getting Started 
> Before getting started , let us learn how to install [Golang](https://go.dev/) and investigate what [Stripe](https://stripe.com/docs) is.
>
> As [Edwin Kimani](https://github.com/edugeezu) asks, What is ?, Why ? Golang or Stripe ? Well let's learn why Golang or Stripe

Clone the repository:
```
$ git clone https://github.com/Crunch-Garage/go-stripe.git
```

Get into project directory:
```
$ cd go-stripe/
```

Install the project dependancies:
``` 
$ go mod tidy
```

## 📒 Stripe Authentication

Before running the project, you need to get your stripe Api key. The Stripe Api uses [API keys](https://stripe.com/docs/keys) to authenticate requests. You can view and manage your API keys in [the Stripe Dashboard.](https://stripe.com/login?redirect=/account/apikeys) 

Global API key example:
```
stripe.Key = "sk_test_51LYSmlLmC3rksDIqGzKxI4eqASwBZi4vZnKzwXIUtNXI0NGtBV2tWSyncdTlDZzdyaaaIBpIuAq64WZae9z6RvAA001TmltJWT"
```

Run project:
```
$ go run main.go
```

## 👨‍💻 API requests

Create Customer API:
```
http://localhost:8080/api/customer/create
```
Method POST

Request Body:
Passing parameters in the the request body is optional 

```
{
  "name" : "Jane Doe",
  "phone": "+254712345678",
  "email": "janedoe@example.com",
  "description":"Jane Doe"
}

```

Response Body:
```
{
  "address": {
    "city": "",
    "country": "",
    "line1": "",
    "line2": "",
    "postal_code": "",
    "state": ""
  },
  "balance": 0,
  "created": 1661086328,
  "currency": "",
  "default_source": null,
  "deleted": false,
  "delinquent": false,
  "description": "Jane Doe",
  "discount": null,
  "email": "janedoe@example.com",
  "id": "cus_MHnDA03YgPP9uS",
  "invoice_prefix": "807DC757",
  "invoice_settings": {
    "custom_fields": null,
    "default_payment_method": null,
    "footer": ""
  },
  "livemode": false,
  "metadata": {},
  "name": "Jane Doe",
  "next_invoice_sequence": 1,
  "phone": "+254712345678",
  "preferred_locales": [],
  "shipping": null,
  "sources": {
    "has_more": false,
    "total_count": 0,
    "url": "/v1/customers/cus_MHnDA03YgPP9uS/sources",
    "data": []
  },
  "subscriptions": {
    "has_more": false,
    "total_count": 0,
    "url": "/v1/customers/cus_MHnDA03YgPP9uS/subscriptions",
    "data": []
  },
  "tax_exempt": "none",
  "tax_ids": {
    "has_more": false,
    "total_count": 0,
    "url": "/v1/customers/cus_MHnDA03YgPP9uS/tax_ids",
    "data": []
  }
}

```
Create Charge API:
```
http://localhost:8080/api/charge/create
```
Method POST

Request Body:
For development or testing mode 'tok_visa' is used in place of Token

```
{
  "amount": 750,
  "description": "Ride hailing",
  "receiptEmail": "johndoe2@example.com",
  "customer": "cus_MH4aPLRtAMUd6o",
  "chargeToken": "tok_visa"
}

```

Response Body:
```
{
  "amount": 750,
  "amount_refunded": 0,
  "application": null,
  "application_fee": null,
  "application_fee_amount": 0,
  "authorization_code": "",
  "balance_transaction": {
    "amount": 0,
    "available_on": 0,
    "created": 0,
    "currency": "",
    "description": "",
    "exchange_rate": 0,
    "id": "txn_3LZECaLmC3rksDIq13QFxLby",
    "fee": 0,
    "fee_details": null,
    "net": 0,
    "recipient": "",
    "reporting_category": "",
    "source": null,
    "status": "",
    "type": ""
  },
  "billing_details": {
    "address": {
      "city": "",
      "country": "",
      "line1": "",
      "line2": "",
      "postal_code": "",
      "state": ""
    },
    "email": "",
    "name": "",
    "phone": ""
  },
  "calculated_statement_descriptor": "Stripe",
  "captured": true,
  "created": 1661088512,
  "currency": "usd",
  "customer": null,
  "description": "Ride hailing",
  "destination": null,
  "dispute": null,
  "disputed": false,
  "failure_code": "",
  "failure_message": "",
  "fraud_details": {
    "user_report": "",
    "stripe_report": ""
  },
  "id": "ch_3LZECaLmC3rksDIq1mUn39dU",
  "invoice": null,
  "level3": {
    "customer_reference": "",
    "line_items": null,
    "merchant_reference": "",
    "shipping_address_zip": "",
    "shipping_from_zip": "",
    "shipping_amount": 0
  },
  "livemode": false,
  "metadata": {},
  "on_behalf_of": null,
  "outcome": {
    "network_status": "approved_by_network",
    "reason": "",
    "risk_level": "normal",
    "risk_score": 61,
    "rule": null,
    "seller_message": "Payment complete.",
    "type": "authorized"
  },
  "paid": true,
  "payment_intent": "",
  "payment_method": "card_1LZECaLmC3rksDIqBrC3dn6J",
  "payment_method_details": {
    "ach_credit_transfer": null,
    "ach_debit": null,
    "alipay": null,
    "bancontact": null,
    "au_becs_debit": null,
    "bitcoin": null,
    "card": {
      "brand": "visa",
      "checks": {
        "address_line1_check": "",
        "address_postal_code_check": "",
        "cvc_check": ""
      },
      "country": "US",
      "exp_month": 8,
      "exp_year": 2023,
      "fingerprint": "kr1bwpyjWoPLL0Pu",
      "funding": "credit",
      "installments": null,
      "last4": "4242",
      "network": "visa",
      "moto": false,
      "three_d_secure": null,
      "wallet": null,
      "description": "",
      "iin": "",
      "issuer": ""
    },
    "card_present": null,
    "eps": null,
    "fpx": null,
    "giropay": null,
    "ideal": null,
    "klarna": null,
    "multibanco": null,
    "p24": null,
    "sepa_debit": null,
    "sofort": null,
    "stripe_account": null,
    "type": "card",
    "wechat": null
  },
  "receipt_email": "johndoe2@example.com",
  "receipt_number": "",
  "receipt_url": "https://pay.stripe.com/receipts/payment/CAcaFwoVYWNjdF8xTFlTbWxMbUMzcmtzRElxKIDmiJgGMgZyLCcYdxA6LBZ4exID55jxTu4kiXOdJHjayGDUzFlod5cEOxckSSF7KX5Rrk7MeML674NT",
  "refunded": false,
  "refunds": {
    "has_more": false,
    "total_count": 0,
    "url": "/v1/charges/ch_3LZECaLmC3rksDIq1mUn39dU/refunds",
    "data": []
  },
  "review": null,
  "shipping": null,
  "source": {
    "address_city": "",
    "address_country": "",
    "address_line1": "",
    "address_line1_check": "",
    "address_line2": "",
    "address_state": "",
    "address_zip": "",
    "address_zip_check": "",
    "available_payout_methods": null,
    "brand": "Visa",
    "cvc_check": "",
    "country": "US",
    "currency": "",
    "default_for_currency": false,
    "deleted": false,
    "description": "",
    "dynamic_last4": "",
    "exp_month": 8,
    "exp_year": 2023,
    "fingerprint": "kr1bwpyjWoPLL0Pu",
    "funding": "credit",
    "id": "card_1LZECaLmC3rksDIqBrC3dn6J",
    "iin": "",
    "issuer": "",
    "last4": "4242",
    "metadata": {},
    "name": "",
    "recipient": null,
    "three_d_secure": null,
    "tokenization_method": "",
    "customer": null,
    "object": "card"
  },
  "source_transfer": null,
  "statement_descriptor": "",
  "statement_descriptor_suffix": "",
  "status": "succeeded",
  "transfer": null,
  "transfer_data": null,
  "transfer_group": ""
}

```
Create Payment Method API:
```
http://localhost:8080/api/paymentmethod/create
```
Method POST

Request Body: 

```
{
  "number":"4242424242424242",
  "expMonth": "8",
  "expYear": "2023",
  "cVC": "314",
  "type": "card"
}

```

Response Body:

```
{
  "au_becs_debit": null,
  "billing_details": {
    "address": {
      "city": "",
      "country": "",
      "line1": "",
      "line2": "",
      "postal_code": "",
      "state": ""
    },
    "email": "",
    "name": "",
    "phone": ""
  },
  "card": {
    "brand": "visa",
    "checks": {
      "address_line1_check": "",
      "address_postal_code_check": "",
      "cvc_check": "unchecked"
    },
    "country": "US",
    "exp_month": 8,
    "exp_year": 2023,
    "fingerprint": "kr1bwpyjWoPLL0Pu",
    "funding": "credit",
    "last4": "4242",
    "three_d_secure_usage": {
      "supported": true
    },
    "wallet": null,
    "description": "",
    "iin": "",
    "issuer": ""
  },
  "card_present": null,
  "created": 1661086317,
  "customer": null,
  "fpx": null,
  "id": "pm_1LZDdBLmC3rksDIqHNuONSHL",
  "ideal": null,
  "livemode": false,
  "metadata": {},
  "object": "payment_method",
  "sepa_debit": null,
  "type": "card"
}