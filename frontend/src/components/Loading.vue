<script lang="ts" setup>
import {NSpin} from 'naive-ui';
</script>

<template>
  <div class="loading-container">
    <!--    <div class="loading-circle"></div>-->
    <n-spin size="large"/>
    <p>正在等待游戏启动...</p>
  </div>
</template>


<style>
.loading-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: calc(100vh - 20px);
  background-color: #1e1e2e;
  color: white;
  text-align: center;
}

.loading-circle {
  width: 50px;
  height: 50px;
  border: 5px solid #2839b0;
  border-top: 5px solid #ffffff;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 20px;
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
