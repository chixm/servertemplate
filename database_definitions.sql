CREATE SCHEMA `atagosample` ;

CREATE TABLE `atagosample`.`user` (
  `userID` VARCHAR(200) NOT NULL,
  `password` VARCHAR(60) NOT NULL,
  `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` TIMESTAMP NULL,
  `deletedAt` TIMESTAMP NULL,
  PRIMARY KEY (`userID`))
COMMENT = 'basic user info for whole application';

