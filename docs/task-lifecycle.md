## Task 生命周期

TaskFlow 中每个任务（Task）都会经历完整的生命周期。平台通过任务状态（Status）描述任务当前所处的阶段，所有模块均以任务状态作为统一的业务依据。

### 生命周期

```text
             Create Task
                  │
                  ▼
             CREATED
                  │
                  ▼
             QUEUED
                  │
                  ▼
             RUNNING
             ┌────┴────┐
             │         │
             ▼         ▼
         SUCCESS    FAILED
                         │
               Retry? ───┤
                 Yes      │ No
                  │       ▼
                  └──► QUEUED
```

------

### 状态说明

| 状态    | 说明                                   |
| ------- | -------------------------------------- |
| CREATED | 任务已创建，并成功写入数据库。         |
| QUEUED  | 任务已进入消息队列，等待 Worker 消费。 |
| RUNNING | Worker 已开始执行任务。                |
| SUCCESS | 任务执行成功。                         |
| FAILED  | 任务执行失败，可根据策略决定是否重试。 |

------

### 状态流转

正常情况下：

```text
CREATED
    │
    ▼
QUEUED
    │
    ▼
RUNNING
    │
    ▼
SUCCESS
```

执行失败：

```text
RUNNING
    │
    ▼
FAILED
```

开启自动重试：

```text
FAILED
    │
    ▼
QUEUED
    │
    ▼
RUNNING
```

------

### 生命周期说明

- API Server 创建任务后，将任务信息写入数据库。
- 任务成功发送到消息队列后，状态更新为 `QUEUED`。
- Worker 消费任务后，状态更新为 `RUNNING`。
- Executor 完成任务后，根据执行结果更新为 `SUCCESS` 或 `FAILED`。
- 当系统启用自动重试策略时，失败任务可重新进入消息队列再次执行。

Task 生命周期是整个系统的核心状态机，后续 Redis、Kafka、gRPC 以及 Executor 等模块均围绕该生命周期进行设计和实现。