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

