CREATE TABLE products
(
    id             UUID       NOT NULL
        CONSTRAINT products_pk
            PRIMARY KEY,
    price_in_cents INTEGER      NOT NULL,
    title          VARCHAR(255) NOT NULL,
    description    TEXT
);

INSERT INTO products (id, price_in_cents, title, description)
VALUES ('ff5d8560-c82e-44f0-9947-930852ca5ecc', 4390, 'Suporte para Notebook', 'Ideal para notebook ou tablet de qualquer tamanho'),
       ('b6eb4c27-fb36-4085-8cf5-7d4966606de6', 2690, 'Mouse Sem Fio', 'Design Anatômico'),
       ('54d1f02b-0e07-44f1-8c4e-17ee37059176', 4590, 'Mouse Pad', 'Tecido de malha fina de alta qualidade'),
       ('e360e76f-bde7-4802-9e47-bb37099660de', 4590, 'Teclado', 'Teclas macias e silenciosas'),
       ('83432a24-39bf-47f0-8ebb-4b1a852e8be2', 39990, 'Fritadeira', 'Controle de temperatura de até 200ºC'),
       ('4151e87c-6a90-4f60-ae20-1643d5205fe3', 14900, 'Panela de Pressão', 'Alumínio com revestimento interno e externo de antiaderente Starflon T1'),
       ('87472e90-4a99-447d-af2f-0b7b15e368b6', 8010, 'Lâmpada', 'Diversas opções de programação da iluminação'),
       ('9ec3e785-9aae-4e93-92f4-155592291efe', 199, 'Lava Louças', 'Rende mais nova fórmula com maior rendimento'),
       ('fc3a452e-552d-4937-9ade-9b631bb90fb9', 19217, 'Aspirador de Pó', 'Sistema 360º que possibilita o alcance à lugares de difícil acesso'),
       ('2d5bdbe3-14fd-42f3-abfc-f1b050b75b05', 7188, 'Luminaria', 'Coloração RGB'),
       ('a49ded07-03c4-4ab0-950a-d5ce6a3bc83f', 8990, 'Ventilador', 'Super potente e silencioso, produzido com 6 pás que garantem grande vazão de ar'),
       ('929b4d32-ba52-441f-8093-91caa1b6b148', 1331, 'Cabide de Plástico', 'Pacote com 10 unidades');
