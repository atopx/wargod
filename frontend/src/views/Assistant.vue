<template>
  <n-layout v-if="hasData">
    <n-card :bordered="false" content-style="padding:0px">
      <div class="champion-container">
        <div class="champion-avatar-container">
          <img
              :src="`/assets/champions/${data!.champion.id}.png`"
              :alt="data!.champion.name"
              class="champion-avatar"
          />
        </div>
        <div class="champion-info">
          <h2 class="champion-title">{{ data!.champion.title }}</h2>
          <p class="champion-name">{{ data!.champion.name }}</p>

          <n-grid x-gap="4" :cols="4">
            <n-gi v-for="(role, idx) in data!.champion.roles" :key="idx">
              <n-tag :bordered="false" type="info" size="small">{{ role }}</n-tag>
            </n-gi>
          </n-grid>
        </div>
        <div class="champion-stats">
          <p class="champ-right stat-win-rate">胜率: {{ (data!.champion.win_rate * 100).toFixed(2) }}%</p>
          <p class="champ-right stat-win-rate">胜率排名: {{ data!.champion.rank }}</p>
          <p class="champ-right stat-kda">场均KDA: {{ data!.champion.kda.toFixed(2) }}</p>
        </div>
      </div>
      <n-divider/>
      <Spell :spells="data!.spells"/>
      <n-divider/>
      <Rune :runes="data!.runes"/>
      <n-divider/>
      <Skill :skill="data!.skill"/>
      <n-divider/>
      <Item :boots="data!.boots" :starter_items="data!.starter_items" :core_items="data!.core_items"/>
    </n-card>
  </n-layout>
  <n-card v-else>
    <n-empty description="暂无数据"></n-empty>
  </n-card>
</template>


<script setup lang="ts">
import {onMounted, ref} from 'vue'
import Spell from '@/components/game/Spell.vue'
import Rune from "@/components/game/Rune.vue";
import Item from "@/components/game/Item.vue";
import Skill from "@/components/game/Skill.vue";
import {EventsOn} from "../../wailsjs/runtime";
import {ChampInfo} from "../types/game";
import {NCard, NEmpty} from "naive-ui";
import Storage from "../utils/storage";

const hasData = ref(false);
const data = ref<ChampInfo | null>(null);


EventsOn("ChampSelect", function (info: ChampInfo) {
  data.value = info;
  hasData.value = true;
  Storage.set('champInfo', info);
})

onMounted(() => {
  const savedData = Storage.get<ChampInfo>('champInfo');
  if (savedData) {
    data.value = savedData;
    hasData.value = true;
  }
})
</script>

<style scoped>
.champion-container {
  display: flex;
  align-items: flex-start; /* Align items to the top */
  justify-content: flex-start; /* Align items to the start */
  padding: 16px;
}

.champ-right {
  width: 100px;
  text-align: center;
}

.champion-avatar-container {
  flex-shrink: 0; /* Prevents avatar from shrinking */
  margin-right: 16px; /* Space between avatar and info */
}

.champion-avatar {
  width: 80px;
  height: 80px;
  border-radius: 10px;
}

.champion-info {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  text-align: left;
}

.champion-title {
  margin: 0;
  font-size: 18px;
  font-weight: bold;
}

.champion-name {
  margin: 0;
  color: #666;
  font-size: 13px;
}


.champion-stats {
  margin-left: auto; /* Push stats to the right */
  margin-right: auto; /* Right margin */
  text-align: right;
}

.champion-stats p {
  margin: 4px 0;
}

.stat-win-rate {
  color: darkblue;
}

.stat-kda {
  color: darkcyan;
}
</style>