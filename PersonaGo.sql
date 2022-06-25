create database GO;

use Go;

DROP TABLE IF EXISTS Persona;

create table Persona (
	IdPersona int not null primary key,
	Nombre VARCHAR(50),
	Apellido VARCHAR(50),
	Email VARCHAR(50),
	Genero VARCHAR(50)
);
insert into Persona (IdPersona, Nombre, Apellido, Email, Genero) values (1, 'Archibaldo', 'Cage', 'acage0@google.co.uk', 'Male');
insert into Persona (IdPersona, Nombre, Apellido, Email, Genero) values (2, 'Jeannie', 'Docksey', 'jdocksey1@xrea.com', 'Female');
insert into Persona (IdPersona, Nombre, Apellido, Email, Genero) values (3, 'Graig', 'Corstan', 'gcorstan2@acquirethisname.com', 'Male');
insert into Persona (IdPersona, Nombre, Apellido, Email, Genero) values (4, 'Gianina', 'Finlaison', 'gfinlaison3@thetimes.co.uk', 'Female');
insert into Persona (IdPersona, Nombre, Apellido, Email, Genero) values (5, 'Dun', 'Kitt', 'dkitt4@deviantart.com', 'Male');
insert into Persona (IdPersona, Nombre, Apellido, Email, Genero) values (6, 'Pat', 'Cowen', 'pcowen5@yale.edu', 'Male');
insert into Persona (IdPersona, Nombre, Apellido, Email, Genero) values (7, 'Cecelia', 'Dockree', 'cdockree6@springer.com', 'Genderqueer');
insert into Persona (IdPersona, Nombre, Apellido, Email, Genero) values (8, 'Windy', 'O''Loghlen', 'wologhlen7@yelp.com', 'Female');
insert into Persona (IdPersona, Nombre, Apellido, Email, Genero) values (9, 'Elsie', 'Shelper', 'eshelper8@dot.gov', 'Female');
insert into Persona (IdPersona, Nombre, Apellido, Email, Genero) values (10, 'Lilian', 'Moehle', 'lmoehle9@homestead.com', 'Female');