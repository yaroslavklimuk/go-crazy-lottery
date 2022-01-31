CREATE TABLE IF NOT EXISTS users (
    id integer primary key,
    name varchar(255) unique,
    passwd varchar(255) not null,
    banc_acc varchar(255) not null,
    address varchar(255) not null,
    balance integer not null
);

CREATE TABLE IF NOT EXISTS money_rewards (
    id integer primary key,
    user_id integer not null,
    amount integer not null,
    sent integer default 0,
    CONSTRAINT fk_users
        FOREIGN KEY (user_id)
            REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS item_rewards (
    id integer primary key,
    user_id integer not null,
    type varchar(30) not null,
    sent integer default 0,
    CONSTRAINT fk_users
        FOREIGN KEY (user_id)
            REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS sessions (
    token varchar(255) primary key,
    user_id integer not null,
    expired_at integer not null,
    CONSTRAINT fk_users
        FOREIGN KEY (user_id)
            REFERENCES users(id)
);