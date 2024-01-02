create table creditcards
(
    id         varchar(100) not null,
    type varchar(100) not null,
    name varchar(100) not null,
    number  varchar(32) not null,
    expired      varchar(100) not null,
    cvv      varchar(100) not null,
    user_id    varchar(100) not null,
    created_at bigint       not null,
    updated_at bigint       not null,
    primary key (id),
    foreign key fk_creditcard_user_id (user_id) references users (id)
) engine = innodb;
