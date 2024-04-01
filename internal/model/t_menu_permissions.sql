CREATE TABLE t_menu_permissions(
   id bigint AUTO_INCREMENT,
   menu_id bigint NOT NULL default 0 COMMENT '菜单ID',
   permission_id bigint NOT NULL default 0 COMMENT '权限ID',
   INDEX idx_menu_id (menu_id),
   INDEX idx_role_id (permission_id),
   PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT '菜单权限关联表';