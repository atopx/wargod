<!-- src/App.vue -->
<script setup lang="ts">
import {onMounted, ref} from 'vue';
import Loading from './components/Loading.vue'; // 加载动画组件
import Game from './Game.vue';
import {EventsOn} from "../wailsjs/runtime";
import {GetState} from "../wailsjs/go/api/Api";
import {useNotification} from "naive-ui";

const GameState = ref("Disconnect")

interface notification {
  level: string;
  message: string;
}

const notice = useNotification();

onMounted(() => {
  GetState().then((value: string) => {
    GameState.value = value
  })
  EventsOn("GameState", function (state: string) {
    GameState.value = state;
  })
  EventsOn("notification", async function (data: notification) {
    switch (data.level) {
      case "info":
        notice.info({content: data.message})
        break
      case "warn":
        notice.warning({content: data.message})
        break
      case "error":
        notice.error({content: data.message})
        break
      case "success":
        notice.error({content: data.message})
        break
    }
  })


});
</script>

<template>
  <n-notification-provider>
    <!--  <Loading v-if="(GameState === 'Disconnect')"/>-->
    <Loading v-if="(GameState === 'Disconnect')"/>
    <Game v-else/>
  </n-notification-provider>
</template>
