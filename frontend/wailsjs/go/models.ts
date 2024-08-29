export namespace conf {
	
	export class AramConfig {
	    priority_champ_ids: number[];
	
	    static createFrom(source: any = {}) {
	        return new AramConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.priority_champ_ids = source["priority_champ_ids"];
	    }
	}
	export class StatusContent {
	    status_message: string;
	    ranked_league_queue: string;
	    ranked_league_tier: string;
	    ranked_league_division: string;
	
	    static createFrom(source: any = {}) {
	        return new StatusContent(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status_message = source["status_message"];
	        this.ranked_league_queue = source["ranked_league_queue"];
	        this.ranked_league_tier = source["ranked_league_tier"];
	        this.ranked_league_division = source["ranked_league_division"];
	    }
	}
	export class Config {
	    auto_next: boolean;
	    auto_accept: boolean;
	    auto_status: boolean;
	    auto_swap: boolean;
	    status_content: StatusContent;
	    aram_config: AramConfig;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.auto_next = source["auto_next"];
	        this.auto_accept = source["auto_accept"];
	        this.auto_status = source["auto_status"];
	        this.auto_swap = source["auto_swap"];
	        this.status_content = this.convertValues(source["status_content"], StatusContent);
	        this.aram_config = this.convertValues(source["aram_config"], AramConfig);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace model {
	
	export class GameBan {
	    championId: number;
	    pickTurn: number;
	
	    static createFrom(source: any = {}) {
	        return new GameBan(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.championId = source["championId"];
	        this.pickTurn = source["pickTurn"];
	    }
	}
	export class GameTeam {
	    teamId: number;
	    bans: GameBan[];
	    baronKills: number;
	    dominionVictoryScore: number;
	    dragonKills: number;
	    firstBaron: boolean;
	    firstBlood: boolean;
	    firstDargon: boolean;
	    firstInhibitor: boolean;
	    firstTower: boolean;
	    hordeKills: number;
	    inhibitorKills: number;
	    riftHeraldKills: number;
	    towerKills: number;
	    vilemawKills: number;
	    win: string;
	
	    static createFrom(source: any = {}) {
	        return new GameTeam(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.teamId = source["teamId"];
	        this.bans = this.convertValues(source["bans"], GameBan);
	        this.baronKills = source["baronKills"];
	        this.dominionVictoryScore = source["dominionVictoryScore"];
	        this.dragonKills = source["dragonKills"];
	        this.firstBaron = source["firstBaron"];
	        this.firstBlood = source["firstBlood"];
	        this.firstDargon = source["firstDargon"];
	        this.firstInhibitor = source["firstInhibitor"];
	        this.firstTower = source["firstTower"];
	        this.hordeKills = source["hordeKills"];
	        this.inhibitorKills = source["inhibitorKills"];
	        this.riftHeraldKills = source["riftHeraldKills"];
	        this.towerKills = source["towerKills"];
	        this.vilemawKills = source["vilemawKills"];
	        this.win = source["win"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GameTimelineItem {
	    additionalProp1: number;
	    additionalProp2: number;
	    additionalProp3: number;
	
	    static createFrom(source: any = {}) {
	        return new GameTimelineItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.additionalProp1 = source["additionalProp1"];
	        this.additionalProp2 = source["additionalProp2"];
	        this.additionalProp3 = source["additionalProp3"];
	    }
	}
	export class GameTimeline {
	    lane: string;
	    role: string;
	    participantId: number;
	    creepsPerMinDeltas: GameTimelineItem;
	    csDiffPerMinDeltas: GameTimelineItem;
	    damageTakenDiffPerMinDeltas: GameTimelineItem;
	    damageTakenPerMinDeltas: GameTimelineItem;
	    goldPerMinDeltas: GameTimelineItem;
	    xpDiffPerMinDeltas: GameTimelineItem;
	    xpPerMinDeltas: GameTimelineItem;
	
	    static createFrom(source: any = {}) {
	        return new GameTimeline(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.lane = source["lane"];
	        this.role = source["role"];
	        this.participantId = source["participantId"];
	        this.creepsPerMinDeltas = this.convertValues(source["creepsPerMinDeltas"], GameTimelineItem);
	        this.csDiffPerMinDeltas = this.convertValues(source["csDiffPerMinDeltas"], GameTimelineItem);
	        this.damageTakenDiffPerMinDeltas = this.convertValues(source["damageTakenDiffPerMinDeltas"], GameTimelineItem);
	        this.damageTakenPerMinDeltas = this.convertValues(source["damageTakenPerMinDeltas"], GameTimelineItem);
	        this.goldPerMinDeltas = this.convertValues(source["goldPerMinDeltas"], GameTimelineItem);
	        this.xpDiffPerMinDeltas = this.convertValues(source["xpDiffPerMinDeltas"], GameTimelineItem);
	        this.xpPerMinDeltas = this.convertValues(source["xpPerMinDeltas"], GameTimelineItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GameStats {
	    assists: number;
	    causedEarlySurrender: boolean;
	    champLevel: number;
	    combatPlayerScore: number;
	    damageDealtToObjectives: number;
	    damageDealtToTurrets: number;
	    damageSelfMitigated: number;
	    deaths: number;
	    doubleKills: number;
	    earlySurrenderAccomplice: boolean;
	    firstBloodAssist: boolean;
	    firstBloodKill: boolean;
	    firstInhibitorAssist: boolean;
	    firstInhibitorKill: boolean;
	    firstTowerAssist: boolean;
	    firstTowerKill: boolean;
	    gameEndedInEarlySurrender: boolean;
	    gameEndedInSurrender: boolean;
	    goldEarned: number;
	    goldSpent: number;
	    inhibitorKills: number;
	    item0: number;
	    item1: number;
	    item2: number;
	    item3: number;
	    item4: number;
	    item5: number;
	    item6: number;
	    killingSprees: number;
	    kills: number;
	    largestCriticalStrike: number;
	    largestKillingSpree: number;
	    largestMultiKill: number;
	    longestTimeSpentLiving: number;
	    magicDamageDealt: number;
	    magicDamageDealtToChampions: number;
	    magicalDamageTaken: number;
	    neutralMinionsKilled: number;
	    neutralMinionsKilledEnemyJungle: number;
	    neutralMinionsKilledTeamJungle: number;
	    objectivePlayerScore: number;
	    participantId: number;
	    pentaKills: number;
	    perk0: number;
	    perk0Var1: number;
	    perk0Var2: number;
	    perk0Var3: number;
	    perk1: number;
	    perk1Var1: number;
	    perk1Var2: number;
	    perk1Var3: number;
	    perk2: number;
	    perk2Var1: number;
	    perk2Var2: number;
	    perk2Var3: number;
	    perk3: number;
	    perk3Var1: number;
	    perk3Var2: number;
	    perk3Var3: number;
	    perk4: number;
	    perk4Var1: number;
	    perk4Var2: number;
	    perk4Var3: number;
	    perk5: number;
	    perk5Var1: number;
	    perk5Var2: number;
	    perk5Var3: number;
	    perkPrimaryStyle: number;
	    perkSubStyle: number;
	    physicalDamageDealt: number;
	    physicalDamageDealtToChampions: number;
	    physicalDamageTaken: number;
	    playerAugment1: number;
	    playerAugment2: number;
	    playerAugment3: number;
	    playerAugment4: number;
	    playerAugment5: number;
	    playerAugment6: number;
	    playerScore0: number;
	    playerScore1: number;
	    playerScore2: number;
	    playerScore3: number;
	    playerScore4: number;
	    playerScore5: number;
	    playerScore6: number;
	    playerScore7: number;
	    playerScore8: number;
	    playerScore9: number;
	    playerSubteamId: number;
	    quadraKills: number;
	    sightWardsBoughtInGame: number;
	    subteamPlacement: number;
	    teamEarlySurrendered: boolean;
	    timeCCingOthers: number;
	    totalDamageDealt: number;
	    totalDamageDealtToChampions: number;
	    totalDamageTaken: number;
	    totalHeal: number;
	    totalMinionsKilled: number;
	    totalPlayerScore: number;
	    totalScoreRank: number;
	    totalTimeCrowdControlDealt: number;
	    totalUnitsHealed: number;
	    tripleKills: number;
	    trueDamageDealt: number;
	    trueDamageDealtToChampions: number;
	    trueDamageTaken: number;
	    turretKills: number;
	    unrealKills: number;
	    visionScore: number;
	    visionWardsBoughtInGame: number;
	    wardsKilled: number;
	    wardsPlaced: number;
	    win: boolean;
	
	    static createFrom(source: any = {}) {
	        return new GameStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.assists = source["assists"];
	        this.causedEarlySurrender = source["causedEarlySurrender"];
	        this.champLevel = source["champLevel"];
	        this.combatPlayerScore = source["combatPlayerScore"];
	        this.damageDealtToObjectives = source["damageDealtToObjectives"];
	        this.damageDealtToTurrets = source["damageDealtToTurrets"];
	        this.damageSelfMitigated = source["damageSelfMitigated"];
	        this.deaths = source["deaths"];
	        this.doubleKills = source["doubleKills"];
	        this.earlySurrenderAccomplice = source["earlySurrenderAccomplice"];
	        this.firstBloodAssist = source["firstBloodAssist"];
	        this.firstBloodKill = source["firstBloodKill"];
	        this.firstInhibitorAssist = source["firstInhibitorAssist"];
	        this.firstInhibitorKill = source["firstInhibitorKill"];
	        this.firstTowerAssist = source["firstTowerAssist"];
	        this.firstTowerKill = source["firstTowerKill"];
	        this.gameEndedInEarlySurrender = source["gameEndedInEarlySurrender"];
	        this.gameEndedInSurrender = source["gameEndedInSurrender"];
	        this.goldEarned = source["goldEarned"];
	        this.goldSpent = source["goldSpent"];
	        this.inhibitorKills = source["inhibitorKills"];
	        this.item0 = source["item0"];
	        this.item1 = source["item1"];
	        this.item2 = source["item2"];
	        this.item3 = source["item3"];
	        this.item4 = source["item4"];
	        this.item5 = source["item5"];
	        this.item6 = source["item6"];
	        this.killingSprees = source["killingSprees"];
	        this.kills = source["kills"];
	        this.largestCriticalStrike = source["largestCriticalStrike"];
	        this.largestKillingSpree = source["largestKillingSpree"];
	        this.largestMultiKill = source["largestMultiKill"];
	        this.longestTimeSpentLiving = source["longestTimeSpentLiving"];
	        this.magicDamageDealt = source["magicDamageDealt"];
	        this.magicDamageDealtToChampions = source["magicDamageDealtToChampions"];
	        this.magicalDamageTaken = source["magicalDamageTaken"];
	        this.neutralMinionsKilled = source["neutralMinionsKilled"];
	        this.neutralMinionsKilledEnemyJungle = source["neutralMinionsKilledEnemyJungle"];
	        this.neutralMinionsKilledTeamJungle = source["neutralMinionsKilledTeamJungle"];
	        this.objectivePlayerScore = source["objectivePlayerScore"];
	        this.participantId = source["participantId"];
	        this.pentaKills = source["pentaKills"];
	        this.perk0 = source["perk0"];
	        this.perk0Var1 = source["perk0Var1"];
	        this.perk0Var2 = source["perk0Var2"];
	        this.perk0Var3 = source["perk0Var3"];
	        this.perk1 = source["perk1"];
	        this.perk1Var1 = source["perk1Var1"];
	        this.perk1Var2 = source["perk1Var2"];
	        this.perk1Var3 = source["perk1Var3"];
	        this.perk2 = source["perk2"];
	        this.perk2Var1 = source["perk2Var1"];
	        this.perk2Var2 = source["perk2Var2"];
	        this.perk2Var3 = source["perk2Var3"];
	        this.perk3 = source["perk3"];
	        this.perk3Var1 = source["perk3Var1"];
	        this.perk3Var2 = source["perk3Var2"];
	        this.perk3Var3 = source["perk3Var3"];
	        this.perk4 = source["perk4"];
	        this.perk4Var1 = source["perk4Var1"];
	        this.perk4Var2 = source["perk4Var2"];
	        this.perk4Var3 = source["perk4Var3"];
	        this.perk5 = source["perk5"];
	        this.perk5Var1 = source["perk5Var1"];
	        this.perk5Var2 = source["perk5Var2"];
	        this.perk5Var3 = source["perk5Var3"];
	        this.perkPrimaryStyle = source["perkPrimaryStyle"];
	        this.perkSubStyle = source["perkSubStyle"];
	        this.physicalDamageDealt = source["physicalDamageDealt"];
	        this.physicalDamageDealtToChampions = source["physicalDamageDealtToChampions"];
	        this.physicalDamageTaken = source["physicalDamageTaken"];
	        this.playerAugment1 = source["playerAugment1"];
	        this.playerAugment2 = source["playerAugment2"];
	        this.playerAugment3 = source["playerAugment3"];
	        this.playerAugment4 = source["playerAugment4"];
	        this.playerAugment5 = source["playerAugment5"];
	        this.playerAugment6 = source["playerAugment6"];
	        this.playerScore0 = source["playerScore0"];
	        this.playerScore1 = source["playerScore1"];
	        this.playerScore2 = source["playerScore2"];
	        this.playerScore3 = source["playerScore3"];
	        this.playerScore4 = source["playerScore4"];
	        this.playerScore5 = source["playerScore5"];
	        this.playerScore6 = source["playerScore6"];
	        this.playerScore7 = source["playerScore7"];
	        this.playerScore8 = source["playerScore8"];
	        this.playerScore9 = source["playerScore9"];
	        this.playerSubteamId = source["playerSubteamId"];
	        this.quadraKills = source["quadraKills"];
	        this.sightWardsBoughtInGame = source["sightWardsBoughtInGame"];
	        this.subteamPlacement = source["subteamPlacement"];
	        this.teamEarlySurrendered = source["teamEarlySurrendered"];
	        this.timeCCingOthers = source["timeCCingOthers"];
	        this.totalDamageDealt = source["totalDamageDealt"];
	        this.totalDamageDealtToChampions = source["totalDamageDealtToChampions"];
	        this.totalDamageTaken = source["totalDamageTaken"];
	        this.totalHeal = source["totalHeal"];
	        this.totalMinionsKilled = source["totalMinionsKilled"];
	        this.totalPlayerScore = source["totalPlayerScore"];
	        this.totalScoreRank = source["totalScoreRank"];
	        this.totalTimeCrowdControlDealt = source["totalTimeCrowdControlDealt"];
	        this.totalUnitsHealed = source["totalUnitsHealed"];
	        this.tripleKills = source["tripleKills"];
	        this.trueDamageDealt = source["trueDamageDealt"];
	        this.trueDamageDealtToChampions = source["trueDamageDealtToChampions"];
	        this.trueDamageTaken = source["trueDamageTaken"];
	        this.turretKills = source["turretKills"];
	        this.unrealKills = source["unrealKills"];
	        this.visionScore = source["visionScore"];
	        this.visionWardsBoughtInGame = source["visionWardsBoughtInGame"];
	        this.wardsKilled = source["wardsKilled"];
	        this.wardsPlaced = source["wardsPlaced"];
	        this.win = source["win"];
	    }
	}
	export class GameParticipant {
	    teamId: number;
	    championId: number;
	    highestAchievedSeasonTier: string;
	    participantId: number;
	    spell1Id: number;
	    spell2Id: number;
	    stats: GameStats;
	    timeline: GameTimeline;
	
	    static createFrom(source: any = {}) {
	        return new GameParticipant(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.teamId = source["teamId"];
	        this.championId = source["championId"];
	        this.highestAchievedSeasonTier = source["highestAchievedSeasonTier"];
	        this.participantId = source["participantId"];
	        this.spell1Id = source["spell1Id"];
	        this.spell2Id = source["spell2Id"];
	        this.stats = this.convertValues(source["stats"], GameStats);
	        this.timeline = this.convertValues(source["timeline"], GameTimeline);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GameParticipantIdentity {
	    participantId: number;
	    // Go type: struct { AccountId int "json:\"accountId\""; CurrentAccountId int "json:\"currentAccountId\""; CurrentPlatformId string "json:\"currentPlatformId\""; GameName string "json:\"gameName\""; MatchHistoryUri string "json:\"matchHistoryUri\""; PlatformId string "json:\"platformId\""; ProfileIcon int "json:\"profileIcon\""; Puuid string "json:\"puuid\""; SummonerId int64 "json:\"summonerId\""; SummonerName string "json:\"summonerName\""; TagLine string "json:\"tagLine\"" }
	    player: any;
	
	    static createFrom(source: any = {}) {
	        return new GameParticipantIdentity(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.participantId = source["participantId"];
	        this.player = this.convertValues(source["player"], Object);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GameInfo {
	    endOfGameResult: string;
	    gameCreation: number;
	    // Go type: time
	    gameCreationDate: any;
	    gameDuration: number;
	    gameId: number;
	    gameMode: string;
	    gameType: string;
	    gameVersion: string;
	    mapId: number;
	    participantIdentities: GameParticipantIdentity[];
	    participants: GameParticipant[];
	    platformId: string;
	    queueId: number;
	    seasonId: number;
	    teams: GameTeam[];
	
	    static createFrom(source: any = {}) {
	        return new GameInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.endOfGameResult = source["endOfGameResult"];
	        this.gameCreation = source["gameCreation"];
	        this.gameCreationDate = this.convertValues(source["gameCreationDate"], null);
	        this.gameDuration = source["gameDuration"];
	        this.gameId = source["gameId"];
	        this.gameMode = source["gameMode"];
	        this.gameType = source["gameType"];
	        this.gameVersion = source["gameVersion"];
	        this.mapId = source["mapId"];
	        this.participantIdentities = this.convertValues(source["participantIdentities"], GameParticipantIdentity);
	        this.participants = this.convertValues(source["participants"], GameParticipant);
	        this.platformId = source["platformId"];
	        this.queueId = source["queueId"];
	        this.seasonId = source["seasonId"];
	        this.teams = this.convertValues(source["teams"], GameTeam);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	export class GameRecord {
	    gameBeginDate: string;
	    gameCount: number;
	    gameEndDate: string;
	    gameIndexBegin: number;
	    gameIndexEnd: number;
	    games: GameInfo[];
	
	    static createFrom(source: any = {}) {
	        return new GameRecord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.gameBeginDate = source["gameBeginDate"];
	        this.gameCount = source["gameCount"];
	        this.gameEndDate = source["gameEndDate"];
	        this.gameIndexBegin = source["gameIndexBegin"];
	        this.gameIndexEnd = source["gameIndexEnd"];
	        this.games = this.convertValues(source["games"], GameInfo);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class GameSummary {
	    accountId: number;
	    platformId: string;
	    games: GameRecord;
	
	    static createFrom(source: any = {}) {
	        return new GameSummary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.accountId = source["accountId"];
	        this.platformId = source["platformId"];
	        this.games = this.convertValues(source["games"], GameRecord);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	export class RankedEntry {
	    losses: number;
	    wins: number;
	    leaguePoints: number;
	    division: string;
	    highestDivision: string;
	    tier: string;
	    highestTier: string;
	    previousSeasonAchievedDivision: string;
	    previousSeasonHighestDivision: string;
	    previousSeasonEndDivision: string;
	    previousSeasonAchievedTier: string;
	    previousSeasonHighestTier: string;
	    previousSeasonEndTier: string;
	    queueType: string;
	    ratedRating: number;
	    ratedTier: string;
	    isProvisional: boolean;
	    miniSeriesProgress: string;
	    provisionalGameThreshold: number;
	    provisionalGamesRemaining: number;
	    // Go type: struct { DaysUntilDecay int "json:\"daysUntilDecay\""; DemotionWarning int "json:\"demotionWarning\""; DisplayDecayWarning bool "json:\"displayDecayWarning\""; TimeUntilInactivityStatusChanges int "json:\"timeUntilInactivityStatusChanges\"" }
	    warnings: any;
	
	    static createFrom(source: any = {}) {
	        return new RankedEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.losses = source["losses"];
	        this.wins = source["wins"];
	        this.leaguePoints = source["leaguePoints"];
	        this.division = source["division"];
	        this.highestDivision = source["highestDivision"];
	        this.tier = source["tier"];
	        this.highestTier = source["highestTier"];
	        this.previousSeasonAchievedDivision = source["previousSeasonAchievedDivision"];
	        this.previousSeasonHighestDivision = source["previousSeasonHighestDivision"];
	        this.previousSeasonEndDivision = source["previousSeasonEndDivision"];
	        this.previousSeasonAchievedTier = source["previousSeasonAchievedTier"];
	        this.previousSeasonHighestTier = source["previousSeasonHighestTier"];
	        this.previousSeasonEndTier = source["previousSeasonEndTier"];
	        this.queueType = source["queueType"];
	        this.ratedRating = source["ratedRating"];
	        this.ratedTier = source["ratedTier"];
	        this.isProvisional = source["isProvisional"];
	        this.miniSeriesProgress = source["miniSeriesProgress"];
	        this.provisionalGameThreshold = source["provisionalGameThreshold"];
	        this.provisionalGamesRemaining = source["provisionalGamesRemaining"];
	        this.warnings = this.convertValues(source["warnings"], Object);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Ranked {
	    previousSeasonSplitPoints: number;
	    earnedRegaliaRewardIds: string[];
	    highestPreviousSeasonAchievedTier: string;
	    highestPreviousSeasonAchievedDivision: string;
	    highestPreviousSeasonEndDivision: string;
	    highestPreviousSeasonEndTier: string;
	    highestRankedEntry: RankedEntry;
	    highestRankedEntrySR: RankedEntry;
	    queueMap: {[key: string]: RankedEntry};
	    queues: RankedEntry[];
	    rankedRegaliaLevel: number;
	    seasons: {[key: string]: RankedEntry};
	    // Go type: struct { AdditionalProp1 int "json:\"additionalProp1\""; AdditionalProp2 int "json:\"additionalProp2\""; AdditionalProp3 int "json:\"additionalProp3\"" }
	    splitsProgress: any;
	
	    static createFrom(source: any = {}) {
	        return new Ranked(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.previousSeasonSplitPoints = source["previousSeasonSplitPoints"];
	        this.earnedRegaliaRewardIds = source["earnedRegaliaRewardIds"];
	        this.highestPreviousSeasonAchievedTier = source["highestPreviousSeasonAchievedTier"];
	        this.highestPreviousSeasonAchievedDivision = source["highestPreviousSeasonAchievedDivision"];
	        this.highestPreviousSeasonEndDivision = source["highestPreviousSeasonEndDivision"];
	        this.highestPreviousSeasonEndTier = source["highestPreviousSeasonEndTier"];
	        this.highestRankedEntry = this.convertValues(source["highestRankedEntry"], RankedEntry);
	        this.highestRankedEntrySR = this.convertValues(source["highestRankedEntrySR"], RankedEntry);
	        this.queueMap = this.convertValues(source["queueMap"], RankedEntry, true);
	        this.queues = this.convertValues(source["queues"], RankedEntry);
	        this.rankedRegaliaLevel = source["rankedRegaliaLevel"];
	        this.seasons = this.convertValues(source["seasons"], RankedEntry, true);
	        this.splitsProgress = this.convertValues(source["splitsProgress"], Object);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class SummonerRerollPoints {
	    currentPoints: number;
	    maxRolls: number;
	    numberOfRolls: number;
	    pointsCostToRoll: number;
	    pointsToReroll: number;
	
	    static createFrom(source: any = {}) {
	        return new SummonerRerollPoints(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.currentPoints = source["currentPoints"];
	        this.maxRolls = source["maxRolls"];
	        this.numberOfRolls = source["numberOfRolls"];
	        this.pointsCostToRoll = source["pointsCostToRoll"];
	        this.pointsToReroll = source["pointsToReroll"];
	    }
	}
	export class Summoner {
	    accountId: number;
	    displayName: string;
	    gameName: string;
	    internalName: string;
	    nameChangeFlag: boolean;
	    percentCompleteForNextLevel: number;
	    privacy: string;
	    profileIconId: number;
	    puuid: string;
	    rerollPoints: SummonerRerollPoints;
	    summonerId: number;
	    summonerLevel: number;
	    tagLine: string;
	    unnamed: boolean;
	    xpSinceLastLevel: number;
	    xpUntilNextLevel: number;
	
	    static createFrom(source: any = {}) {
	        return new Summoner(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.accountId = source["accountId"];
	        this.displayName = source["displayName"];
	        this.gameName = source["gameName"];
	        this.internalName = source["internalName"];
	        this.nameChangeFlag = source["nameChangeFlag"];
	        this.percentCompleteForNextLevel = source["percentCompleteForNextLevel"];
	        this.privacy = source["privacy"];
	        this.profileIconId = source["profileIconId"];
	        this.puuid = source["puuid"];
	        this.rerollPoints = this.convertValues(source["rerollPoints"], SummonerRerollPoints);
	        this.summonerId = source["summonerId"];
	        this.summonerLevel = source["summonerLevel"];
	        this.tagLine = source["tagLine"];
	        this.unnamed = source["unnamed"];
	        this.xpSinceLastLevel = source["xpSinceLastLevel"];
	        this.xpUntilNextLevel = source["xpUntilNextLevel"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace opgg {
	
	export class Rune {
	    play?: number;
	    win?: number;
	    pick_rate?: number;
	    id?: number;
	    primary_page_id?: number;
	    primary_rune_ids?: number[];
	    secondary_page_id?: number;
	    secondary_rune_ids?: number[];
	    stat_mod_ids?: number[];
	
	    static createFrom(source: any = {}) {
	        return new Rune(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.play = source["play"];
	        this.win = source["win"];
	        this.pick_rate = source["pick_rate"];
	        this.id = source["id"];
	        this.primary_page_id = source["primary_page_id"];
	        this.primary_rune_ids = source["primary_rune_ids"];
	        this.secondary_page_id = source["secondary_page_id"];
	        this.secondary_rune_ids = source["secondary_rune_ids"];
	        this.stat_mod_ids = source["stat_mod_ids"];
	    }
	}

}

