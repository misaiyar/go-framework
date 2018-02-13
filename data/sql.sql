/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50711
Source Host           : localhost:3306
Source Database       : blog

Target Server Type    : MYSQL
Target Server Version : 50711
File Encoding         : 65001

Date: 2017-08-29 17:17:10
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for system_access
-- ----------------------------
DROP TABLE IF EXISTS `system_access`;
CREATE TABLE `system_access` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `module_id` int(11) NOT NULL COMMENT '模块ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_index` (`role_id`,`module_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='角色权限';

-- ----------------------------
-- Records of system_access
-- ----------------------------
INSERT INTO `system_access` VALUES ('1', '1', '2');
INSERT INTO `system_access` VALUES ('2', '1', '3');
INSERT INTO `system_access` VALUES ('3', '1', '5');
INSERT INTO `system_access` VALUES ('4', '1', '6');
INSERT INTO `system_access` VALUES ('5', '1', '8');
INSERT INTO `system_access` VALUES ('6', '1', '9');
INSERT INTO `system_access` VALUES ('7', '1', '10');
INSERT INTO `system_access` VALUES ('8', '1', '11');
INSERT INTO `system_access` VALUES ('9', '1', '12');
INSERT INTO `system_access` VALUES ('10', '1', '14');
INSERT INTO `system_access` VALUES ('11', '1', '15');
INSERT INTO `system_access` VALUES ('12', '1', '16');
INSERT INTO `system_access` VALUES ('13', '1', '17');
INSERT INTO `system_access` VALUES ('14', '1', '19');
INSERT INTO `system_access` VALUES ('15', '1', '20');
INSERT INTO `system_access` VALUES ('16', '1', '21');
INSERT INTO `system_access` VALUES ('17', '1', '22');

-- ----------------------------
-- Table structure for system_module
-- ----------------------------
DROP TABLE IF EXISTS `system_module`;
CREATE TABLE `system_module` (
  `module_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '模块ID',
  `module_name` varchar(255) NOT NULL COMMENT '模块名',
  `method_name` varchar(255) NOT NULL DEFAULT '' COMMENT '方法名',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '功能描述',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级模块ID',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `show` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否界面显示',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`module_id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8 COMMENT='系统功能模块';

-- ----------------------------
-- Records of system_module
-- ----------------------------
INSERT INTO `system_module` VALUES ('1', 'Admin', '', '基础服务', '0', '1', '1', '1', '1');
INSERT INTO `system_module` VALUES ('2', 'Admin', 'Main', '主界面', '1', '1', '1', '2', '1');
INSERT INTO `system_module` VALUES ('3', 'Admin', 'Logout', '登出系统', '1', '3', '0', '2', '1');
INSERT INTO `system_module` VALUES ('4', 'Module', '', '模块管理', '0', '3', '1', '1', '1');
INSERT INTO `system_module` VALUES ('5', 'Module', 'Index', '功能列表', '4', '1', '1', '2', '1');
INSERT INTO `system_module` VALUES ('6', 'Module', 'Save', '添加/修改模块', '4', '1', '0', '2', '1');
INSERT INTO `system_module` VALUES ('7', 'Role', '', '角色管理', '0', '2', '1', '1', '1');
INSERT INTO `system_module` VALUES ('8', 'Role', 'Index', '角色列表', '7', '1', '1', '2', '1');
INSERT INTO `system_module` VALUES ('9', 'Role', 'Access', '获取角色权限', '7', '1', '0', '2', '1');
INSERT INTO `system_module` VALUES ('10', 'Role', 'SetAccess', '更新角色权限', '7', '1', '0', '2', '1');
INSERT INTO `system_module` VALUES ('11', 'Role', 'Save', '添加/修改角色', '7', '1', '0', '2', '1');
INSERT INTO `system_module` VALUES ('12', 'Role', 'ViewRole', '查看角色用户', '7', '1', '0', '2', '1');
INSERT INTO `system_module` VALUES ('13', 'System', '', '系统管理', '0', '2', '1', '1', '1');
INSERT INTO `system_module` VALUES ('14', 'System', 'Index', '字典管理', '13', '4', '1', '2', '1');
INSERT INTO `system_module` VALUES ('15', 'System', 'Deldictvalue', '删除字典值', '13', '4', '0', '2', '1');
INSERT INTO `system_module` VALUES ('16', 'System', 'Savedictvalue', '添加字典值', '13', '5', '0', '2', '1');
INSERT INTO `system_module` VALUES ('17', 'System', 'Adddict', '添加字典', '13', '2', '0', '2', '1');
INSERT INTO `system_module` VALUES ('18', 'Sysuser', '', '系统用户管理', '0', '1', '1', '1', '1');
INSERT INTO `system_module` VALUES ('19', 'Sysuser', 'Index', '用户列表', '18', '1', '1', '2', '1');
INSERT INTO `system_module` VALUES ('20', 'Sysuser', 'Save', '添加/修改用户', '18', '2', '0', '2', '1');
INSERT INTO `system_module` VALUES ('21', 'Sysuser', 'updateStatus', '禁用/解禁用户', '18', '3', '0', '2', '1');
INSERT INTO `system_module` VALUES ('22', 'Sysuser', 'resetPasswd', '重置用户密码', '18', '4', '0', '2', '1');

-- ----------------------------
-- Table structure for system_role
-- ----------------------------
DROP TABLE IF EXISTS `system_role`;
CREATE TABLE `system_role` (
  `role_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(255) NOT NULL COMMENT '角色名',
  `description` varchar(255) NOT NULL COMMENT '描述',
  `status` tinyint(4) NOT NULL COMMENT '状态',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='系统角色';

-- ----------------------------
-- Records of system_role
-- ----------------------------
INSERT INTO `system_role` VALUES ('1', '超级管理员', '超级管理员', '1');

-- ----------------------------
-- Table structure for system_setting
-- ----------------------------
DROP TABLE IF EXISTS `system_setting`;
CREATE TABLE `system_setting` (
  `k` varchar(32) NOT NULL COMMENT '键',
  `v` text NOT NULL COMMENT '值',
  `d` varchar(100) NOT NULL DEFAULT '' COMMENT '描述',
  PRIMARY KEY (`k`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统设置';

-- ----------------------------
-- Records of system_setting
-- ----------------------------

-- ----------------------------
-- Table structure for system_user
-- ----------------------------
DROP TABLE IF EXISTS `system_user`;
CREATE TABLE `system_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `realname` varchar(255) NOT NULL DEFAULT '' COMMENT '姓名',
  `role_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '角色',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='系统用户';

-- ----------------------------
-- Records of system_user
-- ----------------------------
INSERT INTO `system_user` VALUES ('1', 'admin', '21232f297a57a5a743894a0e4a801fc3', 'admin@misaiyer.com', 'Admin', '1', '1', '2017-08-28 00:00:00');
SET FOREIGN_KEY_CHECKS=1;
