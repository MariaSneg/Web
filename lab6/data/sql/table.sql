CREATE TABLE `blog`.`post`
(
    `id`           INT          NOT NULL AUTO_INCREMENT,
    `title`        VARCHAR(255) NOT NULL,
    `subtitle`     VARCHAR(255) NULL,
    `author`       VARCHAR(255) NOT NULL,
    `author_url`   VARCHAR(255) NULL,
    `publish_date` VARCHAR(255) NULL,
    `image_url`    VARCHAR(255) NULL,
    `featured`     TINYINT(1) NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE
) ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;