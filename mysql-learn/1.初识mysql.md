# 初识mysql
## mysql架构
客户、服务端架构

## mysql服务端的启动方式
```
mysqld // 很少用这个
mysql_safe
mysql.server // 一般采用这个
mysql_multi // 多个实例用这个
```

## mysql通信方式
```
TCP/IP
unix套接字
管道/共享内存
```

## mysql处理过程
```
1. 客户端
2. 处理连接 // 包含认证处理
3. 查询缓存 // 5.8后就没了
4. 语法解析
5. 查询优化
6. 存储引擎
7. 文件系统

2.3.4.5都是server层，而6.7是引擎层
```

## 引擎
```
innodb // 支持事务，行锁，聚簇索引(数据即索引)
myisam // 不支持事务，表锁，索引和数据分开
memory
```

## 设置引擎
```
show engines;

create table t (
    age int
) engine = innodb;
```