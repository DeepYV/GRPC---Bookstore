


CREATE TABLE IF NOT EXISTS public.book
(
    book_id text NOT NULL,
    author text,
    title text ,
    book_name text,
    CONSTRAINT book_pkey PRIMARY KEY (book_id)
)


CREATE TABLE IF NOT EXISTS page
(
    book_id text ,
    page_number text,
    page_size text ,
    page_content text,
    CONSTRAINT page_bookid_fkey FOREIGN KEY (book_id)
        REFERENCES book (book_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)