-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : sam. 21 mai 2022 à 11:17
-- Version du serveur : 5.7.36
-- Version de PHP : 7.4.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `data-access`
--

-- --------------------------------------------------------

--
-- Structure de la table `data_message`
--

DROP TABLE IF EXISTS `data_message`;
CREATE TABLE IF NOT EXISTS `data_message` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` varchar(128) NOT NULL,
  `author` varchar(255) NOT NULL,
  `date` date NOT NULL,
  `Topic_ID` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `Topic_ID` (`Topic_ID`)
) ENGINE=MyISAM AUTO_INCREMENT=114 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `data_message`
--

INSERT INTO `data_message` (`id`, `content`, `author`, `date`, `Topic_ID`) VALUES
(111, 'Donc en théorie ça pmarche\r\n', 'Random', '2022-05-20', 124),
(112, 'dernier petit test', 'kris', '2022-05-20', 124),
(110, 'pleaaaase', 'Random', '2022-05-20', 124),
(109, 'zzz', 'Random', '2022-05-20', 124),
(108, 'rrrrr', 'Random', '2022-05-20', 124),
(106, 'WtF', 'Random', '2022-05-20', 124),
(100, 'Fonctionne stp', 'Random', '2022-05-20', 120),
(107, 'alllerrrr', 'Random', '2022-05-20', 124),
(104, 'On prie fort', 'Random', '2022-05-20', 124),
(105, 'SUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUU', 'Random', '2022-05-20', 124),
(113, 'yiiiiiiiiiiiiiiiiiiiiiiheah', 'kris', '2022-05-20', 120);

-- --------------------------------------------------------

--
-- Structure de la table `data_tags`
--

DROP TABLE IF EXISTS `data_tags`;
CREATE TABLE IF NOT EXISTS `data_tags` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(128) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `data_topic`
--

DROP TABLE IF EXISTS `data_topic`;
CREATE TABLE IF NOT EXISTS `data_topic` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Nom` varchar(255) NOT NULL,
  `Tag_ID` int(11) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `Tag_ID` (`Tag_ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `data_user`
--

DROP TABLE IF EXISTS `data_user`;
CREATE TABLE IF NOT EXISTS `data_user` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  `Email` varchar(255) NOT NULL,
  `Password` varchar(255) NOT NULL,
  `Admin` varchar(255) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM AUTO_INCREMENT=218 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `data_user`
--

INSERT INTO `data_user` (`ID`, `Name`, `Email`, `Password`, `Admin`) VALUES
(192, 'Alex', 'alex@gmail.com', '123', 'false'),
(212, 'Krisfraude', 'krisfraude@gmail.com', 'test1212', 'false'),
(213, 'Kris', 'Kris@gmail.com', '123', 'false');
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
