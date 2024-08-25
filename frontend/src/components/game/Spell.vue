<template>
<!--  <n-card title="召唤师技能推荐">-->
  <n-card>
    <n-grid x-gap="12" :cols="3">
      <n-gi v-for="(spell, idx) in spells" :key="idx" class="spell-order-item">
        <div class="spell-icons">
          <img :src="`/src/assets/spells/${spell.ids[0]}.png`" alt="spell-left" class="spell-img"/>
          <img :src="`/src/assets/spells/${spell.ids[1]}.png`" alt="spell-right" class="spell-img"/>
        </div>
        <div class="spell-stats">
          <div class="stat win-rate">胜率 {{ (spell.win / spell.play * 100).toFixed(2) }}%</div>
          <div class="stat pick_rate">使用率 {{ (spell.pick_rate * 100).toFixed(2) }}%</div>
        </div>
      </n-gi>
    </n-grid>
  </n-card>
</template>

<script lang="ts" setup>
import {NCard} from 'naive-ui';
import {Item} from "../../types/game";

const props = defineProps<{
  spells: Array<Item>
}>()
</script>

<style scoped>

.spell-order-item {
  display: flex;
  align-items: center;
}

.spell-icons {
  display: flex;
  align-items: center;
  margin-left: 32px;
  margin-right: 32px;
}

.spell-img {
  height: 32px;
  width: 32px;
  margin-right: 6px;
}

.spell-stats {
  text-align: center;
}

.stat {
  margin-bottom: 4px;
}

.win-rate {
  color: darkblue;
  font-size: 12px;
}

.pick_rate {
  color: darkcyan;
  font-size: 12px;
}
</style>
