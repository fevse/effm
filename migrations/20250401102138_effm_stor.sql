-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS people (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    patronymic TEXT,
    age INT,
    sex TEXT,
    nationality TEXT
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX name_surname
ON people (name, surname);
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO people (name, surname)
VALUES ('Daria', 'Morgendorffer'),
('Jane', 'Lane');
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO people (name, surname, patronymic)
VALUES ('Bender', 'Rodriguez', 'Bending');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS people;
-- +goose StatementEnd
