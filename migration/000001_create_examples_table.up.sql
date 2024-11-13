CREATE TABLE IF NOT EXISTS `examples`
(
    `id`          bigint unsigned     NOT NULL AUTO_INCREMENT,
    `field1`      varchar(255)        not null default '',
    `field2`      bigint unsigned     not null default '0',
    `field3`      tinyint(1) unsigned not null default '0',
    `field4`      tinyint(2) unsigned NOT NULL default '1',
    `create_time` bigint unsigned     NOT NULL DEFAULT '0',
    `update_time` bigint unsigned     NOT NULL DEFAULT '0',
    `delete_time` bigint unsigned     NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
