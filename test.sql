CREATE TABLE `userinfo` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64)  DEFAULT NULL,
    `password` VARCHAR(32)  DEFAULT NULL,
    `department` VARCHAR(64)  DEFAULT NULL,
    `email` varchar(64) DEFAULT NULL,
    PRIMARY KEY (`uid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8