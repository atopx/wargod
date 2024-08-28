<template>
  <n-card :title="summoner.gameName">
  </n-card>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";
import {model} from "../../wailsjs/go/models";
import {CurrentSummoner, GetRankInfo, ListGameSummary} from "../../wailsjs/go/api/Api";
import {useNotification} from "naive-ui";

const summoner = ref(new model.Summoner());
const summonerRanked = ref(new model.Ranked());
const summonerGames = ref(new model.GameSummary())

const notification = useNotification();

const refreshSummoner = () => {
  return CurrentSummoner().then((resp) => {
    summoner.value = resp;
    console.log(summoner.value.puuid);
  }).catch((error) => {
    notification.error({
      content: error
    });
  });
}

const refreshSummonerRanked = async (uuid: string) => {
  GetRankInfo(uuid).then((resp) => {
    summonerRanked.value = resp;
  }).catch((error) => {
    notification.error({
      content: error
    });
  });
}

const refreshSummonerGames = async (uuid: string) => {
  ListGameSummary(uuid, 0, 20).then((resp) => {
    summonerGames.value = resp;
  }).catch((error) => {
    notification.error({
      content: error as string
    });
  });
}

onMounted(() => {
  refreshSummoner().then(() => {
    if (summoner.value.puuid) {
      // refreshSummonerRanked(summoner.value.puuid);
      // refreshSummonerGames(summoner.value.puuid);
    } else {
      notification.error({
        content: "查询召唤师数据失败"
      });
    }
  });
});

</script>