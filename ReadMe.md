# kubernetes集群管理项目项目

## 准备
1.设置环境变量
```bash
export SECRET_KEY="YMyn4gXjomxvYBRmheSE2agb3ZsFKQXN"

export PORT=":8888"

export LOG_LEVEL="info"

export JWT_EXPIRE_TIME=3600

export MYSQL_ADDRESS="127.0.0.1"

export MYSQL_PORT=3306

export MYSQL_USERNAME=root

export MYSQL_PASSWORD="xxx"

export MYSQL_DBNAME="xxx"

export MAX_IDLE_CONNECTION=25

export MAX_OPEN_CONNECTION=25
```
SECRET_KEY: jwt secret key. 设置复杂点. debug模式自动使用"joe12345"作为secret_key

PORT: 应用监听端口, 不设置默认为: 8080

LOG_LEVEL: 日志级别, 不设置默认为: debug

JWT_EXPIRE_TIME: jwt token过期时间. 不设置默认为: 86400秒

MYSQL_ADDRESS: 数据库连接地址

MYSQL_PORT: 数据库端口, 默认: 3306

MYSQL_USERNAME: 数据库用户, 默认: root

MYSQL_PASSWORD: 数据库密码

MYSQL_DBNAME: 库名

MAX_IDLE_CONNECTION: 连接池中最大空闲连接数, 默认: 25

MAX_OPEN_CONNECTION: 最大打开连接数, 默认: 25

