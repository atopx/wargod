<template>
  <n-layout has-sider style="height: 100vh">
    <n-layout-sider
        class="main-sider"
        bordered
        :collapsed-width="64"
        :width="100"
    >
      <n-menu
          :indent="16"
          :collapsed-width="64"
          :collapsed-icon-size="22"
          :options="topMenuOptions"
          @update:value="handleMenuChange"
      />
      <n-menu
          class="bottom-menu"
          :indent="16"
          :collapsed-width="64"
          :collapsed-icon-size="22"
          :options="bottomMenuOptions"
          @update:value="handleMenuChange"
      />
    </n-layout-sider>
    <n-layout>
      <Career v-if="currentMenu === 'career'"></Career>
      <Record v-if="currentMenu === 'record'"></Record>
      <Assistant v-if="currentMenu === 'assistant'"></Assistant>
      <!--        <Personal v-if="currentMenu === 'personal'"></Personal>-->
      <Setting v-if="currentMenu === 'setting'"></Setting>
    </n-layout>
  </n-layout>
</template>


<script setup lang="ts">
import {h, ref} from 'vue';
import {NLayout, NLayoutSider, NMenu} from 'naive-ui';
import Career from "./views/Career.vue";
import Record from "./views/Record.vue";
import Assistant from "./views/Assistant.vue";
import Setting from "./views/Setting.vue";
import CareerIcon from '/src/assets/icons/Career.svg';
import RecordIcon from '/src/assets/icons/Record.svg';
import AssistantIcon from '/src/assets/icons/Game.svg';
// import PersonalIcon from '/src/assets/icons/User.svg';
import SettingIcon from '/src/assets/icons/Setting.svg';

// 响应式状态
const currentMenu = ref('assistant');

// 菜单选项
const topMenuOptions = [
  {label: '生涯', key: 'career', icon: () => h('img', {src: CareerIcon, style: 'width: 20px; height: 20px;'})},
  {label: '战绩', key: 'record', icon: () => h('img', {src: RecordIcon, style: 'width: 20px; height: 20px;'})},
  {label: '助手', key: 'assistant', icon: () => h('img', {src: AssistantIcon, style: 'width: 20px; height: 20px;'})},
];

const bottomMenuOptions = [
  // {label: '玩家', key: 'personal', icon: () => h('img', {src: PersonalIcon, style: 'width: 20px; height: 20px;'})},
  {label: '设置', key: 'setting', icon: () => h('img', {src: SettingIcon, style: 'width: 20px; height: 20px;'})},
];

// 处理菜单变化
const handleMenuChange = (key: string) => {
  currentMenu.value = key;
};
</script>

<style scoped>

.main-sider {
  color: darkblue;
  background-color: #a4b3ef;
  transition: background-color 0.3s;
}

.bottom-menu {
  margin-top: calc(100vh - 210px);
}

</style>
