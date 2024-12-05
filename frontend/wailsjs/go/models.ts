export namespace main {
	
	export class BubbleChatConfig {
	    // Go type: struct { Mode string "yaml:\"mode\"" }
	    ShowKey: any;
	    DisplayLimit: number;
	    // Go type: struct { BackgroundColor string "yaml:\"background_color\""; FontColor string "yaml:\"font_color\""; FontSize string "yaml:\"font_size\"" }
	    Style: any;
	    // Go type: struct { ShowUserIcon bool "yaml:\"show_user_icon\""; UserIconPath string "yaml:\"user_icon_path\"" }
	    Content: any;
	
	    static createFrom(source: any = {}) {
	        return new BubbleChatConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ShowKey = this.convertValues(source["ShowKey"], Object);
	        this.DisplayLimit = source["DisplayLimit"];
	        this.Style = this.convertValues(source["Style"], Object);
	        this.Content = this.convertValues(source["Content"], Object);
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
	export class Config {
	    BubbleChat: BubbleChatConfig;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.BubbleChat = this.convertValues(source["BubbleChat"], BubbleChatConfig);
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

