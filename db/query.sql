-- name: CreateCustomer :one
insert into customers(firstName, lastName, email)
values ($1, $2, $3)
returning id;

-- name: UpdateComusmerEmail :exec
update customers set email = $1 where id = $2;

-- name: CreateOrderItem :execresult
insert into orderitems (orderId, productId, quantity, subtotal) values ($1, $2, $3, $4);

-- name: CreateOrder :one
insert into orders (customerId, orderDate, totalAmount)
values ($1, $2, $3)
returning id;

-- name: CreateProduct :one
insert into products (name, price)
values ($1, $2)
returning id;

