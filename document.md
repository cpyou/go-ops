
## grpc 认证

[参考 煎鱼的迷之博客的github开源项目](https://github.com/EDDYCJY/blog)

### TLS认证初始化工作

证书生成
```shell
openssl ecparam -genkey -name secp384r1 -out server.key
```
自签公钥
```shell
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
```

填写信息
```shell
Country Name (2 letter code) []:
State or Province Name (full name) []:
Locality Name (eg, city) []:
Organization Name (eg, company) []:
Organizational Unit Name (eg, section) []:
Common Name (eg, fully qualified host name) []:go-ops
Email Address []:
```
将生成的证书放到 config/server/ 目录下

```shell
├── conf
│   ├── app.ini
│   ├── auth_model.conf
│   ├── auth_policy.csv
│   └── server
│       ├── server.key
│       └── server.pem

```

## cron定时任务
[参考](https://www.jianshu.com/p/626acb9549b1) 


>用过 linux 的应该对 cron 有所了解。linux 中可以通过 crontab -e 来配置定时任务。不过，linux 中的 cron 只能精确到分钟。而我们这里要讨论的 Go 实现的 cron 可以精确到秒，除了这点比较大的区别外，cron 表达式的基本语法是类似的。（如果使用过 Java 中的 Quartz，对 cron 表达式应该比较了解，而且它和这里我们将要讨论的 Go 版 cron 很像，也都精确到秒）
cron 表达式代表了一个时间集合，使用 6 个空格分隔的字段表示。

`秒(Seconds)   分(Minutes)	时(Hours)	日(Day of month)	  月(Month)	  星期(Day of week)`
1. 星号(*)
表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月

2. 斜线(/)
表示增长间隔，如第1个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后每隔 15 分钟执行一次（即 3、18、33、48 这些时间点执行），这里也可以表示为：3/15

3. 逗号(,)
用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行

4. 连字号(-)
表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 直接每个小时（包括9和17）

5. 问号(?)
只用于日(Day of month)和星期(Day of week)，表示不指定值，可以用于代替 *
