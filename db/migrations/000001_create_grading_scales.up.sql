CREATE TABLE IF NOT EXISTS `grading_scales`
(
    `id`                  bigint           NOT NULL AUTO_INCREMENT,
    `grade`               VARCHAR(2)       NOT NULL DEFAULT '' COMMENT 'grade',
    `min_percentage`      FLOAT            NOT NULL DEFAULT 0.0,
    `gpa_value`           FLOAT            NOT NULL DEFAULT 0.0,
    `create_time`         datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
    `update_time`         datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
    PRIMARY KEY (`id`),
    KEY `idx_grade_on_grading_scales` (`grade`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
