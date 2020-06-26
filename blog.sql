/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50643
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 50643
 File Encoding         : 65001

 Date: 26/06/2020 11:03:29
*/

CREATE DATABASE blog;
USE blog;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_admin
-- ----------------------------
DROP TABLE IF EXISTS `tb_admin`;
CREATE TABLE `tb_admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` char(60) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_admin
-- ----------------------------
BEGIN;
INSERT INTO `tb_admin` VALUES (1, 'admin@qq.com', '$2a$10$p0S0FckuXVASxXHXqO00Mez1rZl3wvn6jbRraOTVvhqilKbjb4vtC');
COMMIT;

-- ----------------------------
-- Table structure for tb_article
-- ----------------------------
DROP TABLE IF EXISTS `tb_article`;
CREATE TABLE `tb_article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_id` int(11) NOT NULL DEFAULT '0',
  `title` varchar(255) NOT NULL DEFAULT '',
  `content` mediumtext NOT NULL,
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '0:正常 1:删除',
  `top` tinyint(2) NOT NULL DEFAULT '1' COMMENT '1: 正常 2:置顶',
  `depiction` text NOT NULL COMMENT '描述',
  `image` varchar(50) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_category_id` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_article
-- ----------------------------
BEGIN;
INSERT INTO `tb_article` VALUES (1, 1, 'go语言gin框架写的个人博客', '&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	gin写的个人博客, github地址(&lt;a href=&#34;https://github.com/yangwanbo/myblog&#34; target=&#34;_blank&#34;&gt;https://github.com/luckywb13/blog&lt;/a&gt;),如有问题请多多指教.\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	博客db:mysql5.6, v&amp;gt;=5.6\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	使用方式：\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	git clone https://github.com/luckywb13/blog\n&amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	cd luckywb13/blog\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	make linux_build\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	&lt;br /&gt;\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	使用docker启动\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	docker-compose up -d --build\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	&lt;br /&gt;\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	手动启动\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	./main -debug=false -conf=app.yaml\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	&lt;br /&gt;\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	手动启动时需要手动创建数据表，表结构文件为blog.sql，然后修改app.yaml中db的相关配置。\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	&lt;br /&gt;\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	后端管理页面:{you domain}/admin/login\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	博客主页面{you domain}/home/index\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	后台管理账号初始账号:admin@qq.com 密码:123456\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	&lt;br /&gt;\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	留言版引用了畅言平台, 需要注册并且审核备案号后获取信息\n&lt;/p&gt;\n&lt;p style=&#34;text-indent:2em;&#34;&gt;\n	&lt;br /&gt;\n&lt;/p&gt;', 0, 2, '利用gin框架写的个人博客', '', '2020-06-25 23:14:53', '2020-06-25 23:25:30');
COMMIT;

-- ----------------------------
-- Table structure for tb_category
-- ----------------------------
DROP TABLE IF EXISTS `tb_category`;
CREATE TABLE `tb_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '1:正常 2:删除',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_category
-- ----------------------------
BEGIN;
INSERT INTO `tb_category` VALUES (1, 'golang', 1, '2020-06-25 23:03:50', '2020-06-25 23:03:50');
COMMIT;

-- ----------------------------
-- Table structure for tb_comment
-- ----------------------------
DROP TABLE IF EXISTS `tb_comment`;
CREATE TABLE `tb_comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL DEFAULT '0' COMMENT '文章id',
  `email` varchar(100) NOT NULL DEFAULT '',
  `name` varchar(100) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `ip` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `index_article_id` (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_comment
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tb_num
-- ----------------------------
DROP TABLE IF EXISTS `tb_num`;
CREATE TABLE `tb_num` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL DEFAULT '0',
  `comment_num` int(11) NOT NULL DEFAULT '0',
  `read_num` int(11) NOT NULL DEFAULT '0',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_article_id` (`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_num
-- ----------------------------
BEGIN;
INSERT INTO `tb_num` VALUES (1, 1, 0, 22, '2020-06-25 23:25:32', '2020-06-25 23:15:11');
COMMIT;

-- ----------------------------
-- Table structure for tb_talk
-- ----------------------------
DROP TABLE IF EXISTS `tb_talk`;
CREATE TABLE `tb_talk` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `icon` varchar(255) NOT NULL DEFAULT '',
  `status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '1:正常 2:删除',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_talk
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
