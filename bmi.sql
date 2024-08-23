CREATE DATABASE  IF NOT EXISTS `bmi`;

USE `bmi`;
drop table bmi_table;

create table  bmi_table
(
    id int NOT NULL AUTO_INCREMENT,
    Name varchar(255) NOT NULL,
    BMI decimal(14, 2),
    PRIMARY KEY (id)
);
