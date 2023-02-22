# 冬季小合作-极简版抖音

## 基于gin框架的极简版抖音青训营项目

## 运行方式：

-   首先你需要有 go 的环境

```shell
brew install go
```

-   克隆本项目

```shell
git clone https://github.com/yulinJoseph/ByteDanceTeenageCamp.git
cd ByteDanceTeenageCamp
```

-   安装依赖

```shell
go mod tidy
```

-   或者选择在本机安装(Optional)

```shell
brew install etcd
brew install minio
brew install mysql

brew services start etcd
brew services start minio
brew services start mysql
```

- 对 minio 进行配置

  1. 进入 minio 的控制台，默认地址为 `http://localhost:9000`，默认账号密码为 `minioadmin` 和 `minioadmin`
  2. 创建一个名为 `douyin` 的 `bucket`
  3. 将 `bucket` 的权限(`Access Policy`)设置为 `public`

- 运行

  -   如果你不会查找端口被哪个占用的话，我建议你开三个终端，分别运行下面三条指令。关闭时只需 `Ctrl + C` 即可

  ```shell
  make run_api
  make run_user
  make run_video
  ```

  -   反之

  ```shell
  make run
  ```

## 模块说明：

#### Controllers

对于Controllerss层级的所有函数实现有如下规范：

所有的逻辑由代理对象进行，完成以下两个逻辑

- 解析得到参数
- 开始调用下层逻辑

#### Service

对于service层级的函数实现由如下规范：

同样由一个代理对象进行，完成以下三个或两个逻辑

当上层需要返回数据信息，则进行三个逻辑：

- 检查参数。
- 准备数据。
- 打包数据。

当上层不需要返回数据信息，则进行两个逻辑：

- 检查参数。
- 执行上层指定的动作。

#### Models

对于models层的各个操作，没有像service和Controller层针对前端发来的请求就行对应的处理，models层是面向于数据库的增删改查，不需要考虑和上层的交互。

而service层根据上层的需要来调用models层的不同代码请求数据库内的内容。

