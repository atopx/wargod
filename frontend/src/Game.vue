<template>
  <n-layout has-sider style="height: 100vh">
    <n-layout-sider
        class="main-sider"
        bordered
        :collapsed-width="64"
        :width="100"
    >
      <n-menu
          :collapsed-width="64"
          :collapsed-icon-size="22"
          :options="topMenuOptions"
          @update:value="handleMenuChange"
      />
      <n-menu
          class="bottom-menu"
          :collapsed-width="64"
          :collapsed-icon-size="22"
          :options="bottomMenuOptions"
          @update:value="handleMenuChange"
      />
    </n-layout-sider>
    <n-layout>
      <n-card style="height: 100vh" :bordered="false" content-style="padding: 0;">
        <Career v-if="currentMenu === 'career'"></Career>
        <Record v-if="currentMenu === 'record'"></Record>
        <Assistant v-if="currentMenu === 'assistant'"></Assistant>
        <Personal v-if="currentMenu === 'personal'"></Personal>
        <Setting v-if="currentMenu === 'settings'"></Setting>
      </n-card>
    </n-layout>
  </n-layout>
</template>


<script setup lang="ts">
import {ref} from 'vue';
import {NCard, NLayout, NLayoutSider, NMenu} from 'naive-ui';
import Career from "./views/Career.vue";
import Record from "./views/Record.vue";
import Assistant from "./views/Assistant.vue";
import Personal from "./views/Personal.vue";
import Setting from "./views/Setting.vue";

// 响应式状态
const currentMenu = ref('assistant');

// 菜单选项
const topMenuOptions = [
  {label: '生涯', key: 'career'},
  {label: '战绩', key: 'record'},
  {label: '助手', key: 'assistant'},
];

const bottomMenuOptions = [
  {label: '玩家', key: 'personal'},
  {label: '设置', key: 'settings'},
];

// 处理菜单变化
const handleMenuChange = (key: string) => {
  currentMenu.value = key;
};
</script>

<style scoped>

.main-sider {
  background-color: #a4b3ef;
  transition: background-color 0.3s;
}

.bottom-menu {
  margin-top: calc(100vh - 260px);
}

</style>
