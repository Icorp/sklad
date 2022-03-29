create DATABASE "sklad";

create table product_category
(
    id   uuid default gen_random_uuid() not null
        primary key,
    name text                           not null
);

create table products
(
    id                  uuid                     default gen_random_uuid() not null
        primary key,
    name                text,
    price               numeric,
    count               integer,
    product_category_id uuid                                               not null
        references product_category,
    created_at          timestamp with time zone default now(),
    updated_at          timestamp with time zone default now()
);

create table providers
(
    id         uuid                     default gen_random_uuid() not null
        primary key,
    name       text,
    address    text,
    phone      text,
    email      text,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create table clients
(
    id         uuid                     default gen_random_uuid() not null
        primary key,
    name       text,
    email      text,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create table employees
(
    id         uuid                     default gen_random_uuid() not null
        primary key,
    name       text,
    email      text,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create table orders
(
    id          uuid                     default gen_random_uuid() not null
        primary key,
    client_id   uuid                                               not null
        references clients,
    product_id  uuid                                               not null
        references products,
    employee_id uuid                                               not null
        references employees,
    count       integer,
    address     text,
    created_at  timestamp with time zone default now(),
    updated_at  timestamp with time zone default now()
);

