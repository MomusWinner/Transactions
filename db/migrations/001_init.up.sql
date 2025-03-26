create table customers(
    id serial primary key,
    firstName text not null,
    lastName text not null,
    email text not null
);

-- name: UpdateCustomerEmail :exec
update custoemer set email = $1;

create table products(
    id serial primary key,
    name text not null,
    price float not null
);

create table orders(
    id serial primary key,
    customerId int references customers(id) not null,
    orderDate date not null,
    totalAmount float not null
);

create table orderItems(
    id serial primary key,
    orderId int references orders(id) not null,
    productId int references products(id) not null,
    quantity int not null,
    subtotal float not null
);
