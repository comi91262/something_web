DROP TABLE IF EXISTS `books`;
CREATE TABLE `books` (
  `ID` int(64) NOT NULL AUTO_INCREMENT,
  `Title` varchar(40) NOT NULL DEFAULT '',
  `Content` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

