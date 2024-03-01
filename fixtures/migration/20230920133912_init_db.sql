-- +goose Up
-- +goose StatementBegin
UPDATE "setting" SET value = 'https://site.com' WHERE key = 'site_domain';
UPDATE "setting" SET value = 'Lime license' WHERE key = 'site_name';
UPDATE "setting" SET value = 'Lime Team' WHERE key = 'site_signature';
UPDATE "setting" SET value = 'support@site.com' WHERE key = 'site_email_support';
UPDATE "setting" SET value = 'localhost' WHERE key = 'smtp_host';
UPDATE "setting" SET value = '1025' WHERE key = 'smtp_port';
UPDATE "setting" SET value = 'username' WHERE key = 'smtp_username';
UPDATE "setting" SET value = 'password' WHERE key = 'smtp_password';
UPDATE "setting" SET value = 'None' WHERE key = 'smtp_encryption';
UPDATE "setting" SET value = 'User Name from site' WHERE key = 'mail_sender_name';
UPDATE "setting" SET value = 'admin@mail.com' WHERE key = 'mail_sender_email';

INSERT INTO "customer" ("id", "email", "status") VALUES 
('7v38n58hXHVsNxS', 'user@mail.com', 't');

INSERT INTO "template" ("id", "name", "limit", "price", "term", "hide", "status", "check") VALUES 
('4dDADaT1t0m71Md', 'template 1', '{"servers":5,"companies":5,"users":5}', 500, 'd', 'f', 't', '{"ip":1,"mac":1,"country":1}'),
('t7Kl9LT44Xki3Ix', 'template 2', '{"servers":10,"companies":10,"users":10}', 1000, 'w', 'f', 't', '{"ip":1,"mac":1,"country":1}'),
('0BDbBZ10d3Jb9jz', 'template 3', '{"servers":15,"companies":15,"users":15}', 1500, 'm', 'f', 't', '{"ip":1,"mac":1,"country":1}'),
('AFQG1faC02qc2g6', 'template 4', '{"servers":20,"companies":20,"users":20}', 2000, 'y', 't', 't', '{"ip":1,"mac":1,"country":1}');

INSERT INTO "license" ("id", "template_id", "customer_id", "payment", "type", "status") VALUES 
('jG2Jg83WLw4Dd1l', '4dDADaT1t0m71Md', '7v38n58hXHVsNxS', '{"stripe_id":"cus_APBaLDeqQoVy8m"}', 'test', 't');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM "license";
DELETE FROM "template";
DELETE FROM "customer";
UPDATE "setting" SET value = '' WHERE key = 'smtp_host';
UPDATE "setting" SET value = '' WHERE key = 'smtp_port';
UPDATE "setting" SET value = '' WHERE key = 'smtp_username';
UPDATE "setting" SET value = '' WHERE key = 'smtp_password';
UPDATE "setting" SET value = '' WHERE key = 'smtp_encryption';
UPDATE "setting" SET value = '' WHERE key = 'mail_sender_email';
UPDATE "setting" SET value = '' WHERE key = 'mail_sender_name';
UPDATE "setting" SET value = '' WHERE key = 'site_domain';
UPDATE "setting" SET value = '' WHERE key = 'site_name';
UPDATE "setting" SET value = '' WHERE key = 'site_signature';
UPDATE "setting" SET value = '' WHERE key = 'site_email_support';
-- +goose StatementEnd