CREATE TABLE IF NOT EXISTS drinkers (uid INTEGER PRIMARY KEY AUTOINCREMENT, name varchar NOT NULL, can_make int default 1, headshot_path varchar default '');
INSERT INTO drinkers (name, can_make, headshot_path)
VALUES ('Space Rob', 1, '/img/space-sloth.jpg'),
    ('Earth Rob', 1, '/img/yates.png'),
     ('Marcel', 1, '/img/marcel.png'),
     ('Patrick', 1, '/img/patrick.png'),
     ('Aarti', 0, '/img/aarti.png'),
     ('Anjela', 0, '/img/anjela.png'),
     ('Jonathan', 1, '/img/jonathan.png'),
     ('Sam', 1, '/img/sam.png'),
     ('Emma', 0, '/img/emma.png'),
     ('Robby', 1, '/img/robby.png');

CREATE TABLE  IF NOT EXISTS history (id INTEGER PRIMARY KEY AUTOINCREMENT, barista varchar NOT NULL, cleaner varchar not null, drawn_at text not null);



