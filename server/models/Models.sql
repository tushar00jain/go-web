CREATE TYPE PhoneType AS ENUM ('MOBILE', 'HOME', 'WORK');

CREATE TYPE PhoneNumber AS (
  phoneNumber VARCHAR,
  phoneType PhoneType
);

CREATE TABLE Person (
  id SERIAL PRIMARY KEY NOT NULL,
  name VARCHAR NOT NULL,
  email VARCHAR,
  phoneNumber PhoneNumber
);

CREATE TABLE AddressBook (
  self INT REFERENCES Person(id) ON DELETE CASCADE,
  people INT REFERENCES Person(id) ON DELETE CASCADE,
  PRIMARY KEY(self, people)
);

INSERT INTO Person (name, email, phoneNumber.phoneNumber, phoneNumber.phoneType) VALUES
  ('test', 'test@test.com', '1111111111', 'MOBILE'),
  ('test', 'test@test.com', '1111111111', 'MOBILE'),
  ('test', 'test@test.com', '1111111111', 'MOBILE'),
  ('test', 'test@test.com', '1111111111', 'MOBILE'),
  ('test', 'test@test.com', '1111111111', 'MOBILE'),
  ('test', 'test@test.com', '1111111111', 'MOBILE'),
  ('test', 'test@test.com', '1111111111', 'MOBILE'),
  ('test', 'test@test.com', '1111111111', 'MOBILE'),
  ('test', 'test@test.com', '1111111111', 'MOBILE'),
  ('test', 'test@test.com', '1111111111', 'MOBILE');

INSERT INTO AddressBook (self, people) VALUES
  (1, 2),
  (2, 3);
