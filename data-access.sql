-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : dim. 22 mai 2022 à 11:43
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
) ENGINE=MyISAM AUTO_INCREMENT=135 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `data_message`
--

INSERT INTO `data_message` (`id`, `content`, `author`, `date`, `Topic_ID`) VALUES
(134, 'Franchement ez', 'Random', '2022-05-22', 7),
(133, 'wxcvbn,', 'Random', '2022-05-22', 6),
(132, 'qsdfg', 'Random', '2022-05-22', 5),
(126, 'aller la', 'Random', '2022-05-21', 2),
(127, 'fefe', 'kris', '2022-05-21', 3),
(128, 'Yo KRis', 'kris', '2022-05-21', 3);

-- --------------------------------------------------------

--
-- Structure de la table `data_topic`
--

DROP TABLE IF EXISTS `data_topic`;
CREATE TABLE IF NOT EXISTS `data_topic` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  `IsAide` tinyint(1) NOT NULL,
  `IsBug` tinyint(1) NOT NULL,
  `IsBoss` tinyint(1) NOT NULL,
  `IsLore` tinyint(1) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `data_topic`
--

INSERT INTO `data_topic` (`ID`, `Name`, `IsAide`, `IsBug`, `IsBoss`, `IsLore`) VALUES
(2, 'Bon', 0, 0, 0, 0),
(3, 'grg', 0, 0, 0, 0),
(5, 'azerty', 1, 0, 0, 0),
(6, 'qsdfghj', 1, 0, 1, 0),
(7, 'O&S  ne sont absolument pas compliqués', 0, 0, 1, 0);

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
) ENGINE=MyISAM AUTO_INCREMENT=219 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `data_user`
--

INSERT INTO `data_user` (`ID`, `Name`, `Email`, `Password`, `Admin`) VALUES
(192, 'Alex', 'alex@gmail.com', '123', 'false'),
(212, 'Krisfraude', 'krisfraude@gmail.com', 'test1212', 'false'),
(213, 'Kris', 'Kris@gmail.com', '123', 'false'),
(218, 'caca', 'caca@gmail.Com', 'caca', 'false');
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
