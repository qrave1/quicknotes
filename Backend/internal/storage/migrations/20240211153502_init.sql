/* noinspection SqlNoDataSourceInspectionForFile */

-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id  serial PRIMARY KEY,
    username varchar(50)  NOT NULL,
    email    VARCHAR(50)  NOT NULL,
    password varchar(128) NOT NULL
);

CREATE TABLE folders
(
    id serial PRIMARY KEY,
    name      varchar(50) NOT NULL,
    user_id   integer,
    CONSTRAINT fk_folder_user
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);


CREATE TABLE notes
(
    id  serial PRIMARY KEY,
    title    varchar(50) not null,
    body     text        not null,
    folder_id integer,
    CONSTRAINT fk_note_folder
        FOREIGN KEY (folder_id)
            REFERENCES folders (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notes;
DROP TABLE IF EXISTS folders;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
