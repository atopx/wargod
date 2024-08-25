<!-- src/App.vue -->
<script setup lang="ts">
import {onMounted, ref} from 'vue';
import Loading from './components/Loading.vue'; // 加载动画组件
import Game from './Game.vue';
import {EventsOn} from "../wailsjs/runtime"; // Home 组件


const GameState = ref("Disconnect")

onMounted(() => {
  EventsOn("GameState", function (state: string) {
    GameState.value = state;
  })
});
</script>

<template>
  <!-- 根据 isLoading 的状态决定显示加载动画还是 Home 组件 -->
  <Loading v-if="(GameState === 'Disconnect')"/>
  <Game v-else/>
</template>
