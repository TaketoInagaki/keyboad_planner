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

/* タスク */
CREATE TABLE planner.task (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'タスクの識別子',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT 'タスクを作成したユーザーの識別子',
    `title` VARCHAR(128) NOT NULL COMMENT 'タスクのタイトル',
    `status` INT(4) NOT NULL default(0) COMMENT 'タスクのステータス',
    `date` DATETIME(6) NOT NULL COMMENT 'タスク日程',
    `date_type` VARCHAR(20) NOT NULL COMMENT 'タスク日程の種類',
    `week_number` SMALLINT NOT NULL default(0) COMMENT 'タスクの週数',
    `delete_flg` INT(4) NOT NULL default(0) COMMENT '削除フラグ',
    `created` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'タスク';

/* 振り返り */
CREATE TABLE planner.reflection (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '振り返りの識別子',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '振り返りを作成したユーザーの識別子',
    `content` VARCHAR(255) NOT NULL COMMENT '振り返り内容',
    `date` DATETIME(6) NOT NULL COMMENT '振り返り日程',
    `date_type` VARCHAR(20) NOT NULL COMMENT '振り返り日程の種類',
    `week_number` SMALLINT NOT NULL default(0) COMMENT '振り返りの週数',
    `delete_flg` INT(4) NOT NULL default(0) COMMENT '削除フラグ',
    `created` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_reflection_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '振り返り';

/* 振り返り */
CREATE TABLE planner.checklist (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'チェックの識別子',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT 'チェックを作成したユーザーの識別子',
    `content` VARCHAR(255) NOT NULL COMMENT 'チェック内容',
    `date` DATETIME(6) NOT NULL COMMENT 'チェック日程',
    `date_type` VARCHAR(20) NOT NULL COMMENT 'チェック日程の種類',
    `week_number` SMALLINT NOT NULL default(0) COMMENT 'チェックの週数',
    `delete_flg` INT(4) NOT NULL default(0) COMMENT '削除フラグ',
    `created` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_check_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'チェック';

/* 振り返り */
CREATE TABLE planner.actionlist (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'アクションの識別子',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT 'アクションを作成したユーザーの識別子',
    `content` VARCHAR(255) NOT NULL COMMENT 'アクション内容',
    `status` INT(4) NOT NULL default(0) COMMENT 'アクションのステータス',
    `date` DATETIME(6) NOT NULL COMMENT 'アクション日程',
    `date_type` VARCHAR(20) NOT NULL COMMENT 'アクション日程の種類',
    `week_number` SMALLINT NOT NULL default(0) COMMENT 'アクションの週数',
    `delete_flg` INT(4) NOT NULL default(0) COMMENT '削除フラグ',
    `created` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_action_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'アクション';

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

/* やりたいことリスト */
CREATE TABLE planner.wish (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'やりたいことリストの識別子',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT 'やりたいことリストを作成したユーザーの識別子',
    `content` VARCHAR(255) NOT NULL COMMENT 'やりたいことリスト内容',
    `delete_flg` INT(4) NOT NULL default(0) COMMENT '削除フラグ',
    `created` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_wish_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'やりたいことリスト';