create table product_category
(
    id   uuid default gen_random_uuid() not null
        constraint product_category_pkey
            primary key,
    name text                           not null
);

alter table product_category
    owner to postgres;

create table products
(
    id                  uuid                     default gen_random_uuid() not null
        constraint products_pkey
            primary key,
    name                text,
    price               numeric,
    count               integer,
    product_category_id uuid                                               not null
        constraint products_product_category_id_fkey
            references product_category,
    created_at          timestamp with time zone default now(),
    updated_at          timestamp with time zone default now()
);

alter table products
    owner to postgres;

create table providers
(
    id         uuid                     default gen_random_uuid() not null
        constraint providers_pkey
            primary key,
    name       text,
    address    text,
    phone      text,
    email      text,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

alter table providers
    owner to postgres;

create table clients
(
    id         uuid                     default gen_random_uuid() not null
        constraint clients_pkey
            primary key,
    name       text,
    email      text,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

alter table clients
    owner to postgres;

create table employees
(
    id         uuid                     default gen_random_uuid() not null
        constraint employees_pkey
            primary key,
    name       text,
    email      text,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

alter table employees
    owner to postgres;

create table orders
(
    id            uuid                     default gen_random_uuid() not null
        constraint orders_pkey
            primary key,
    client_id     uuid                                               not null
        constraint orders_client_id_fkey
            references clients,
    product_id    uuid                                               not null
        constraint orders_product_id_fkey
            references products,
    employee_id   uuid                                               not null
        constraint orders_employee_id_fkey
            references employees,
    count         integer,
    address       text,
    created_at    timestamp with time zone default now(),
    updated_at    timestamp with time zone default now(),
    current_price integer                                            not null
);

alter table orders
    owner to postgres;

