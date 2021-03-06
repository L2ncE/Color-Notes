# 全国大学生微信小程序大赛-多彩笔记

## 技术实现方案

### 语言以及相关框架

使用Go语言配合Gin框架完成开发

### 数据库选型

数据库使用了MySQL、Redis、Mongo。

#### MySQL
MySQL用于储存常规信息，例如用户的信息，笔记的相关信息（不包含笔记内容）

#### Redis
Redis作为缓存防止用户量过大使MySQL崩溃，除此之外笔记的点赞收藏等相关操作均用Redis集合完成，并且设置定时任务，每隔一段时间
就会将数据储存进MySQL数据库中，并且使用延时双删来保证数据一致。RDB冷备、AOF做热备使Redis持久化

#### MongoDB
由于MySQL存储大文本数据慢，所以在笔记存储上使用MongoDB，可以在不丢失速度的情况下对大文本数据进行增删改查

### 性能优化
使用pprof工具对项目进行了性能优化并改善

### 项目部署
使用了Docker容器来对项目进行部署，并且生成仅有18m的镜像

### 配置相关
使用第三方框架Viper对项目进行配置，并实现配置热重载