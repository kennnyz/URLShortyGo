drop table if exists urls;

create table if not exists urls (
                                    id bigint primary key,
                                    short_url varchar(255) not null,
                                    long_url varchar(255) not null
);