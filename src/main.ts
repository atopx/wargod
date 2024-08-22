import {listen} from "@tauri-apps/api/event";
import {invoke} from "@tauri-apps/api/core";
import {WebviewWindow} from "@tauri-apps/api/webviewWindow";

let opggWindow: WebviewWindow | null = null;

function setupEventListeners() {
    listen<string>("game-status-change", (event) => {
        const status = event.payload;
        console.log(`收到游戏状态事件: ${status}`);

        switch (status) {
            case "Connecting":
                console.log("状态：连接中，跳转到初始加载页面");
                // window.location.href = '/index.html'; // 使用绝对路径
                window.location.replace("/index.html");
                break;

            case "Connected":
                console.log("状态：已连接，跳转到首页");
                window.location.replace("/pages/index.html"); // 使用绝对路径
                break;

            case "Gaming":
                console.log("状态：游戏中，跳转到游戏页面并打开OPGG窗口");
                window.location.replace("/pages/game.html"); // 使用绝对路径
                openOPGGWindow();
                break;

            case "ConnectedFromGaming":
                console.log("状态：从游戏返回，跳转到生涯页面并关闭OPGG窗口");
                window.location.replace("/pages/stat.html"); // 使用绝对路径
                closeOPGGWindow();
                break;

            case "ConnectingFromGaming":
                console.log("状态：从游戏中断，跳转到初始加载页面并提示是否重连");
                window.location.replace("/index.html"); // 使用绝对路径
                openReconnectDialog();
                break;

            default:
                console.warn(`未识别的状态: ${status}`);
        }
    });
}

// document.addEventListener('DOMContentLoaded', () => {

// });

setupEventListeners(); // 启动事件监听器

// 在组件加载完成后，调用后端方法，要求重新发送当前状态
invoke("mock_connected")
    .then(() => {
        console.log("已请求后端重新发送当前游戏状态");
    })
    .catch((err) => {
        console.error("请求重新发送游戏状态失败:", err);
    });

async function openOPGGWindow() {
    if (!opggWindow) {
        opggWindow = new WebviewWindow("opgg", {
            url: "./pages/opgg.html",
            width: 1134,
            height: 826,
            resizable: false,
        });

        opggWindow.once("tauri://error", (e) => {
            console.error("OPGG窗口创建失败:", e);
        });
    }
}

async function closeOPGGWindow() {
    if (opggWindow) {
        await opggWindow.close();
        opggWindow = null;
    }
}

function openReconnectDialog() {
    const reconnect = confirm("是否快速重连？");
    if (reconnect) {
        // 调用后端重连方法
        invoke("reconnect_game").catch((err) => console.error("重连失败:", err));
    }
}
