CREATE TABLE users
(
    id            serial       primary key ,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE brawlers_list
(
    id            serial      primary key ,
    brawlers_name varchar(255) not null,
    current_level int          not null,
    available_PP  int          not null,
    new_level     int          not null,
    gold          int          not null,
    pp            int          not null,
    cp_for_gold   int          not null,
    cp_gold       int          not null,
    cp_for_pp     int          not null,
    cp_pp         int          not null,
    cp_total      int          not null
);

CREATE TABLE users_lists
(
    id       serial                                               primary key ,
    user_id  int references users (id) on delete cascade          not null,
    lists_id int references brawlers_list (id) on delete cascade not null

);