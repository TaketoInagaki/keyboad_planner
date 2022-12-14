/* ユーザー */
CREATE TABLE planner.user (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `name` VARCHAR(20) NOT NULL COMMENT 'ユーザー名',
    `password` VARCHAR(80) NOT NULL COMMENT 'パスワードハッシュ',
    `role` VARCHAR(80) NOT NULL COMMENT 'ロール',
    `delete_flg` INT(4) NOT NULL default(0) COMMENT '削除フラグ',
    `created` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'ユーザー';

-- INSERT INTO
--     user(
--         id,
--         name,
--         password,
--         role,
--         delete_flg,
--         created,
--         modified
--     )
-- VALUES
--     (
--         1,
--         '織田信長',
--         '$2a$10$SfNkdmsCDglWXC0eSbvsPeccgvUB49xxwSBYgVL4LrBShLs.f5f26',
--         'user',
--         0,
--         '2022-12-07 12:50:00.558124',
--         '2022-12-07 12:50:00.558124'
--     );

-- INSERT INTO
--     user(
--         id,
--         name,
--         password,
--         role,
--         delete_flg,
--         created,
--         modified
--     )
-- VALUES
--     (
--         2,
--         '豊臣秀吉',
--         '$2a$10$SfNkdmsCDglWXC0eSbvsPeccgvUB49xxwSBYgVL4LrBShLs.f5f26',
--         'user',
--         0,
--         '2022-12-07 12:50:00.558124',
--         '2022-12-07 12:50:00.558124'
--     );

-- INSERT INTO
--     user(
--         id,
--         name,
--         password,
--         role,
--         delete_flg,
--         created,
--         modified
--     )
-- VALUES
--     (
--         3,
--         '徳川家康',
--         '$2a$10$SfNkdmsCDglWXC0eSbvsPeccgvUB49xxwSBYgVL4LrBShLs.f5f26',
--         'user',
--         0,
--         '2022-12-07 12:50:00.558124',
--         '2022-12-07 12:50:00.558124'
--     );

/* タスク */
CREATE TABLE planner.task (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'タスクの識別子',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT 'タスクを作成したユーザーの識別子',
    `title` VARCHAR(128) NOT NULL COMMENT 'タスクのタイトル',
    `date` DATETIME(6) NOT NULL COMMENT 'タスク日程',
    `date_type` VARCHAR(20) NOT NULL COMMENT 'タスク日程の種類',
    `week_number` SMALLINT NOT NULL default(0) COMMENT 'タスクの週数',
    `delete_flg` INT(4) NOT NULL default(0) COMMENT '削除フラグ',
    `created` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'タスク';

-- INSERT INTO
--     user(
--         id,
--         user_id,
--         title,
--         date,
--         date_type,
--         week_number,
--         delete_flg,
--         created,
--         modified
--     )
-- VALUES
--     (
--         1,
--         1,
--         '牛乳買う',
--         '2022-12',
--         'Monthly',
--         0,
--         0,
--         '2022-12-07 12:50:00.558124',
--         '2022-12-07 12:50:00.558124'
--     );

-- INSERT INTO
--     user(
--         id,
--         user_id,
--         title,
--         date,
--         date_type,
--         week_number,
--         delete_flg,
--         created,
--         modified
--     )
-- VALUES
--     (
--         2,
--         1,
--         '洗濯物干す',
--         '2022',
--         'Weekly',
--         5,
--         0,
--         '2022-12-07 12:50:00.558124',
--         '2022-12-07 12:50:00.558124'
--     );

-- INSERT INTO
--     user(
--         id,
--         user_id,
--         title,
--         date,
--         date_type,
--         week_number,
--         delete_flg,
--         created,
--         modified
--     )
-- VALUES
--     (
--         3,
--         1,
--         'アメリカ旅行に行く',
--         '2022',
--         'Yearly',
--         0,
--         0,
--         '2022-12-07 12:50:00.558124',
--         '2022-12-07 12:50:00.558124'
--     );

/* 振り返り */
CREATE TABLE planner.reflection (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '振り返りの識別子',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '振り返りを作成したユーザーの識別子',
    `content` VARCHAR(255) NOT NULL COMMENT '振り返り内容',
    `content_type` VARCHAR(20) NOT NULL COMMENT '振り返りの種類',
    `date` DATETIME(6) NOT NULL COMMENT '振り返り日程',
    `date_type` VARCHAR(20) NOT NULL COMMENT '振り返り日程の種類',
    `week_number` SMALLINT NOT NULL default(0) COMMENT '振り返りの週数',
    `delete_flg` INT(4) NOT NULL default(0) COMMENT '削除フラグ',
    `created` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_reflection_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '振り返り';

-- INSERT INTO
--     user(
--         id,
--         user_id,
--         content,
--         content_type,
--         date,
--         date_type,
--         week_number,
--         delete_flg,
--         created,
--         modified
--     )
-- VALUES
--     (
--         1,
--         1,
--         '睡眠の質が悪い',
--         'Check',
--         '2022-12-07',
--         'Daily',
--         0,
--         0,
--         '2022-12-07 12:50:00.558124',
--         '2022-12-07 12:50:00.558124'
--     );

-- INSERT INTO
--     user(
--         id,
--         user_id,
--         content,
--         content_type,
--         date,
--         date_type,
--         week_number,
--         delete_flg,
--         created,
--         modified
--     )
-- VALUES
--     (
--         2,
--         1,
--         '日光を浴びるために朝散歩する',
--         'Action',
--         '2022-12-07',
--         'Daily',
--         0,
--         0,
--         '2022-12-07 12:50:00.558124',
--         '2022-12-07 12:50:00.558124'
--     );

-- INSERT INTO
--     user(
--         id,
--         user_id,
--         content,
--         content_type,
--         date,
--         date_type,
--         week_number,
--         delete_flg,
--         created,
--         modified
--     )
-- VALUES
--     (
--         3,
--         1,
--         'メダリスト第３巻激アツ',
--         'Note',
--         '2022-12-07',
--         'Daily',
--         0,
--         0,
--         '2022-12-07 12:50:00.558124',
--         '2022-12-07 12:50:00.558124'
--     );

/* 継続リスト */
CREATE TABLE planner.continuation (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '継続リストの識別子',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '継続リストを作成したユーザーの識別子',
    `content` VARCHAR(255) NOT NULL COMMENT '継続リスト内容',
    `content_type` VARCHAR(20) NOT NULL COMMENT '継続リストの種類',
    `delete_flg` INT(4) NOT NULL default(0) COMMENT '削除フラグ',
    `created` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_continuation_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '継続リスト';

-- INSERT INTO
--     user(
--         id,
--         user_id,
--         content,
--         content_type,
--         delete_flg,
--         created,
--         modified
--     )
-- VALUES
--     (
--         1,
--         1,
--         '朝マック',
--         'Continue',
--         0,
--         '2022-12-07 12:50:00.558124',
--         '2022-12-07 12:50:00.558124'
--     );

-- INSERT INTO
--     user(
--         id,
--         user_id,
--         content,
--         content_type,
--         delete_flg,
--         created,
--         modified
--     )
-- VALUES
--     (
--         2,
--         1,
--         'アプリ開発',
--         'Begin',
--         0,
--         '2022-12-07 12:50:00.558124',
--         '2022-12-07 12:50:00.558124'
--     );

-- INSERT INTO
--     user(
--         id,
--         user_id,
--         content,
--         content_type,
--         delete_flg,
--         created,
--         modified
--     )
-- VALUES
--     (
--         3,
--         1,
--         '英語学習',
--         'Quit',
--         0,
--         '2022-12-07 12:50:00.558124',
--         '2022-12-07 12:50:00.558124'
--     );