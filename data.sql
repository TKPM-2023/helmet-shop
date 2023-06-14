drop table if exists users;
drop table if exists categories;
drop table if exists products;
drop table if exists product_ratings;
drop table if exists orders;
drop table if exists order_details;
drop table if exists contacts;
drop table if exists carts;
drop table if exists cart_products;
drop table if exists images;

create table users
(
    id         int auto_increment primary key,
    email      varchar(50)                                                                  not null,
    password   varchar(50)                                                                  not null,
    last_name  varchar(50)                                                                  not null,
    first_name varchar(50)                                                                  not null,
    phone      varchar(20)                                                                  null,
    role       enum ('user', 'admin') default 'user'                                        not null,
    salt       varchar(50)                                                                  null,
    avatar     json                                                                         null,
    status     int                    default 1                                             not null,
    created_at timestamp              default current_timestamp                             null,
    updated_at timestamp              default current_timestamp on update current_timestamp null,
    unique (email)
);


create table categories
(
    id          int auto_increment primary key,
    name        varchar(100)                                                    not null,
    description text                                                            null,
    icon        json                                                            null,
    status      int       default 1                                             not null,
    created_at  timestamp default current_timestamp                             null,
    updated_at  timestamp default current_timestamp on update current_timestamp null,
    total_product int     default 0                                             not null
);


create table products
(
    id          int auto_increment primary key,
    name        varchar(255)                                                    not null,
    description text                                                            null,
    price       float                                                           not null,
    quantity    int                                                             not null,
    images      json                                                            not null,
    status      int       default 1                                             not null,
    created_at  timestamp default current_timestamp                             null,
    updated_at  timestamp default current_timestamp on update current_timestamp null,
    total_rating int     default 0                                             not null,
    category_id         int not null
);

create table product_ratings
(
    id         int auto_increment primary key,
    user_id    int                                                             not null,
    product_id    int                                                             not null,
    point      float     default 0                                             null,
    comment    text                                                            null,
    status     int       default 1                                             not null,
    created_at timestamp default current_timestamp                             null,
    updated_at timestamp default current_timestamp on update current_timestamp null
);

create table orders
(
    id          int auto_increment primary key,
    user_id     int                                                             not null,
    contact_id int                                                              not null,
    total_price float                                                           not null,
    comment     text                                                            null,
    status      int       default 1                                             not null,
    created_at  timestamp default current_timestamp                             null,
    updated_at  timestamp default current_timestamp on update current_timestamp null,
    order_status varchar(50)  not null,
);


create table order_details
(
    id          int auto_increment primary key,
    order_id    int                                                             not null,
    product_origin json                                                            null,
    price       float                                                           not null,
    quantity    int                                                             not null,
    discount    float     default 0                                             null,
    status      int       default 1                                             not null,
    created_at  timestamp default current_timestamp                             null,
    updated_at  timestamp default current_timestamp on update current_timestamp null
);


create table contacts (
                          id int auto_increment primary key,
                          user_id int not null,
                          name varchar(255) not null,
                          addr varchar(255) not null,
                          phone varchar(20) null,
                          status int default 1 not null,
                          created_at timestamp default current_timestamp null,
                          updated_at timestamp default current_timestamp on update current_timestamp null
);

create table carts
(
    id          int auto_increment primary key,
    user_id        int                                                             not null,
    status      int       default 1                                             not null,
    created_at  timestamp default current_timestamp                             null,
    updated_at  timestamp default current_timestamp on update current_timestamp null,
    total_product int     default 0                                             not null
);

create table cart_products
(
    cart_id          int auto_increment primary key,
    product_id        int                                                             not null,
    quantity    int                                                             not null,
    created_at  timestamp default current_timestamp                             null,
    updated_at  timestamp default current_timestamp on update current_timestamp null
);


create table favorites
(
    id          int auto_increment primary key,
    user_id        int                                                             not null,
    status      int       default 1                                             not null,
    created_at  timestamp default current_timestamp                             null,
    updated_at  timestamp default current_timestamp on update current_timestamp null,
    total_product int     default 0                                             not null
);

create table favorite_products
(
    favorite_id          int auto_increment primary key,
    product_id        int                                                             not null,
    created_at  timestamp default current_timestamp                             null,
    updated_at  timestamp default current_timestamp on update current_timestamp null
);
create table images (
                        id int auto_increment primary key,
                        url text not null,
                        width int not null,
                        height int not null,
                        cloud_name char(50) not null
);
