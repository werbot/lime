-- +goose Up
-- +goose StatementBegin
CREATE TABLE "setting" (
	"id"     TEXT PRIMARY KEY NOT NULL,
	"key"    TEXT UNIQUE NOT NULL,
	"value"  TEXT DEFAULT NULL
);
CREATE INDEX idx_setting_key ON "setting" ("key");
INSERT INTO "setting" VALUES ('klA5PpR36La0Kr4', 'installed', 0);
INSERT INTO "setting" VALUES ('aAW8wL4Dd01Lll4', 'site_domain', '');
INSERT INTO "setting" VALUES ('Aj1WWd0JDw3w88a', 'site_name', '');
INSERT INTO "setting" VALUES ('7ooOSOO5os5N55n', 'site_signature', '');
INSERT INTO "setting" VALUES ('jJA0UuTIa337i7t', 'site_email_support', '');
INSERT INTO "setting" VALUES ('BCc8A0whWaH3b00', 'smtp_host', '');
INSERT INTO "setting" VALUES ('zOZ71doNtD5Tn95', 'smtp_port', '0');
INSERT INTO "setting" VALUES ('9kkA0aK3KXx3Yy9', 'smtp_username', '');
INSERT INTO "setting" VALUES ('Sq51dDsQoO61D6d', 'smtp_password', '');
INSERT INTO "setting" VALUES ('VfI89zlF3vLiZ42', 'smtp_encryption', '');
INSERT INTO "setting" VALUES ('7v38n58hXHVsNxS', 'mail_sender_name', '');
INSERT INTO "setting" VALUES ('6kK99i33PZXzpxI', 'mail_sender_email', '');
INSERT INTO "setting" VALUES ('6lABbUqQ70Lau40', 'mail_letter_access_link', '{"subject":"License Manager access link","text":"Hi,\n\nHere is your access link to manage your license key, it will be expired in {{.Expire}}:\n\n{{.Domain}}/signin?token={{.Token}}\n\nIf you did not requested for this link, please contact us immediately. Your license key may has been compromised.\nThank you again for your support!\n\nEmail Support: {{.EmailSupport}}\n\nBest regard,\n{{.Signature}}","html":""}');

CREATE TABLE "customer" (
  "id" varchar(15) PRIMARY KEY NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "status" bool NOT NULL DEFAULT true,
  "updated_at" timestamp DEFAULT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_customer_id ON "customer" ("id");
CREATE INDEX idx_customer_email ON "customer" ("email");

CREATE TABLE "template" (
  "id" varchar(15) PRIMARY KEY NOT NULL,
  "name" varchar(255) UNIQUE NOT NULL,
  "limit" json DEFAULT '{}' NOT NULL,
  "term"  varchar(1) NOT NULL CHECK ("term" = 'd' OR "term" = 'w' OR "term" = 'm' OR "term" = 'y'),
  "price" varchar(10) NOT NULL,
  "check" json DEFAULT '{}' NOT NULL,
  "hide" bool NOT NULL DEFAULT false,
  "status" bool NOT NULL DEFAULT true,
  "updated_at" timestamp DEFAULT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_template_id ON "template" ("id");

CREATE TABLE "license" (
  "id" varchar(15) PRIMARY KEY NOT NULL,
  "template_id" varchar(15) NOT NULL,
  "customer_id" varchar(15) NOT NULL,
  "payment" json DEFAULT '{}' NOT NULL,
  "type" varchar(15) NOT NULL,
  "status" bool NOT NULL DEFAULT true,
  "metadata" json DEFAULT '{}' NOT NULL,
  "updated_at" timestamp DEFAULT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("template_id") REFERENCES "template"("id") ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY ("customer_id") REFERENCES "customer"("id") ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE INDEX idx_license_id ON "license" ("id");
CREATE INDEX idx_license_template_id ON "license" ("template_id");
CREATE INDEX idx_license_customer_id ON "license" ("customer_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "license";
DROP TABLE IF EXISTS "template";
DROP TABLE IF EXISTS "customer";
DROP TABLE IF EXISTS "setting";
-- +goose StatementEnd