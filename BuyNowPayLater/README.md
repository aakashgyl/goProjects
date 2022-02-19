# Buy Now Pay Later Service
As a pay later service we allow our users to buy goods from a merchant now, and then allow them to pay for those goods at a later date.
The service works inside the boundary of following simple constraints -
1. Let's say that for every transaction paid through us, we deduct some transaction fees
while paying back to the merchant..
   - For example, if the transaction amount is Rs.100, and the transaction fee is 10%,
   we pay Rs. 90 back to the merchant.
   - The fee varies from merchant to merchant.
   - A merchant can decide to change the fee, at any point in time.
2. All users get onboarded with a credit limit, beyond which they can't transact.
   - If a transaction value crosses this credit limit, we reject the transaction.
   
##Use Cases
There are various use cases our service is intended to fulfil
- allow merchants to be onboarded with the fee
- allow merchants fee to be changed
- allow users to be onboarded (name, email-id and credit-limit)
- allow a user to carry out a transaction of some amount with a merchant.
- allow a user to pay back their dues (full or partial)
- for inputs like merchant fee percentage changes or credit limit changes for a user, the
  system adapts itself.
- Reporting:
  - how much fee we’ve collected from a merchant till date
  - dues for a user so far
  - which users have reached their credit limit
  - total dues from all users together
  
## CLI
Here is how the command line interface, corresponding to the use-cases mentioned above, looks like
- `new user u1 u1@email.in 1000` 
- `new merchant m1 m1@gmail.com 2%`
- `new txn u1 m2 400`
- `update merchant m1 1% payback u1 300` 
- `report fee m1`
- `report dues u1`
- `report users-at-credit-limit` 
- `report total-dues`

##Example Flow
### Flow 1:
```
new user u1 u1 300
new user u2 u2 400
new user u3 u3 500
new merchant m1 m1 0.5
new merchant m2 m2 1.5
new merchant m3 m3 1.25
new txn u2 m1 500
new txn u1 m2 300
new txn u1 m3 10
report users-at-credit-limit
new txn u3 m3 200
new txn u3 m3 300
report users-at-credit-limit
report fee m3
payback u3 400
report total-dues
```

### Flow 2:
```
new user u1 u1 300
new merchant m1 m1 0.5
new txn u1 m1 500
new txn u1 m1 300
report users-at-credit-limit
payback u1 200
report total-dues
report fee m1
report users-at-credit-limit
```