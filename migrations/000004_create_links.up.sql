CREATE TABLE links (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    uuid CHAR(36) NOT NULL,
    title VARCHAR(255) NOT NULL,
    url TEXT NOT NULL,
    url_hash CHAR(64) NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),

    PRIMARY KEY (id),
    UNIQUE INDEX idx_links_uuid (uuid),
    UNIQUE INDEX idx_user_url (user_id, url_hash),

    CONSTRAINT fk_links_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_general_ci;

CREATE TABLE link_tags (
    link_id BIGINT UNSIGNED NOT NULL,
    tag_id BIGINT UNSIGNED NOT NULL,

    PRIMARY KEY (link_id, tag_id),
    INDEX idx_link_tags_tag_id (tag_id),

    CONSTRAINT fk_link_tags_link
        FOREIGN KEY (link_id)
        REFERENCES links(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_link_tags_tag
        FOREIGN KEY (tag_id)
        REFERENCES tags(id)
        ON DELETE CASCADE
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_general_ci;
