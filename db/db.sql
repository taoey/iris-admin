CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `uname` varchar(100) DEFAULT '' COMMENT '用户名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0-无效 1有效',
  `email` varchar(200) NOT NULL DEFAULT '' COMMENT '邮箱',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';


CREATE TABLE `auth` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `auth_id` tinyint(4) DEFAULT '0' COMMENT '权限ID',
  `auth_name` varchar(100) DEFAULT '' COMMENT '权限名称',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0-无效 1有效',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_id` (`auth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限表';

CREATE TABLE `use_auth` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `uid` bigint(20) DEFAULT '0' COMMENT '用户ID',
  `uname` varchar(100) DEFAULT '' COMMENT '用户名',
  `auth_id` tinyint(4) DEFAULT '0' COMMENT '权限ID',
  `auth_name` varchar(100) DEFAULT '' COMMENT '权限名称',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0-无效 1有效',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户权限表';