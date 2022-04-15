-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE "surveys" (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "brand" varchar(255) NOT NULL,
    "template" uuid NOT NULL, 
    "name" varchar NOT NULL,
    "mothodology" int NOT NULL,
    "status" int NOT NULL,
    "description" TEXT,
    "response_redirect_url" varchar(220),
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS surveys;
-- +goose StatementEnd
