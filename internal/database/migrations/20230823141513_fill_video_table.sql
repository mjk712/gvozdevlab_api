-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO Video(url)
VALUES('https://www.youtube.com/watch?v=dQw4w9WgXcQ'),
('https://www.youtube.com/watch?v=D220J3AR-E8&t=27s'),
('https://www.youtube.com/watch?v=JkHHWlm6Cns'),
('https://www.youtube.com/watch?v=Sagg08DrO5U'),
('https://www.youtube.com/watch?v=VZTnBXAwuUA&list=PL5pQf7svQVJWngGEG7P84MvzcRZcIrvc_');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd