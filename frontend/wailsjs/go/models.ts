export namespace scraping {
	
	export enum Platform {
	    _NES = 0,
	    _GB = 1,
	    _SNES = 2,
	    _BSX = 3,
	    _VB = 4,
	    _N64 = 5,
	    _GBC = 6,
	    _GBA = 7,
	    _GCN = 8,
	    _NDS = 9,
	    _Wii = 10,
	    _DSi = 11,
	    _3DS = 12,
	    _WiiU = 13,
	    _Switch = 14,
	    _SMS = 15,
	    _GEN = 16,
	    _GG = 17,
	    _SCD = 18,
	    _32X = 19,
	    _SAT = 20,
	    _DC = 21,
	    _PS1 = 22,
	    _PS2 = 23,
	    _PS3 = 24,
	    _PS4 = 25,
	    _PS5 = 26,
	    _PSP = 27,
	    _PSVita = 28,
	    _XBOX = 29,
	    _XB360 = 30,
	    _XBLA = 31,
	    _XBONE = 32,
	    _3DO = 33,
	    _ARCADE = 34,
	    _A2600 = 35,
	    _A5200 = 36,
	    _A7800 = 37,
	    _AJ = 38,
	    _JCD = 39,
	    _AL = 40,
	    _MOB = 41,
	    _NEO = 42,
	    _PC = 43,
	    _CD_i = 44,
	    _TG_16 = 45,
	    _TGCD = 46,
	    _MUL = 47,
	}
	export class GameHash {
	    SHA1: string;
	    MD5: string;
	    CRC: string;
	
	    static createFrom(source: any = {}) {
	        return new GameHash(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.SHA1 = source["SHA1"];
	        this.MD5 = source["MD5"];
	        this.CRC = source["CRC"];
	    }
	}
	export class Rom {
	    Title: string;
	    Id: number;
	    Platform: Platform;
	    GameHash: GameHash;
	    CoverUrl: string;
	    CoverImg: any;
	    DownloadUrl: string;
	    PageUrl: string;
	    RomName: string;
	
	    static createFrom(source: any = {}) {
	        return new Rom(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Title = source["Title"];
	        this.Id = source["Id"];
	        this.Platform = source["Platform"];
	        this.GameHash = this.convertValues(source["GameHash"], GameHash);
	        this.CoverUrl = source["CoverUrl"];
	        this.CoverImg = source["CoverImg"];
	        this.DownloadUrl = source["DownloadUrl"];
	        this.PageUrl = source["PageUrl"];
	        this.RomName = source["RomName"];
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

export namespace userconfig {
	
	export class Config {
	    RetroachievmentKey: string;
	    AutoValidate: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.RetroachievmentKey = source["RetroachievmentKey"];
	        this.AutoValidate = source["AutoValidate"];
	    }
	}
	export class ConfigStore {
	
	
	    static createFrom(source: any = {}) {
	        return new ConfigStore(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

