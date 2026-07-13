# TaskFlow

> A Distributed Task Scheduling Platform built with Go.

## 项目简介

TaskFlow 是一个基于 Go 构建的分布式任务调度平台（Distributed Task Scheduling Platform）。

平台本身只负责任务的创建、管理、调度以及生命周期维护，不关心任务的具体实现。所有任务能力均由独立的 Executor 提供，TaskFlow 根据任务类型（TaskType）将任务分发给对应的 Executor 执行，实现平台与业务能力的完全解耦。

---

## 项目目标

* 学习企业级 Go 后端开发
* 学习 RESTful API 设计
* 学习 Gin + GORM + MySQL
* 学习 Redis 在缓存、限流、幂等等场景中的应用
* 学习 Kafka 消息队列及异步任务处理
* 学习 gRPC 服务间通信
* 学习 Docker 容器化部署
* 学习 Kubernetes 集群部署
* 学习分布式任务调度系统设计
* 构建一个可持续迭代、可扩展、实用的工作平台



## 技术栈

### Backend

* Go
* Gin
* GORM
* MySQL

### Frontend

* Vue 3
* Vite
* Element Plus

### Middleware

* Redis
* Kafka
* JWT

### RPC

* gRPC

### Deployment

* Docker
* Kubernetes

---

## 当前开发进度

* [x] 项目初始化
* [x] 用户登录
* [x] 用户注册
* [x] JWT 身份认证
* [ ] Task 模块
* [ ] Dashboard
* [ ] Redis
* [ ] Kafka
* [ ] gRPC
* [ ] Docker
* [ ] Kubernetes
* [ ] Executor



---

## 文档

详细设计文档位于 `docs/` 目录：

* architecture.md —— 系统架构设计
* database.md —— 数据库设计
* api.md —— RESTful API 设计
* task-lifecycle.md —— Task 生命周期
* executor.md —— Executor 设计
* roadmap.md —— 项目开发路线
* roadmap.md —— 部署方案

---

## 项目规划

**v0.1**

* 用户系统
* JWT
* Task CRUD

**v0.2**

* Redis
* 缓存
* 限流
* 幂等

**v0.3**

* Kafka
* Worker
* 异步任务
* Retry

**v0.4**

* gRPC
* Executor
* 服务拆分

**v0.5**

* Docker
* Docker Compose

**v0.6**

* Kubernetes

**v1.0**

* Executor
* 完整分布式任务调度平台

---

## License

This project is licensed under the MIT License.
