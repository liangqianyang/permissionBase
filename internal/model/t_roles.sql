CREATE TABLE t_roles (
 id bigint AUTO_INCREMENT,
 name varchar(32) NOT NULL DEFAULT '' COMMENT '角色名称',
 description varchar(128) NOT NULL DEFAULT '' COMMENT '角色描述',
 state tinyint NOT NULL DEFAULT '1' COMMENT '状态 1:正常 -1:禁用',
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 INDEX idx_name (name),
 PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT '角色表';