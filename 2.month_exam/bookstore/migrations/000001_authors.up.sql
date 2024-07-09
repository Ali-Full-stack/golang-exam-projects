create table if not exists author(
	author_id serial primary key,
	name varchar(65) not null,
	birth_date timestamp not null,
	biography text  unique,
	created_at timestamp default now(),
	updated_at timestamp default now()
);

create index author_name_idx
on author(name) ;