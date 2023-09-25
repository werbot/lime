-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
    "id" varchar(32) PRIMARY KEY NOT NULL,
    "login" varchar(255) UNIQUE NOT NULL,
    "email" varchar(255) UNIQUE NOT NULL,
    "password" varchar(255) NOT NULL,
    "enabled" bool NOT NULL DEFAULT true,
    "role" int4 NOT NULL DEFAULT 1,
    "Updated" timestamp DEFAULT NULL,
    "created" timestamp NOT NULL
);
CREATE INDEX idx_user_id ON "user" ("id");
CREATE INDEX idx_user_login ON "user" ("login");
CREATE INDEX idx_user_email ON "user" ("email");
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd