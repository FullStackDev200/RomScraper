export namespace scraping {
	
	export class Game {
	    Title: string;
	    Id: number;
	    Link: string;
	    CoverImg: any;
	    CoverUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new Game(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Title = source["Title"];
	        this.Id = source["Id"];
	        this.Link = source["Link"];
	        this.CoverImg = source["CoverImg"];
	        this.CoverUrl = source["CoverUrl"];
	    }
	}
	export class Rom {
	    Title: string;
	    Id: number;
	    Platform: string;
	    // Go type: struct { SHA1 string; MD5 string; CRC string }
	    GameHash: any;
	    CoverUrl: string;
	    DownloadUrl: string;
	    PageUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new Rom(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Title = source["Title"];
	        this.Id = source["Id"];
	        this.Platform = source["Platform"];
	        this.GameHash = this.convertValues(source["GameHash"], Object);
	        this.CoverUrl = source["CoverUrl"];
	        this.DownloadUrl = source["DownloadUrl"];
	        this.PageUrl = source["PageUrl"];
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

