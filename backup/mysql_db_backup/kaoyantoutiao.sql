-- MySQL dump 10.13  Distrib 5.7.18, for Linux (x86_64)
--
-- Host: localhost    Database: kaoyantoutiao
-- ------------------------------------------------------
-- Server version	5.7.18

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `kytt_answer`
--

DROP TABLE IF EXISTS `kytt_answer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_answer` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `question_id` int(10) unsigned NOT NULL,
  `question_title` varchar(128) NOT NULL,
  `user_nickname` varchar(64) NOT NULL,
  `cotent` mediumtext NOT NULL,
  `like_count` int(11) NOT NULL,
  `comment_count` int(11) NOT NULL,
  `answer_date` datetime NOT NULL,
  `title_image` varchar(128) DEFAULT NULL,
  `summary` varchar(512) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_answer`
--

LOCK TABLES `kytt_answer` WRITE;
/*!40000 ALTER TABLE `kytt_answer` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_answer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_feedback`
--

DROP TABLE IF EXISTS `kytt_feedback`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_feedback` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `time` datetime NOT NULL,
  `content` mediumtext CHARACTER SET utf8,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `kytt_feedback_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `kytt_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_feedback`
--

LOCK TABLES `kytt_feedback` WRITE;
/*!40000 ALTER TABLE `kytt_feedback` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_feedback` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_headline`
--

DROP TABLE IF EXISTS `kytt_headline`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_headline` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `user_nickname` varchar(64) CHARACTER SET utf8 NOT NULL,
  `title` varchar(128) CHARACTER SET utf8 NOT NULL,
  `content` mediumtext CHARACTER SET utf8 NOT NULL,
  `post_date` datetime NOT NULL,
  `like_count` int(11) unsigned NOT NULL,
  `comment_count` int(11) unsigned NOT NULL,
  `forward_count` int(11) unsigned NOT NULL,
  `view_count` int(11) unsigned NOT NULL,
  `tag` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `title_image` varchar(128) CHARACTER SET utf8 DEFAULT NULL,
  `summary` varchar(512) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `kytt_headline_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `kytt_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_headline`
--

LOCK TABLES `kytt_headline` WRITE;
/*!40000 ALTER TABLE `kytt_headline` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_headline` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_headline_comment`
--

DROP TABLE IF EXISTS `kytt_headline_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_headline_comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `headline_id` int(10) unsigned NOT NULL,
  `user_nickname` varchar(64) NOT NULL,
  `content` mediumtext NOT NULL,
  `like_count` int(11) NOT NULL,
  `time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `headline_id` (`headline_id`),
  CONSTRAINT `kytt_headline_comment_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `kytt_user` (`id`),
  CONSTRAINT `kytt_headline_comment_ibfk_2` FOREIGN KEY (`headline_id`) REFERENCES `kytt_headline` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_headline_comment`
--

LOCK TABLES `kytt_headline_comment` WRITE;
/*!40000 ALTER TABLE `kytt_headline_comment` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_headline_comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_official_headline`
--

DROP TABLE IF EXISTS `kytt_official_headline`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_official_headline` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `user_nickname` varchar(64) CHARACTER SET utf8 NOT NULL,
  `title` varchar(128) CHARACTER SET utf8 NOT NULL,
  `content` mediumtext NOT NULL,
  `post_date` datetime NOT NULL,
  `like_count` int(11) NOT NULL,
  `comment_count` int(11) NOT NULL,
  `forward_count` int(11) NOT NULL,
  `view_count` int(11) NOT NULL,
  `tag` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `title_image` varchar(128) CHARACTER SET utf8 DEFAULT NULL,
  `summary` varchar(512) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_official_headline`
--

LOCK TABLES `kytt_official_headline` WRITE;
/*!40000 ALTER TABLE `kytt_official_headline` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_official_headline` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_official_headline_comment`
--

DROP TABLE IF EXISTS `kytt_official_headline_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_official_headline_comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `official_headline_id` int(10) unsigned NOT NULL,
  `official_name` varchar(64) NOT NULL,
  `content` mediumtext NOT NULL,
  `like_count` int(10) unsigned NOT NULL,
  `time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_official_headline_comment`
--

LOCK TABLES `kytt_official_headline_comment` WRITE;
/*!40000 ALTER TABLE `kytt_official_headline_comment` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_official_headline_comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_organization_auth`
--

DROP TABLE IF EXISTS `kytt_organization_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_organization_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `auth_file` varchar(128) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `kytt_organization_auth_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `kytt_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_organization_auth`
--

LOCK TABLES `kytt_organization_auth` WRITE;
/*!40000 ALTER TABLE `kytt_organization_auth` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_organization_auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_question`
--

DROP TABLE IF EXISTS `kytt_question`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_question` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `user_nickname` varchar(64) NOT NULL,
  `title` varchar(128) NOT NULL,
  `content` mediumtext NOT NULL,
  `post_date` datetime NOT NULL,
  `like_count` int(11) NOT NULL,
  `comment_count` int(11) NOT NULL,
  `forward_count` int(11) NOT NULL,
  `view_count` int(11) NOT NULL,
  `tag` varchar(255) CHARACTER SET armscii8 DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_question`
--

LOCK TABLES `kytt_question` WRITE;
/*!40000 ALTER TABLE `kytt_question` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_question` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_tag`
--

DROP TABLE IF EXISTS `kytt_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag` varchar(255) CHARACTER SET utf8 NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_tag`
--

LOCK TABLES `kytt_tag` WRITE;
/*!40000 ALTER TABLE `kytt_tag` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_user`
--

DROP TABLE IF EXISTS `kytt_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `nickname` varchar(64) NOT NULL,
  `realname` varchar(64) DEFAULT NULL,
  `gender` tinyint(1) DEFAULT NULL,
  `birthday` date DEFAULT NULL,
  `avater_url` varchar(128) DEFAULT NULL,
  `telephone` varchar(16) NOT NULL,
  `email` varchar(64) NOT NULL,
  `type` tinyint(1) NOT NULL,
  `signup_date` datetime NOT NULL,
  `last_signin_date` datetime DEFAULT NULL,
  `last_signin_location` varchar(128) DEFAULT NULL,
  `last_sigin_ip` varchar(16) DEFAULT NULL,
  `active_time` int(11) NOT NULL,
  `is_auth` tinyint(1) NOT NULL,
  `user_state` tinyint(1) NOT NULL,
  `follower_count` int(11) unsigned NOT NULL,
  `following_count` int(11) unsigned NOT NULL,
  `answer_count` int(11) unsigned NOT NULL,
  `headline_count` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_user`
--

LOCK TABLES `kytt_user` WRITE;
/*!40000 ALTER TABLE `kytt_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_user_activity`
--

DROP TABLE IF EXISTS `kytt_user_activity`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_user_activity` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `headline_id` int(10) unsigned NOT NULL,
  `type` varchar(16) CHARACTER SET utf8 NOT NULL,
  `time` int(11) NOT NULL,
  `action_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `headline_id` (`headline_id`),
  CONSTRAINT `kytt_user_activity_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `kytt_user` (`id`),
  CONSTRAINT `kytt_user_activity_ibfk_2` FOREIGN KEY (`headline_id`) REFERENCES `kytt_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_user_activity`
--

LOCK TABLES `kytt_user_activity` WRITE;
/*!40000 ALTER TABLE `kytt_user_activity` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_user_activity` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_user_auth_local`
--

DROP TABLE IF EXISTS `kytt_user_auth_local`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_user_auth_local` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `username` varchar(64) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `telephone` varchar(16) DEFAULT NULL,
  `password` varchar(128) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `kytt_user_auth_local_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `kytt_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_user_auth_local`
--

LOCK TABLES `kytt_user_auth_local` WRITE;
/*!40000 ALTER TABLE `kytt_user_auth_local` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_user_auth_local` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_user_auth_oauth`
--

DROP TABLE IF EXISTS `kytt_user_auth_oauth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_user_auth_oauth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `oauth_name` varchar(128) CHARACTER SET utf8 NOT NULL,
  `oauth_id` varchar(255) CHARACTER SET utf8 NOT NULL,
  `oauth_access_token` varchar(255) CHARACTER SET utf8 NOT NULL,
  `oauth_expires` int(11) DEFAULT NULL,
  `avatar_url` varchar(128) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `kytt_user_auth_oauth_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `kytt_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_user_auth_oauth`
--

LOCK TABLES `kytt_user_auth_oauth` WRITE;
/*!40000 ALTER TABLE `kytt_user_auth_oauth` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_user_auth_oauth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_user_tag`
--

DROP TABLE IF EXISTS `kytt_user_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_user_tag` (
  `int` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `school_tag` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `major_tag` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `other` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`int`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `kytt_user_tag_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `kytt_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_user_tag`
--

LOCK TABLES `kytt_user_tag` WRITE;
/*!40000 ALTER TABLE `kytt_user_tag` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_user_tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kytt_user_tag_score`
--

DROP TABLE IF EXISTS `kytt_user_tag_score`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kytt_user_tag_score` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `tag_id` int(10) unsigned NOT NULL,
  `score` float NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `tag_id` (`tag_id`),
  CONSTRAINT `kytt_user_tag_score_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `kytt_user` (`id`),
  CONSTRAINT `kytt_user_tag_score_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `kytt_tag` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kytt_user_tag_score`
--

LOCK TABLES `kytt_user_tag_score` WRITE;
/*!40000 ALTER TABLE `kytt_user_tag_score` DISABLE KEYS */;
/*!40000 ALTER TABLE `kytt_user_tag_score` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `new_table`
--

DROP TABLE IF EXISTS `new_table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `new_table` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `answer_id` int(10) unsigned NOT NULL,
  `user_nickname` varchar(64) NOT NULL,
  `content` mediumtext NOT NULL,
  `like_count` int(10) unsigned NOT NULL,
  `time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `new_table`
--

LOCK TABLES `new_table` WRITE;
/*!40000 ALTER TABLE `new_table` DISABLE KEYS */;
/*!40000 ALTER TABLE `new_table` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-06-11  9:00:17
