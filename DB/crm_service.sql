-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: localhost    Database: crm_service
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `actors`
--

DROP TABLE IF EXISTS `actors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `actors` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role_id` int DEFAULT NULL,
  `verified` enum('Yes','No') DEFAULT NULL,
  `active` enum('Yes','No') DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `role_id_index` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actors`
--

LOCK TABLES `actors` WRITE;
/*!40000 ALTER TABLE `actors` DISABLE KEYS */;
INSERT INTO `actors` VALUES (1,'Arga','$2a$10$gaBOFre1oCcyYHAVhklmVOQlFqXE4Bis.bnYVoxILO8DdOKXzlx92',1,'Yes','Yes','2023-06-06 10:08:42','2023-06-06 10:08:42',NULL),(2,'Arga','$2a$10$kOigLH3QSijwvqDfbq1z6.rXCDbqNgzdjBJK..BU3z.tV4LL0fRSm',1,'Yes','Yes','2023-06-06 13:49:07','2023-06-06 13:49:07',NULL),(3,'Robi','$2a$10$C0WrR.2b.ol27Fj2AcYzN.6a75dTmssYgRr2a2HMXJngqslV3xsci',2,'Yes','Yes','2023-06-08 05:13:55','2023-06-08 05:13:55',NULL),(4,'Robi','$2a$10$C0WrR.2b.ol27Fj2AcYzN.6a75dTmssYgRr2a2HMXJngqslV3xsci',2,'Yes','Yes','2023-06-08 05:13:55','2023-06-09 02:37:56',NULL),(5,'Robi','$2a$10$C0WrR.2b.ol27Fj2AcYzN.6a75dTmssYgRr2a2HMXJngqslV3xsci',2,'No','Yes','2023-06-08 05:13:55','2023-06-09 02:38:22',NULL),(6,'Robi','$2a$10$C0WrR.2b.ol27Fj2AcYzN.6a75dTmssYgRr2a2HMXJngqslV3xsci',2,'Yes','No','2023-06-08 05:13:55','2023-06-09 02:41:18',NULL);
/*!40000 ALTER TABLE `actors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `firstname` varchar(255) DEFAULT NULL,
  `lastname` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` VALUES (1,'',NULL,NULL,NULL,'2023-06-04 06:15:20','2023-06-04 06:15:20',NULL),(2,'Arga NEW','Bimantara NEW','argab@gmail NEW','bima.png NEW','2023-06-04 06:30:56','2023-06-04 10:52:55',NULL),(3,'Arga','Bimantara','argab@gmail','bima.png','2023-06-04 06:32:26','2023-06-04 14:53:28','2023-06-04 14:53:28'),(4,'','','','','2023-06-04 07:32:16','2023-06-04 07:32:16',NULL),(5,'kemal','joyo','  qwerty@gmaul','mboh.png','2023-06-04 07:33:22','2023-06-04 07:33:22',NULL),(6,'kemal','joyo','  qwerty@gmaul','mboh.png','2023-06-04 07:33:56','2023-06-04 07:33:56',NULL),(7,'Arga NEW','Bimantara NEW','argab@gmail NEW','bima.png NEW','2023-06-04 14:53:58','2023-06-04 14:53:58',NULL),(8,'Arga','Bimantara','argab@gmail','bima.png','2023-06-06 14:23:34','2023-06-06 14:23:34',NULL),(9,'Arga','Bimantara','argab@gmail','bima.png','2023-06-06 14:24:44','2023-06-06 14:24:44',NULL),(10,'Arga','Bimantara','argab@gmail','bima.png','2023-06-06 14:25:05','2023-06-06 14:25:05',NULL),(11,'Arga','Bimantara','argab@gmail','bima.png','2023-06-06 14:25:09','2023-06-06 14:25:09',NULL),(12,'Arga','Bimantara','argab@gmail','bima.png','2023-06-06 14:25:38','2023-06-06 14:25:38',NULL),(13,'Arga','Bimantara','argab@gmail','bima.png','2023-06-06 14:26:16','2023-06-06 14:26:16',NULL),(14,'Arga','Bimantara','argab@gmail','bima.png','2023-06-06 14:27:03','2023-06-06 14:27:03',NULL),(15,'Arga','Bimantara','argab@gmail','bima.png','2023-06-06 14:32:56','2023-06-06 14:32:56',NULL);
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `register`
--

DROP TABLE IF EXISTS `register`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `register` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` bigint unsigned DEFAULT NULL,
  `super_admin_id` bigint unsigned DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `register`
--

LOCK TABLES `register` WRITE;
/*!40000 ALTER TABLE `register` DISABLE KEYS */;
/*!40000 ALTER TABLE `register` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role` (
  `id` int NOT NULL,
  `role_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `role_ibfk_1` FOREIGN KEY (`id`) REFERENCES `actors` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'Admin');
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-09 10:56:45
