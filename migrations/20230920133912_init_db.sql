-- +goose Up
-- +goose StatementBegin
CREATE TABLE "audit" (
  "id" varchar(15) PRIMARY KEY NOT NULL,
  "section" varchar(1) NOT NULL,
  "customer_id" varchar(15) NOT NULL,
  "action" varchar(2) NOT NULL,
  "metadata" json DEFAULT '{"request":{"user_agent":null,"user_ip":null},"data":null}' NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_audit_id ON "audit" ("id");

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

CREATE TABLE "pattern" (
  "id" varchar(15) PRIMARY KEY NOT NULL,
  "name" varchar(255) UNIQUE NOT NULL,
  "limit" json DEFAULT '{}' NOT NULL,
  "term"  varchar(1) NOT NULL,
  "price" varchar(10) NOT NULL,
  "currency" varchar(1) NOT NULL,
  "check" json DEFAULT '{}' NOT NULL,
  "private" bool NOT NULL DEFAULT false,
  "status" bool NOT NULL DEFAULT true,
  "updated_at" timestamp DEFAULT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_pattern_id ON "pattern" ("id");

CREATE TABLE "payment" (
  "id" varchar(15) PRIMARY KEY NOT NULL,
  "pattern_id" varchar(15) NOT NULL,
  "customer_id" varchar(15) NOT NULL,
  "provider" varchar(15) NOT NULL,
  "status" varchar(1) NOT NULL,
  "metadata" json DEFAULT '{}' NOT NULL,
  "updated_at" timestamp DEFAULT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("pattern_id") REFERENCES "pattern"("id") ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY ("customer_id") REFERENCES "customer"("id") ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE INDEX idx_payment_id ON "payment" ("id");
CREATE INDEX idx_payment_pattern_id ON "payment" ("pattern_id");
CREATE INDEX idx_payment_customer_id ON "payment" ("customer_id");

CREATE TABLE "license" (
  "id" varchar(15) PRIMARY KEY NOT NULL,
  "payment_id" varchar(15) NOT NULL,
  "hash" varchar(32) NOT NULL,
  "status" bool NOT NULL DEFAULT true,
  "metadata" json DEFAULT '{}' NOT NULL,
  "updated_at" timestamp DEFAULT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("payment_id") REFERENCES "payment"("id") ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE INDEX idx_license_id ON "license" ("id");
CREATE INDEX idx_license_payment_id ON "license" ("payment_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "license";
DROP TABLE IF EXISTS "payment";
DROP TABLE IF EXISTS "pattern";
DROP TABLE IF EXISTS "customer";
DROP TABLE IF EXISTS "setting";
DROP TABLE IF EXISTS "audit";
-- +goose StatementEnd