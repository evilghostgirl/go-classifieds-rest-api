

CREATE TABLE Provinces(
	ID VARCHAR(255) PRIMARY KEY NOT NULL ,
	Name VARCHAR(129) NOT NULL
);

CREATE TABLE Users (
  ID VARCHAR(255) PRIMARY KEY NOT NULL,
  first_name varchar(100) NOT NULL,
  last_name varchar(100) NOT NULL

);

CREATE TABLE Categories (
	ID VARCHAR(255) PRIMARY KEY NOT NULL,
	Title VARCHAR(129),
	Description VARCHAR
);


CREATE TABLE Localizations(
	ID VARCHAR (255) PRIMARY KEY NOT NULL,
	Name VARCHAR (129) NOT NULL,
	ProvinceID VARCHAR(255) REFERENCES Provinces(ID)
	);
	
CREATE TABLE Classifieds(
	ID VARCHAR(255) PRIMARY KEY NOT NULL,
	Title VARCHAR(129),
	Description VARCHAR,
	-- ,
	AddedTime TIMESTAMP,
	EndTime TIMESTAMP,
	Price DECIMAL,
	LocalizationID VARCHAR(255) REFERENCES Localizations(ID),
	OwnerID VARCHAR(255) REFERENCES Users(ID),
	CategoryID VARCHAR(255) REFERENCES Categories(ID)
);


INSERT INTO "provinces" ("id", "name")
VALUES ('23340f27-fb9b-43e1-b624-662c6d416e88', 'Kujawsko-Pomorskie');

INSERT INTO "provinces" ("id", "name")
VALUES ('23340f27-fb9b-43e1-b624-662c6d416e99', 'Pomorskie');


INSERT INTO "localizations" ("id", "name", "provinceid")
VALUES ('23340f27-fb9b-43e1-b624-662c5d416e99', 'Szczecinek', '23340f27-fb9b-43e1-b624-662c6d416e99');



INSERT INTO "localizations" ("id", "name", "provinceid")
VALUES ('23330f27-fb9b-43e1-b624-662c5d416e99', 'Toruń', '23340f27-fb9b-43e1-b624-662c6d416e88');
-- INSERT INTO "localizations" ("id", "name", "province_id")
-- VALUES ('23340f27-fb9b-43e1-c224-332c6d416e99', 'Toruń', '23340f27-fb9b-43e1-b624-662c6d416e88');

INSERT INTO "users" ("id", "first_name", "last_name")
VALUES ('23340f27-fb9b-43e1-b624-662c6d216e99', 'Michał', 'Radziwiłł');

INSERT INTO "users" ("id", "first_name", "last_name")
VALUES ('23340f27-fb9b-43e1-b624-662c6d411e99', 'Lea', 'Szafi');


INSERT INTO "categories" ("id", "title", "description")
VALUES ('23340f27-fb9b-43e1-b624-662c6d411a99', 'Ubrania', 'ciuszki');

INSERT INTO "categories" ("id", "title", "description")
VALUES ('23340f22-fb9b-43e1-b624-662c6d411a99', 'Dom i Ogród', 'rzeczy do domu');




INSERT INTO "classifieds" ("id", "title", "description", "addedtime", "endtime", "price", "localizationid", "ownerid", "categoryid")
VALUES ('26640f27-fe9b-43e1-b624-662c6d416e88', 'Koszulka zielona', 'Jest to zielona koszulka', '2004-10-19 10:23:54', '2010-10-19 10:23:54', '300', '23330f27-fb9b-43e1-b624-662c5d416e99', '23340f27-fb9b-43e1-b624-662c6d411e99', '23340f27-fb9b-43e1-b624-662c6d411a99');

INSERT INTO "classifieds" ("id", "title", "description", "addedtime", "endtime", "price", "localizationid", "ownerid", "categoryid")
VALUES ('26640f27-fa9b-43e1-b624-662c6d416e88', 'Koszulka biała', 'Jest to zielona koszulka', '2004-10-19 10:23:54', '2010-10-19 10:23:54', '300', '23340f27-fb9b-43e1-b624-662c5d416e99', '23340f27-fb9b-43e1-b624-662c6d411e99', '23340f27-fb9b-43e1-b624-662c6d411a99');

INSERT INTO "classifieds" ("id", "title", "description", "addedtime", "endtime", "price", "localizationid", "ownerid", "categoryid")
VALUES ('26640f27-fb9b-43e1-b624-662c6d416e88', 'Spodnie', 'spodnie', '2004-10-19 10:23:54', '2010-10-19 10:23:54', '300', '23340f27-fb9b-43e1-b624-662c5d416e99', '23340f27-fb9b-43e1-b624-662c6d411e99', '23340f27-fb9b-43e1-b624-662c6d411a99');

INSERT INTO "classifieds" ("id", "title", "description", "addedtime", "endtime", "price", "localizationid", "ownerid", "categoryid")
VALUES ('26640f27-fc9b-43e1-b624-662c6d416e88', 'Legginsy zielone', 'nic', '2004-10-19 10:23:54', '2010-10-19 10:23:54', '300', '23330f27-fb9b-43e1-b624-662c5d416e99', '23340f27-fb9b-43e1-b624-662c6d411e99', '23340f27-fb9b-43e1-b624-662c6d411a99');
