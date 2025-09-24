# E-Commerce Backend 
## A basic E-commerce backend with user management(jwt, role-based access, protected routes), full CRUD product routes, order management, cart management, stripe and paypal integration for payments. 

## Routes 
### Auth 
> Post `/auth/register` 
> Post `/auth/login`

### User 
> Get `/users/me` (auth)
> Put `/users/me` (auth)

### Products 
> Get `/products`
> Get `/products/:id`
> Post `/products/` (only admin) 
> Put `/products/:id` (only admin)
> Delete `/products/:id (only admin)

### Cart 
> Post `/cart` 
> Delete `/cart/:itemId` 

### Orders 
> Post `/orders` (Stripe/Paypal PaymentIntent)
> Get `/orders` 
> Get `/orders/:id`

### Webhooks 
Post `/webhooks/stripe` 
Post `/webhooks/paypal`