create table if not exists order_dish
(
    order_id varchar(100) not null,
    user_id  varchar(100) not null,
    dish_id  varchar(100) not null,
    amount   int          not null,
    constraint order_dish_dish_id_fk foreign key (dish_id) references dish (id),
    constraint order_dish_order_id_fk foreign key (order_id) references "order" (id),
    constraint order_dish_user_id_fk foreign key (user_id) references "user" (id)
);
