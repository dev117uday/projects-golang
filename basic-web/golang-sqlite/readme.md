# Golang SQLite

- Create table (even if doesnt exist) 
```sqlite3
create table if not exists newsfeed
(
	ID integer not null
		constraint newsfeed_pk
			primary key autoincrement,
	CONTENT text
);
```

- simple insertion
```sqlite
INSERT INTO newsfeed (content) values (?)
```
