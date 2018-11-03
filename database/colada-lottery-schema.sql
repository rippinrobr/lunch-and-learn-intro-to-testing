CREATE TABLE IF NOT EXISTS drinkers (uid INTEGER PRIMARY KEY AUTOINCREMENT, name varchar NOT NULL, can_make int default 1, headshot_path varchar default '');
INSERT INTO drinkers (name, can_make)
VALUES ('Space Rob', 1),
    ('Earth Rob', 1),
     ('Marcel', 1),
     ('Patrick', 1),
     ('Aarti', 0),
     ('Anjela', 0),
     ('Jonathan', 1),
     ('Sam', 1),
     ('Emma', 0),
     ('Robby', 1);

CREATE TABLE  IF NOT EXISTS history (id INTEGER PRIMARY KEY AUTOINCREMENT, barista varchar NOT NULL, cleaner varchar not null);



