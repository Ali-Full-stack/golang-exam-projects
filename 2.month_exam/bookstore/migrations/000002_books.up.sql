create table if not exists books(
	book_id serial primary key,
	title varchar(255) not null,
	category varchar(65) not null,
	author_id int not null references author(author_id) on delete cascade,
	publication_date timestamp not null,
	isbn varchar(65) unique ,
	description text not null,
	created_at timestamp default now(),
	updated_at timestamp default now()
);

create index books_composite_idx 
on books (title, category, publication_date);