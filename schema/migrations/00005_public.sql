-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS public;

create table if not exists public.hram_lib
(
    id            integer                                    not null
        constraint idx_36209_primary
            primary key,
    id_subsection integer      default 0                     not null,
    name          varchar(255) default ''::character varying not null,
    filename      varchar(255) default ''::character varying not null,
    flag_new      integer,
    weight        integer      default 0                     not null
);

create table if not exists public.hram_lib_section
(
    id     integer                                    not null
        constraint idx_36218_primary
            primary key,
    name   varchar(255) default ''::character varying not null,
    weight integer      default 0                     not null
);

create table if not exists public.hram_lib_subsection
(
    id         integer                                    not null
        constraint idx_36223_primary
            primary key,
    id_section integer      default 0                     not null,
    name       varchar(255) default ''::character varying not null,
    weight     integer      default 0                     not null,
    header     bytea,
    footer     bytea
);

create table if not exists public.hram_mus
(
    id            integer                                    not null
        constraint idx_36231_primary
            primary key,
    id_subsection integer      default 0                     not null,
    name          varchar(255) default ''::character varying not null,
    filename      varchar(255) default ''::character varying not null,
    flag_new      integer,
    weight        integer      default 0                     not null
);

create table if not exists public.hram_mus_section
(
    id     integer                                    not null
        constraint idx_36240_primary
            primary key,
    name   varchar(255) default ''::character varying not null,
    weight integer      default 0                     not null
);

create table if not exists public.hram_mus_subsection
(
    id         integer                                    not null
        constraint idx_36245_primary
            primary key,
    id_section integer      default 0                     not null,
    name       varchar(255) default ''::character varying not null,
    weight     integer      default 0                     not null,
    header     bytea,
    footer     bytea
);

create table if not exists public.hram_news
(
    id      integer                                    not null
        constraint idx_36253_primary
            primary key,
    data    date,
    news    text                                       not null,
    url     varchar(255) default ''::character varying not null,
    chapter varchar(30)  default 'ОБЩЕЦЕРКОВНЫЕ'::character varying
);

create table if not exists public.hram_news_new
(
    id      integer      default 0                     not null,
    data    date,
    news    text                                       not null,
    url     varchar(255) default ''::character varying not null,
    chapter varchar(30)  default 'ОБЩЕЦЕРКОВНЫЕ'::character varying
);

create table if not exists public.hram_talk
(
    id       integer                                    not null
        constraint idx_36268_primary
            primary key,
    data_q   date,
    name     varchar(255) default ''::character varying not null,
    email    varchar(255) default ''::character varying not null,
    question text                                       not null,
    answer   text                                       not null,
    data_a   date,
    flag     integer
);

create table if not exists public.hram_talk_new
(
    id       integer      default 0                     not null,
    data_q   date,
    name     varchar(255) default ''::character varying not null,
    email    varchar(255) default ''::character varying not null,
    question text                                       not null,
    answer   text                                       not null,
    data_a   date,
    flag     integer
);

create table if not exists public.hram_talk_old
(
    id       integer      default 0                     not null,
    data_q   date,
    name     varchar(255) default ''::character varying not null,
    email    varchar(255) default ''::character varying not null,
    question text                                       not null,
    answer   text                                       not null,
    data_a   date,
    flag     integer
);

create table if not exists public.hram_talk_recover
(
    id       integer                                    not null
        constraint idx_36291_primary
            primary key,
    data_q   date,
    name     varchar(255) default ''::character varying not null,
    email    varchar(255) default ''::character varying not null,
    question text                                       not null,
    answer   text                                       not null,
    data_a   date,
    flag     integer
);

create table if not exists public."user"
(
    id       integer      not null
        constraint idx_36298_primary
            primary key,
    name     varchar(30)  not null,
    fullname text         not null,
    hash     varchar(240) not null
);

comment on table public."user" is 'Пользователи системы';

create unique index if not exists idx_36298_users_hash_uindex
    on public."user" (hash);

create unique index if not exists idx_36298_users_username_uindex
    on public."user" (name);

create table if not exists public.goose_db_version
(
    id         integer generated by default as identity
        primary key,
    version_id bigint                  not null,
    is_applied boolean                 not null,
    tstamp     timestamp default now() not null
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd