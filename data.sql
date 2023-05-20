drop table if exists users;
drop table if exists images;

create table users
(
    id         int auto_increment primary key,
    email      varchar(50)                                                                  not null,
    password   varchar(50)                                                                  not null,
    last_name  varchar(50)                                                                  not null,
    first_name varchar(50)                                                                  not null,
    phone      varchar(20)                                                                  null,
    role      enum ('user', 'admin') default 'user'                                        not null,
    salt       varchar(50)                                                                  null,
    avatar     json                                                                         null,
    status     int                    default 1                                             not null,
    created_at timestamp              default current_timestamp                             null,
    updated_at timestamp              default current_timestamp on update current_timestamp null,
    unique (email)
);


create table images (
    id int auto_increment primary key,
    url text not null,
    width int not null,
    height int not null
);