CREATE TABLE t_menus (
     id bigint AUTO_INCREMENT,
     parent_id bigint NOT NULL DEFAULT '0' COMMENT '父级菜单id',
     name varchar(32) NOT NULL DEFAULT '' COMMENT '菜单名称',
     icon varchar(32) NOT NULL DEFAULT '' COMMENT '菜单图标',
     url varchar(128) NOT NULL DEFAULT '' COMMENT '菜单url',
     menu_type tinyint NOT NULL DEFAULT '1' COMMENT '菜单类型 1:菜单 2:按钮',
     sort int NOT NULL DEFAULT '0' COMMENT '排序',
     state tinyint NOT NULL DEFAULT '1' COMMENT '状态 1:正常 -1:禁用',
     operator_id bigint NOT NULL DEFAULT '0' COMMENT '操作人id',
     operator varchar(32) NOT NULL DEFAULT '' COMMENT '操作人名称',
     created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
     updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     INDEX idx_name (name),
     INDEX idx_parent_id (parent_id),
     PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT '菜单表';