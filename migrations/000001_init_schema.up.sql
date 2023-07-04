drop table if exists urls;



create extension if not exists "uuid-ossp";

create table if not exists urls (
                                    id uuid primary key default uuid_generate_v4(),
                                    short_url varchar(255) not null,
                                    long_url varchar(255) not null
);