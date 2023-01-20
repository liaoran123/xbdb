# 小白数据库系统（考据级全文索引数据库系统）  
对goleveldb进行数据系统封装  

## 功能核心亮点  
内置**考据级全文索引**，无需分词库，却永远没有搜索不到的词。  

## 功能
创建表。 
默认第一个字段必须是主键 
字段类型：string,int,int64,float64,time
其他字段可以设置为索引、组合索引，全文索引等。
添加、删除、修改、查询功能。
查询，可以通过主键查询，索引查询。
索引查询，可以是精准匹配和前缀匹配。


## 案例 
[千佛网](http://www.soufoshuo.com)是由客户端和[research](https://github.com/liaoran123/research)、[jianjie](https://github.com/liaoran123/jianjie) 两个微服务系统构建。
所有数据皆是使用小白数据库系统。

 
 
 
