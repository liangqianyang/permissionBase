CREATE TABLE t_role_menus(
    id bigint AUTO_INCREMENT,
    role_id bigint NOT NULL default 0 COMMENT '角色ID',
    menu_id bigint NOT NULL default 0 COMMENT '菜单ID',
    INDEX idx_role_id (role_id),
    INDEX idx_menu_id (menu_id),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT '角色菜单关联表';
