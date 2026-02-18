CREATE TABLE author (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  first_name VARCHAR NOT NULL,
  last_name VARCHAR NOT NULL
);

CREATE TABLE article (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title VARCHAR NOT NULL,
  date_published DATE NOT NULL,
  page_range_start INT,
  page_range_end INT,
  url VARCHAR,
  journal_name VARCHAR,
  annotation VARCHAR,
  author_id INT,
  CONSTRAINT fk_authorarticle FOREIGN KEY (author_id) REFERENCES author (id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS article;
DROP TABLE IF EXISTS author;