CREATE TABLE t_role_menus (
    id bigint PRIMARY KEY AUTO_INCREMENT,
    role_id int NOT NULL default 0 COMMENT '角色ID',
    menu_id int NOT NULL COMMENT '菜单ID',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_role_id (role_id),
    INDEX idx_menu_id (menu_id),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT '角色菜单关联表';
