# 阿里云RDS MySQL数据库连接
⏰2026/02/22
## 前期准备
[下载DataGrip](https://www.jetbrains.com/datagrip/)
## 在阿里云上创建数据库实例
#### 按[操作指南](https://developer.aliyun.com/adc/tutorial/2922079?spm=a2c6h.28997733.0.0.3a7c69a2X9jQRl)进行，无需下载MySQL Workbench
- 完成参数配置
- 创建账号
- 设置白名单，**加载本机公网IP**
- 开通外网地址，外网地址和外网端口将在连接实例时使用
## 用DataGrip连接数据库
- 获取连接信息：操作台进入->云数据库RDS/实例列表/数据库连接，复制**外网地址和端口**
- 在DataGrip中新建连接，填入相关信息，如“主机”填写外网地址，“端口”填写外网端口，“用户名”和“密码”填写之前创建的账号的用户名和密码，“数据库”填写数据库名，**点击“测试连接”测试连接是否成功**

## 用Go代码连接数据库
#### 安装Go语言的MySQL驱动

- 在本地创建项目文件夹
    ```
    # 如果你想放在桌面（以Windows为例）
    cd C:\Users\你的用户名\Desktop

    # 创建一个专门放这个项目的文件夹
    mkdir aliyun-rds-demo

    # 进入这个文件夹
    cd aliyun-rds-demo
    ```

- 初始化 Go 项目
路径在 `aliyun-rds-demo` 下，输入`go mod init aliyun-rds-demo`

- 安装 MySQL 驱动
输入`go get -u github.com/go-sql-driver/mysql`

- 编写go代码
需要修改几处信息：阿里云账号、密码、公网地址、数据库名（新建数据库）

- 运行代码## 协作测试
这是同事添加的内容
这是同事的第二次修改
