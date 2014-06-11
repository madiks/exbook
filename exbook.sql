-- phpMyAdmin SQL Dump
-- version 4.0.4.1
-- http://www.phpmyadmin.net
--
-- 主机: 127.0.0.1
-- 生成日期: 2014 �?06 �?11 �?03:00
-- 服务器版本: 5.6.11
-- PHP 版本: 5.5.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- 数据库: `exbook`
--
CREATE DATABASE IF NOT EXISTS `exbook` DEFAULT CHARACTER SET utf8 COLLATE utf8_bin;
USE `exbook`;

-- --------------------------------------------------------

--
-- 表的结构 `activity`
--

CREATE TABLE IF NOT EXISTS `activity` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '活动id',
  `author` varchar(100) COLLATE utf8_bin DEFAULT NULL COMMENT '创建者',
  `aid` int(11) DEFAULT NULL COMMENT '创建者id',
  `title` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '标题',
  `condition` varchar(700) COLLATE utf8_bin DEFAULT NULL COMMENT '条件',
  `content` text COLLATE utf8_bin COMMENT '内容',
  `like` int(11) DEFAULT NULL COMMENT '喜欢数量',
  `limit` int(11) DEFAULT NULL COMMENT '人数限制',
  `join` int(11) DEFAULT NULL COMMENT '已参加人数',
  `region` varchar(100) COLLATE utf8_bin DEFAULT NULL COMMENT '地域',
  `status` int(1) DEFAULT NULL COMMENT '状态',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `activity_log`
--

CREATE TABLE IF NOT EXISTS `activity_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '活动参加记录',
  `aid` int(11) DEFAULT NULL COMMENT '活动id',
  `uid` int(11) DEFAULT NULL COMMENT '用户id',
  `create_time` datetime DEFAULT NULL COMMENT '参加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `admin`
--

CREATE TABLE IF NOT EXISTS `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uname` char(255) COLLATE utf8_bin DEFAULT NULL,
  `pwd` varchar(500) COLLATE utf8_bin DEFAULT NULL,
  `lastlogin` datetime DEFAULT NULL COMMENT '最近登陆时间',
  `ip` varchar(300) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `admin`
--

INSERT INTO `admin` (`id`, `uname`, `pwd`, `lastlogin`, `ip`) VALUES
(1, 'madiks', '939314105ce8701e67489642ef4d49e8', '2014-05-15 16:36:53', '127.0.0.1');

-- --------------------------------------------------------

--
-- 表的结构 `bookmark`
--

CREATE TABLE IF NOT EXISTS `bookmark` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '收藏id',
  `uid` int(11) DEFAULT NULL COMMENT '用户id',
  `type` int(1) DEFAULT NULL COMMENT '收藏类型，1是交易收藏，2是图书收藏',
  `tid` int(11) DEFAULT NULL COMMENT '交易id',
  `bid` int(11) DEFAULT NULL COMMENT '图书id',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `status` int(11) DEFAULT NULL COMMENT '可用状态，1可用，0不可用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=30 ;

--
-- 转存表中的数据 `bookmark`
--

INSERT INTO `bookmark` (`id`, `uid`, `type`, `tid`, `bid`, `create_time`, `status`) VALUES
(3, 4, 1, 2, 0, '2014-05-20 09:51:26', 1),
(4, 4, 2, 0, 3, '2014-05-20 09:51:34', 1),
(6, 4, 2, 0, 1, '2014-05-20 10:04:55', 1),
(8, 1, 2, 0, 3, '2014-05-26 14:25:36', 1),
(10, 1, 2, 0, 1, '2014-05-26 14:32:49', 1),
(13, 1, 2, 0, 8, '2014-05-27 19:34:13', 1),
(14, 4, 1, 7, 0, '2014-05-28 09:31:45', 1),
(17, 4, 2, 0, 17, '2014-05-28 11:59:27', 1),
(18, 1, 2, 0, 19, '2014-05-29 10:04:22', 1),
(19, 1, 1, 9, 0, '2014-05-29 10:05:31', 1),
(20, 1, 2, 0, 5, '2014-05-29 23:08:18', 1),
(21, 1, 2, 0, 20, '2014-06-01 13:58:44', 1),
(22, 1, 1, 13, 0, '2014-06-01 14:02:02', 1),
(23, 1, 1, 7, 0, '2014-06-04 11:35:38', 1),
(24, 6, 2, 0, 10, '2014-06-04 12:59:40', 1),
(25, 4, 2, 0, 5, '2014-06-04 13:09:35', 1),
(26, 1, 2, 0, 10, '2014-06-04 15:07:05', 1),
(27, 1, 2, 0, 34, '2014-06-05 12:51:38', 1),
(29, 1, 1, 16, 0, '2014-06-07 21:12:47', 1);

-- --------------------------------------------------------

--
-- 表的结构 `books`
--

CREATE TABLE IF NOT EXISTS `books` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '图书id',
  `title` varchar(300) COLLATE utf8_bin DEFAULT NULL COMMENT '书名',
  `subtitle` varchar(300) COLLATE utf8_bin DEFAULT NULL COMMENT '副标题',
  `author` varchar(300) COLLATE utf8_bin DEFAULT NULL COMMENT '作者',
  `pubdate` varchar(200) COLLATE utf8_bin DEFAULT NULL COMMENT '出版时间',
  `image` varchar(400) COLLATE utf8_bin DEFAULT NULL COMMENT '图书图片',
  `binding` varchar(50) COLLATE utf8_bin DEFAULT NULL COMMENT '印刷档次',
  `translator` varchar(100) COLLATE utf8_bin DEFAULT NULL COMMENT '译者',
  `pages` int(11) DEFAULT NULL COMMENT '页数',
  `dbalt` varchar(400) COLLATE utf8_bin DEFAULT NULL COMMENT '豆瓣图书地址',
  `publisher` varchar(200) COLLATE utf8_bin DEFAULT NULL COMMENT '出版社',
  `isbn10` varchar(200) COLLATE utf8_bin DEFAULT NULL,
  `isbn13` varchar(200) COLLATE utf8_bin DEFAULT NULL,
  `dbid` varchar(300) COLLATE utf8_bin DEFAULT NULL COMMENT '豆瓣id',
  `price` varchar(50) COLLATE utf8_bin DEFAULT NULL COMMENT '标价',
  `isprivate` int(1) DEFAULT NULL COMMENT '是否是豆瓣未找到的图书，1是，0不是',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=36 ;

--
-- 转存表中的数据 `books`
--

INSERT INTO `books` (`id`, `title`, `subtitle`, `author`, `pubdate`, `image`, `binding`, `translator`, `pages`, `dbalt`, `publisher`, `isbn10`, `isbn13`, `dbid`, `price`, `isprivate`) VALUES
(1, '当幸福来敲门（英文版）', NULL, '克里斯·加德纳', '2010-9', 'http://img3.douban.com/mpic/s6219903.jpg', '平装', NULL, 281, 'http://book.douban.com/subject/5269101/', '清华大学出版社', '7302236712', '9787302236719', '5269101', '45.80元', 0),
(2, '激情晨读英语美文', '英语美文', '高林显', '2007-1', 'http://img5.douban.com/mpic/s6014269.jpg', '平装', NULL, 297, 'http://book.douban.com/subject/2221912/', '山东电子音像出版社', '7894810193', '9787894810199', '2221912', '12.80元', 0),
(3, '一分钟的你自己', NULL, '斯宾塞.约翰逊', '2002-10', 'http://img3.douban.com/mpic/s23562144.jpg', '平装', NULL, 152, 'http://book.douban.com/subject/1034572/', '延边人民出版社', '7806487352', '9787806487358', '1034572', '16.80', 0),
(4, '算法（第5部分）', '（第5部分）图算法', '塞奇威克', '2010-1', 'http://img5.douban.com/mpic/s4139749.jpg', NULL, NULL, 303, 'http://book.douban.com/subject/4191525/', '机械工业出版社', '7111285050', '9787111285052', '4191525', '59.00元', 0),
(5, 'MacTalk 人生元编程', '', '池建强', '2014-2-1', 'http://img3.douban.com/mpic/s27219901.jpg', '平装', '', 0, 'http://book.douban.com/subject/25826578/', '人民邮电出版社', '7115342237', '9787115342232', '25826578', '45', 0),
(8, '全世界人民都知道', '', '李承鹏', '2013-1-10', 'http://img3.douban.com/mpic/s24581281.jpg', '平装', '', 0, 'http://book.douban.com/subject/20497889/', '新星出版社', '751331036X', '9787513310369', '20497889', '29.00元', 0),
(9, '大数据时代', '生活、工作与思维的大变革', '[英] 维克托•迈尔•舍恩伯格（Viktor Mayer-Schönberger）', '2012-12', 'http://img3.douban.com/mpic/s24574862.jpg', '平装', '', 0, 'http://book.douban.com/subject/20429677/', '浙江人民出版社', '7213052543', '9787213052545', '20429677', '49.90元', 0),
(10, '你凭什么领导别人', '', '(英)罗布·戈菲/加雷思·琼斯', '2008', 'http://img3.douban.com/mpic/s21956105.jpg', '其他', '', 0, 'http://book.douban.com/subject/2996037/', '商务印书馆', '7100055121', '9787100055123', '2996037', '43.00元', 0),
(11, '记住你是谁', '15位哈佛教授震撼心灵的人生故事', '[美]黛西·韦德曼', '2006-2', 'http://img5.douban.com/mpic/s25683128.jpg', '', '', 0, 'http://book.douban.com/subject/1802576/', '商务印书馆', '7100048095', '9787100048095', '1802576', '39.00元', 0),
(14, '大数据云图', '如何在大数据时代寻找下一个大机遇', '', '2013-12-1', 'http://img5.douban.com/mpic/s27153366.jpg', '', '', 0, 'http://book.douban.com/subject/25773128/', '浙江人民出版社', '7213058428', '9787213058424', '25773128', '45.9元', 0),
(15, '货币的非国家化', '对多元货币的理论与实践的分析', '弗里德里希·冯·哈耶克', '2007-8', 'http://img3.douban.com/mpic/s2611882.jpg', '平装', '', 0, 'http://book.douban.com/subject/2155342/', '新星出版社', '7802253160', '9787802253162', '2155342', '25.00元', 0),
(16, '中国人的焦虑从哪里来', '论财富与地位的不平等', '', '2013-1', 'http://img5.douban.com/mpic/s24876157.jpg', '', '', 0, 'http://book.douban.com/subject/20563189/', '群言出版社', '7802563518', '9787802563513', '20563189', '38.00元', 0),
(17, 'Android和PHP开发最佳实践', '', '黄隽实', '2013-3-20', 'http://img5.douban.com/mpic/s25804527.jpg', '平装', '', 0, 'http://book.douban.com/subject/21677501/', '机械工业出版社华章公司', '7111410505', '9787111410508', '21677501', '79.00元', 0),
(18, '深入浅出Node.js', '', '', '2013-12-1', 'http://img5.douban.com/mpic/s27269296.jpg', '平装', '', 0, 'http://book.douban.com/subject/25768396/', '人民邮电出版社', '7115335508', '9787115335500', '25768396', 'CNY 69.00', 0),
(19, '善用佳软', '高效能人士的软件应用之道', '张玉新', '2013-4-25', 'http://img5.douban.com/mpic/s26334086.jpg', '', '', 0, 'http://book.douban.com/subject/24152664/', '人民邮电出版社', '7115313075', '9787115313072', '24152664', '29.00元', 0),
(20, '图解TCP/IP : 第5版', '', '[日]竹下隆史', '2013-6', 'http://img5.douban.com/mpic/s26676928.jpg', '平装', '', 0, 'http://book.douban.com/subject/24737674/', '人民邮电出版社', '7115318972', '9787115318978', '24737674', '69.00', 0),
(21, '社会工程', '安全体系中的人性漏洞', '', '2013-12-1', 'http://img3.douban.com/mpic/s27273253.jpg', '平装', '', 0, 'http://book.douban.com/subject/25768304/', '人民邮电出版社', '7115335389', '9787115335388', '25768304', 'CNY 59.00', 0),
(22, 'Android应用UI设计模式', '', '', '2013-12', 'http://img3.douban.com/mpic/s27126934.jpg', '平装', '', 0, 'http://book.douban.com/subject/25764505/', '人民邮电出版社', '7115334684', '9787115334688', '25764505', '69.00元', 0),
(23, '上海书评选萃：穿透历史', '', '', '2013-8', 'http://img5.douban.com/mpic/s26935417.jpg', '精装', '', 0, 'http://book.douban.com/subject/24919004/', '译林出版社', '7544740706', '9787544740708', '24919004', '35.00', 0),
(24, '重点中学', '', '何天白', '2010-11', 'http://img3.douban.com/mpic/s6483563.jpg', '', '', 0, 'http://book.douban.com/subject/5381386/', '花山文艺出版社', '7807559144', '9787807559146', '5381386', '28.00元', 0),
(25, 'Android编程权威指南', '', '[美] Bill Phillips', '2014-4', 'http://img5.douban.com/mpic/s27240339.jpg', '平装', '', 0, 'http://book.douban.com/subject/25848404/', '人民邮电出版社', '7115346437', '9787115346438', '25848404', '99.00', 0),
(26, '拖延心理学', '向与生俱来的行为顽症宣战', '[美] 简·博克(Jane B. Burka)', '2009-12', 'http://img3.douban.com/mpic/s4355925.jpg', '平装', '', 0, 'http://book.douban.com/subject/4180711/', '中国人民大学出版社', '7300113907', '9787300113906', '4180711', '39.80元', 0),
(27, 'Python入门经典', '以解决计算问题为导向的Python编程实践', '(美)William F    PunchRichard Enbody', '2012-8', 'http://img5.douban.com/mpic/s23111469.jpg', '', '', 0, 'http://book.douban.com/subject/11610789/', '机械工业出版社', '7111394135', '9787111394136', '11610789', '79.00元', 0),
(28, '狼图腾', '', '', '2004-4', 'http://img3.douban.com/mpic/s1466042.jpg', '平装', '', 0, 'http://book.douban.com/subject/1022060/', '长江文艺出版社', '7535427308', '9787535427304', '1022060', '32.00元', 0),
(29, '努尔哈赤全传', '', '', '2014-1', 'http://img3.douban.com/mpic/s27203482.jpg', '平装', '', 0, 'http://book.douban.com/subject/25812865/', '江苏文艺出版社', '7539969474', '9787539969473', '25812865', '39.80', 0),
(30, '30天写小说', '30天写小说', '', '2013-5', 'http://img3.douban.com/mpic/s26387343.jpg', '平装', '', 0, 'http://book.douban.com/subject/24527521/', '中国人民大学出版社', '7300173594', '9787300173597', '24527521', '28.00元', 0),
(31, '创意生活情趣168', '', '', '2013-1', 'http://img3.douban.com/mpic/s25109862.jpg', '', '', 0, 'http://book.douban.com/subject/20396722/', '中国摄影出版社', '7802368316', '9787802368316', '20396722', '68.00元', 0),
(32, '唐代衣食住行（插图珍藏本）', '', '', '2013-4-1', 'http://img3.douban.com/mpic/s27018245.jpg', '平装', '', 0, 'http://book.douban.com/subject/25704808/', '中華書局', '7101091679', '9787101091670', '25704808', '', 0),
(33, '驯狮记', '—Mac OS X 10.8 Mountain Lion使用手册', '', '2013-2', 'http://img3.douban.com/mpic/s26241422.jpg', '', '', 0, 'http://book.douban.com/subject/21719986/', '人民邮电出版社', '7115306974', '9787115306975', '21719986', '79.00元', 0),
(34, '图表说服力', 'Excel与PowerPoint的图表终极活用术', '', '2011-10', 'http://img3.douban.com/mpic/s6988792.jpg', '', '', 0, 'http://book.douban.com/subject/6873593/', '', '7302263728', '9787302263722', '6873593', '69.00元', 0),
(35, 'Kinect人机交互开发实践', '', '', '2012-12', 'http://img5.douban.com/mpic/s24456749.jpg', '平装', '', 0, 'http://book.douban.com/subject/20423598/', '人民邮电出版社', '7115300291', '9787115300294', '20423598', '39.00元', 0);

-- --------------------------------------------------------

--
-- 表的结构 `inventory`
--

CREATE TABLE IF NOT EXISTS `inventory` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '库存id',
  `uid` int(11) DEFAULT NULL COMMENT '用户id',
  `bid` int(11) DEFAULT NULL COMMENT '图书id',
  `newoldrate` varchar(50) COLLATE utf8_bin DEFAULT NULL COMMENT '新旧程度',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `status` int(1) DEFAULT NULL COMMENT '可用状态，1可用，0不可用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=26 ;

--
-- 转存表中的数据 `inventory`
--

INSERT INTO `inventory` (`id`, `uid`, `bid`, `newoldrate`, `create_time`, `update_time`, `status`) VALUES
(1, 1, 1, '9成新', '2014-05-16 16:36:21', '2014-05-16 16:36:27', 1),
(3, 1, 3, '6成新', '2014-05-19 15:55:14', '2014-05-19 15:55:17', 1),
(4, 1, 2, '6成新', '2014-05-27 12:24:00', '2014-05-27 12:24:00', 1),
(5, 1, 4, '全新', '2014-05-27 12:38:25', '2014-05-27 12:38:25', 1),
(7, 1, 5, '全新', '2014-05-27 13:20:26', '2014-05-27 13:20:26', 1),
(8, 1, 5, '9成新', '2014-05-27 13:22:39', '2014-05-27 13:22:39', 1),
(9, 1, 8, '7成新', '2014-05-27 13:24:00', '2014-05-27 13:24:00', 1),
(10, 1, 9, '7成新', '2014-05-27 13:25:01', '2014-05-27 13:25:01', 1),
(11, 1, 10, '7成新', '2014-05-27 15:13:56', '2014-05-27 15:13:56', 1),
(12, 1, 15, '9成新', '2014-05-27 16:19:23', '2014-05-27 16:19:23', 1),
(14, 4, 11, '全新', '2014-05-28 09:30:15', '2014-05-28 09:30:15', 1),
(15, 4, 19, '8成新', '2014-05-28 09:30:45', '2014-05-28 09:30:45', 1),
(16, 4, 20, '7成新', '2014-05-28 09:31:07', '2014-05-28 09:31:07', 1),
(17, 1, 19, '8成新', '2014-05-30 10:22:19', '2014-05-30 10:22:19', 1),
(18, 6, 24, '7成新', '2014-06-04 10:24:19', '2014-06-04 10:24:19', 1),
(19, 6, 25, '全新', '2014-06-04 10:24:41', '2014-06-04 10:24:41', 1),
(20, 6, 18, '8成新', '2014-06-04 10:25:03', '2014-06-04 10:25:03', 1),
(21, 6, 17, '7成新', '2014-06-04 10:25:36', '2014-06-04 10:25:36', 1),
(22, 6, 26, '8成新', '2014-06-04 10:26:26', '2014-06-04 10:26:26', 1),
(23, 6, 27, '6成新', '2014-06-04 10:27:03', '2014-06-04 10:27:03', 1),
(24, 1, 34, '7成新', '2014-06-04 15:05:51', '2014-06-04 15:05:51', 1),
(25, 4, 1, '全新', '2014-06-05 13:12:01', '2014-06-05 13:12:01', 1);

-- --------------------------------------------------------

--
-- 表的结构 `message`
--

CREATE TABLE IF NOT EXISTS `message` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `from_uid` int(11) DEFAULT NULL COMMENT '发送人id',
  `from_uname` varchar(200) COLLATE utf8_bin DEFAULT NULL,
  `for_uid` int(11) DEFAULT NULL COMMENT '接收人id',
  `for_uname` varchar(200) COLLATE utf8_bin DEFAULT NULL,
  `msg` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '内容',
  `create_time` datetime DEFAULT NULL,
  `status` int(11) DEFAULT NULL COMMENT '私信状态',
  `tid` int(11) DEFAULT NULL COMMENT '交易id',
  `istrade` int(1) DEFAULT NULL COMMENT '是否是交易私信，0不是，1是',
  `tag` varchar(200) COLLATE utf8_bin DEFAULT NULL COMMENT '附加信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `message_reply`
--

CREATE TABLE IF NOT EXISTS `message_reply` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '回复私信id',
  `mid` int(11) DEFAULT NULL COMMENT '私信id',
  `uid` int(11) DEFAULT NULL COMMENT '用户id',
  `uname` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `msg` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '消息内容',
  `tag` varchar(200) COLLATE utf8_bin DEFAULT NULL COMMENT '附加提示信息',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `status` int(1) DEFAULT NULL COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `offers`
--

CREATE TABLE IF NOT EXISTS `offers` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '报价id',
  `tid` int(11) DEFAULT NULL COMMENT '交易id',
  `tcid` int(11) DEFAULT NULL COMMENT '所属交易创建这的id',
  `uid` int(11) DEFAULT NULL COMMENT '报价用户id',
  `uname` varchar(200) COLLATE utf8_bin DEFAULT NULL COMMENT '报价用户名',
  `offerlist` varchar(1000) COLLATE utf8_bin DEFAULT NULL COMMENT '报价项目，序列化存储',
  `offerstatus` int(11) DEFAULT NULL COMMENT '报价的状态,1为默认，2为已读，3为已被接受，4为已回复',
  `tstatus` int(11) DEFAULT NULL COMMENT '交易人对于报价这的反馈,0为无反馈，1为同意',
  `msg` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '附加消息',
  `create_time` datetime DEFAULT NULL,
  `ohasmoney` int(1) DEFAULT NULL COMMENT '是否有钱',
  `omoney` int(11) DEFAULT NULL COMMENT '钱数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=17 ;

--
-- 转存表中的数据 `offers`
--

INSERT INTO `offers` (`id`, `tid`, `tcid`, `uid`, `uname`, `offerlist`, `offerstatus`, `tstatus`, `msg`, `create_time`, `ohasmoney`, `omoney`) VALUES
(5, 7, 1, 4, 'adiks', '[,15,]', 2, 1, '1换1.', '2014-05-28 13:46:15', 1, 5),
(9, 9, 4, 1, 'madiks', '[,3,12,]', 3, 1, '1:1', '2014-05-28 22:56:44', 0, 0),
(10, 9, 4, 1, 'madiks', '[,13,]', 4, 1, 'yuuu.', '2014-05-29 00:09:00', 0, 0),
(13, 9, 4, 1, 'madiks', '[,12,]', 2, 1, '1:1', '2014-05-30 10:24:42', 1, 10),
(14, 14, 1, 1, 'madiks', '[,24,]', 2, 1, '', '2014-06-04 15:08:16', 0, 0),
(16, 16, 0, 1, 'madiks', '[,3,4,10,]', 2, 1, 'new 2 test.', '2014-06-05 12:50:25', 0, 0);

-- --------------------------------------------------------

--
-- 表的结构 `regionset`
--

CREATE TABLE IF NOT EXISTS `regionset` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `region` varchar(200) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=7 ;

--
-- 转存表中的数据 `regionset`
--

INSERT INTO `regionset` (`id`, `region`) VALUES
(1, '长春'),
(2, '杭州'),
(3, '西安'),
(4, '北京'),
(5, '深圳'),
(6, '重庆');

-- --------------------------------------------------------

--
-- 表的结构 `reply_offer`
--

CREATE TABLE IF NOT EXISTS `reply_offer` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '回复报价id',
  `oid` int(11) DEFAULT NULL COMMENT '报价id',
  `toid` int(11) DEFAULT NULL COMMENT '所属报价创建者的id',
  `uid` int(11) DEFAULT NULL COMMENT '用户id',
  `uname` varchar(200) COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `tid` int(11) DEFAULT NULL COMMENT '交易id',
  `msg` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '附加信息',
  `status` int(11) DEFAULT NULL COMMENT '报价状态，1为正常，2为取消',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `isprivate` int(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=10 ;

--
-- 转存表中的数据 `reply_offer`
--

INSERT INTO `reply_offer` (`id`, `oid`, `toid`, `uid`, `uname`, `tid`, `msg`, `status`, `create_time`, `isprivate`) VALUES
(1, 10, 1, 4, 'adiks', 9, 'uuiiiopp', 2, '2014-05-29 09:03:59', 0),
(8, 10, 1, 1, 'madiks', 9, 'woshi', 2, '2014-05-29 11:28:07', 0),
(9, 12, 1, 4, 'adiks', 9, 'kkkkkkkkkkkkkkkkkkkk', 2, '2014-05-29 11:58:58', 0);

-- --------------------------------------------------------

--
-- 表的结构 `trades`
--

CREATE TABLE IF NOT EXISTS `trades` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '交换id',
  `uname` varchar(200) COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `uid` int(11) DEFAULT NULL COMMENT '用户id',
  `selllist` varchar(1000) COLLATE utf8_bin DEFAULT NULL COMMENT '交换列表，序列化存储',
  `hasanyoffer` int(1) DEFAULT NULL,
  `smoney` int(11) DEFAULT NULL,
  `shasmoney` int(1) DEFAULT NULL,
  `wantlist` varchar(1000) COLLATE utf8_bin DEFAULT NULL COMMENT '目标列表，序列化存储',
  `wmoney` int(11) DEFAULT NULL,
  `whasmoney` int(1) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `region` varchar(100) COLLATE utf8_bin DEFAULT NULL COMMENT '区域',
  `status` int(1) DEFAULT NULL COMMENT '交换状态',
  `remark` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '备注',
  `dealuid` int(11) DEFAULT NULL COMMENT '交易成功用户id',
  `dealuname` varchar(200) COLLATE utf8_bin DEFAULT NULL COMMENT '交易成功用户名',
  `t_commment` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '对于交易人的评价',
  `o_commment` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '对于报价人的评价',
  `t_estimation` int(11) DEFAULT NULL COMMENT '对于交易人的五星评价',
  `o_estimation` int(11) DEFAULT NULL COMMENT '对于报价人的五星评价',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=17 ;

--
-- 转存表中的数据 `trades`
--

INSERT INTO `trades` (`id`, `uname`, `uid`, `selllist`, `hasanyoffer`, `smoney`, `shasmoney`, `wantlist`, `wmoney`, `whasmoney`, `create_time`, `update_time`, `region`, `status`, `remark`, `dealuid`, `dealuname`, `t_commment`, `o_commment`, `t_estimation`, `o_estimation`) VALUES
(1, 'madiks', 1, '[,1,]', 1, 0, 0, '[,3,4,]', 10, 1, '2014-05-16 16:34:36', '2014-05-26 16:04:26', '长春', 1, '求1换1', 0, '', '', '', 0, 0),
(2, 'madiks', 1, '[,3,]', 0, 12, 1, '[,1,4,]', 0, 0, '2014-05-19 15:49:24', '2014-05-26 16:04:17', '长春', 1, '随便换换', 0, '', '', '', 0, 0),
(6, 'madiks', 1, '[,12,10,9,7,]', 1, 0, 0, '[,3,1,]', 0, 0, '2014-05-27 19:50:43', '2014-05-27 20:28:22', '长春', 1, 'djaksda', 0, '', '', '', 0, 0),
(7, 'madiks', 1, '[,13,10,11,]', 1, 0, 0, '[,8,18,]', 0, 0, '2014-05-27 20:38:09', '2014-05-27 20:38:09', '长春', 1, 'sajsgdjasjaj.', 0, '', '', '', 0, 0),
(8, 'adiks', 4, '[,15,]', 1, 0, 0, '[,3,]', 0, 0, '2014-05-28 09:31:28', '2014-05-28 09:31:28', '长春', 1, '', 0, '', '', '', 0, 0),
(9, 'adiks', 4, '[,16,14,15,]', 0, 0, 0, '[,1,3,]', 0, 0, '2014-05-28 22:56:16', '2014-05-30 10:24:43', '长春', 1, 'sssss.', 0, '', '', '', 0, 0),
(10, 'madiks', 1, '[,7,8,]', 0, 0, 0, '[,19,20,]', 0, 0, '2014-05-29 23:08:10', '2014-05-29 23:08:10', '长春', 1, '', 0, '', '', '', 0, 0),
(13, 'madiks', 1, '[,17,11,]', 0, 0, 0, '[,20,23,]', 0, 0, '2014-06-01 14:00:17', '2014-06-01 14:00:59', '长春', 1, 'jaHSHaa.', 0, '', '', '', 0, 0),
(14, 'madiks', 1, '[,24,]', 0, 0, 0, '[,10,]', 0, 0, '2014-06-04 15:07:51', '2014-06-04 15:08:16', '长春', 1, '', 0, '', '', '', 0, 0),
(16, 'madiks', 1, '[,1,5,24,]', 0, 0, 0, '[,10,5,]', 0, 0, '2014-06-05 12:48:03', '2014-06-05 12:50:25', '长春', 1, 'hello new.', 0, '', '', '', 0, 0);

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uname` char(200) COLLATE utf8_bin DEFAULT NULL,
  `pwd` varchar(500) COLLATE utf8_bin DEFAULT NULL,
  `region` varchar(200) COLLATE utf8_bin DEFAULT NULL COMMENT '区域',
  `tname` varchar(300) COLLATE utf8_bin DEFAULT NULL COMMENT '真实姓名',
  `tel` varchar(300) COLLATE utf8_bin DEFAULT NULL COMMENT '电话',
  `img` varchar(400) COLLATE utf8_bin DEFAULT NULL COMMENT '头像',
  `remark` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '备注',
  `lastlogin` datetime DEFAULT NULL,
  `trade_rank` int(11) DEFAULT NULL COMMENT '交易用户评级',
  `offer_rank` int(11) DEFAULT NULL COMMENT '报价用户评级',
  `address` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '现居住地址',
  `email` varchar(500) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=7 ;

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`id`, `uname`, `pwd`, `region`, `tname`, `tel`, `img`, `remark`, `lastlogin`, `trade_rank`, `offer_rank`, `address`, `email`) VALUES
(1, 'madiks', '827ccb0eea8a706c4c34a16891f84e7b', '长春', '马迪a', '18744033062', '', '', '2014-06-10 19:18:08', 0, 0, 'changchun province renming street NO.4245', 'madiks@163.com'),
(4, 'adiks', '827ccb0eea8a706c4c34a16891f84e7b', '长春', '', '15736236789', '', '', '2014-06-07 18:33:29', 0, 0, '长春市卫星广场', ''),
(5, 'masvs', '827ccb0eea8a706c4c34a16891f84e7b', '杭州', '', '', '', '', '2014-06-04 13:07:09', 0, 0, '', ''),
(6, 'windruner', '827ccb0eea8a706c4c34a16891f84e7b', '长春', 'WR', '18325666223', '', '', '2014-06-04 12:59:28', 0, 0, '长春市西安大路', '');

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
