create table if not exists boilerplate.party (
    id uuid not null
        constraint party_id_pkey primary key,
    party_name text not null
);