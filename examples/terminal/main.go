package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"strings"
	"wargod/model"
)

const p = `魂锁典狱长
未来守护者
探险家
狂野女猎手
圣枪游侠
虚空之女
赏金猎人
残月之肃
法外狂徒
皮城女警
逆羽
寒冰射手
战争女神
英勇投弹手
影哨
暴走萝莉
荣耀行刑官
惩戒之箭
瘟疫之源
麦林炮手
虚空掠夺者
机械先驱
诡术妖姬
暗夜猎手
卡牌大师
戏命师
永猎双子
天启者
德玛西亚之翼
堕落天使
光辉女郎
解脱者
盲僧
机械公敌
邪恶小法师
沙漠玫瑰
复仇之矛
愁云使者
发条魔灵
异画师
刀锋之影
疾风剑豪
爆破鬼才
沙漠皇帝
复仇焰魂
血港鬼影
蒸汽机器人
迅捷斥候
灵罗娃娃
深渊巨口
冰晶凤凰
远古巫灵
九尾妖狐
德玛西亚皇子
双界灵兔
符文法师
酒桶
熔岩巨兽
纳祖芒荣耀
时光守护者
猩红收割者
虚空之眼
正义天使
无极剑圣
众星之子
牛头酋长`

func main() {
	ids := []int{}
	data := strings.Split(p, "\n")

	mp := make(map[string]int)
	for id, champ := range model.ChampMap {
		mp[champ.Title] = id
	}
	for _, v := range data {
		ids = append(ids, mp[strings.TrimSpace(v)])
	}
	r, _ := json.Marshal(ids)
	fmt.Println(string(r))
}
