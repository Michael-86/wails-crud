export namespace main {
	
	export class Kabel {
	    first: string;
	    last: string;
	
	    static createFrom(source: any = {}) {
	        return new Kabel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.first = source["first"];
	        this.last = source["last"];
	    }
	}

}

