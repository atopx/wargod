<template>
  <n-card>

    <n-flex justify="space-between">
      <div class="skill-row">
        <div v-for="(id, index) in skill.ids" :key="index" class="skill-item">
          <span :class="['keycap', id.toLowerCase()]">{{ id }}</span>
          <img src="/src/assets/icons/icon-arrow-right.svg" alt="arrow" v-if="index < skill.ids.length - 1"/>
        </div>
      </div>
      <n-space justify="end">
        <div class="stat win-rate">胜率 {{ (skill.win / skill.play * 100).toFixed(2) }}%</div>
        <div class="stat pick_rate">使用率 {{ (skill.pick_rate * 100).toFixed(2) }}%</div>
      </n-space>
    </n-flex>·
    <div class="skill-order">
      <n-grid :x-gap="8" :cols="24">
        <n-gi v-for="(id, index) in skill.order" :key="index" :class="['keycap', id.toLowerCase()]"
              style="width: 24px;height: 24px;line-height: 24px">
          {{ id }}
        </n-gi>

      </n-grid>
    </div>
  </n-card>

</template>

<script lang="ts" setup>
import {NCard} from 'naive-ui';
import {Skill} from "../../types/game";

const props = defineProps<{
  skill: Skill
}>()
</script>

<style scoped>

.skill-order {
  margin-top: 5px;
  margin-left: 20px;
}

.stat {
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

.skill-row {
  display: flex;
  align-items: center; /* Aligns items vertically in the center */
  margin-left: 20px;
}

.skill-item {
  display: flex;
  align-items: center; /* Ensures each key and arrow are centered in their own container */
}

.keycap {
  display: inline-block;
  width: 30px;
  height: 30px;
  line-height: 30px;
  margin: 2px;
  text-align: center;
  border-radius: 5px;
  color: white;
  font-weight: bold;
  font-family: Arial, sans-serif;
}

/* Define colors for each character */
.q {
  background-color: #1e90ff; /* Light Blue for Q */
}

.w {
  background-color: #2ecc71; /* Green for W */
}

.e {
  background-color: #e67e22; /* Orange for E */
}

.r {
  background-color: #3498db; /* Blue for R */
}


</style>
