CREATE TYPE PhoneType AS ENUM ('MOBILE', 'HOME', 'WORK');

CREATE TYPE PhoneNumber AS (
  "Number" VARCHAR,
  "Type" PhoneType
);

CREATE TABLE Person (
  "Id" SERIAL PRIMARY KEY NOT NULL,
  "Name" VARCHAR NOT NULL,
  "Email" VARCHAR,
  "Number" PhoneNumber
);

CREATE TABLE AddressBook (
  "Self" INT REFERENCES Person("Id") ON DELETE CASCADE,
  "People" INT REFERENCES Person("Id") ON DELETE CASCADE,
  PRIMARY KEY("Self", "People")
);

INSERT INTO Person ("Name", "Email", "Number") VALUES
  ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  ('test', 'test@test.com', ('1111111111', 'MOBILE'));

INSERT INTO AddressBook ("Self", "People") VALUES
  (1, 2),
  (2, 3);
