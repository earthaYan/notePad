# notePad
notePad后端
.
├── api
├── build
├── cmd
├── configs
├── deployments
├── docs
├── init
├── internal
│   └── pkg
├── pkg
├── README.md
├── scripts
├── third_party
├── tools
└── vendor
# 项目目录结构
## Go目录
### /cmd
项目主干，每个应用程序的目录名应该与可执行文件的名称相匹配。通常有一个小的 main 函数，从 /internal 和 /pkg 目录导入和调用代码，除此之外没有别的东西
### /internal
私有应用程序和库代码。不希望其他人在其应用程序或库中导入的代码，不局限于顶级 internal 目录。在项目树的任何级别上都可以有多个内部目录。应用程序共享的代码可以放在 /internal/pkg 目录下(
### /pkg
外部应用程序可以使用的库代码
### /vendor
应用程序依赖项(手动管理或使用你喜欢的依赖项管理工具)。`go mod vendor` 命令创建 /vendor 目录。如果未使用默认情况下处于启用状态的 Go 1.14，则可能需要在 go build 命令中添加 -mod=vendor 标志

## 服务应用程序目录
### /api
OpenAPI/Swagger 规范，JSON 模式文件，协议定义文件

## Web应用程序目录
### /web
特定于 Web 应用程序的组件:静态 Web 资产、服务器端模板和 SPAs
## 通用应用目录
### /configs
配置文件模板或默认配置。比如confd 或 consul-template 模板文件
### /init
System init（systemd，upstart，sysv）和 process manager/supervisor（runit，supervisor）配置。
### /scripts
执行各种构建、安装、分析等操作的脚本和功能分开的各个makefile
### /build
打包和持续集成。将云( AMI )、容器( Docker )、操作系统( deb、rpm、pkg )包配置和脚本放在 /build/package 目录下。 CI (travis、circle、drone)配置和脚本放在 /build/ci 目录中。请注意，有些 CI 工具(例如 Travis CI)对配置文件的位置非常挑剔。尝试将配置文件放在 /build/ci 目录中，将它们链接到 CI 工具期望它们的位置(如果可能的话)。
### /deployments或者/deploy
IaaS、PaaS、系统和容器编排部署配置和模板(docker-compose、kubernetes/helm、mesos、terraform、bosh)
### /test
额外的外部测试应用程序和测试数据
## 其他目录
### /docs
设计和用户文档(除了 godoc 生成的文档之外)
### /tools
项目的支持工具。注意，这些工具可以从 /pkg 和 /internal 目录导入代码。
### /examples
应用程序和/或公共库的示例
### /third_party
外部辅助工具，分叉代码和其他第三方工具(例如 Swagger UI)。
### /githooks
Git hooks。
### /assets
静态资源，如图像、徽标
### /website
前端项目