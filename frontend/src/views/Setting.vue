<template>
  <n-card title="默认设置">
    <n-form>
      <n-space vertical size="large">
        <n-form-item label="自动应战">
          <n-switch
              size="large"
              v-model:value="formRef.auto_accept"
              @update:value="updateConfig('auto_accept', $event)"
          ></n-switch>
        </n-form-item>
        <n-form-item label="段位状态">
          <n-switch
              size="large"
              v-model:value="formRef.auto_status"
              @update:value="updateConfig('auto_status', $event)"
          ></n-switch>
        </n-form-item>
        <template v-if="formRef.auto_status">
          <n-form-item label="状态消息">
            <n-input
                v-model:value="formRef.status_content.status_message"
            ></n-input>
          </n-form-item>
          <n-form-item label="队列">
            <n-select
                v-model:value="formRef.status_content.ranked_league_queue"
                :options="queueOpts"
            ></n-select>
          </n-form-item>
          <n-form-item label="段位">
            <n-select
                v-model:value="formRef.status_content.ranked_league_tier"
                :options="tierOpts"
            ></n-select>
          </n-form-item>
          <n-form-item label="段位级别">
            <n-select
                v-model:value="formRef.status_content.ranked_league_division"
                :options="filteredDivOpts"
                :disabled="isDivisionDisabled"
            ></n-select>
          </n-form-item>
        </template>
        <!-- 手动保存按钮 -->
        <n-form-item>
          <n-button type="primary" @click="saveConfig">手动保存</n-button>
        </n-form-item>
      </n-space>
    </n-form>
  </n-card>
</template>

<script setup lang="ts">
import {computed, onMounted, ref} from 'vue';
import {conf} from "../../wailsjs/go/models";
import {GetConfig, SetConfig} from "../../wailsjs/go/api/Api";

const queueOpts = [
  {label: '单双排', value: 'RANKED_SOLO_5x5'},
  {label: '灵活组排', value: 'RANKED_FLEX_SR'},
  {label: '云顶之弈', value: 'RANKED_TFT'},
];

const tierOpts = [
  {label: '无', value: 'NONE'},
  {label: '黑铁', value: 'IRON'},
  {label: '青铜', value: 'BRONZE'},
  {label: '白银', value: 'SILVER'},
  {label: '黄金', value: 'GOLD'},
  {label: '铂金', value: 'PLATINUM'},
  {label: '翡翠', value: 'EMERALD'},
  {label: '钻石', value: 'DIAMOND'},
  {label: '超凡大师', value: 'MASTER'},
  {label: '傲世宗师', value: 'GRANDMASTER'},
  {label: '最强王者', value: 'CHALLENGER'},
];

const divOpts = [
  {label: 'I', value: 'I'},
  {label: 'II', value: 'II'},
  {label: 'III', value: 'III'},
  {label: 'IV', value: 'IV'},
  {label: 'V', value: 'V'},
];


const formRef = ref(new conf.Config());

// 计算属性，用于确定 divOpts 是否应该被禁用
const isDivisionDisabled = computed(() => {
  return ['NONE', 'MASTER', 'GRANDMASTER', 'CHALLENGER'].includes(formRef.value.status_content.ranked_league_tier);
});

// 计算属性，用于根据段位筛选可用的段位级别选项
const filteredDivOpts = computed(() => {
  return isDivisionDisabled.value ? [] : divOpts;
});

const getConfig = async () => {
  try {
    formRef.value = await GetConfig();
  } catch (error) {
    console.error('Failed to get configuration:', error);
  }
};

const updateConfig = (key: keyof conf.Config, value: any) => {
  formRef.value[key] = value;
};

const saveConfig = async () => {
  try {
    await SetConfig(formRef.value);
  } catch (error) {
    console.error('Failed to save configuration:', error);
  }
  await getConfig();
};

onMounted(() => {
  getConfig();
});
</script>
