-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO Preview(url)
VALUES('https://imgur.com/J71bAQJ'),
('https://imgur.com/0XEanXw'),
('https://imgur.com/C8uoUMm'),
('https://imgur.com/CkXnahk'),
('https://imgur.com/ZKjHry7');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
