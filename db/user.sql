create extension if not exists "uuid-ossp";

drop table if exists user_base, user_role, problem, contest, judger, file;

create table if not exists user_base
(
    uuid      uuid         not null default uuid_generate_v4() primary key,
    username  varchar(255) not null unique,
    password  varchar(255) not null,
    nickname  varchar(255) not null,
    profile   jsonb        not null default '{}'::JSONB,
    role      varchar(255) not null default 'user',
    solved    integer      not null default 0,
    submitted integer      not null default 0,
    rating    integer      not null default 1500,
    create_at timestamp    not null default now(),
    update_at timestamp    not null default now(),
    delete_at timestamp
);

create table if not exists user_role
(
    uuid       uuid         not null default uuid_generate_v4() primary key,
    name       varchar(255) not null unique,
    permission jsonb        not null default '{}'::JSONB,
    create_at  timestamp    not null default now(),
    update_at  timestamp    not null default now(),
    delete_at  timestamp
);

create table if not exists problem
(
    uuid      uuid         not null default uuid_generate_v4() primary key,
    pid       varchar(255) not null unique,
    title     varchar(255) not null,
    type      varchar(255) not null,
    content   text         not null,
    create_by varchar(255) not null,
    rating    integer,
    config    jsonb        not null,
    create_at timestamp    not null default now(),
    update_at timestamp    not null default now(),
    delete_at timestamp
);

create table if not exists file
(
    uuid      uuid         not null default uuid_generate_v4() primary key,
    path      varchar(255) not null unique,
    type      varchar(255) not null,
    create_at timestamp    not null default now(),
    update_at timestamp    not null default now(),
    delete_at timestamp
);

create table if not exists contest
(
    uuid        uuid         not null default uuid_generate_v4() primary key,
    cid         integer      not null unique,
    title       varchar(255) not null,
    password    varchar(255),
    problems    jsonb        not null,
    description text,
    start_time  timestamp    not null,
    end_time    timestamp    not null,
    lock_time   timestamp    not null,
    create_by   varchar(255) not null,
    create_at   timestamp    not null default now(),
    update_at   timestamp    not null default now(),
    delete_at   timestamp
);

create table if not exists submission
(
    uuid      uuid         not null default uuid_generate_v4() primary key,
    sid       integer      not null unique,
    pid       varchar(255) not null,
    uid       varchar(255) not null,
    cid       varchar(255),
    result    varchar(255) not null,
    language  varchar(255) not null,
    memory    integer      not null,
    time      integer      not null,
    code      text         not null,
    create_at timestamp    not null default now(),
    update_at timestamp    not null default now(),
    delete_at timestamp
);

create table if not exists judger
(
    uuid      uuid         not null default uuid_generate_v4() primary key,
    jid       varchar(255) not null unique,
    status    integer      not null default 0,
    create_at timestamp    not null default now(),
    update_at timestamp    not null default now(),
    delete_at timestamp
);