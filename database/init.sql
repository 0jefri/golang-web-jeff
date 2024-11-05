create database golang_web;

create table users(id serial primary key not null, username varchar(250) unique not null, password varchar(250) not null );
ALTER TABLE users
ADD COLUMN email VARCHAR(250) not null unique;

insert into users(username, password, email) values('jefri', '123', 'jefri@mail.com');

create table task(id serial primary key not null, name varchar(100) unique not null, status varchar(100) not null default 'progres');