

CREATE TABLE Provinces(
	ID VARCHAR(255) PRIMARY KEY NOT NULL ,
	Name VARCHAR(129) NOT NULL
);

CREATE TABLE Localizations(
	ID VARCHAR (255) PRIMARY KEY NOT NULL,
	Name VARCHAR (129) NOT NULL,
	ProvinceID VARCHAR(255) REFERENCES Provinces(ID)
	);
	
CREATE TABLE Users (
  ID VARCHAR(255) PRIMARY KEY NOT NULL,
  first_name varchar(100) NOT NULL,
  last_name varchar(100) NOT NULL,
  LocalizationID VARCHAR(255) REFERENCES Localizations(ID)

);

CREATE TABLE Categories (
	ID VARCHAR(255) PRIMARY KEY NOT NULL,
	Title VARCHAR(129),
	Description VARCHAR
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

INSERT INTO "users" ("id", "first_name", "last_name", "localizationid")
VALUES ('23340f27-fb9b-43e1-b624-662c6d216e99', 'Michał', 'Radziwiłł', '23330f27-fb9b-43e1-b624-662c5d416e99');

INSERT INTO "users" ("id", "first_name", "last_name", "localizationid")
VALUES ('23340f27-fb9b-43e1-b624-662c6d411e99', 'Lea', 'Szafi', '23340f27-fb9b-43e1-b624-662c5d416e99');


INSERT INTO "categories" ("id", "title", "description")
VALUES ('23340f27-fb9b-43e1-b624-662c6d411a99', 'Ubrania', 'ciuszki');

INSERT INTO "categories" ("id", "title", "description")
VALUES ('23340f22-fb9b-43e1-b624-662c6d411a99', 'Dom i Ogród', 'rzeczy do domu');

INSERT INTO "categories" ("id", "title", "description")
VALUES ('bc15d1d3-fb6c-4695-94b3-00dd1d76e4d4', 'Motoryzacja', 'Szeroka oferta używanych samochodów. Motocykle, przyczepy i opony.');

INSERT INTO "categories" ("id", "title", "description")
VALUES ('3b9a9c97-3d5a-4074-aee6-be7191897adb', 'Zdrowie', 'Suplementy diety, urządzenia do masażu, leki bez recepty i wiele innych.');


INSERT INTO "classifieds" ("id", "title", "description", "addedtime", "endtime", "price", "localizationid", "ownerid", "categoryid")
VALUES ('26640f27-fe9b-43e1-b624-662c6d416e88', 'KOSZULKA SERCE DAMSKA MODNY T-SHIRT - M', 'Koszulką damska biała - Linia życia - M

trwały nadruk technologią FLEX
dowolny kolor koszulki
nadruk nie traci kolorów, odporny na uszkodzenia
gramatura koszulki w zależności od dost. ok. 170g/m2
bawełna zmiękczana 100%
Ściągacz wzmocniony lycrą. Taśma wzmacniająca pod ściągaczem. Podwójne szwy na ramionach i spodzie koszulki', '2004-10-19 10:23:54', '2010-10-19 10:23:54', '300', '23330f27-fb9b-43e1-b624-662c5d416e99', '23340f27-fb9b-43e1-b624-662c6d411e99', '23340f27-fb9b-43e1-b624-662c6d411a99');

INSERT INTO "classifieds" ("id", "title", "description", "addedtime", "endtime", "price", "localizationid", "ownerid", "categoryid")
VALUES ('26640f27-fa9b-43e1-b624-662c6d416e88', '4F T-shirt KOSZULKA Męska TSM003 Bawełna Biała', 'Parametry:
Model: NOSH4-TSM003
Producent: 4F, Polska
Materiał: 100% bawełna
Krój: Regular
Przewiewna, świetnie leży
Ergonomiczny krój gwarantujący wygodę
Model niekrępujący ruchów
Uniwersalna koszulka idealna na co dzień oraz do uprawiania różnych sportów
Kolor: Biały / Biała', '2004-10-19 10:23:54', '2010-10-19 10:23:54', '300', '23340f27-fb9b-43e1-b624-662c5d416e99', '23340f27-fb9b-43e1-b624-662c6d411e99', '23340f27-fb9b-43e1-b624-662c6d411a99');

INSERT INTO "classifieds" ("id", "title", "description", "addedtime", "endtime", "price", "localizationid", "ownerid", "categoryid")
VALUES ('26640f27-fb9b-43e1-b624-662c6d416e88', 'MODNE SPODNIE DAMSKIE BOJÓWKI WIĄZANE KOLORY', 'Wymiary podane u dołu aukcji w Tabeli Rozmiarów

Tolerancja +/-3cm

spodnie wykonane z elastycznego materiału typu rurki
bardzo wygodne
sznurowane w talii
posiadają dwie kieszenie
W informacji dla sprzedającego proszę dopisać wybrany kolor w przeciwnym razie wysyłamy losowy', 
'2004-10-19 10:23:54', '2010-10-19 10:23:54', '300', 
'23340f27-fb9b-43e1-b624-662c5d416e99', '23340f27-fb9b-43e1-b624-662c6d411e99', '23340f27-fb9b-43e1-b624-662c6d411a99');

INSERT INTO "classifieds" ("id", "title", "description", "addedtime", "endtime", "price", "localizationid", "ownerid", "categoryid")
VALUES ('26640f27-fc9b-43e1-b624-662c6d416e88', 'MODNE I STYLOWE SPODNIE DRESOWE LAMPAS S/36 L315', 'Świetne spodnie dresowe.
W pasie gumka, którą można regulować szerokość.
Posiadają kieszenie.
Nogawki wykończone ściągaczem.
Po bokach modne lampasy, które są delikatnie dłuższe od nogawki.
Bardzo wygodny i praktyczny model.
Materiał lekko uciągliwy.
Najnowsza kolekcja.
Satysfakcja gwarantowana.
Podane wymiary są orientacyjne, mogą się nieznacznie różnić +/- 2cm.
Mierzone na płasko, podane szerokości należy pomnożyć razy dla uzyskania obwodu.
Modelka ma 161 cm wzrostu.
Na zdjęciu rozmiar S. Zapraszam do zakupu.', 
'2019-10-19 10:23:54', '2019-10-20 10:23:54', '300', 
'23330f27-fb9b-43e1-b624-662c5d416e99', '23340f27-fb9b-43e1-b624-662c6d411e99', '23340f27-fb9b-43e1-b624-662c6d411a99');
