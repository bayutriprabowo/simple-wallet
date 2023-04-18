## Endpoint
### Add User Method POST
- Url
```endpoint
localhost:8090/account
```
- Request body
```
{
    "customer_name": "Siti Husna",
    "balance": 10000
}
```
### Get Account Method GET
- url
```endpoit
localhost:8090/account
```
### Get Account Detail Methode GET
- url
```endpoint
localhost:8090/account/{{account_number}}
```
### Transfer method PUT
- url
```endpoint
localhost:8090/transfer
```
- Request header
```header
account_number = {{account_number}}
```
- Request body
```body
{
    "to_account_number": "555002",
    "amount": 100
}
```
