# wargod
The God of War in League of Legends games, It's legitimate game assistant.



OP.GG 窗口大小:  640 x 975

main窗口大小: 1134, 826



我正在模仿这个桌面程序自己开发一个轻量级且高性能的tauri重制版，我使用的技术栈是后端是rust tauri, 前端是typescript和原生html/css。

需要你为我的前端编写html/css/ts代码，，UI字体全局使用MiSans

- 尺寸：这个窗口的布局需要是弹性的，最小尺寸 1134 * 826
- 字体：全局使用MiSans
- 颜色：
  - 游戏对局胜卡片背景色：#2839b01b
  - 游戏对局负卡片背景色：#28d3190c
  - 游戏对局重开作废背景色：#28a2a2a2

我们假设现在所有的图片和UI资源都存在，路径如下：
  - 非游戏内UI图标放在 /src/asset/icons/*.svg
  - 物品道具图片在 /src/asset/game/item/*.png
  - 英雄技能图片放在 /src/asset/game/spell/*.png
  - 英雄头像图片放在 /src/asset/game/champion/*.png
  - 玩家的头像图片放在 /src/asset/game/profile/*.png

要尽可能把前端效果做成与截图一致，所有的图片和UI设计资源你可以假设已经存在(你可以自由命名并添加注释, 我会处理)，如果你有问题或需要可以告诉我。


# 首先我要明确两个细节：

1. 我们使用的tauri技术，你作为前端可以假设所有的后端函数和事件都已经实现，只需要调用后端的异步函数或者是事件，所有前端所使用的你都可以先定义，我稍后会按照你定义的函数、入参和出参编写对应的逻辑。

2. 前端使用的技术是TypeScript(5.5)，不要使用JS，我需要你先游览一下TS5.5的文档：https://www.typescriptlang.org/docs/

   
# 接下来我们要调整多页面的路由和布局，我们一步一步来实现，首先是全局固定(在所有页面)的左侧垂直导航菜单栏(每个菜单都使用ICON+文字)，分上下两部分：

### 垂直导航菜单栏靠上布局的包含：
  
1. 首页，src/pages/index.html
2. 生涯：src/pages/stat.html
3. 查询：src/pages/search.html
4. 对局：src/pages/game.html
5. OPGG：这个比较特殊，会弹出一个新窗口(FixSize: 1134*826)

### 垂直导航菜单栏靠下布局的只有一个`设置(src/pages/settings.html)`


# 现在请你编写导航菜单和布局设计，所有页面的主内容暂时先不要编写，我们晚一些处理。


# 现在我们需要编写src/index.html和src/main.ts, 用来处理游戏是否启动的逻辑：

1. 我们会有一个简单的loading动画，可以是任意用户友好的形式，你来决定动画样式
2. 我们还会有一个全局的事件监听器，用来监听后端与游戏连接的状态
   - 当收到后端从状态 `Connecting` 到 `Connected` 的消息时，跳转到 `src/pages/index.html`
   - 当收到后端从 `Connected` 到 `Connecting` 的消息时，跳转到 `src/index.html`，继续loading
   - 当收到后端从 `Connected` 到 `Gaming` 的消息时，, 跳转到 `src/pages/game.html`,同时异步打开 `OPGG` 窗口
   - 当收到后端从 `Gaming` 到 `Conected` 的消息时，跳转到 `src/pages/stat.html`, 同时关闭 `OPGG` 窗口
   - 当收到后端从 `Gaming` 到 `Connecting` 的消息时，跳转到 `src/index.html`，同时打开一个对话框询问“是否快速重连”


# 我们需要对刚才的实现做一些简单的优化

1. main.ts的位置是在`src/main.ts`
2. 当loading时除了动画还需要补充一段用户友好的提示：“正在等待游戏启动...”
3. 由于我们使用的是tauri-v2.rc, 打开新窗口的方法已经变化，请从官方文档中寻找并告诉我如何修复，文档地址：https://v2.tauri.app/start/


# 修复几个小BUG

1. WebviewWindow是从"@tauri-apps/api/webviewWindow"导入。
2. 由于main.rs的位置是在`src/main.rs`, 所以跳转到`src/pages/*.html`时的路径需要修正
3. opgg的窗口不是打开外部，而是`src/pages/opgg.html`
4. 请把`src/index.html`中的loading-container高度调整为98vh,这样不会超出全局高出而出现滚动条


# 由于现在是开发阶段，并没有实际打开游戏，请帮我编写一个后端的模拟事件，让其在启动3s后发送 connecting->connected 的状态，使前端可以跳转到首页

# 你需要修改前端：
1. import { invoke } from '@tauri-apps/api/core';
2. import { WebviewWindow } from '@tauri-apps/api/webviewWindow';
3. 在收到事件消息中增加日志输出
