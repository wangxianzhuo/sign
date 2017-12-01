CREATE TABLE sign (
    id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    reference_id VARCHAR(36) DEFAULT NULL,
    tag VARCHAR(50) DEFAULT 'signed',
    created_time TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY(id)
);

CREATE TABLE user (
    id VARCHAR(36) NOT NULL,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    PRIMARY KEY(id)
)