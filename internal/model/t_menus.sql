CREATE TABLE t_menus (
 id bigint AUTO_INCREMENT,
 parent_id bigint NOT NULL DEFAULT '0' COMMENT '父级菜单id',
 name varchar(32) NOT NULL DEFAULT '' COMMENT '菜单名称',
 icon varchar(32) NOT NULL DEFAULT '' COMMENT '菜单图标',
 path varchar(128) NOT NULL DEFAULT '' COMMENT '菜单地址',
 sort int NOT NULL DEFAULT '0' COMMENT '排序',
 state tinyint NOT NULL DEFAULT '1' COMMENT '状态 1:正常 -1:禁用',
 created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 INDEX idx_name (name),
 PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT '菜单表';