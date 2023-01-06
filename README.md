# Product Transaction API

API ini ditujukan untuk kebutuhan backend pada Aplikasi Kasir Devan Fried Chicken

## Table of Contents

- [Setup](#setup)
- [Routes](#routes)
- [API Documentation](#api-documentation)
- [Contributor](#contributor)

## Setup

To run this project, you must have Golang and postgreSQL

- go run server.go

## Routes

| HTTP METHOD                   |      POST       |            GET             |           PUT            |           DELETE           |
| ----------------------------- | :-------------: | :------------------------: | :----------------------: | :------------------------: |
| /login                        |      Login      |             -              |            -             |             -              |
| /register                     |  Register User  |                            |            -             |             -              |
| /users                        |                 |        Detail User         |        Edit User         |        Delete User         |
| /products                     |   Add Product   |      List of Products      |            -             |             -              |
| /products/<int:id>            |                 |       Detail Product       |       Edit Product       |       Delete Product       |
| /transactions                 | Add Transaction |    List of Transaction     |            -             |                            |
| /transactions/<int:id>        |                 |     Detail Transaction     |     Edit Transaction     |     Delete Transaction     |
| /product-transaction          |        -        |  List Product Transaction  |            -             |             -              |
| /product-transaction/<int:id> |        -        | Detail Product Transaction | Edit Product Transaction | Delete Product Transaction |

## API Documentation

### List of Endpoints

- [Product](#Product)
  - [Get All Product](#get-all-product)
  - [Get Product by id](#get-product-by-id)
- [Transaction](#Transaction)
  - [Get Transaction by id](#get-transaction-by-id)
  - [Add Transaction](#add-transaction)
  - [Edit Transaction](#edit-transaction)
  - [Delete Transaction](#delete-transaction)

## Product

### Get All Product

- Method : GET
- URL : `/products`
- Request body : -
- Response body :

```json
{
  "message": "OK",
  "data": [
    {
      "id": 1,
      "name": "Nasi Putih",
      "price": 4000
    },

    {
      "id": 2,
      "name": "Ayam Dada",
      "price": 10000
    },

    {
      "id": 3,
      "name": "Ayam Paha Bawah",
      "price": 7000
    }
  ]
}
```

### Get Product by id

- Method : GET
- URL : `/products/<int:id>`
- Request body : -
- Response body :

```json
{
  "message": "OK",
  "data": {
    "id": 1,
    "name": "Nasi Putih",
    "price": 4000
  }
}
```

## Transaction

### Add Transaction

- Method : POST
- URL : `/transactions`
- Request body :

```json
[
  {
    "id_product": 1,
    "quantity": 2
  },
  {
    "id_product": 2,
    "quantity": 1
  }
]
```

- Response body :

```json
{
  "message": "Transaction Added",
  "data": {
    "id": 1,
    "total_price": 18000,
    "date": "17-12-2022",
    "detail_transaction": [
      {
        "id": 1,
        "id_product": 1,
        "quantity": 2,
        "total_price": 8000
      },
      {
        "id": 2,
        "id_product": 2,
        "quantity": 1,
        "total_price": 10000
      }
    ]
  }
}
```

### Get Transaction By Id

- Method : GET
- URL : `/transactions/<int:id>`
- Request body : -

- Response body :

```json
{
  "message": "OK",
  "data": {
    "id": 1,
    "total_price": 18000,
    "date": "17-12-2022",
    "detail_transaction": [
      {
        "id": 1,
        "id_product": 1,
        "quantity": 2,
        "total_price": 8000
      },
      {
        "id": 2,
        "id_product": 2,
        "quantity": 1,
        "total_price": 10000
      }
    ]
  }
}
```

### Delete Transaction

- Method : DELETE
- URL : `/transactions/<int:id>`
- Request body : -
- Response body:

```json
{
  "message": "Transaction 1 has been deleted",
  "data": ""
}
```

## Contributor

- Yulyano Thomas Djaya - 56419764
