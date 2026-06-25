# Takeout Backend / 外卖后端服务

An entry-level food delivery backend built with Go + Gin.
This project splits the backend into three lightweight services so the frontend can call shop list, menu, and order APIs independently.

一个使用 Go + Gin 搭建的入门级外卖后端项目。
当前后端已经拆分成 3 个轻量服务，方便前端分别调用商家列表、菜单、下单接口。

## Project Highlights / 项目亮点

- Three independent Go services, each focused on one business responsibility.
- Clean directory structure, easy to understand and easy to extend.
- Ready to match the current Vue frontend ports and request paths.
- Uses mock business data so the frontend can run without a database.
- Great as a learning demo for service splitting and frontend-backend integration.

- 3 个独立 Go 服务，每个服务只负责一个业务能力。
- 目录结构清晰，容易看懂，也方便后续扩展。
- 已适配当前 Vue 前端使用的端口和请求路径。
- 目前使用模拟业务数据，不接数据库也能直接联调。
- 很适合拿来学习服务拆分和前后端对接。

## Tech Stack / 技术栈

| Technology | Version | Description |
|------------|---------|-------------|
| Go | 1.25.5 | Backend language |
| Gin | 1.12.0 | HTTP web framework |

| 技术 | 版本 | 说明 |
|------|------|------|
| Go | 1.25.5 | 后端开发语言 |
| Gin | 1.12.0 | HTTP Web 框架 |

## Project Structure / 项目结构

```text
backend-go/
├── shop-service/
│   └── main.go
├── menu-service/
│   └── main.go
├── order-service/
│   └── main.go
├── go.mod
└── go.sum
```

### Service Overview / 服务说明

- `shop-service`: provides the shop list API on port `8080`.
- `menu-service`: provides the menu API on port `8002`.
- `order-service`: provides the order API on port `8003`.

- `shop-service`：提供商家列表接口，监听端口 `8080`。
- `menu-service`：提供菜单接口，监听端口 `8002`。
- `order-service`：提供下单接口，监听端口 `8003`。

## API Mapping / 接口对应关系

| Service | Port | Method | Path | Purpose |
|---------|------|--------|------|---------|
| shop-service | 8080 | GET | `/shop` | Return shop list |
| menu-service | 8002 | GET | `/menu/:id` | Return menu by shop id |
| order-service | 8003 | POST | `/order` | Create a mock order |

| 服务 | 端口 | 方法 | 路径 | 作用 |
|------|------|------|------|------|
| shop-service | 8080 | GET | `/shop` | 返回商家列表 |
| menu-service | 8002 | GET | `/menu/:id` | 根据商家 id 返回菜单 |
| order-service | 8003 | POST | `/order` | 创建模拟订单 |

## What Is Implemented / 实现了什么

### 1. Shop Service / 商家服务

- Returns shop name, rating, sales, delivery time, tags, and image.
- Matches the homepage card data used by the frontend.

- 返回商家名称、评分、销量、配送时间、标签、图片等字段。
- 已适配前端首页商家卡片的数据展示需求。

### 2. Menu Service / 菜单服务

- Returns menu data based on shop id.
- Includes `shop_name`, price, description, badge, image, rating, and sales.

- 根据商家 id 返回对应菜单数据。
- 包含 `shop_name`、价格、描述、角标、图片、评分、销量等字段。

### 3. Order Service / 下单服务

- Accepts frontend cart payload with `user_id` and `items`.
- Returns mock order info such as order id, total amount, delivery fee, and estimated time.

- 接收前端购物车提交的 `user_id` 和 `items`。
- 返回模拟订单信息，包括订单号、总金额、配送费和预计送达时间。

## Why This Structure / 为什么这样设计

- Splitting services makes responsibilities clearer.
- Each service can be started, tested, and extended independently.
- This structure is closer to real microservice thinking than putting everything into one file.
- It keeps the demo simple while still showing service boundaries.

- 拆成多个服务后，职责更清晰。
- 每个服务都可以单独启动、单独测试、单独扩展。
- 比把所有接口堆在一个文件里，更接近真实项目中的微服务思路。
- 同时又保留了演示项目该有的简单清爽。

## Quick Start / 快速开始

### Install Dependencies / 安装依赖

```bash
go mod tidy
```

### Run Services / 启动服务

Run each service in a separate terminal.

分别在不同终端中启动下面 3 个服务。

```bash
cd shop-service
go run .
```

```bash
cd menu-service
go run .
```

```bash
cd order-service
go run .
```

### Build All Services / 构建全部服务

```bash
go build ./...
```

## Example Requests / 请求示例

### Get Shop List / 获取商家列表

```bash
curl http://localhost:8080/shop
```

### Get Menu / 获取菜单

```bash
curl http://localhost:8002/menu/1
```

### Create Order / 创建订单

```bash
curl -X POST http://localhost:8003/order \
  -H "Content-Type: application/json" \
  -d "{\"user_id\":1,\"items\":[{\"id\":1,\"name\":\"Burger\",\"price\":32,\"count\":2}]}"
```

## What Problems It Solves / 解决了什么问题

- Allows the frontend to connect to stable APIs without waiting for a real database.
- Separates shop, menu, and order responsibilities for easier learning and maintenance.
- Makes local development clearer because each port has one obvious purpose.
- Provides a simple backend reference for beginners building delivery or e-commerce demos.

- 让前端在没有真实数据库的情况下，也能先接上稳定接口。
- 将商家、菜单、下单职责拆开，便于学习和维护。
- 本地开发更清晰，每个端口只负责一个明确功能。
- 为做外卖、电商演示项目的初学者提供一个简单后端参考。

## Who Can Use This as Reference / 适合哪些人参考

- Go beginners learning how to write simple HTTP services.
- Frontend developers who want a lightweight backend for local integration.
- Students building coursework, demo projects, or graduation practice.
- Developers exploring service splitting before introducing databases and middleware.

- 学 Go 后端的初学者。
- 需要一个轻量后端联调前端的前端开发者。
- 做课程设计、练手项目、毕业设计的同学。
- 想在接入数据库和中间件前，先理解服务拆分思路的开发者。

## Future Improvements / 后续可扩展方向

- Add shared models and common middleware.
- Replace mock data with MySQL or MongoDB.
- Add configuration files, logging, and environment-based ports.
- Add Docker and deployment support.

- 增加公共模型和公共中间件。
- 将模拟数据替换成 MySQL 或 MongoDB。
- 增加配置文件、日志能力和环境变量端口管理。
- 增加 Docker 和部署支持。

## License / 许可证

MIT License
