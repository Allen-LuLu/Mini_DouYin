# `cmd/publish`
> Write kitex code here.
```
.
|-conf  # 该微服务的配置初始化 
|-dao # 数据库相关操作逻辑
|-example    # 测试代码
|-model # 该微服务的数据结构定义
|-script # kitex生成的代码，可以不用管
|-build.sh # kitex生成的脚本，用于编译微服务，不用修改
|-handler.go # kitex生成的代码，微服务的业务逻辑在此实现
|-kite.yaml # kitex生成的文件，可以不管
|-main.go # 该微服务的main文件，用于配置微服务的地址，服务注册，中间件等
```
**Reference:** Using `kitex -service demo demo.thrift` will get:
```
.
├── build.sh                     // 服务的构建脚本，会创建一个名为 output 的目录并生成启动服务所需的文件到里面
├── handler.go                   // 用户在该文件里实现 IDL service 定义的方法
├── kitex_gen                    // IDL 内容相关的生成代码
│   ├── base                     // base.thrift 的生成代码
│   │   ├── base.go              // thriftgo 的产物，包含 base.thrift 定义的内容的 go 代码
│   │   └── k-base.go            // kitex 在 thriftgo 的产物之外生成的代码
│   └── demo                     // demo.thrift 的生成代码
│       ├── demo.go              // thriftgo 的产物，包含 demo.thrift 定义的内容的 go 代码
│       ├── k-demo.go            // kitex 在 thriftgo 的产物之外生成的代码
│       └── demoservice          // kitex 为 demo.thrift 里定义的 demo service 生成的代码
│           ├── demoservice.go   // 提供了 client.go 和 server.go 共用的一些定义
│           ├── client.go        // 提供了 NewClient API
│           └── server.go        // 提供了 NewServer API
├── main.go                      // 程序入口
└── script                       // 构建脚本
    └── bootstrap.sh             // 服务的启动脚本，会被 build.sh 拷贝至 output 下
```
使用minio做对象存储
https://github.com/minio/minio
