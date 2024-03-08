#!/usr/bin/env bash

printf "\n MYSQL CREATE DATABASE: START ----------------------------------------------------------\n"
mysql -uroot -proot -e "CREATE DATABASE testlocal"
mysql -uroot -proot -e "USE testlocal"
printf "\n MYSQL CREATE DATABASE: END ----------------------------------------------------------\n"

printf "\n MYSQL CREATE USER: START ----------------------------------------------------------\n"
mysql -uroot -proot -e "CREATE USER 'testlocal_WPROD'@'%' IDENTIFIED BY 'root'"
mysql -uroot -proot -e "GRANT ALL PRIVILEGES ON * . * TO 'testlocal_WPROD'@'%'"
mysql -uroot -proot -e "FLUSH PRIVILEGES"
mysql -uroot -proot -e "CREATE USER 'testlocal_RPROD'@'%' IDENTIFIED BY 'root'"
mysql -uroot -proot -e "GRANT ALL PRIVILEGES ON * . * TO 'testlocal_RPROD'@'%'"
mysql -uroot -proot -e "FLUSH PRIVILEGES"
printf "\n MYSQL CREATE USER: END ----------------------------------------------------------\n"