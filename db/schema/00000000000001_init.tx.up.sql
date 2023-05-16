create schema if not exists boilerplate;

create table if not exists boilerplate.user
(
    id            uuid      not null
        constraint user_id_pkey primary key,
    token         text      not null,
    platform      smallint  not null,
    last_login_at timestamp not null,
    username      text      not null
);

create table if not exists boilerplate.user_settings
(
    id                     uuid      not null
        constraint user_settings_id_pkey primary key,
    user_id                uuid      not null
        constraint user_settings_user_id_fk references boilerplate.user,
    is_sound_enabled       boolean   not null default true,
    is_music_enabled       boolean   not null default true,
    is_left_handed_enabled boolean   not null default false
);

