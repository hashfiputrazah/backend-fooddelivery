create table if not exists dish
(
    id          varchar(100)    not null primary key,
    name        varchar(150)    not null,
    description text            null,
    price       decimal         not null,
    image       text            null,
    vegetarian  boolean         not null,
    rating      float default 0 not null,
    category    varchar(50)     not null
);
