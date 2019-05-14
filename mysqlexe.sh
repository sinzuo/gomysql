#!/bin/sh

HOSTNAME="127.0.0.1"                                           #数据库信息
PORT="4306"
USERNAME="root"
PASSWORD="jiangyibo"

DBNAME="mytest"                                                       #数据库名称
TABLENAME="test"                                            #数据库中表的名称

#也可以写 HOSTNAME="localhost"，端口号 PORT可以不设定

 #创建数据库
create_db_sql="create database IF NOT EXISTS ${DBNAME}"
mysql -h${HOSTNAME}  -P${PORT}  -u${USERNAME} -p${PASSWORD} -e "${create_db_sql}"

# 注意：-p${PASSWORD}中间不能有空格

#导入数据库
mysql -h127.0.0.1 -u root -pjiangyibo -P${PORT}  ${DBNAME} < ${TABLENAME}.sql

#导出数据库
#mysqldump  -h 112.74.65.117  -P5306 -uroot   wifi  -pifidc2403 > wifi_2017-10-26.sql


 #创建表
 #create_table_sql="create table IF NOT EXISTS ${TABLENAME} (  name varchar(20), id int(11) default 0 )"
 #mysql -h${HOSTNAME}  -P${PORT}  -u${USERNAME} -p${PASSWORD}  -D ${DBNAME} -e "${create_db_sql}"

 #插入数据
 #insert_sql="insert into ${TABLENAME} values('billchen',2)"
 #mysql -h${HOSTNAME}  -P${PORT}  -u${USERNAME} -p${PASSWORD} ${DBNAME} -e  "${insert_sql}"

 #查询
 #select_sql="select * from ${TABLENAME}"
 #mysql -h${HOSTNAME}  -P${PORT}  -u${USERNAME} -p${PASSWORD} ${DBNAME} -e "${select_sql}"

 #更新数据
 #update_sql="update ${TABLENAME} set id=3"
 #mysql -h${HOSTNAME}  -P${PORT}  -u${USERNAME} -p${PASSWORD} ${DBNAME} -e "${update_sql}"
 #mysql -h${HOSTNAME}  -P${PORT}  -u${USERNAME} -p${PASSWORD} ${DBNAME} -e  "${select_sql}"

 #删除数据
 #delete_sql="delete from ${TABLENAME}"
 #mysql -h${HOSTNAME}  -P${PORT}  -u${USERNAME} -p${PASSWORD} ${DBNAME} -e  "${delete_sql}"
 #mysql -h${HOSTNAME}  -P${PORT}  -u${USERNAME} -p${PASSWORD} ${DBNAME} -e  "${select_sql}"


