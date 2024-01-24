create table if not exists rating
(
    id      varchar(100) not null primary key,
    user_id varchar(100) not null,
    dish_id varchar(100) not null,
    rating  int          not null,
    constraint rating_user_id_fk foreign key (user_id) references "user" (id)
);
