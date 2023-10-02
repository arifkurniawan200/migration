-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS source_product (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(255),
    qty INT,
    selling_price DECIMAL,
    promo_price DECIMAL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE source_product;
-- +goose StatementEnd
