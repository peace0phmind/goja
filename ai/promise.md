# Promise

## Properties
- `state`: Promise 的当前状态 (pending/fulfilled/rejected)
- `result`: Promise 的结果值
- `constructor`: 指向 Promise 构造函数

## Static Methods
- `all`: 等待所有 Promise 完成，返回所有结果的数组
- `allSettled`: 等待所有 Promise 完成，返回所有 Promise 的状态和结果
- `any`: 返回第一个成功的 Promise 的值，如果全部失败则返回 AggregateError
- `race`: 返回第一个完成的 Promise 的结果
- `reject`: 返回一个被拒绝的 Promise
- `resolve`: 返回一个解决的 Promise

## Instance Methods
- `then`: 添加成功和失败的回调函数，返回新的 Promise
- `catch`: 添加失败的回调函数，返回新的 Promise
- `finally`: 添加一个在 Promise 完成时执行的回调，无论成功或失败

## States
- `pending`: 初始状态，既不是成功，也不是失败状态
- `fulfilled`: 操作成功完成
- `rejected`: 操作失败

## Internal Features
- `PromiseCapability`: 用于创建和解决 Promise 的内部结构
- `PromiseReaction`: 处理 Promise 状态变化的内部机制
- `RejectionTracking`: 用于跟踪未处理的 Promise 拒绝
- `AsyncContextTracking`: 用于跟踪异步上下文
- `JobQueue`: 用于管理 Promise 微任务队列 