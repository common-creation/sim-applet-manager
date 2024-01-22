export namespace db {
	
	export class Key {
	    name: string;
	    aid: string;
	    encKey: string;
	    macKey: string;
	    kekKey: string;
	
	    static createFrom(source: any = {}) {
	        return new Key(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.aid = source["aid"];
	        this.encKey = source["encKey"];
	        this.macKey = source["macKey"];
	        this.kekKey = source["kekKey"];
	    }
	}
	export class Sim {
	    keys: Key[];
	
	    static createFrom(source: any = {}) {
	        return new Sim(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.keys = this.convertValues(source["keys"], Key);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

export namespace gp {
	
	export class HexFingerPrint {
	    hex: string;
	    fingerPrint: string;
	
	    static createFrom(source: any = {}) {
	        return new HexFingerPrint(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hex = source["hex"];
	        this.fingerPrint = source["fingerPrint"];
	    }
	}
	export class ListResult {
	    package: HexFingerPrint;
	    applets: HexFingerPrint[];
	
	    static createFrom(source: any = {}) {
	        return new ListResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.package = this.convertValues(source["package"], HexFingerPrint);
	        this.applets = this.convertValues(source["applets"], HexFingerPrint);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class Result {
	    success: boolean;
	    output: string;
	    status: string;
	
	    static createFrom(source: any = {}) {
	        return new Result(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.output = source["output"];
	        this.status = source["status"];
	    }
	}

}

export namespace main {
	
	export class SimInfo {
	    iccid: string;
	    config: db.Sim;
	
	    static createFrom(source: any = {}) {
	        return new SimInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.iccid = source["iccid"];
	        this.config = this.convertValues(source["config"], db.Sim);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

