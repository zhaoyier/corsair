
CREATE TABLE `gd_niusan` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
	`security_code` INT(11) NOT NULL DEFAULT '0',
	`niusan` VARCHAR(100) NOT NULL DEFAULT '',
	`disabled` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0',
	`create_date` BIGINT(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
	`update_date` BIGINT(20) NOT NULL DEFAULT '0' COMMENT '更新时间',
	PRIMARY KEY(`id`,`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'gd_niusan';
CREATE INDEX `security_code_of_gdniusan_idx` ON `gd_niusan`(`security_code`);
CREATE INDEX `niusan_of_gdniusan_idx` ON `gd_niusan`(`niusan`);

