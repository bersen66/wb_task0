DROP DOMAIN IF EXISTS LOCALE;
DROP TABLE IF EXISTS orders;

CREATE DOMAIN LOCALE
    AS VARCHAR(25) NOT NULL
    CHECK (
            VALUE IN (
                      'en',
                      'ru'
            ) );



CREATE TABLE orders
(
    order_uid          uuid PRIMARY KEY,
    track_number       VARCHAR(25) UNIQUE NOT NULL,
    entry              VARCHAR(25)        NOT NULL,
    delivery           JSON               NOT NULL,
    payment            JSON               NOT NULL,
    items              JSON               NOT NULL,
    locale             LOCALE             NOT NULL,
    internal_signature VARCHAR(15)        NOT NULL,
    customer_id        VARCHAR(15)        NOT NULL, -- FOREIGN KEY
    delivery_service   VARCHAR(15)        NOT NULL,
    shardkey           BIGINT             NOT NULL,
    sm_id              BIGINT             NOT NULL,
    date_created       TIMESTAMPTZ        NOT NULL,
    oof_shard          BIGINT             NOT NULL
);

SELECT *
FROM orders;