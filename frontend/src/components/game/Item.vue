<template>
  <n-card content-style="padding: 0;">
    <n-table :bordered="false" :single-line="false" striped>

      <tbody>
      <!-- 使用 v-for 遍历 starter_items 和 boots 数组 -->
      <tr v-for="(starter, index) in starter_items" :key="index">
        <td>
          <!-- 显示初始装备信息 -->
          <n-flex justify="space-between">
            <div class="item-icons">
              <div v-for="id in starter.ids" :key="id">
                <img :src="`/src/assets/items/${id}.png`" :alt="`item-${id}`" class="item-img"/>
              </div>
            </div>
            <div class="item-stats">
              <div class="win-rate">胜率 {{ (starter.win / starter.play * 100).toFixed(2) }}%</div>
              <div class="pick_rate">使用率 {{ (starter.pick_rate * 100).toFixed(2) }}%</div>
            </div>
          </n-flex>
        </td>
        <td v-if="boots[index]">
          <!-- 显示鞋子信息，如果 boots 数组有相应的索引 -->
          <n-flex justify="space-between">
            <div v-for="id in boots[index].ids" :key="id">
              <img :src="`/src/assets/items/${id}.png`" :alt="`item-${id}`" class="item-img"/>
            </div>
            <div class="item-stats">
              <div class="win-rate">胜率 {{ (boots[index].win / boots[index].play * 100).toFixed(2) }}%</div>
              <div class="pick_rate">使用率 {{ (boots[index].pick_rate * 100).toFixed(2) }}%</div>
            </div>
          </n-flex>
        </td>
        <td v-else>
          <!-- 如果 boots 数组没有相应的索引，显示空单元格 -->
          <div>无数据</div>
        </td>
      </tr>
      </tbody>
    </n-table>

    <n-card>
      <div v-for="(item, index) in core_items" :key="index">
        <n-flex justify="space-between">
          <div class="item-icons core-icons">
            <div v-for="(id, idx) in item.ids" :key="id">
              <img :src="`/src/assets/items/${id}.png`" :alt="`item-${id}`" class="item-img"/>
              <img src="/src/assets/icons/icon-arrow-right.svg" alt="arrow" v-if="idx < item.ids.length - 1">
            </div>
          </div>
          <n-space justify="end">
            <div class="stat win-rate">胜率 {{ (item.win / item.play * 100).toFixed(2) }}%</div>
            <div class="stat pick_rate">使用率 {{ (item.pick_rate * 100).toFixed(2) }}%</div>
          </n-space>
        </n-flex>
      </div>
    </n-card>

  </n-card>
</template>

<script setup lang="ts">

import {NCard, NTable} from "naive-ui";
import {Item} from "../../types/game";

const props = defineProps<{
  starter_items: Array<Item>;
  boots: Array<Item>;
  core_items: Array<Item>;
}>()
</script>

<style scoped>

.item-icons {
  display: flex;
  align-items: center;
  margin-left: 32px;
}

.core-icons {
  margin-left: 20px;
}

.item-img {
  height: 32px;
  width: 32px;
  margin-right: 6px;
}

.item-stats, .stat {
  text-align: center;
  margin-right: 100px;
}


.win-rate {
  color: darkblue;
  font-size: 13px;
}

.pick_rate {
  color: darkcyan;
  font-size: 13px;
}

</style>

