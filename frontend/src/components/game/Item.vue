<template>
    <n-table :bordered="false" :single-line="false" striped>
        <tbody>
            <!-- 使用 v-for 遍历 starter_items 和 boots 数组 -->
            <tr v-for="(starter, index) in starter_items" :key="index">
                <td>
                    <!-- 显示初始装备信息 -->
                    <n-flex justify="space-between">
                        <div class="item-icons icon-layer">
                            <div
                                v-for="(item, index) in unionCounter(starter.ids)"
                                :key="index"
                                class="item-container"
                            >
                                <img
                                    :src="`/assets/items/${item.id}.png`"
                                    :alt="`item-${item.id}`"
                                    class="item-img"
                                />
                                <!-- 重复物品叠放, 只有当数量大于1时才显示数量 -->
                                <span v-if="item.count > 1" class="item-count">{{ item.count }}</span>
                            </div>
                        </div>

                        <div class="item-stats">
                            <div class="win-rate">
                                胜率 {{ ((starter.win / starter.play) * 100).toFixed(2) }}%
                            </div>
                            <div class="pick_rate">使用率 {{ (starter.pick_rate * 100).toFixed(2) }}%</div>
                        </div>
                    </n-flex>
                </td>
                <td v-if="boots[index]">
                    <!-- 显示鞋子信息，如果 boots 数组有相应的索引 -->
                    <n-flex justify="space-between">
                        <div v-for="id in boots[index].ids" :key="id">
                            <img :src="`/assets/items/${id}.png`" :alt="`item-${id}`" class="item-img" />
                        </div>
                        <div class="item-stats">
                            <div class="win-rate">
                                胜率 {{ ((boots[index].win / boots[index].play) * 100).toFixed(2) }}%
                            </div>
                            <div class="pick_rate">
                                使用率 {{ (boots[index].pick_rate * 100).toFixed(2) }}%
                            </div>
                        </div>
                    </n-flex>
                </td>
                <td v-else>
                    <div>无数据</div>
                </td>
            </tr>
        </tbody>
    </n-table>

    <n-card :bordered="false">
        <div v-for="(item, index) in core_items" :key="index">
            <n-flex justify="space-between">
                <div class="item-icons core-icons">
                    <div v-for="(id, idx) in item.ids" :key="id">
                        <img :src="`/assets/items/${id}.png`" :alt="`item-${id}`" class="item-img" />
                        <img
                            src="/assets/icons/icon-arrow-right.svg"
                            alt="arrow"
                            v-if="idx < item.ids.length - 1"
                        />
                    </div>
                </div>
                <div class="stat win-rate">胜率 {{ ((item.win / item.play) * 100).toFixed(2) }}%</div>
                <div class="stat pick_rate">使用率 {{ (item.pick_rate * 100).toFixed(2) }}%</div>
            </n-flex>
        </div>
    </n-card>
</template>

<script setup lang="ts">
import { NCard, NTable } from "naive-ui";
import { Item } from "../../types/game";

const props = defineProps<{
    starter_items: Array<Item>;
    boots: Array<Item>;
    core_items: Array<Item>;
}>();

type UnionItems = {
    id: number;
    count: number;
};

function unionCounter(ids: Array<number>): Array<UnionItems> {
    const itemMap: { [key: number]: number } = {};

    // 统计每个物品的数量
    ids.forEach((id) => {
        if (itemMap[id]) {
            itemMap[id] += 1;
        } else {
            itemMap[id] = 1;
        }
    });

    // 将结果转换为 UnionItems 数组
    return Object.keys(itemMap).map((id) => ({
        id: Number(id),
        count: itemMap[Number(id)],
    }));
}
</script>

<style scoped>
.item-icons {
    display: flex;
    align-items: center;
    margin-left: 10px;
}

.core-icons {
    width: 270px;
    margin-left: 0;
}

.item-img {
    height: 32px;
    width: 32px;
    margin-right: 4px;
}

.item-stats {
    text-align: center;
    margin-right: 70px;
    width: 90px;
}

.stat {
    margin-right: 60px;
    text-align: center;
    width: 90px;
}

.win-rate {
    color: darkblue;
    font-size: 13px;
}

.pick_rate {
    color: darkcyan;
    font-size: 13px;
}

.icon-layer {
    width: 214px;
}

.item-container {
    position: relative;
    display: inline-block;
}

.item-count {
    position: absolute;
    bottom: 7px; /* 调整到你想要的位置 */
    right: 7px; /* 调整到你想要的位置 */
    background-color: rgba(0, 0, 0, 0.9); /* 半透明背景 */
    color: white;
    font-size: 7px; /* 调整字体大小 */
    padding: 2px 5px;
    border-radius: 3px;
    font-weight: 900;
}
</style>
