INSERT INTO product (id, name)
VALUES
    ('43c14dec-9a45-4257-8fed-2aa1a5a15c91', 'Product 1'),
    ('4bfa2d06-1709-46e9-9262-5e54857294e6', 'Product 2'),
    ('e21d443c-ceba-4479-ab43-5bcb2fe335dc', 'Product 3'),
    ('1806e505-1e79-4b0f-89ed-e1535c8983b4', 'Product 4'),
    ('9e28484c-c914-4c01-9f5d-fb2fae163599', 'Product 5');

INSERT INTO size (name)
VALUES
    ('Small'),
    ('Medium'),
    ('Large');

INSERT INTO warehouse (id, office_name)
VALUES
    ('e62f6003-7392-47cf-bd40-73b0371d23b3', 'Warehouse 1'),
    ('a494f72f-4044-43ed-b5c2-1145aa88352d', 'Warehouse 2'),
    ('5691d71b-9ca1-4298-ad8d-733bfe72e88c', 'Warehouse 3');

INSERT INTO option (product_id, size, id) VALUES
    ('43c14dec-9a45-4257-8fed-2aa1a5a15c91', 'Small', 'OPT1'),
    ('43c14dec-9a45-4257-8fed-2aa1a5a15c91', 'Medium', 'OPT2'),
    ('43c14dec-9a45-4257-8fed-2aa1a5a15c91', 'Large', 'OPT3'),
    ('4bfa2d06-1709-46e9-9262-5e54857294e6', 'Small', 'OPT4'),
    ('4bfa2d06-1709-46e9-9262-5e54857294e6', 'Medium', 'OPT5'),
    ('4bfa2d06-1709-46e9-9262-5e54857294e6', 'Large', 'OPT6'),
    ('e21d443c-ceba-4479-ab43-5bcb2fe335dc', 'Small', 'OPT7'),
    ('e21d443c-ceba-4479-ab43-5bcb2fe335dc', 'Medium', 'OPT8'),
    ('e21d443c-ceba-4479-ab43-5bcb2fe335dc', 'Large', 'OPT9'),
    ('1806e505-1e79-4b0f-89ed-e1535c8983b4', 'Small', 'OPT10'),
    ('1806e505-1e79-4b0f-89ed-e1535c8983b4', 'Medium', 'OPT11'),
    ('1806e505-1e79-4b0f-89ed-e1535c8983b4', 'Large', 'OPT12'),
    ('9e28484c-c914-4c01-9f5d-fb2fae163599', 'Small', 'OPT13'),
    ('9e28484c-c914-4c01-9f5d-fb2fae163599', 'Medium', 'OPT14'),
    ('9e28484c-c914-4c01-9f5d-fb2fae163599', 'Small', 'OPT15');

INSERT INTO inventory (option_id, warehouse_id, quantity) VALUES
    ('OPT1', 'e62f6003-7392-47cf-bd40-73b0371d23b3', 0),
    ('OPT1', 'a494f72f-4044-43ed-b5c2-1145aa88352d', 150),
    ('OPT1', '5691d71b-9ca1-4298-ad8d-733bfe72e88c', 100),
    ('OPT2', 'e62f6003-7392-47cf-bd40-73b0371d23b3', 0),
    ('OPT2', 'a494f72f-4044-43ed-b5c2-1145aa88352d', 0),
    ('OPT2', '5691d71b-9ca1-4298-ad8d-733bfe72e88c', 0),
    ('OPT3', 'e62f6003-7392-47cf-bd40-73b0371d23b3', 100),
    ('OPT3', 'a494f72f-4044-43ed-b5c2-1145aa88352d', 200),
    ('OPT3', '5691d71b-9ca1-4298-ad8d-733bfe72e88c', 300),
    ('OPT4', 'e62f6003-7392-47cf-bd40-73b0371d23b3', 100),
    ('OPT4', 'a494f72f-4044-43ed-b5c2-1145aa88352d', 200),
    ('OPT5', 'e62f6003-7392-47cf-bd40-73b0371d23b3', 100),
    ('OPT6', 'e62f6003-7392-47cf-bd40-73b0371d23b3', 0),
    ('OPT7', 'a494f72f-4044-43ed-b5c2-1145aa88352d', 1),
    ('OPT8', 'a494f72f-4044-43ed-b5c2-1145aa88352d', 150),
    ('OPT8', '5691d71b-9ca1-4298-ad8d-733bfe72e88c', 150),
    ('OPT9', 'e62f6003-7392-47cf-bd40-73b0371d23b3', 100),
    ('OPT10', 'e62f6003-7392-47cf-bd40-73b0371d23b3', 150),
    ('OPT11', 'e62f6003-7392-47cf-bd40-73b0371d23b3', 100),
    ('OPT12', 'e62f6003-7392-47cf-bd40-73b0371d23b3', 150),
    ('OPT13', 'e62f6003-7392-47cf-bd40-73b0371d23b3', 100),
    ('OPT14', 'e62f6003-7392-47cf-bd40-73b0371d23b3', 150),
    ('OPT15', 'e62f6003-7392-47cf-bd40-73b0371d23b3', 150);
