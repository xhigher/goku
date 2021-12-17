CREATE DATABASE IF NOT EXISTS goku_user DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;
CREATE DATABASE IF NOT EXISTS goku_game DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;

use goku_user;

CREATE TABLE `user_detail` (
  `mid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `sex` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  `ct` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `ut` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`mid`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4; 

CREATE TABLE `user_address` (
  `addr_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `mid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `name` varchar(100) NOT NULL DEFAULT '',
  `phoneno` varchar(50)  NOT NULL DEFAULT '',
  `locations` varchar(20)  NOT NULL DEFAULT '',
  `ut` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`addr_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4; 

insert into user_detail (`name`,`sex`, `status`,`ct`, `ut`) VALUES
('xhigher',1, 1, 1639559202, 1639559202),
('goku',0, 1, 1639559300, 1639559300);

insert into user_address (`mid`,`name`, `phoneno`,`locations`, `ut`) VALUES
(10000,'Tom', '13715789456', '广东省 深圳市 南山区 滨海大厦1001', 1639559202),
(10000,'Jack', '13715789123', '广东省 深圳市 南山区 前海大厦520', 1639559300);

use goku_game;

CREATE TABLE `game_config` (
  `game_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `icons` varchar(200) NOT NULL DEFAULT '',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  `ct` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `ut` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`game_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6000 DEFAULT CHARSET=utf8mb4;