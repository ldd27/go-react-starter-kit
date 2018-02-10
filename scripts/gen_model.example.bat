@echo off
xorm reverse -s mysql root:123456@tcp(localhost:3306)/go_react_starter_kit?charset=utf8 goxorm  models
pause