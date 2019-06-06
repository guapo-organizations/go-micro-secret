# go-micro-secret
golang微服务，网上找的一些公用工具库，自己在封装一下，如连接数据库、格式校验之类的



- help   里面存放的是一些工具，比如验证手机号、验证邮箱
- databse 是连接数据库的，目前只写了mysql的连接
- frame_toole 是我自己对grpc服务开启的时候所需要的一些组件导入的封装，不用也可以，不用的话自己写mysql、redis、等东西
- cache 缓存是redis
- consul 服务发现，里面提供了注册服务和查询服务
- tls 是tls/ssl加密文件