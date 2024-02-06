CREATE TABLE t_users (
 id bigint AUTO_INCREMENT,
 login_name varchar(64) NOT NULL DEFAULT '' COMMENT '登录名',
 nickname varchar(128) NOT NULL DEFAULT '' COMMENT '昵称',
 password varchar(128) NOT NULL DEFAULT '' COMMENT '登录密码',
 mobile varchar(32) NOT NULL DEFAULT '' COMMENT '手机号码',
 email varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱',
 gender tinyint NOT NULL DEFAULT '0' COMMENT '性别 0:未知 1:男 2:女',
 state tinyint NOT NULL DEFAULT '1' COMMENT '状态 1:正常 -1:禁用',
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 UNIQUE uniq_mobile (mobile),
 UNIQUE uniq_email (email),
 UNIQUE uniq_login_name (login_name(10)),
 PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT '用户表';