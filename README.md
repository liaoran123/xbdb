# 小白数据库系统（考据级全文索引数据库系统）  
对goleveldb进行系统封装  

## 功能核心亮点  
内置**考据级全文索引**，无需分词库，却永远没有搜索不到的词。  

## 功能
**创建表**。 <br>
默认第一个字段必须是主键 <br>
**字段类型**：string,int,int64,float64,time<br>
其他字段可以设置为**索引、组合索引，全文索引**等。<br>
**添加、删除、修改、查询**功能。<br>
**查询**，可以通过主键查询，索引查询，可以是精准匹配和前缀匹配。<br>




 
 
