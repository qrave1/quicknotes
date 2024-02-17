/* noinspection SqlNoDataSourceInspectionForFile */

-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    user_id  serial PRIMARY KEY,
    username varchar(50)  NOT NULL,
    email    VARCHAR(50)  NOT NULL,
    password varchar(128) NOT NULL
);

CREATE TABLE folders
(
    folder_id serial PRIMARY KEY,
    name      varchar(50) NOT NULL,
    user_id   integer,
    CONSTRAINT fk_group_user
        FOREIGN KEY (user_id)
            REFERENCES users (user_id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);


CREATE TABLE notes
(
    note_id  serial PRIMARY KEY,
    title    varchar(50) not null,
    body     text        not null,
    folder_id integer,
    CONSTRAINT fk_note_group
        FOREIGN KEY (folder_id)
            REFERENCES folders (folder_id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
