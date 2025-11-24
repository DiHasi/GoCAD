export namespace parser {
	
	export class Element {
	    name: number;
	    x: number;
	    y: number;
	
	    static createFrom(source: any = {}) {
	        return new Element(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.x = source["x"];
	        this.y = source["y"];
	    }
	}
	export class ParseResult {
	    elem_names: string[];
	    net_names: string[];
	    Q: Record<number, any>;
	    R: Record<number, any>;
	    plot: Element[];
	    opt_plot: Element[];
	    D: number[][];
	    NetElements: Record<number, Array<Element>>;
	    F_before: number;
	    F_after: number;
	
	    static createFrom(source: any = {}) {
	        return new ParseResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.elem_names = source["elem_names"];
	        this.net_names = source["net_names"];
	        this.Q = source["Q"];
	        this.R = source["R"];
	        this.plot = this.convertValues(source["plot"], Element);
	        this.opt_plot = this.convertValues(source["opt_plot"], Element);
	        this.D = source["D"];
	        this.NetElements = this.convertValues(source["NetElements"], Array<Element>, true);
	        this.F_before = source["F_before"];
	        this.F_after = source["F_after"];
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

