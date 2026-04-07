# User Redeem Code System

一个基于 Go + Gin 框架开发的用户兑换码系统，提供用户管理和兑换码生成、使用、查询等功能。

## 技术栈

- **语言**：Go
- **框架**：Gin
- **数据库**：MySQL
- **缓存**：Redis
- **认证**：JWT
- **配置管理**：YAML

## 项目结构

```
├── cmd/                # 命令行工具
│   ├── gen/            # 代码生成器
│   └── test/           # 测试工具
├── config/             # 配置文件
├── docs/               # 文档
├── entity/             # 实体定义
├── global/             # 全局变量
├── handler/            # API 处理器
├── initialize/         # 初始化组件
├── logic/              # 业务逻辑
├── middleware/         # 中间件
├── model/              # 数据模型
├── pkg/                # 公共包
│   ├── helper/         # 辅助函数
│   ├── jwt/            # JWT 工具
│   ├── result/         # 统一返回格式
│   └── util/           # 工具函数
├── router/             # 路由配置
├── .gitignore          # Git 忽略文件
├── Makefile            # 构建脚本
├── config-example.yaml # 配置文件示例
├── go.mod              # Go 模块依赖
├── go.sum              # Go 模块校验和
└── main.go             # 项目入口
```

## 功能特性

### 用户管理

- 用户注册
- 用户登录
- 获取用户信息
- 更新用户信息
- 删除用户
- 获取用户列表

### 兑换码批次管理

- 创建兑换码批次
- 获取兑换码批次详情
- 获取兑换码批次列表
- 更新兑换码批次
- 删除兑换码批次

### 兑换码管理

- 生成兑换码
- 获取兑换码详情
- 获取兑换码列表
- 更新兑换码
- 删除兑换码
- 使用兑换码

## 安装和运行

### 环境要求

- Go 1.18+
- MySQL 5.7+
- Redis 5.0+

### 安装步骤

1. **克隆项目**

```bash
git clone https://github.com/1348453525/user-redeem-code-gin.git
cd user-redeem-code-gin
```

1. **安装依赖**

```bash
go mod tidy
```

1. **配置文件**
   - 复制配置文件示例

   ```bash
   cp config-example.yaml config.yaml
   ```

   - 根据实际情况修改 `config.yaml` 文件中的配置项

2. **数据库初始化**
   - 执行 `docs/sql.sql` 文件中的 SQL 语句创建数据库和表

3. **运行项目**

   ```bash
   go run main.go
   ```

   或使用 Makefile

   ```bash
   make run
   ```

## API 文档

### 用户相关接口

| 接口 | 方法 | 路径 | 认证 | 描述 |
| ------ | ---- | ---- | ------ | ------ |
| 用户注册 | POST | /Register | 否 | 注册新用户 |
| 用户登录 | POST | /Login | 否 | 用户登录 |
| 用户退出 | GET | /Logout | 是 | 用户退出登录 |
| 获取用户信息 | GET | /User/Info | 是 | 获取当前用户信息 |
| 获取用户列表 | GET | /User/GetList | 是 | 获取用户列表 |
| 更新用户信息 | PUT | /User/Update | 是 | 更新用户信息 |
| 删除用户 | DELETE | /User/Delete | 是 | 删除用户 |

### 兑换码批次相关接口

| 接口 | 方法 | 路径 | 认证 | 描述 |
| ------ | ---- | ---- | ------ | ------ |
| 创建兑换码批次 | POST | /RedeemCodeBatch/Create | 是 | 创建兑换码批次 |
| 获取兑换码批次详情 | GET | /RedeemCodeBatch/Detail | 是 | 获取兑换码批次详情 |
| 获取兑换码批次列表 | GET | /RedeemCodeBatch/GetList | 是 | 获取兑换码批次列表 |
| 更新兑换码批次 | PUT | /RedeemCodeBatch/Update | 是 | 更新兑换码批次 |
| 删除兑换码批次 | DELETE | /RedeemCodeBatch/Delete | 是 | 删除兑换码批次 |

### 兑换码相关接口

| 接口 | 方法 | 路径 | 认证 | 描述 |
| ------ | ---- | ---- | ------ | ------ |
| 获取兑换码详情 | GET | /RedeemCode/Detail | 是 | 获取兑换码详情 |
| 获取兑换码列表 | GET | /RedeemCode/GetList | 是 | 获取兑换码列表 |
| 更新兑换码 | PUT | /RedeemCode/Update | 是 | 更新兑换码 |
| 删除兑换码 | DELETE | /RedeemCode/Delete | 是 | 删除兑换码 |
| 使用兑换码 | POST | /RedeemCode/Use | 是 | 使用兑换码 |

## 配置文件说明

```yaml
# 项目启动端口
app:
  name: user-redeem-code-gin
  version: v0.0.1
  addr: :8080
  mode: dev # release, dev

# 数据库配置
db:
  driver: mysql
  host: 127.0.0.1
  port: 3306
  user: root
  password: root
  db: user_redeem_code
  charset: utf8mb4
  max_idle_conns: 10 # 最大空闲连接数
  max_open_conns: 100 # 最大连接数

# redis 配置
redis:
  addr: 127.0.0.1:6379
  password:
```

## 开发说明

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request！
