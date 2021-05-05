CREATE TABLE products
(
    id             UUID       NOT NULL
        CONSTRAINT products_pk
            PRIMARY KEY,
    price_in_cents INTEGER      NOT NULL
);

INSERT INTO products (id, price_in_cents)
VALUES ('ff5d8560-c82e-44f0-9947-930852ca5ecc', 4390),
       ('b6eb4c27-fb36-4085-8cf5-7d4966606de6', 2690),
       ('54d1f02b-0e07-44f1-8c4e-17ee37059176', 4590),
       ('e360e76f-bde7-4802-9e47-bb37099660de', 4590),
       ('83432a24-39bf-47f0-8ebb-4b1a852e8be2', 39990),
       ('4151e87c-6a90-4f60-ae20-1643d5205fe3', 14900),
       ('87472e90-4a99-447d-af2f-0b7b15e368b6', 8010),
       ('9ec3e785-9aae-4e93-92f4-155592291efe', 199),
       ('fc3a452e-552d-4937-9ade-9b631bb90fb9', 19217),
       ('2d5bdbe3-14fd-42f3-abfc-f1b050b75b05', 7188),
       ('a49ded07-03c4-4ab0-950a-d5ce6a3bc83f', 8990),
       ('929b4d32-ba52-441f-8093-91caa1b6b148', 1331);

CREATE TABLE users
(
    id             UUID       NOT NULL
        CONSTRAINT users_pk
            PRIMARY KEY,
    date_of_birth DATE        NOT NULL
);

INSERT INTO users (id, date_of_birth)
VALUES ('11054c65-89dd-46a6-86ab-c603195100a5', '2000-11-25'),
       ('b6eda5dc-8197-498a-b41b-7841b2b728f9', '1997-01-24'),
       ('512135bd-efe2-44af-afbc-6ddc072a41f6', '1993-11-01'),
       ('a87a888f-d2d8-4697-8777-1bb9ac858507', '1991-10-06');
