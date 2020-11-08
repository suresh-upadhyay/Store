CREATE TABLE Product
(
    ID SERIAL,
    Name TEXT NOT NULL,
    Price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)

CREATE TABLE Store
(
    StoreId s_id SERIAL,
    ProductId SERIAL,
    IsAvailable boolean
 )