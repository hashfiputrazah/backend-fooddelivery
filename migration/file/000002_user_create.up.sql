create table if not exists "user"
(
    id           varchar(100) not null primary key,
    full_name    varchar(150) not null,
    password     text         not null,
    email        varchar(150) not null,
    address      varchar(150) null,
    birth_date   date         null,
    gender       varchar(150) not null,
    phone_number varchar(150) null,
    constraint user_pk unique (email)
);
