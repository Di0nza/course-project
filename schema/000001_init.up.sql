CREATE TABLE users
    (
        id serial not null unique,
        email varchar(255) not null unique,
        password_hash varchar(255) not null
    );

CREATE TABLE brawlers_list
    (
        id serial not null unique,
        brawlers_name varchar(255) not null,
        current_level int not null,
        available_PP int not null,
        new_level int not null,
        gold int not null,
        pp int not null,
        cp_for_gold int not null,
        cp_gold int not null,
        cp_for_pp int not null,
        cp_pp int not null,
        cp_total int not null
);

CREATE TABLE users_lists(
    id serial not null unique,
    user_id int references users(id) on delete cascade  not null,
    lists_id int references brawlers_lists (id) on delete cascade not null
);