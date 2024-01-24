create table if not exists basket
(
    id      varchar(100) not null primary key,
    user_id varchar(100) not null,
    dish_id varchar(100) not null,
    amount  int          not null,
    constraint basket_dish_id_fk foreign key (dish_id) references dish (id),
    constraint basket_user_id_fk foreign key (user_id) references "user" (id)
);
