# 常用API整理

## Get Me

GET /lol-chat/v1/me

## 设置段位

```python
@retry()
async def setTierShowed(self, queue, tier, division):
    data = {
        "lol": {
            "rankedLeagueQueue": queue,
            "rankedLeagueTier": tier,
            "rankedLeagueDivision": division,
        }
    }
    res = await self.__put("/lol-chat/v1/me", data=data)
    return await res.json()
```

## 自动应答对局

POST /lol-matchmaking/v1/ready-check/accept

```python
@retry()
async def acceptMatchMaking(self):
    res = await self.__post("/lol-matchmaking/v1/ready-check/accept")
    return res
```

## 获取当前选择英雄

GET  /lol-champ-select/v1/current-champion

```python
@retry()
async def getCurrentChampion(self):
    res = await self.__get("/lol-champ-select/v1/current-champion")
    return await res.json()
```

## 摇骰子
POST /lol-champ-select/v1/session/my-selection/reroll

```python
async def reroll(self):
    res = await self.__post("/lol-champ-select/v1/session/my-selection/reroll")
    return await res.json()
```

## 选择英雄

PATCH /lol-champ-select/v1/session/actions/{actionsId}

```python
@retry()
async def selectChampion(self, actionsId, championId, completed=None):
    data = {
        "championId": championId,
        "type": "pick",
    }

    if completed:
        data["completed"] = True

    res = await self.__patch(f"/lol-champ-select/v1/session/actions/{actionsId}", data=data)

    return await res.read()
```

## 设置符文页

```python
# 获取当前符文页
@retry()
async def getCurrentRunePage(self):
    res = await self.__get("/lol-perks/v1/currentpage")
    return await res.json()

# 删除符文页
@retry()
async def deleteCurrentRunePage(self):
    page = await self.getCurrentRunePage()

    res = None
    if page.get("isDeletable"):
        id = page["id"]

        res = await self.__delete(f"/lol-perks/v1/pages/{id}")
        res = await res.json()

# 创建符文页
@retry()
async def createRunePage(self, name, primaryId, secondaryId, perks):
    body = {
        "name": name,
        "primaryStyleId": primaryId,
        "subStyleId": secondaryId,
        "selectedPerkIds": perks,
        "current": True,
    }

    res = await self.__post("/lol-perks/v1/pages", data=body)
    return await res.json()

@asyncSlot(bool)
async def __onSetRunePageButtonClicked(self, _):
    data = self.data[self.selectedIndex]
    name = "Seraphine" + self.tr(": ") + self.summary["name"]

    await connector.deleteCurrentRunePage()
    await connector.createRunePage(name, data["primaryId"], data["secondaryId"], data["perks"])
```


## 选皮肤、召唤师技能

```python
@retry()
async def selectConfig(self, skinId, spell1Id=None, spell2Id=None, wardSkinId=None):
    data = {"selectedSkinId": skinId}

    # 4-点燃 12-闪现 14-传送 **推测未验证**
    if spell1Id:
        data["spell1Id"] = spell1Id
    if spell2Id:
        data["spell2Id"] = spell2Id
    if wardSkinId:
        # 不知道是什么，默认-1
        data["wardSkinId"] = wardSkinId

    res = await self.__patch("/lol-champ-select/v1/session/my-selection", data)
    return await res.json()
```


## 获取当前阵营

```python
@retry()
async def getMapSide(self):
    res = await self.__get("/lol-champ-select/v1/pin-drop-notification")
    res = await res.json()

    return res.get("mapSide", "")
```


## 查询游戏状态

```python
async def getGameStatus(self):
    res = await self.__get("/lol-gameflow/v1/gameflow-phase")
    res = await res.text()

    return res[1:-1]
```


## 事件

### 进入房间
Event: LcuEvent { subscription_type: JsonApiEvent("lol-gameflow_v1_gameflow-phase"), data: String("Lobby"), event_type: "Update" }

### 开始匹配
Event: LcuEvent { subscription_type: JsonApiEvent("lol-gameflow_v1_gameflow-phase"), data: String("Matchmaking"), event_type: "Update" }

## 匹配到对局
Event: LcuEvent { subscription_type: JsonApiEvent("lol-gameflow_v1_gameflow-phase"), data: String("ReadyCheck"), event_type: "Update" }

### 解散房间
Event: LcuEvent { subscription_type: JsonApiEvent("lol-gameflow_v1_gameflow-phase"), data: String("None"), event_type: "Update" }

### 英雄选择: ChampSelect
### 游戏开始: GameStart
### 载入游戏: InProgress
### 等待重新连接: Reconnect
### 等待游戏结果: WaitingForStatus
### 准备退出游戏: PreEndOfGame
### 游戏结束：EndOfGame
