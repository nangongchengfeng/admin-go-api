/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.102.20-root
 Source Server Type    : MySQL
 Source Server Version : 50739
 Source Host           : 192.168.102.20:3306
 Source Schema         : admin-go-api

 Target Server Type    : MySQL
 Target Server Version : 50739
 File Encoding         : 65001

 Date: 28/04/2024 21:32:01
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_admin
-- ----------------------------
DROP TABLE IF EXISTS `sys_admin`;
CREATE TABLE `sys_admin`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `post_id`     int(11) NULL DEFAULT NULL COMMENT '岗位id',
    `dept_id`     int(11) NULL DEFAULT NULL COMMENT '部门id',
    `username`    varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账号',
    `password`    varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
    `nickname`    varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '昵称',
    `icon`        varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
    `email`       varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
    `phone`       varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机',
    `note`        varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注信息',
    `create_time` datetime                                               NOT NULL COMMENT '创建时间',
    `status`      int(11) NOT NULL DEFAULT 1 COMMENT '帐号启用状态：1->启用,2->禁用',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `username`(`username`) USING BTREE,
    INDEX         `create_time`(`create_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 98 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_admin
-- ----------------------------
INSERT INTO `sys_admin`
VALUES (89, 1, 3, 'admin', '21232f297a57a5a743894a0e4a801fc3', '南宫乘风',
        'http://127.0.0.1:2002/admin-api-go/upload/2024040028/149834600.jpg', '1794748404@qq.com', '15129460556',
        '后端研发', '2023-05-23 22:15:50', 1);

-- ----------------------------
-- Table structure for sys_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_admin_role`;
CREATE TABLE `sys_admin_role`
(
    `admin_id` int(11) NOT NULL COMMENT '管理员id',
    `role_id`  int(11) NOT NULL COMMENT '角色id'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员和角色关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_admin_role
-- ----------------------------
INSERT INTO `sys_admin_role`
VALUES (89, 1);

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `parent_id`   int(11) NOT NULL COMMENT '父id',
    `dept_type`   int(11) NOT NULL COMMENT '部门类型（1->公司, 2->中心，3->部门）',
    `dept_name`   varchar(30) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '部门名称',
    `dept_status` int(11) NOT NULL DEFAULT 1 COMMENT '部门状态（1->正常 2->停用）',
    `create_time` datetime                                        NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `dept_name`(`dept_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_bin COMMENT = '部门表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept`
VALUES (1, 0, 1, '深圳架构研发中心', 1, '2023-06-14 17:53:23');
INSERT INTO `sys_dept`
VALUES (2, 1, 2, '罗湖区研发中心', 1, '2023-06-14 17:53:55');
INSERT INTO `sys_dept`
VALUES (3, 2, 3, '架构设计部门', 1, '2023-06-14 17:54:15');
INSERT INTO `sys_dept`
VALUES (4, 2, 3, '前端研发部门', 1, '2023-06-14 17:55:17');
INSERT INTO `sys_dept`
VALUES (5, 2, 3, '后端研发部门', 1, '2023-06-14 17:55:25');
INSERT INTO `sys_dept`
VALUES (6, 2, 3, '系统测试部门', 1, '2023-06-14 17:55:31');
INSERT INTO `sys_dept`
VALUES (7, 2, 3, '视觉交互部门', 1, '2023-06-14 17:55:39');
INSERT INTO `sys_dept`
VALUES (8, 2, 3, '产品体验部门', 1, '2023-06-14 17:55:46');
INSERT INTO `sys_dept`
VALUES (11, 2, 3, 'Ownit自研', 1, '2024-04-27 09:33:49');

-- ----------------------------
-- Table structure for sys_login_info
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_info`;
CREATE TABLE `sys_login_info`
(
    `id`             int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `username`       varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '' COMMENT '用户账号',
    `ip_address`     varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '' COMMENT '登录IP地址',
    `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '' COMMENT '登录地点',
    `browser`        varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '' COMMENT '浏览器类型',
    `os`             varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '' COMMENT '操作系统',
    `login_status`   int(11) NULL DEFAULT NULL COMMENT '登录状态（1-成功 2-失败）',
    `message`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '' COMMENT '提示消息',
    `login_time`     datetime NULL DEFAULT NULL COMMENT '访问时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '登录日志记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_login_info
-- ----------------------------
INSERT INTO `sys_login_info`
VALUES (4, '答案是多', '127.0.0.1', '服务器登录', 'Chrome/124.0.0.0', 'Windows 10', 2, '验证码已过期',
        '2024-04-28 18:29:26');
INSERT INTO `sys_login_info`
VALUES (5, '答案是多', '127.0.0.1', '服务器登录', 'Chrome/124.0.0.0', 'Windows 10', 2, '密码不正确',
        '2024-04-28 18:29:58');
INSERT INTO `sys_login_info`
VALUES (6, 'admin', '127.0.0.1', '服务器登录', 'Chrome/124.0.0.0', 'Windows 10', 2, '密码不正确',
        '2024-04-28 18:30:09');
INSERT INTO `sys_login_info`
VALUES (7, 'admin', '127.0.0.1', '服务器登录', 'Chrome/124.0.0.0', 'Windows 10', 1, '登录成功', '2024-04-28 18:30:16');
INSERT INTO `sys_login_info`
VALUES (8, 'admin', '127.0.0.1', '服务器登录', 'Chrome/124.0.0.0', 'Windows 10', 1, '登录成功', '2024-04-28 21:27:22');

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `parent_id`   int(11) NULL DEFAULT NULL COMMENT '父级菜单id',
    `menu_name`   varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单名称',
    `icon`        varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图标',
    `value`       varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '接口权限值',
    `menu_type`   int(11) NULL DEFAULT NULL COMMENT '菜单类型：1->目录；2->菜单；3->按钮（接口绑定权限）',
    `url`         varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单url',
    `menu_status` int(11) NULL DEFAULT 2 COMMENT '启用状态；1->禁用；2->启用',
    `sort`        int(11) NULL DEFAULT NULL COMMENT '排序',
    `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 78 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu`
VALUES (4, 0, '基础管理', 'el-icon-setting', '', 1, '', 2, 2, '2022-09-04 13:57:39');
INSERT INTO `sys_menu`
VALUES (6, 4, '用户信息', 'el-icon-user-solid', 'base:admin:list', 2, 'base/admin', 2, 1, '2022-09-04 13:59:39');
INSERT INTO `sys_menu`
VALUES (7, 4, '角色信息', 'el-icon-s-operation', 'base:role:list', 2, 'base/role', 2, 2, '2022-09-04 14:00:12');
INSERT INTO `sys_menu`
VALUES (8, 4, '菜单信息', 'el-icon-s-grid', 'base:menu:list', 2, 'base/menu', 2, 3, '2022-09-04 14:00:17');
INSERT INTO `sys_menu`
VALUES (9, 4, '部门信息', 'el-icon-s-data', 'base:dept:list', 2, 'base/dept', 2, 4, '2022-09-04 14:01:58');
INSERT INTO `sys_menu`
VALUES (10, 4, '岗位信息', 'el-icon-aim', 'base:post:list', 2, 'base/post', 2, 5, '2022-09-04 14:02:06');
INSERT INTO `sys_menu`
VALUES (16, 6, '新增用户', '', 'base:admin:add', 3, '', 2, 1, '2022-09-04 18:32:55');
INSERT INTO `sys_menu`
VALUES (17, 6, '修改用户', '', 'base:admin:edit', 3, '', 2, 2, '2022-09-04 18:33:29');
INSERT INTO `sys_menu`
VALUES (18, 6, '删除用户', '', 'base:admin:delete', 3, '', 2, 3, '2022-09-04 18:33:51');
INSERT INTO `sys_menu`
VALUES (21, 7, '新增角色', '', 'base:role:add', 3, '', 2, 1, '2022-09-04 18:44:47');
INSERT INTO `sys_menu`
VALUES (22, 7, '修改角色', '', 'base:role:edit', 3, '', 2, 2, '2022-09-04 18:45:10');
INSERT INTO `sys_menu`
VALUES (23, 7, '删除角色', '', 'base:role:delete', 3, '', 2, 3, '2022-09-04 18:45:46');
INSERT INTO `sys_menu`
VALUES (24, 7, '分配权限', '', 'base:role:assign', 3, '', 2, 4, '2022-09-04 18:46:20');
INSERT INTO `sys_menu`
VALUES (26, 8, '新增菜单', '', 'base:menu:add', 3, '', 2, 1, '2022-09-04 18:49:51');
INSERT INTO `sys_menu`
VALUES (27, 8, '修改菜单', '', 'base:menu:edit', 3, '', 2, 2, '2022-09-04 18:50:24');
INSERT INTO `sys_menu`
VALUES (28, 8, '删除菜单', '', 'base:menu:delete', 3, '', 2, 3, '2022-09-04 18:50:53');
INSERT INTO `sys_menu`
VALUES (29, 9, '新增部门', '', 'base:dept:add', 3, '', 2, 1, '2022-09-04 18:52:16');
INSERT INTO `sys_menu`
VALUES (30, 9, '修改部门', '', 'base:dept:edit', 3, '', 2, 2, '2022-09-04 18:52:37');
INSERT INTO `sys_menu`
VALUES (31, 9, '删除部门', '', 'base:dept:delete', 3, '', 2, 3, '2022-09-04 18:52:50');
INSERT INTO `sys_menu`
VALUES (32, 10, '新增岗位', '', 'base:post:add', 3, '', 2, 1, '2022-09-04 18:53:28');
INSERT INTO `sys_menu`
VALUES (33, 10, '修改岗位', '', 'base:post:edit', 3, '', 2, 2, '2022-09-04 18:53:48');
INSERT INTO `sys_menu`
VALUES (34, 10, '删除岗位', '', 'base:post:delete', 3, '', 2, 3, '2022-09-04 18:54:00');
INSERT INTO `sys_menu`
VALUES (44, 0, '日志管理', 'el-icon-open', '', 1, '', 2, 3, '2022-09-05 11:06:57');
INSERT INTO `sys_menu`
VALUES (45, 44, '操作日志', 'el-icon-notebook-2', 'monitor:operator:list', 2, 'monitor/operator', 2, 1,
        '2022-09-05 11:10:54');
INSERT INTO `sys_menu`
VALUES (46, 44, '登录日志', 'el-icon-notebook-1', 'monitor:loginLog:list', 2, 'monitor/loginLog', 2, 2,
        '2022-09-05 11:11:31');
INSERT INTO `sys_menu`
VALUES (47, 45, '清空操作日志', '', 'monitor:operator:clean', 3, '', 2, 1, '2022-09-05 11:12:36');
INSERT INTO `sys_menu`
VALUES (49, 46, '清空登录日志', '', 'monitor:loginLog:clean', 3, '', 2, 1, '2022-09-05 11:16:01');
INSERT INTO `sys_menu`
VALUES (60, 6, '重置密码', NULL, 'base:admin:reset', 3, NULL, 2, 6, '2022-12-01 16:33:34');
INSERT INTO `sys_menu`
VALUES (62, 46, '删除登录日志', '', 'monitor:loginLog:delete', 3, '', 2, 4, '2022-12-02 15:41:56');
INSERT INTO `sys_menu`
VALUES (72, 0, '首页', 'el-icon-s-home', '', 1, 'welcome', 2, 1, '2023-05-24 22:11:13');
INSERT INTO `sys_menu`
VALUES (73, 45, '删除操作日志', '', 'monitor:operator:delete', 3, '', 2, 3, '2023-06-02 10:09:38');

-- ----------------------------
-- Table structure for sys_operation_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_log`;
CREATE TABLE `sys_operation_log`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `admin_id`    int(11) NOT NULL COMMENT '管理员id',
    `username`    varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '管理员账号',
    `method`      varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '请求方式',
    `ip`          varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'IP',
    `url`         varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'URL',
    `create_time` datetime                                               NOT NULL COMMENT '操作时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX         `username`(`username`) USING BTREE,
    INDEX         `create_time`(`create_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '操作日志记录表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_operation_log
-- ----------------------------
INSERT INTO `sys_operation_log`
VALUES (1, 89, 'admin', 'get', '127.0.0.1', '/api/admin/list', '2024-04-28 21:29:30');
INSERT INTO `sys_operation_log`
VALUES (2, 89, 'admin', 'get', '127.0.0.1', '/api/role/vo/list', '2024-04-28 21:29:51');

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
    `post_code`   varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '岗位编码',
    `post_name`   varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '岗位名称',
    `post_status` int(11) NOT NULL DEFAULT 1 COMMENT '状态（1->正常 2->停用）',
    `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
    `remark`      varchar(500) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '备注',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_bin COMMENT = '岗位信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post`
VALUES (1, '2020', '2020', 1, '2023-06-14 20:08:22', '2020');
INSERT INTO `sys_post`
VALUES (2, 'BBB', '经理', 1, '2023-06-14 20:08:45', '主管部分部门');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `role_name`   varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色名称',
    `role_key`    varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色权限字符串',
    `status`      int(11) NOT NULL DEFAULT 1 COMMENT '启用状态：1->启用；2->禁用',
    `description` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '描述',
    `create_time` datetime                                               NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX         `name`(`role_name`) USING BTREE,
    INDEX         `create_time`(`create_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role`
VALUES (1, '超级管理员', 'admin', 1, '最大权限', '2023-06-12 20:04:53');
INSERT INTO `sys_role`
VALUES (9, '管理员', 'admin2', 2, '最大权限', '2024-04-28 10:42:21');

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`
(
    `role_id` int(11) NULL DEFAULT NULL COMMENT '角色ID',
    `menu_id` int(11) NULL DEFAULT NULL COMMENT '菜单ID'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色和菜单关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu`
VALUES (1, 16);
INSERT INTO `sys_role_menu`
VALUES (1, 17);
INSERT INTO `sys_role_menu`
VALUES (1, 18);
INSERT INTO `sys_role_menu`
VALUES (1, 60);
INSERT INTO `sys_role_menu`
VALUES (1, 7);
INSERT INTO `sys_role_menu`
VALUES (1, 21);
INSERT INTO `sys_role_menu`
VALUES (1, 22);
INSERT INTO `sys_role_menu`
VALUES (1, 23);
INSERT INTO `sys_role_menu`
VALUES (1, 24);
INSERT INTO `sys_role_menu`
VALUES (1, 8);
INSERT INTO `sys_role_menu`
VALUES (1, 26);
INSERT INTO `sys_role_menu`
VALUES (1, 27);
INSERT INTO `sys_role_menu`
VALUES (1, 28);
INSERT INTO `sys_role_menu`
VALUES (1, 9);
INSERT INTO `sys_role_menu`
VALUES (1, 29);
INSERT INTO `sys_role_menu`
VALUES (1, 30);
INSERT INTO `sys_role_menu`
VALUES (1, 31);
INSERT INTO `sys_role_menu`
VALUES (1, 10);
INSERT INTO `sys_role_menu`
VALUES (1, 32);
INSERT INTO `sys_role_menu`
VALUES (1, 33);
INSERT INTO `sys_role_menu`
VALUES (1, 34);
INSERT INTO `sys_role_menu`
VALUES (1, 44);
INSERT INTO `sys_role_menu`
VALUES (1, 45);
INSERT INTO `sys_role_menu`
VALUES (1, 47);
INSERT INTO `sys_role_menu`
VALUES (1, 73);
INSERT INTO `sys_role_menu`
VALUES (1, 46);
INSERT INTO `sys_role_menu`
VALUES (1, 49);
INSERT INTO `sys_role_menu`
VALUES (1, 62);
INSERT INTO `sys_role_menu`
VALUES (1, 72);
INSERT INTO `sys_role_menu`
VALUES (1, 4);
INSERT INTO `sys_role_menu`
VALUES (1, 6);
INSERT INTO `sys_role_menu`
VALUES (9, 62);

SET
FOREIGN_KEY_CHECKS = 1;
