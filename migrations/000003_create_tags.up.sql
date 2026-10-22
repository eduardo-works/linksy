CREATE TABLE tags (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    uuid CHAR(36) NOT NULL,
    title VARCHAR(255) NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    color CHAR(6) NOT NULL DEFAULT '9FBE9A',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),

    PRIMARY KEY (id),
    UNIQUE INDEX idx_tags_uuid (uuid),
    UNIQUE INDEX idx_user_tag_title (user_id, title),

    CONSTRAINT fk_tags_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_general_ci;

CREATE TABLE page_tags (
    page_id BIGINT UNSIGNED NOT NULL,
    tag_id BIGINT UNSIGNED NOT NULL,

    PRIMARY KEY (page_id, tag_id),
    INDEX idx_page_tags_tag_id (tag_id),

    CONSTRAINT fk_page_tags_page
        FOREIGN KEY (page_id)
        REFERENCES pages(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_page_tags_tag
        FOREIGN KEY (tag_id)
        REFERENCES tags(id)
        ON DELETE CASCADE
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_general_ci;
