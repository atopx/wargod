# WARGOD

The God of War in ***League of Legends*** games, It's legitimate game assistant.

### 预期核心功能

1. 修改状态
2. 修改生涯背景
3. 修改段位
4. 战绩查询 / 分析
5. 自动接受对局
6. 选择英雄事件 -> opgg -> 符文导入 -> 出装、技能

### 技术开发项

1. 游戏进程发现，后台提权
2. 状态机定义
3. 游戏客户端keepalive检测
4. LCU接口模型定义
5. websocket监听游戏事件
6. 处理游戏事件
7. 三方数据采集

### 备注：

项目处于初始阶段，感兴趣的可以一起参与

- 前端技术栈不限(优先使用原生ts/css/html保证轻量和性能)
- 后端rust或者有其他编程经验想学习rust的

### 参考：

1. lcu-api文档: [lcu-schema](https://www.mingweisamuel.com/lcu-schema/tool/#/Plugin%20lol-chat/PutLolChatV1Me)
2. 不错的项目, 可以作为目标: [Seraphine](https://github.com/Zzaphkiel/Seraphine)

### 目标:

1. 功能完整
2. 性能，兼顾轻量
3. 代码可读
