# golib 提供自己的第三方包. 
> 目前主要包含如下一些三方库,一定要将lib和common区分开来. 
> 和业务没有任何关联. 
### 一、常见的数据集合.
  - 树
  - 栈
  - 堆
  - 等
  - fetch

### 二、APP 常用结构体封装. 
- `app` 结构体封装
  - HTTP Svr、TCP svr、RTC Svr.
  - 初始化. 
  - 服务启动. 
- 服务相关结构体封装. 
  - 日志封装. 使用zap进行封装. 
  - 配置文件封装. 考虑动态监听文件模式.



# Args
- FormatArgs2Struct 传入的结构体参数必须是指针类型. 
- FormatArgs2Struct 参数结构体字段必须全部是string类型.