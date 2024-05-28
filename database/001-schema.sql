-- schema.sql
create table
     if not exists plant (
        id serial primary key,
        name varchar(255) not null,
        description varchar(255) not null,
        image varchar(255) not null,
        water int not null, -- millilitres per day
        sun int not null, -- hours of sunshine per day
        germination int not null, -- days to germinate
        flowering int not null, -- days to flower
        harvest int not null, -- days to harvest
        seed int not null, -- seeds per plant
        created_at timestamp not null default now (),
        updated_at timestamp not null default now ()
    );

create table
    if not exists users (
        id serial primary key,
        name varchar(255) not null,
        email varchar(255) not null,
        password varchar(255) not null,
        created_at timestamp not null default now (),
        updated_at timestamp not null default now ()
    );

create table
    if not exists plant_users (
        id serial primary key,
        plant_id int not null references plant (id),
        users_id int not null references users (id),
        created_at timestamp not null default now (),
        updated_at timestamp not null default now ()
    );

create table
    if not exists store (
        id serial primary key,
        name varchar(255) not null,
        address varchar(255) not null,
        city varchar(255) not null,
        postcode varchar(255) not null,
        country varchar(255) not null,
        created_at timestamp not null default now (),
        updated_at timestamp not null default now ()
    );

create table
    if not exists plant_store (
        id serial primary key,
        quantity int not null,
        plant_id int not null references plant (id),
        store_id int not null references store (id),
        created_at timestamp not null default now (),
        updated_at timestamp not null default now ()
    );
