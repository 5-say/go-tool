-- Active: 1727011634447@@mysql@3306@dev_local_platform

DROP TABLE IF EXISTS `jwtx_tokens`;
CREATE TABLE `jwtx_tokens` (
  `id`               bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'token ID',
  `account_id`       bigint unsigned NOT NULL                COMMENT '账户 ID',
  `login_group`      varchar(255)    NOT NULL                COMMENT '登录的分组',
  `login_terminal`   varchar(255)    NOT NULL                COMMENT '登录的终端',
  `make_token_ip`    varchar(50)     NOT NULL                COMMENT '首次请求生成 token 的 IP 地址',
  `created_at`       datetime        NOT NULL                COMMENT '创建时间',
  `last_refresh_at`  datetime        NOT NULL                COMMENT '上次的刷新时间',
  `final_refresh_at` datetime        NOT NULL                COMMENT '最后的刷新时间',
  `expiration_at`    datetime        NOT NULL                COMMENT '过期时间',
  --
  KEY `account_id` (`account_id`),
  KEY `login_group` (`login_group`),
  KEY `login_terminal` (`login_terminal`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='jwtx token 信息表';