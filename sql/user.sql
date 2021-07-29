create table auth."user"
(
    id         serial                                       not null
        constraint user_pk
            primary key,
    full_name  varchar(255) default NULL::character varying,
    name       varchar(100) default NULL::character varying,
    last_name  varchar(100) default NULL::character varying,
    nickname   varchar(50)  default NULL::character varying,
    id_card    varchar(30)  default NULL::character varying,
    email      varchar(100) default NULL::character varying,
    password   varchar(100) default NULL::character varying not null,
    created_at timestamp    default now(),
    updated_at timestamp    default now()
);

alter table auth."user"
    owner to postgres;