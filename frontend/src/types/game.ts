export interface Skill {
    ids: Array<string>;
    order: Array<string>;
    win: number;
    play: number;
    pick_rate: number;
}


export interface Rune {
    primary_page_id: number;
    primary_rune_ids: Array<number>;
    secondary_page_id: number;
    secondary_rune_ids: Array<number>;
    stat_mod_ids: Array<number>;
    win: number;
    play: number;
    pick_rate: number;
}

export interface Item {
    ids: Array<number>;
    play: number;
    win: number;
    pick_rate: number;
}

export interface Champion {
    id: number;
    title: string;
    name: string;
    roles: Array<string>;
    win_rate: number;
    pick_rate: number;
    kda: number;
}

export interface ChampInfo {
    champion: Champion;
    runes: Array<Rune>;
    skill: Skill
    spells: Array<Item>;
    starter_items: Array<Item>;
    boots: Array<Item>;
    core_items: Array<Item>;
}