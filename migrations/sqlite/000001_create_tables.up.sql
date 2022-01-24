CREATE TABLE IF NOT EXISTS users (
    id integer primary key,
    name varchar(255) unique,
    banc_acc varchar(255) not null,
    address varchar(255) not null,
    balance integer not null
);

CREATE TABLE IF NOT EXISTS rewards (
    id integer primary key,
    user_id integer not null,
    reward_id integer not null,
    type varchar(20) not null
);

CREATE TABLE IF NOT EXISTS money_rewards (
    id integer primary key,
    amount integer not null,
    sent integer default 0
);

CREATE TABLE IF NOT EXISTS item_rewards (
    id integer primary key,
    type varchar(30) not null,
    sent integer default 0
);

CREATE TABLE IF NOT EXISTS sessions (
    token varchar(255) primary key,
    user_id integer not null,
    expired_at integer not null
);