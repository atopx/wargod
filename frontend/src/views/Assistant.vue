<template>
  <n-card>
    <n-layout>
      <n-card bordered>
        <div class="champion-container">
          <div class="champion-avatar-container">
            <img
                :src="`/src/assets/champions/${data.champion.id}.png`"
                :alt="data.champion.name"
                class="champion-avatar"
            />
          </div>

          <div class="champion-info">
            <h2 class="champion-name">{{ data.champion.name }}</h2>
            <p class="champion-title">{{ data.champion.title }}</p>
          </div>

          <div class="champion-stats">
            <p class="stat-win-rate">胜率: {{ (data.champion.win_rate * 100).toFixed(2) }}%</p>
            <p class="stat-pick-rate">选取率: {{ (data.champion.pick_rate * 100).toFixed(2) }}%</p>
            <p class="stat-kda">KDA: {{ data.champion.kda.toFixed(2) }}</p>
          </div>
        </div>
      </n-card>
      <n-space vertical>
        <Spell :spells="data.spells"/>
        <Rune :runes="data.runes"/>
        <Skill :skill="data.skill"/>
        <Item :boots="data.boots" :starter_items="data.starter_items" :core_items="data.core_items"/>
      </n-space>
    </n-layout>
  </n-card>

</template>


<script setup lang="ts">
import {ref} from 'vue'
import Spell from '@/components/game/Spell.vue'
import Rune from "@/components/game/Rune.vue";
import Item from "@/components/game/Item.vue";
import Skill from "@/components/game/Skill.vue";
import {EventsOn} from "../../wailsjs/runtime";
import {ChampInfo} from "../types/game";

const data = ref<ChampInfo>({
  champion: {
    id: 96,
    title: "深渊巨口",
    name: "克格莫",
    roles: ["射手", "法师"],
    win_rate: 0.5268,
    pick_rate: 0.0644092,
    kda: 3.5004,
  },
  spells: [
    {ids: [4, 6], "win": 88454, "play": 169252, "pick_rate": 0.5871},
    {ids: [4, 13], "win": 17900, "play": 32339, "pick_rate": 0.1122},
    {ids: [4, 32], "win": 12268, "play": 24084, "pick_rate": 0.0835},
  ],
  runes: [
    {
      "primary_page_id": 8000,
      "primary_rune_ids": [
        8005,
        8009,
        9104,
        8014
      ],
      "secondary_page_id": 8100,
      "secondary_rune_ids": [
        8135,
        8139
      ],
      "stat_mod_ids": [
        5005,
        5008,
        5001
      ],
      "play": 100567,
      "win": 50561,
      "pick_rate": 0.3447
    },
    {
      "primary_page_id": 8100,
      "primary_rune_ids": [
        8128,
        8139,
        8138,
        8135
      ],
      "secondary_page_id": 8000,
      "secondary_rune_ids": [
        8009,
        8014
      ],
      "stat_mod_ids": [
        5008,
        5008,
        5001
      ],
      "play": 69411,
      "win": 37745,
      "pick_rate": 0.2379
    },
    {
      "primary_page_id": 8000,
      "primary_rune_ids": [
        8005,
        9111,
        9103,
        8299
      ],
      "secondary_page_id": 8400,
      "secondary_rune_ids": [
        8429,
        8451
      ],
      "stat_mod_ids": [
        5005,
        5008,
        5001
      ],
      "play": 41437,
      "win": 22547,
      "pick_rate": 0.142
    },
  ],
  skill: {
    "ids": [
      "W",
      "Q",
      "E"
    ],
    "play": 142389,
    "win": 75553,
    "pick_rate": 0.5468,
    "order": [
      "Q",
      "W",
      "E",
      "W",
      "W",
      "R",
      "W",
      "Q",
      "W",
      "Q",
      "R",
      "Q",
      "Q",
      "E",
      "E"
    ]
  },
  core_items: [
    {
      "ids": [
        3153,
        3124,
        3085
      ],
      "play": 37747,
      "win": 20064,
      "pick_rate": 0.1407
    },
    {
      "ids": [
        3070,
        3118,
        3040,
        6653
      ],
      "play": 16137,
      "win": 8894,
      "pick_rate": 0.0602
    },
    {
      "ids": [
        3153,
        3124,
        3091
      ],
      "play": 14124,
      "win": 7316,
      "pick_rate": 0.0526
    },
    {
      "ids": [
        3153,
        3085,
        3124
      ],
      "play": 12978,
      "win": 7015,
      "pick_rate": 0.0484
    },
    {
      "ids": [
        3153,
        3124,
        3302
      ],
      "play": 8615,
      "win": 4411,
      "pick_rate": 0.0321
    },
  ],
  boots: [
    {
      "ids": [
        3006
      ],
      "play": 168607,
      "win": 86619,
      "pick_rate": 0.6268
    },
    {
      "ids": [
        3020
      ],
      "play": 84475,
      "win": 47157,
      "pick_rate": 0.314
    },
    {
      "ids": [
        3111
      ],
      "play": 8033,
      "win": 3972,
      "pick_rate": 0.0299
    },
  ],
  starter_items: [
    {
      "ids": [
        1042,
        3006
      ],
      "play": 106674,
      "win": 55006,
      "pick_rate": 0.3638
    },
    {
      "ids": [
        1027,
        1052,
        2022,
        3070
      ],
      "play": 33072,
      "win": 18235,
      "pick_rate": 0.1128
    },
    {
      "ids": [
        2031,
        3802
      ],
      "play": 24658,
      "win": 13404,
      "pick_rate": 0.0841
    },
  ],
})

EventsOn("ChampSelect", function (info: ChampInfo) {
  data.value = info;
})

</script>

<style scoped>
.champion-container {
  display: flex;
  align-items: flex-start; /* Align items to the top */
  justify-content: flex-start; /* Align items to the start */
  padding: 16px;
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

.champion-name {
  margin: 0;
  font-size: 18px;
  font-weight: bold;
}

.champion-title {
  margin: 0;
  color: #666;
  font-size: 14px;
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

.stat-pick-rate {
  color: darkcyan;
}

.stat-kda {
  color: darkorange;
}
</style>