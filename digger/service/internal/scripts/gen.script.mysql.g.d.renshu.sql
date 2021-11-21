
CREATE TABLE `gd_renshu` (
	`id` BIGINT(20) NOT NULL DEFAULT '0',
	`secucode` VARCHAR(100) NOT NULL DEFAULT '',
	`security_code` INT(11) NOT NULL DEFAULT '0',
	`end_date` VARCHAR(100) NOT NULL DEFAULT '',
	`holder_total_num` FLOAT NOT NULL DEFAULT '0' COMMENT '总人数',
	`total_num_ratio` FLOAT NOT NULL DEFAULT '0',
	`avg_free_shares` FLOAT NOT NULL DEFAULT '0',
	`avg_freeshares_ratio` FLOAT NOT NULL DEFAULT '0',
	`hold_focus` VARCHAR(100) NOT NULL DEFAULT '',
	`price` FLOAT NOT NULL DEFAULT '0',
	`avg_hold_amt` FLOAT NOT NULL DEFAULT '0',
	`hold_ratio_total` FLOAT NOT NULL DEFAULT '0',
	`freehold_ratio_total` FLOAT NOT NULL DEFAULT '0',
	`create_date` BIGINT(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
	`update_date` BIGINT(20) NOT NULL DEFAULT '0' COMMENT '更新时间',
	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'gd_renshu';
CREATE INDEX `security_code_of_gdrenshu_idx` ON `gd_renshu`(`security_code`);
CREATE INDEX `secucode_of_gdrenshu_idx` ON `gd_renshu`(`secucode`);
CREATE INDEX `holder_total_num_of_gdrenshu_idx` ON `gd_renshu`(`holder_total_num`);
CREATE INDEX `end_date_of_gdrenshu_idx` ON `gd_renshu`(`end_date`);

