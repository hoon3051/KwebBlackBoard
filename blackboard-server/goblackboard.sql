-- MariaDB dump 10.19-11.3.0-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: goblackboard
-- ------------------------------------------------------
-- Server version	11.3.0-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `applies`
--

DROP TABLE IF EXISTS `applies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `applies` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `course_id` bigint(20) unsigned DEFAULT NULL,
  `student_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_applies_deleted_at` (`deleted_at`),
  KEY `fk_courses_apply` (`course_id`),
  KEY `fk_users_apply` (`student_id`),
  CONSTRAINT `fk_courses_apply` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`),
  CONSTRAINT `fk_users_apply` FOREIGN KEY (`student_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `applies`
--

LOCK TABLES `applies` WRITE;
/*!40000 ALTER TABLE `applies` DISABLE KEYS */;
INSERT INTO `applies` VALUES
(51,'2024-02-16 06:20:27.870','2024-02-16 06:20:27.870',NULL,18,25),
(52,'2024-02-16 06:20:29.377','2024-02-16 06:20:29.377',NULL,19,25),
(53,'2024-02-16 06:20:30.886','2024-02-16 06:20:30.886',NULL,21,25),
(55,'2024-02-18 22:40:38.140','2024-02-18 22:40:38.140','2024-02-18 22:41:39.660',18,26),
(56,'2024-02-18 22:40:39.752','2024-02-18 22:40:39.752',NULL,19,26),
(57,'2024-02-18 22:40:41.086','2024-02-18 22:40:41.086',NULL,21,26),
(58,'2024-02-18 22:40:43.807','2024-02-18 22:40:43.807',NULL,22,26),
(59,'2024-02-18 22:40:45.206','2024-02-18 22:40:45.206',NULL,23,26),
(60,'2024-02-18 22:40:48.831','2024-02-18 22:40:48.831',NULL,20,26);
/*!40000 ALTER TABLE `applies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `boards`
--

DROP TABLE IF EXISTS `boards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `boards` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `course_id` bigint(20) unsigned DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `body` longtext DEFAULT NULL,
  `desc` longtext DEFAULT NULL,
  `content` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_boards_deleted_at` (`deleted_at`),
  KEY `fk_courses_board` (`course_id`),
  CONSTRAINT `fk_courses_board` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `boards`
--

LOCK TABLES `boards` WRITE;
/*!40000 ALTER TABLE `boards` DISABLE KEYS */;
INSERT INTO `boards` VALUES
(28,'2024-02-16 06:14:55.950','2024-02-16 06:14:55.950',NULL,18,'이산수학 1주차 공지',NULL,NULL,'<p><strong>휴강합니다다</strong></p>'),
(29,'2024-02-16 06:15:31.867','2024-02-16 06:15:31.867',NULL,19,'알고리즘 1주차 공지',NULL,NULL,'<p>건강상의 이유로 <u>녹화 강의</u>를 올려드리겠습니다</p>'),
(30,'2024-02-16 06:16:22.274','2024-02-16 06:16:22.274',NULL,20,'전산학특강 1주차 공지',NULL,NULL,'<p>안녕하세요</p>'),
(31,'2024-02-16 06:21:17.440','2024-02-16 06:21:17.440',NULL,21,'컴퓨터구조 1주차 공지',NULL,NULL,'<p>반갑습니다</p>'),
(32,'2024-02-16 06:21:32.372','2024-02-16 06:21:32.372',NULL,22,'컴퓨터시스템설계 1주차 공지',NULL,NULL,'<p>컴시설입니다</p>'),
(33,'2024-02-16 06:21:46.312','2024-02-16 06:21:46.312',NULL,23,'임베디드시스템 1주차 공지',NULL,NULL,'<p>아두이노를 다룰 예정입니다</p>'),
(34,'2024-02-16 06:22:02.155','2024-02-16 06:22:02.155',NULL,23,'임베디드 시스템 1주차 과제',NULL,NULL,'<p>아두이노 IDE 깔아오세요</p>');
/*!40000 ALTER TABLE `boards` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `courses`
--

DROP TABLE IF EXISTS `courses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `courses` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `coursenumber` longtext DEFAULT NULL,
  `coursename` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_courses_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `courses`
--

LOCK TABLES `courses` WRITE;
/*!40000 ALTER TABLE `courses` DISABLE KEYS */;
INSERT INTO `courses` VALUES
(18,'2024-02-16 06:14:06.789','2024-02-16 06:14:06.789',NULL,'COSE211','이산수학'),
(19,'2024-02-16 06:14:26.865','2024-02-16 06:14:26.865',NULL,'COSE214','알고리즘'),
(20,'2024-02-16 06:14:36.027','2024-02-16 06:14:36.027',NULL,'COSE490','전산학특강'),
(21,'2024-02-16 06:18:29.038','2024-02-16 06:18:29.038',NULL,'COSE222','컴퓨터구조'),
(22,'2024-02-16 06:18:45.525','2024-02-16 06:18:45.525',NULL,'COSE321','컴퓨터시스템설계'),
(23,'2024-02-16 06:19:02.666','2024-02-16 06:19:02.666',NULL,'COSE421','임베디드시스템');
/*!40000 ALTER TABLE `courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `teaches`
--

DROP TABLE IF EXISTS `teaches`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `teaches` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `course_id` bigint(20) unsigned DEFAULT NULL,
  `professor_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_teaches_deleted_at` (`deleted_at`),
  KEY `fk_courses_teach` (`course_id`),
  KEY `fk_users_teach` (`professor_id`),
  CONSTRAINT `fk_courses_teach` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`),
  CONSTRAINT `fk_users_teach` FOREIGN KEY (`professor_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teaches`
--

LOCK TABLES `teaches` WRITE;
/*!40000 ALTER TABLE `teaches` DISABLE KEYS */;
INSERT INTO `teaches` VALUES
(16,'2024-02-16 06:14:06.795','2024-02-16 06:14:06.795',NULL,18,23),
(17,'2024-02-16 06:14:26.869','2024-02-16 06:14:26.869',NULL,19,23),
(18,'2024-02-16 06:14:36.031','2024-02-16 06:14:36.031',NULL,20,23),
(19,'2024-02-16 06:18:29.050','2024-02-16 06:18:29.050',NULL,21,24),
(20,'2024-02-16 06:18:45.530','2024-02-16 06:18:45.530',NULL,22,24),
(21,'2024-02-16 06:19:02.669','2024-02-16 06:19:02.669',NULL,23,24);
/*!40000 ALTER TABLE `teaches` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `displayname` longtext DEFAULT NULL,
  `studentnumber` longtext DEFAULT NULL,
  `isprofessor` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES
(23,'2024-02-16 06:04:46.630','2024-02-16 06:04:46.630',NULL,'hyperspace@korea.ac.kr','$2a$10$vk3QWqH3O67/DWaBf.8A/ePrIooUQHAAjjpbQqhHy17jcW/gLq5gG','박성빈','2000320001',1),
(24,'2024-02-16 06:18:04.496','2024-02-16 06:18:04.496',NULL,'seo@korea.ac.kr','$2a$10$38Lt4wncH78qVV4qwkzdJOpmBBaim8ViEeWaBHYfNoWNKzTLqRNPO','서태원','2000320002',1),
(25,'2024-02-16 06:19:32.309','2024-02-16 06:19:32.309',NULL,'park@korea.ac.kr','$2a$10$vS5qhrahrfo3ayf.5qZ8V.AjMRtTWzRmDRs6EO4nGGMkfWSwB8s66','박하영','2020320021',0),
(26,'2024-02-18 22:40:19.588','2024-02-18 22:40:19.588',NULL,'hoon_3051@korea.ac.kr','$2a$10$yO89yrBphLQCyRAopi/3derwoNIhQf8nAAvvJnUMWBp8vIiEfjlg6','정상훈','2019320021',0);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-02-18 22:46:01
