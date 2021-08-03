> **请克隆至 GOPATH/src/NetworkDisk** 目录下并启用 **GO111MODULE**

# NetworkDisk

## PKG

| 包名        | 作用              | 描述                         |
| ----------- | ----------------- | ---------------------------- |
| config      | 处理部分配置信息  | 处理一些配置                 |
| Database    | 负责MySQL相关操作 | 历史遗留                     |
| Handlers    | 路由处理          | 登录，注册，注销，上传，下载 |
| Middlewares | 中间件            | 信息验证，令牌验证           |
| Models      | 模型              | JWT，Redis，User             |
| Routers     | 路由设置          | ….                           |



### config

> json序列化

```go
func Init(ConfPath string) (result Conf)
```

```go
type Sql struct {
	SqlName string		//数据库名
	SqlUserName string	//数据库登录用账户名
	SqlUserPwd string	//数据库登录用账户密码
	SqlAddr string		//数据库地址
}

type Conf struct {
	Sql
	Addr string	//服务器地址
}
```



```json
{
  "SqlName":"networkdisk",
  "SqlUserName": "root",
  "SqlUserPwd": "root",
  "SqlAddr":"127.0.0.1:3306",
  "Addr":"localhost:8080"
}
```

### Database

```go
func InitGorm(sql *config.Sql) *gorm.DB 
```

### Handlers

| 定义                                                         | 描述             |
| :----------------------------------------------------------- | ---------------- |
| func Download() gin.HandlerFunc                              | 文件下载         |
| funcs(封装了一些常用的功能)                                  | 从JWT中获得ID    |
| func Login(redis redis.RedisPool,template jwt.Jwt) gin.HandlerFunc | 登录             |
| func Logout(pool redis.RedisPool) gin.HandlerFunc            | 注销             |
| func Register(db *gorm.DB) gin.HandlerFunc                   | 注册             |
| func Upload() gin.HandlerFunc                                | 上传文件         |
| func Getsharelinks(pool *Redis.RedisPool) gin.HandlerFunc    | 生成分享令牌     |
| func Usesharedlinks(db *gorm.DB,pool *Redis.RedisPool) gin.HandlerFunc | 使用分享令牌     |
| func Filelist(pool *Redis.RedisPool) gin.HandlerFunc         | 可下载的文件清单 |

### Middleware

| 定义                                                         | 描述                 |
| ------------------------------------------------------------ | -------------------- |
| func CheakJWT(pool redis.RedisPool,template jwt.Jwt) gin.HandlerFunc | 验证令牌正确性       |
| func CheakUserInfo(db *gorm.DB) gin.HandlerFunc              | 验证用户信息的正确性 |



#### CheakJWT

```mermaid
graph LR
获取token==>解析token==>获取Redis记录==>验证记录与token是否一致==>将ID存入上下文
```

#### CheakUserInfo

```mermaid
graph LR
绑定反序列化==>验证ID与密码==>将ID存入上下文
```

### Models

| 位置            | 描述             |
| --------------- | ---------------- |
| JWT/Jwt         | JWT相关定义      |
| Redis/RedisPool | Redis相关定义    |
| User            | User相关定义     |
| File/File       | File相关定义     |
| File/Privilege  | 文件权限相关定义 |
| funcs           | 一些通用函数     |

#### Jwt

```go
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Iss string 	`json:"iss"`
	Exp uint 	`json:"exp"`
	Sub string 	`json:"sub"`
	Aud int 	`json:"aud"` //用户ID
	Ndf uint 	`json:"ndf"`
	Iat uint 	`json:"iat"`
	Jti uint 	`json:"jti"`
}

type Jwt struct{
	Header Header
	Payload Payload
	Secret string
}
```

| 方法                                    | 描述                             |
| --------------------------------------- | -------------------------------- |
| func (j Jwt)Encoding() string           | 基于当前的Header与Payload计算JWT |
| func (j *Jwt)Decoding(jwt string) error | 基于jwt刷新Payload的值           |

#### RedisPool

```go
type RedisPool struct {
	Read 		string
	Write 		string
	IdLeTimeout	int
	MaxIdle		int
	MaxActive	int
	rpool *redis.Pool
	wpool *redis.Pool
}
```

| 方法                                                         | 描述          |
| ------------------------------------------------------------ | ------------- |
| func (r *RedisPool)Init()                                    | 初始化连接池  |
| func (r RedisPool)SET(args ...interface{}) error             | 执行set指令   |
| func (r RedisPool)GET(key string) (string,error)             | 执行get指令   |
| func (r RedisPool)DEL(key string) error                      | 执行del指令   |
| func (r RedisPool)SADD(args ...interface{}) error            | 执行sadd      |
| func (r RedisPool)SISMEMBER(args ...interface{}) (int64,error) | 执行sismember |
| func (r RedisPool)SMEMBERS(key string) (re []string,err error) | 执行smembers  |

#### User

```go
type User struct {
	Uid int `gorm:"primaryKey"`
	Name string	`gorm:"string not null"`//用户名
	Pwd string `gorm:"string not null"`//用户密码
}
```

| 方法                                                | 描述                                  |
| :-------------------------------------------------- | ------------------------------------- |
| func (u *User)Save(db *gorm.DB) (err error)         | 在数据库中创建用户                    |
| func (u *User)Load(db *gorm.DB,uid int) (err error) | 根据提供的UID读取用户信息             |
| func (u *User)PwdIsRight(db *gorm.DB) bool          | 判断密码是否正确，如果不正确返回false |
| func (u *User)IsExist(db *gorm.DB) bool             | 判断是否存在，如果不存在返回false     |

#### File

```go
type File struct {
	PathName string `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Size     int64   `gorm:"not null"`
	Owner    int    `gorm:"not null"` //拥有权用户ID
}
```

| 方法                                  | 描述         |
| ------------------------------------- | ------------ |
| func (f *File)Save(db *gorm.DB) error | 保存到数据库 |

#### Privilege

```go
type Privilege struct {
	Pri 	 uint 	`gorm:"primaryKey"`
	PathName string `gorm:"not null"`
	Owner    int    `gorm:"not null"` //拥有权用户ID
	User	 int//有使用权的用户
}
```


| 方法                                       | 描述             |
| ------------------------------------------ | ---------------- |
| func (p *Privilege)Save(db *gorm.DB) error | 将记录存入数据库 |

#### funcs

| 方法                                                    | 描述                          |
| ------------------------------------------------------- | ----------------------------- |
| func OwnerKey(id int) string                            | 根据id获取所有权在Redis中的键 |
| func UserKey(id int) string                             | 根据id获取使用权在Redis中的键 |
| func Out(db *gorm.DB,redis Redis.RedisPool) (err error) | 将使用权与拥有权缓存至Redis   |

### Routers

```go
group:=server.Group("/", Middlewares.CheakJWT(pool,template))
{
	group.POST("/logout", Handlers.Logout(pool))
	group.POST("/upload", Handlers.Upload(db,pool))
	group.POST("/usesharedlinks", Handlers.Usesharedlinks(db,pool))
	group.GET("/getsharelinks", Handlers.Getsharelinks(pool))
	group.GET("/download", Handlers.Download(pool))
	group.GET("/filelist",Handlers.Filelist(pool))
}
server.POST("/register", Handlers.Register(db))
server.POST("/login", Middlewares.CheakUserInfo(db),Handlers.Login(pool,template))
```

## 路由

### /login

| Body=>form-data | 描述   |
| --------------- | ------ |
| UID             | 用户ID |
| Pwd             | 密码   |

### /register

| Body=>form-data | 描述 |
| --------------- | ---- |
| Pwd             | 密码 |
| Name            | 昵称 |

### /logout

| Headers       | 描述     |
| ------------- | -------- |
| Authorization | 身份令牌 |

### /upload

| Body=>form-data | 描述       |
| --------------- | ---------- |
| file            | 上传的文件 |

| Headers       | 描述                          |
| ------------- | ----------------------------- |
| Authorization | 令牌                          |
| Content-Type  | 值统一为: multipart/form-data |

### /download

| Body=>form-data | 描述                                |
| --------------- | ----------------------------------- |
| filepath        | 当前用户目录下的文件路径(以“/”开头) |
| filename        | 文件名                              |

| Headers       | 描述     |
| ------------- | -------- |
| Authorization | 身份令牌 |

### /getsharelinks

| Headers       | 描述     |
| ------------- | -------- |
| Authorization | 身份令牌 |

| Body=>form-data | 描述                         |
| --------------- | ---------------------------- |
| filepath        | 文件所在目录                 |
| filename        | 文件名                       |
| uid             | 要分享的用户ID,如果为0则通用 |

### /usesharedlinks

| Headers       | 描述     |
| ------------- | -------- |
| Authorization | 身份令牌 |

| Body=>form-data | 描述       |
| --------------- | ---------- |
| link            | 分享的令牌 |

### /filelist

| Headers       | 描述     |
| ------------- | -------- |
| Authorization | 身份令牌 |

## 数据库

> 理想状态下应该是下面的情况

```mysql
mysql> desc users;
+-------+----------+------+-----+---------+----------------+
| Field | Type     | Null | Key | Default | Extra          |
+-------+----------+------+-----+---------+----------------+
| uid   | bigint   | NO   | PRI | NULL    | auto_increment |
| name  | longtext | YES  |     | NULL    |                |
| pwd   | longtext | YES  |     | NULL    |                |
+-------+----------+------+-----+---------+----------------+
3 rows in set (0.01 sec)
```

```mysql
mysql> desc files;
+-----------+--------------+------+-----+---------+-------+
| Field     | Type         | Null | Key | Default | Extra |
+-----------+--------------+------+-----+---------+-------+
| path_name | varchar(191) | NO   | PRI | NULL    |       |
| name      | longtext     | NO   |     | NULL    |       |
| size      | bigint       | NO   |     | NULL    |       |
| owner     | bigint       | NO   |     | NULL    |       |
+-----------+--------------+------+-----+---------+-------+
4 rows in set (0.00 sec)
```

```mysql
mysql> desc privileges;
+-----------+--------------+------+-----+---------+-------+
| Field     | Type         | Null | Key | Default | Extra |
+-----------+--------------+------+-----+---------+-------+
| pri       | bigint       | NO   | PRI | NULL    |       |
| path_name | varchar(191) | NO   | MUL | NULL    |       |
| owner     | bigint       | YES  |     | NULL    |       |
| user      | bigint       | YES  |     | NULL    |       |
+-----------+--------------+------+-----+---------+-------+
4 rows in set (0.00 sec)
```

