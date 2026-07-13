# TaskFlow



## 项目目录

```text
TaskFlow
├── frontend/          # Vue3 前端
├── backend/           # Go 后端
├── docs/              # 工程文档
└── README.md
```



---

## 系统架构

```text
                    Client
              (Web / Desktop)
                     │
                     ▼
               RESTful API
                     │
             TaskFlow API Server
                     │
         Task Lifecycle Management
                     │
            Scheduler / Dispatcher
                     │
          Task Queue (Kafka)
                     │
     ┌───────────────┼───────────────┐
     ▼               ▼               ▼
 AI Executor    Mail Executor   Image Executor
                     │
                     ▼
              Update Task Status
```



---

## 系统组成

TaskFlow 采用平台与执行器（Executor）分离的设计思想，平台只负责任务管理与调度，不负责具体任务执行。所有业务能力均由不同类型的 Executor 实现，平台与业务逻辑完全解耦。

```text
                           ┌──────────────────────┐
                           │      Frontend        │
                           │     Vue3 + Vite      │
                           └──────────┬───────────┘
                                      │ REST API
                                      ▼
                    ┌──────────────────────────────────┐
                    │        TaskFlow API Server        │
                    │        Go + Gin + GORM            │
                    └───────┬───────────────┬───────────┘
                            │               │
                            │               │
                            ▼               ▼
                     ┌────────────┐   ┌────────────┐
                     │   MySQL    │   │   Redis    │
                     │ Task/User  │   │ Cache等    │
                     └────────────┘   └────────────┘
                            │
                            ▼
                     ┌────────────┐
                     │   Kafka    │
                     │ Task Queue │
                     └──────┬─────┘
                            │
               ┌────────────┼────────────┐
               ▼            ▼            ▼
      AI Executor   Mail Executor   Image Executor
               │            │            │
               └────────────┴────────────┘
                            │
                            ▼
                     更新任务执行状态
```

各模块职责如下：

| 模块       | 职责                                                       |
| ---------- | ---------------------------------------------------------- |
| Frontend   | 提供 Web 管理界面，负责用户交互。                          |
| API Server | 提供 RESTful API，完成认证、任务管理、任务调度等核心业务。 |
| MySQL      | 存储用户、任务、执行记录等持久化数据。                     |
| Redis      | 提供缓存、限流、幂等控制、热点数据存储等能力。             |
| Kafka      | 作为任务消息队列，实现异步处理与服务解耦。                 |
| Executor   | 根据 TaskType 执行具体业务，并回写任务执行结果。           |

------

## 核心业务流程

TaskFlow 的核心思想是**平台只管理任务，不负责执行任务**。任务真正的执行由不同类型的 Executor 完成，因此平台可以不断扩展新的任务类型，而无需修改核心调度逻辑。

整体流程如下：

```text
Client
   │
   │ 创建任务
   ▼
API Server
   │
   ├── JWT 鉴权
   ├── 参数校验
   ├── 写入 MySQL
   ├── 写入 Kafka
   ▼
立即返回任务创建成功
   │
   ▼
Kafka
   │
   ▼
Worker
   │
   ├── 消费任务消息
   ├── 查询任务详情
   ├── 根据 TaskType 选择 Executor
   ▼
Executor
   │
   ├── 执行具体业务
   ├── 返回执行结果
   ▼
更新 MySQL 中的任务状态
```

整个流程采用异步设计。API Server 在任务写入数据库并发送到 Kafka 后立即返回，不等待任务执行完成，从而提高接口响应速度并降低系统耦合度。

Executor 不直接暴露给客户端，所有任务均通过 TaskFlow 平台统一管理。未来无论增加 AI、邮件、图片处理、文件转换等功能，都只需要新增对应的 Executor，而无需修改平台核心逻辑。