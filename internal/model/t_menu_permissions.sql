CREATE TABLE t_menu_permissions (
    id bigint PRIMARY KEY AUTO_INCREMENT,
    menu_id int NOT NULL default 0 COMMENT '菜单ID',
    permission_id int NOT NULL COMMENT '权限ID',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_menu_id (menu_id),
    INDEX idx_permission_id (permission_id),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT '菜单权限关联表';
