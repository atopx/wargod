<template>
  <n-grid x-gap="12" :cols="1">
    <n-gi v-for="(rune, idx) in runes" :key="idx">
      <div class="rune-order-item">
        <div class="primary">
          <img :src="`/assets/runes/${rune.primary_rune_ids[0]}.png`" alt="0" class="rune-img"/>
          <img :src="`/assets/runes/${rune.primary_rune_ids[1]}.png`" alt="1" class="rune-img"/>
          <img :src="`/assets/runes/${rune.primary_rune_ids[2]}.png`" alt="2" class="rune-img"/>
          <img :src="`/assets/runes/${rune.primary_rune_ids[3]}.png`" alt="3" class="rune-img"/>
        </div>
        <div class="secondary">
          <img :src="`/assets/runes/${rune.secondary_rune_ids[0]}.png`" alt="0" class="rune-img"/>
          <img :src="`/assets/runes/${rune.secondary_rune_ids[1]}.png`" alt="1" class="rune-img"/>
        </div>
        <div class="stat-mod">
          <img :src="`/assets/runes/${rune.stat_mod_ids[0]}.png`" alt="0" class="rune-img"/>
          <img :src="`/assets/runes/${rune.stat_mod_ids[1]}.png`" alt="1" class="rune-img"/>
          <img :src="`/assets/runes/${rune.stat_mod_ids[2]}.png`" alt="2" class="rune-img"/>
        </div>
        <div class="stat win-rate">胜率 {{ (rune.win / rune.play * 100).toFixed(2) }}%</div>
        <div class="stat pick_rate">使用率 {{ (rune.pick_rate * 100).toFixed(2) }}%</div>
        <n-button type="info" @click="setRune(idx)">应用</n-button>
      </div>
    </n-gi>
  </n-grid>
</template>

<script lang="ts" setup>
import {Rune} from "../../types/game";
import {SetRune} from "../../../wailsjs/go/api/Api";

const props = defineProps<{
  runes: Array<Rune>
}>()

function setRune(idx: number) {
  const rune = props.runes[idx];
  SetRune(rune);
}


</script>

<style scoped>

.rune-order-item {
  display: flex;
  align-items: center;
}

.rune-order-item {
  margin-top: 6px;
}

.primary, .secondary, .stat-mod {
  display: flex;
  align-items: center;
  margin-left: 20px;
  margin-right: 10px;
}

.stat-mod {
  margin-right: 10px;
}

.rune-img {
  height: 24px;
  width: 24px;
  margin-right: 6px;
}

.stat {
  margin-right: 42px;
  width: 90px;
  font-size: 13px;
  text-align: center;
}

.win-rate {
  color: darkblue;

}

.pick_rate {
  color: darkcyan;
}
</style>
