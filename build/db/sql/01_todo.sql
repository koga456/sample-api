CREATE TABLE IF NOT EXISTS todo (
    id             BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title          VARCHAR(40) NOT NULL,
    content       VARCHAR(100) NOT NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at     TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
)