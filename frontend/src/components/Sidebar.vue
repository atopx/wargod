<!-- src/components/Sidebar.vue -->
<template>
  <div class="sidebar">
    <div class="menu top">
      <div @click="navigateTo('/career')" class="menu-item">
        <img src="@/assets/icons/Trophy.svg" alt="Career"/>
        <span>生涯</span>
      </div>
      <div @click="navigateTo('/search')" class="menu-item">
        <img src="@/assets/icons/Search.svg" alt="Search"/>
        <span>搜索</span>
      </div>
      <div @click="navigateTo('/game')" class="menu-item">
        <img src="@/assets/icons/Game.svg" alt="Game"/>
        <span>对局</span>
      </div>
    </div>
    <div class="menu bottom">
      <div @click="navigateTo('/opgg')" class="menu-item">
        <img src="@/assets/icons/opgg.svg" alt="OPGG"/>
        <span>OPGG</span>
      </div>
      <div @click="navigateTo('/setting')" class="menu-item">
        <img src="@/assets/icons/Setting.svg" alt="Settings"/>
        <span>设置</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {useRouter} from 'vue-router';

const router = useRouter();

function navigateTo(path: string) {
  router.push(path);
}
</script>

<style scoped>
.sidebar {
  width: 200px;
  background-color: #f8f9fa;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 20px;
}

.menu {
  display: flex;
  flex-direction: column;
}

.menu-item {
  display: flex;
  align-items: center;
  cursor: pointer;
  margin-bottom: 15px;
}

.menu-item img {
  width: 24px;
  margin-right: 10px;
}
</style>
