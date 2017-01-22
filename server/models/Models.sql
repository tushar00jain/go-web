-- CREATE TYPE PhoneType AS ENUM ('MOBILE', 'HOME', 'WORK');

CREATE TABLE Person (
  "Id" SERIAL PRIMARY KEY NOT NULL,
  "Name" VARCHAR NOT NULL,
  "Email" VARCHAR
  -- "Number" PhoneNumber[]
);

CREATE TABLE PhoneNumber (
  "PersonId" INT REFERENCES Person("Id") ON DELETE CASCADE,
  "Number" VARCHAR,
  "Type" INT,
  PRIMARY KEY("PersonId", "Type")
);

CREATE TABLE AddressBook (
  "Self" INT REFERENCES Person("Id") ON DELETE CASCADE,
  "People" INT REFERENCES Person("Id") ON DELETE CASCADE,
  PRIMARY KEY("Self", "People")
);

INSERT INTO Person ("Name", "Email") VALUES 
  ('test', 'test@test.com'),
  ('test', 'test@test.com'),
  ('test', 'test@test.com'),
  ('test', 'test@test.com'),
  ('test', 'test@test.com'),
  ('test', 'test@test.com'),
  ('test', 'test@test.com'),
  ('test', 'test@test.com'),
  ('test', 'test@test.com'),
  ('test', 'test@test.com');

INSERT INTO PhoneNumber ("PersonId", "Number", "Type") VALUES
  (1, '1111111111', 0),
  (1, '1111111111', 1),
  (2, '1111111111', 0),
  (3, '1111111111', 0),
  (4, '1111111111', 0),
  (5, '1111111111', 2),
  (6, '1111111111', 0),
  (7, '1111111111', 0),
  (8, '1111111111', 0),
  (9, '1111111111', 0),
  (10, '1111111111', 0);
  -- (1, '1111111111', 'MOBILE'),
  -- (1, '1111111111', 'HOME'),
  -- (2, '1111111111', 'MOBILE'),
  -- (3, '1111111111', 'MOBILE'),
  -- (4, '1111111111', 'MOBILE'),
  -- (5, '1111111111', 'MOBILE'),
  -- (6, '1111111111', 'WORK'),
  -- (7, '1111111111', 'MOBILE'),
  -- (8, '1111111111', 'HOME'),
  -- (9, '1111111111', 'MOBILE'),
  -- (10, '1111111111', 'MOBILE');

  -- ('test', 'test@test.com', ARRAY[('1111111111', 'MOBILE')::PhoneNumber]),
  -- ('test', 'test@test.com', ARRAY[('1111111111', 'MOBILE')::PhoneNumber]),
  -- ('test', 'test@test.com', ARRAY[('1111111111', 'MOBILE')::PhoneNumber]),
  -- ('test', 'test@test.com', ARRAY[('1111111111', 'HOME')::PhoneNumber]),
  -- ('test', 'test@test.com', ARRAY[('1111111111', 'MOBILE')::PhoneNumber]),
  -- ('test', 'test@test.com', ARRAY[('1111111111', 'WORK')::PhoneNumber]),
  -- ('test', 'test@test.com', ARRAY[('1111111111', 'MOBILE')::PhoneNumber]),
  -- ('test', 'test@test.com', ARRAY[('1111111111', 'MOBILE')::PhoneNumber]),
  -- ('test', 'test@test.com', ARRAY[('1111111111', 'MOBILE')::PhoneNumber]),
  -- ('test', 'test@test.com', ARRAY[('1111111111', 'MOBILE')::PhoneNumber]);
  -- ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  -- ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  -- ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  -- ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  -- ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  -- ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  -- ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  -- ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  -- ('test', 'test@test.com', ('1111111111', 'MOBILE')),
  -- ('test', 'test@test.com', ('1111111111', 'MOBILE'));

INSERT INTO AddressBook ("Self", "People") VALUES
  (1, 2),
  (2, 3);
