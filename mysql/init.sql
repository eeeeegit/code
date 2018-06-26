drop TABLE if exists userinfo;

CREATE TABLE `userinfo`(
    `account` VARCHAR(16) NOT NULL,
    `password` VARCHAR(32) NULL,
    `score` INT NOT NULL DEFAULT 0,
    `nickname` VARCHAR(64) NULL DEFAULT NULL,
    `createtime` DATETIME NULL DEFAULT NULL,
    PRIMARY KEY (`account`)
) CHARACTER SET utf8mb4;

CREATE TABLE `gobang` (

)