create table if not exists "order"
(
    id            varchar(100) not null primary key,
    user_id       varchar(100) not null,
    delivery_time timestamp    not null,
    order_time    timestamp    not null,
    status        varchar(100) not null,
    price         decimal      not null,
    address       text         not null,
    constraint order_user_id_fk foreign key (user_id) references "user" (id)
);
