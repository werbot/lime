-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
    "id" varchar(32) PRIMARY KEY NOT NULL,
    "login" varchar(255) UNIQUE NOT NULL,
    "email" varchar(255) UNIQUE NOT NULL,
    "password" varchar(255) NOT NULL,
    "status" bool NOT NULL DEFAULT true,
    "role" varchar(5) CHECK ("role" == 'admin' OR "role" == 'user'),
    "updated_at" timestamp DEFAULT NULL,
    "created_at" timestamp NOT NULL
);
CREATE INDEX idx_user_id ON "user" ("id");
CREATE INDEX idx_user_login ON "user" ("login");
CREATE INDEX idx_user_email ON "user" ("email");
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd