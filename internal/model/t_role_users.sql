CREATE TABLE t_role_users(
   id bigint AUTO_INCREMENT,
   role_id bigint NOT NULL default 0 COMMENT '角色ID',
   user_id bigint NOT NULL default 0 COMMENT '用户ID',
   INDEX idx_role_id (role_id),
   INDEX idx_user_id (user_id),
   PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT '角色用户关联表';