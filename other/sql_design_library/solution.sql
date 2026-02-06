-- 1 task

create table Author (
    id bigserial primary key,
    name text not null
);

create table Book ( -- abstract book
    id bigserial primary key,
    name text not null
);


create table Reader (
    id bigserial primary key,
    name text not null
);

create table AuthorBook (
    author_id bigint,
    book_id bigint,
    primary key (author_id, book_id),
    foreign key (author_id) references Author(id),
    foreign key (book_id) references Book(id)
);

create table BookCopy ( -- unique physical book
    id bigserial primary key,
    book_id bigint not null,
    foreign key (book_id) references Book(id)
);

create table BookLoan( -- book that was booked by smn
    book_copy_id bigint not null unique,
    reader_id    bigint not null,
    loaned_at    timestamp not null default now(),

    primary key (book_book_id),
    foreign key (book_book_id) references BookCopy(id),
    foreign key (reader_id)    references Reader(id)
);


-- 2 task

select any_value(b.name) "name", count(bc.id) - count(bl.book_copy_id) "available"
from Book b
left join BookCopy bc on b.id = bc.book_id
left join BookLoan bl on bl.book_copy_id = bc.id
group by b.id;


-- for inserting and checking
INSERT INTO Author (name) VALUES
('Лев Толстой'),
('Фёдор Достоевский');

INSERT INTO Book (name) VALUES
('Война и мир'),
('Преступление и наказание');

INSERT INTO AuthorBook (author_id, book_id) VALUES
(1, 1),
(2, 2);

INSERT INTO Reader (name) VALUES
('Иван'),
('Мария');

INSERT INTO BookCopy (copy_id) VALUES
(1),
(1),
(1),
(2),
(2);

INSERT INTO BookLoan (book_copy_id, reader_id) VALUES
(1, 1),  -- Война и мир у Ивана
(4, 2);  -- Преступление и наказание у Марии
