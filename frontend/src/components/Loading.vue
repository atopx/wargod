<script lang="ts" setup>
import {NSpin, useNotification} from 'naive-ui';
import {GameLauncher} from "../../wailsjs/go/api/Api";
import {ref} from "vue";


async function gameLauncher() {
  disableLauncher.value = true
  GameLauncher().then(() => {
  }).catch((err) => {
    const notification = useNotification()
    notification.error({
      content: err,
    })
  })
}

const disableLauncher = ref(false)

</script>

<template>
  <div class="loading-container">
    <n-spin size="large"/>
    <p>正在等待游戏启动...</p>
    <n-button :disabled="disableLauncher" type="info" @click="gameLauncher">快捷启动</n-button>
  </div>
</template>


<style>
.loading-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #1c565b;
  color: white;
  text-align: center;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
