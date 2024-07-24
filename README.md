# MySQL Status Watcher

MySQL Status Watcher 是一个用来监控 MySQL 数据库状态的工具，能够在Keepalived中检测 MySQL 的健康状态并做出相应的处理。

## 功能

- 监控 MySQL 服务是否正常运行
- 提供自定义脚本以便在检测到 MySQL 服务异常时执行特定操作

## 安装

1. 克隆仓库到本地：
   ```sh
   git clone https://gitee.com/hongshenghe/my-sqlstatus-watcher.git
   cd my-sqlstatus-watcher
   ```

2. 编译项目：
   ```sh
   make release
   ```

## 配置

运行一次程序，会生成一个默认的配置文件 `config.yaml`。你可以根据需要自定义配置文件。
编辑 `config.yaml` 文件，添加 MySQL 实例的配置信息。例如：
```yaml
mysql:
    type: socket
    host: /var/run/mysqld/mysqld.sock
    port: 3306
    user: root
    pass: root
    db: mysql

```
- `type`: MySQL 实例的连接方式，支持 `socket` 和 `tcp` 
- `host`: MySQL  实例的地址，如果 `type` 为 `socket`，则为 socket 文件的路径
- `port`: MySQL  实例的端口，如果 `type` 为 `tcp`，则为端口号
- `user`: MySQL  用户名
- `password`: MySQL 密码 ,这里需要填写加密后的密码（是的，符合各种安全检查要求）
- `db`: MySQL 数据库名 ,可以不填

## 加密密码

使用 `encrypt` 命令对密码进行加密：
```bash
./watcher encrypt -c hello
Encrypted code: 751fa0cfa59314ba9a39752b36611bdb
```
使用加密后的密码填写到配置文件中。

## 使用

1. 运行监控程序：
   ```sh
   cd build;./watcher connect 
   ```

2. 在Keepalived的配置文件中，添加检测脚本：
   ```sh
   vrrp_script chk_mysql {
       script "/path/to/watcher connect"
       interval 2
       weight 2
   }

   vrrp_instance VI_1 {
       ...
       track_script {
           chk_mysql
       }
   }
   ```

## 贡献

欢迎提交 Pull Requests 和 Issues 以改进此项目。

## 许可证

此项目遵循 MIT 许可证。详情请参阅 LICENSE 文件。
 