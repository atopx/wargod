<!-- src/App.vue -->
<script setup lang="ts">
import { onMounted, ref } from "vue";
import Loading from "./components/Loading.vue"; // 加载动画组件
import Game from "./Game.vue";
import { EventsOn } from "../wailsjs/runtime";
import { GetState } from "../wailsjs/go/api/Api";

const GameState = ref("Disconnect");

onMounted(() => {
    GetState().then((value: string) => {
        GameState.value = value;
    });
    EventsOn("GameState", function (state: string) {
        GameState.value = state;
    });
});
</script>

<template>
    <n-notification-provider>
        <!--  <Loading v-if="(GameState === 'Disconnect')"/>-->
        <Loading v-if="GameState === 'Disconnect'" />
        <Game v-else />
    </n-notification-provider>
</template>
