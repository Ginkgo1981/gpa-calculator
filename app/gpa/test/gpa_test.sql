DROP TABLE IF EXISTS `grading_scales`;
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

DROP TABLE IF EXISTS `student_grades`;
CREATE TABLE `student_grades`
(
    `id`                  bigint           NOT NULL AUTO_INCREMENT,
    `student_id`          bigint           NOT NULL DEFAULT 0,
    `course_id`           bigint           NOT NULL DEFAULT 0,
    `grade_received`      VARCHAR(2)       NOT NULL DEFAULT '',
    `create_time`         datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
    `update_time`         datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
    PRIMARY KEY (`id`),
    KEY `idx_student_id_on_grades` (`student_id`),
    KEY `idx_course_id_on_grades` (`course_id`),
    KEY `idx_grade_received_on_grades` (`grade_received`),
    CONSTRAINT UNIQUE `uc_student_id_course_id_on_grades` (`student_id`, `course_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
