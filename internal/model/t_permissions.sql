CREATE TABLE t_permissions (
 id bigint AUTO_INCREMENT,
 name varchar(32) NOT NULL DEFAULT '' COMMENT '权限名称',
 identify_name varchar(32) NOT NULL DEFAULT '' COMMENT '权限标识',
 controller varchar(32) NOT NULL DEFAULT '' COMMENT '控制器',
 action varchar(32) NOT NULL DEFAULT '' COMMENT '方法',
 description varchar(128) NOT NULL DEFAULT '' COMMENT '权限描述',
 state tinyint NOT NULL DEFAULT '1' COMMENT '状态 1:正常 -1:禁用',
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 INDEX idx_name (name),
 UNIQUE uk_identify_name (identify_name),
 PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT '权限表';