create extension if not exists "uuid-ossp";

drop table if exists "user";

create table "user"
(
    uuid        uuid    default uuid_generate_v4() not null
        primary key,
    username    varchar(255)                       not null
        constraint user_pk_username
            unique,
    password    varchar(255)                       not null,
    solved      integer default 0                  not null,
    submitted   integer default 0                  not null,
    locked      boolean default false              not null,
    sex         varchar(255),
    nickname    varchar(255),
    school      varchar(255),
    department  varchar(255),
    major       varchar(255),
    email       varchar(255),
    create_at   date,
    update_at   date,
    delete_at   date,
    description varchar(255)
);

alter table "user"
    owner to postgres;


