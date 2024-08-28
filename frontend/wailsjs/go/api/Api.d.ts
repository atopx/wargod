// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {model} from '../models';
import {conf} from '../models';
import {opgg} from '../models';
import {types} from '../models';

export function CurrentSummoner():Promise<model.Summoner>;

export function GameLauncher():Promise<void>;

export function GetConfig():Promise<conf.Config>;

export function GetGameInfo(arg1:number):Promise<model.GameInfo>;

export function GetRankInfo(arg1:string):Promise<model.Ranked>;

export function GetState():Promise<string>;

export function GetSummoner(arg1:string):Promise<model.Summoner>;

export function ListGameSummary(arg1:string,arg2:number,arg3:number):Promise<model.GameSummary>;

export function SearchSummoner(arg1:string):Promise<model.Summoner>;

export function SetConfig(arg1:conf.Config):Promise<void>;

export function SetRune(arg1:opgg.Rune):Promise<void>;

export function TierMeta():Promise<Array<types.RankedQueue>>;
