create table users
(
    id         varchar(100) not null,
    name       varchar(100) not null,
    password   varchar(100) not null,
    address      varchar(100),
    photos      varchar(100),
    token varchar(100) null,
    created_at bigint       not null,
    updated_at bigint       not null,
    primary key (id)
) engine = InnoDB;
