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
('7v38n58hXHVsNxS', 'user@mail.com', 't'),
('8Rz96VZEeHrv1h2', 'user1@mail.com', 't'),
('bkS4BCcQqK01s67', 'user2@mail.com', 't'),
('8Uul6gGZ49P2pzL', 'user3@mail.com', 'f'),
('Y1ycF9Cj0J32fdD', 'user4@mail.com', 't'),
('u5O57U8r6nNRXxo', 'user5@mail.com', 't'),
('m1MeQvVE63q84jJ', 'user6@mail.com', 't'),
('674wsKSt8DdTWk1', 'user7@mail.com', 't'),
('4lLLn3N5jUu74lJ', 'user8@mail.com', 'f'),
('E7sOofi13S2I5eF', 'user9@mail.com', 't'),
('2WwO5T58nNto7gG', 'user10@mail.com', 't');

INSERT INTO "pattern" ("id", "name", "limit", "price", "currency", "term", "private", "status", "check") VALUES 
('4dDADaT1t0m71Md', 'pattern 1', '{"servers":5,"companies":5,"users":5}', 500, 2, 2, 'f', 't', '{"ip":true,"mac":false,"country":true}'),
('t7Kl9LT44Xki3Ix', 'pattern 2', '{"servers":10,"companies":10,"users":10}', 1000, 2,3, 'f', 't', '{"ip":true,"mac":true,"country":true}'),
('0BDbBZ10d3Jb9jz', 'pattern 3', '{"servers":15,"companies":15,"users":15}', 1500, 2, 4, 'f', 't', '{"ip":true,"mac":true,"country":true}'),
('AFQG1faC02qc2g6', 'pattern promo', '{"servers":20,"companies":20,"users":20}', 2000, 2,5, 't', 't', '{"ip":true,"mac":true,"country":true}'),
('cX1ACDa81do05xO', 'pattern 4', '{"servers":99,"companies":99,"users":99}', 2000, 2, 5, 'f', 't', '{"ip":true,"mac":true,"country":true}'),
('A0P0Bwap5c8C0bW', 'pattern 5', '{"servers":99,"companies":1000,"users":20}', 2010, 2,5, 't', 't', '{"ip":false,"mac":true,"country":true}'),
('eZB0E5F921zNfbn', 'pattern 6', '{"servers":20,"companies":99,"users":20}', 2020, 2, 5, 't', 'f', '{"ip":true,"mac":true,"country":true}'),
('tT1lnD7mM44L5dN', 'pattern 7', '{"servers":20,"companies":20,"users":99}', 2030, 2, 5, 't', 'f', '{"ip":true,"mac":false,"country":true}'),
('0X4lLKxC38kTt7c', 'pattern 8', '{"servers":20,"companies":99,"users":1000}', 3000, 2, 5, 'f', 't', '{"ip":true,"mac":true,"country":true}'),
('4lLLn3F2Ggfc6Q4', 'pattern 9', '{"servers":99,"companies":20,"users":20}', 99000, 2, 5, 't', 'f', '{"ip":false,"mac":true,"country":true}'),
('ReU41LFrf2El76u', 'pattern 10', '{"servers":1000,"companies":99,"users":20}', 12010, 2, 5, 'f', 't', '{"ip":true,"mac":true,"country":false}');

INSERT INTO "payment" ("id", "pattern_id", "customer_id", "provider", "status", "metadata") VALUES 
('A0sB9XpSx0P65ab', '4dDADaT1t0m71Md', '7v38n58hXHVsNxS', 'stripe', 1, '{"stripe_id":"cus_APBaLDeqQoVy8m"}'),
('qa4AAQmaMn00N65', 't7Kl9LT44Xki3Ix', '8Rz96VZEeHrv1h2', 'stripe', 1, '{"stripe_id":"cus_APBaLDeqQoVy7m"}'),
('5OkK93H3I2Yyhio', '0BDbBZ10d3Jb9jz', 'bkS4BCcQqK01s67', 'stripe', 1, '{"stripe_id":"cus_APBaLDeqQoVy6m"}'),
('E1gG8gw26rR2eGW', 'AFQG1faC02qc2g6', '8Uul6gGZ49P2pzL', 'stripe', 2, '{"stripe_id":"cus_APBaLDeqQoVy5m"}'),
('Nv8851QD6dnWVwq', 'cX1ACDa81do05xO', 'Y1ycF9Cj0J32fdD', 'stripe', 1, '{"stripe_id":"cus_APBaLDeqQoVy4m"}'),
('4AkwFd0KD12aW8f', '4dDADaT1t0m71Md', 'u5O57U8r6nNRXxo', 'stripe', 1, '{"stripe_id":"cus_APBaLDeqQoVy3m"}'),
('Hve8AaEo5O0hV21', 't7Kl9LT44Xki3Ix', 'm1MeQvVE63q84jJ', 'stripe', 1, '{"stripe_id":"cus_APBaLDeqQoVy2m"}'),
('4klLY5O489xyKXo', 't7Kl9LT44Xki3Ix', '674wsKSt8DdTWk1', 'stripe', 1, '{"stripe_id":"cus_APBaLDeqQoVy1m"}'),
('l63siQ4LI9Xx7Sq', 't7Kl9LT44Xki3Ix', '4lLLn3N5jUu74lJ', 'stripe', 2, '{"stripe_id":"cus_APBaLDeqQoVy0m"}'),
('B066SsSRuUrs76b', '4dDADaT1t0m71Md', 'E7sOofi13S2I5eF', 'stripe', 1, '{"stripe_id":"cus_APBaLDeqQoVy9m"}'),
('FofOoO52oO5H1h5', '4dDADaT1t0m71Md', '2WwO5T58nNto7gG', 'stripe', 1, '{"stripe_id":"cus_APBaLDeqQoVy11"}'),
('9576nYtyRuN8rUT', '4dDADaT1t0m71Md', '2WwO5T58nNto7gG', 'stripe', 4, '{}');

INSERT INTO "license" ("id", "payment_id", "status", "hash") VALUES 
('jG2Jg83WLw4Dd1l', 'A0sB9XpSx0P65ab', 't', '81dc9bdb52d04dc20036dbd8313ed055'),
('4kMV8O54Km7vTto', 'qa4AAQmaMn00N65', 'f', 'c4ca4238a0b923820dcc509a6f75849b'),
('7OU1SGsc2o7g5uC', '5OkK93H3I2Yyhio', 't', 'c20ad4d76fe97759aa27a0c99bff6710'),
('aN5As7SnhHN350n', 'E1gG8gw26rR2eGW', 't', '5455f4c3847d506a0be7afb78eb0e618'),
('45l98Z7LTzpwWPt', 'Nv8851QD6dnWVwq', 't', 'de121459f979a1dafd43ebe02785f0e0'),
('a7tA9nyY3T0JN5j', '4AkwFd0KD12aW8f', 't', 'a3026b0a6849f749c489cd798654a809'),
('ju7J9423mgUxXGM', 'Hve8AaEo5O0hV21', 't', 'a7653fad4df83288ed8888663f8ff585'),
('B36w8WrZzbI9i0R', '4klLY5O489xyKXo', 't', '15de21c670ae7c3f6f3f1f37029303c9'),
('z2kgL9lGZNK454n', 'l63siQ4LI9Xx7Sq', 't', '386854131f58a556343e056f03626e00'),
('qeCEmC00c1MQ4c6', 'B066SsSRuUrs76b', 't', '84adae6beecaa29029addfc1371b29f9'),
('77imoSIM5s3OSs4', '9576nYtyRuN8rUT', 'f', '0a4441b52449ea8b49697438a26be13e');

INSERT INTO "audit" ("id", "section", "customer_id", "action", "metadata") VALUES 
('k3Kmm4R6GM4rMg2', 0, 'admin', 1, '{"request":{"user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4 Safari/605.1.15","user_ip":"127.0.0.1"},"data":null}'),
('bE0oH0nM5L9hY9l', 1, 'admin', 4, '{"request":{"user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4 Safari/605.1.15","user_ip":"127.0.0.1"},"data":null}'),
('14Icm3VziMC89Zv', 2, 'admin', 3, '{"request":{"user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4 Safari/605.1.15","user_ip":"127.0.0.1"},"data":null}'),
('j8VmU4u8u73JMUv', 1, 'admin', 4, '{"request":{"user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4 Safari/605.1.15","user_ip":"127.0.0.1"},"data":null}'),
('LAlM47Ug24Gma0u', 0, '2WwO5T58nNto7gG', 1, '{"request":{"user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4 Safari/605.1.15","user_ip":"127.0.0.1"},"data":{"id":"2WwO5T58nNto7gG"}}');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM "audit";
DELETE FROM "license";
DELETE FROM "payment";
DELETE FROM "pattern";
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