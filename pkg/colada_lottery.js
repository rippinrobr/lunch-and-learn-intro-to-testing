/* tslint:disable */
import * as wasm from './colada_lottery_bg';

let cachedTextDecoder = new TextDecoder('utf-8');

let cachegetUint8Memory = null;
function getUint8Memory() {
    if (cachegetUint8Memory === null || cachegetUint8Memory.buffer !== wasm.memory.buffer) {
        cachegetUint8Memory = new Uint8Array(wasm.memory.buffer);
    }
    return cachegetUint8Memory;
}

function getStringFromWasm(ptr, len) {
    return cachedTextDecoder.decode(getUint8Memory().subarray(ptr, ptr + len));
}

export function __wbg_alert_7b535868f880270c(arg0, arg1) {
    let varg0 = getStringFromWasm(arg0, arg1);
    alert(varg0);
}
/**
* @returns {void}
*/
export function greet() {
    return wasm.greet();
}

const stack = [];

const slab = [{ obj: undefined }, { obj: null }, { obj: true }, { obj: false }];

function getObject(idx) {
    if ((idx & 1) === 1) {
        return stack[idx >> 1];
    } else {
        const val = slab[idx >> 1];

        return val.obj;

    }
}

let slab_next = slab.length;

function dropRef(idx) {

    idx = idx >> 1;
    if (idx < 4) return;
    let obj = slab[idx];

    obj.cnt -= 1;
    if (obj.cnt > 0) return;

    // If we hit 0 then free up our space in the slab
    slab[idx] = slab_next;
    slab_next = idx;
}

function takeObject(idx) {
    const ret = getObject(idx);
    dropRef(idx);
    return ret;
}
/**
* @returns {any}
*/
export function init() {
    return takeObject(wasm.init());
}

/**
* @returns {any}
*/
export function draw_barista_and_cleaner() {
    return takeObject(wasm.draw_barista_and_cleaner());
}

const __widl_f_set_Headers_target = typeof Headers === 'undefined' ? null : Headers.prototype.set || function() {
    throw new Error(`wasm-bindgen: Headers.set does not exist`);
};

let cachegetUint32Memory = null;
function getUint32Memory() {
    if (cachegetUint32Memory === null || cachegetUint32Memory.buffer !== wasm.memory.buffer) {
        cachegetUint32Memory = new Uint32Array(wasm.memory.buffer);
    }
    return cachegetUint32Memory;
}

function addHeapObject(obj) {
    if (slab_next === slab.length) slab.push(slab.length + 1);
    const idx = slab_next;
    const next = slab[idx];

    slab_next = next;

    slab[idx] = { obj, cnt: 1 };
    return idx << 1;
}

export function __widl_f_set_Headers(arg0, arg1, arg2, arg3, arg4, exnptr) {
    let varg1 = getStringFromWasm(arg1, arg2);
    let varg3 = getStringFromWasm(arg3, arg4);
    try {
        __widl_f_set_Headers_target.call(getObject(arg0), varg1, varg3);
    } catch (e) {
        const view = getUint32Memory();
        view[exnptr / 4] = 1;
        view[exnptr / 4 + 1] = addHeapObject(e);

    }
}

export function __widl_f_new_with_str_and_init_Request(arg0, arg1, arg2, exnptr) {
    let varg0 = getStringFromWasm(arg0, arg1);
    try {
        return addHeapObject(new Request(varg0, getObject(arg2)));
    } catch (e) {
        const view = getUint32Memory();
        view[exnptr / 4] = 1;
        view[exnptr / 4 + 1] = addHeapObject(e);

    }
}

function GetOwnOrInheritedPropertyDescriptor(obj, id) {
    while (obj) {
        let desc = Object.getOwnPropertyDescriptor(obj, id);
        if (desc) return desc;
        obj = Object.getPrototypeOf(obj);
    }
return {}
}

const __widl_f_headers_Request_target = GetOwnOrInheritedPropertyDescriptor(typeof Request === 'undefined' ? null : Request.prototype, 'headers').get || function() {
    throw new Error(`wasm-bindgen: Request.headers does not exist`);
};

export function __widl_f_headers_Request(arg0) {
    return addHeapObject(__widl_f_headers_Request_target.call(getObject(arg0)));
}

export function __widl_instanceof_Response(idx) {
    return getObject(idx) instanceof Response ? 1 : 0;
}

const __widl_f_json_Response_target = typeof Response === 'undefined' ? null : Response.prototype.json || function() {
    throw new Error(`wasm-bindgen: Response.json does not exist`);
};

export function __widl_f_json_Response(arg0, exnptr) {
    try {
        return addHeapObject(__widl_f_json_Response_target.call(getObject(arg0)));
    } catch (e) {
        const view = getUint32Memory();
        view[exnptr / 4] = 1;
        view[exnptr / 4 + 1] = addHeapObject(e);

    }
}

export function __widl_instanceof_Window(idx) {
    return getObject(idx) instanceof Window ? 1 : 0;
}

const __widl_f_fetch_with_request_Window_target = function(x0) {
    return this.fetch(x0);
};

export function __widl_f_fetch_with_request_Window(arg0, arg1) {
    return addHeapObject(__widl_f_fetch_with_request_Window_target.call(getObject(arg0), getObject(arg1)));
}

export function __wbg_newnoargs_b5dbe629f3c72f37(arg0, arg1) {
    let varg0 = getStringFromWasm(arg0, arg1);
    return addHeapObject(new Function(varg0));
}

const __wbg_call_80c8cb20bdc473db_target = typeof Function === 'undefined' ? null : Function.prototype.call || function() {
    throw new Error(`wasm-bindgen: Function.call does not exist`);
};

export function __wbg_call_80c8cb20bdc473db(arg0, arg1, exnptr) {
    try {
        return addHeapObject(__wbg_call_80c8cb20bdc473db_target.call(getObject(arg0), getObject(arg1)));
    } catch (e) {
        const view = getUint32Memory();
        view[exnptr / 4] = 1;
        view[exnptr / 4 + 1] = addHeapObject(e);

    }
}

const __wbg_call_9a6451120c0e32db_target = typeof Function === 'undefined' ? null : Function.prototype.call || function() {
    throw new Error(`wasm-bindgen: Function.call does not exist`);
};

export function __wbg_call_9a6451120c0e32db(arg0, arg1, arg2, exnptr) {
    try {
        return addHeapObject(__wbg_call_9a6451120c0e32db_target.call(getObject(arg0), getObject(arg1), getObject(arg2)));
    } catch (e) {
        const view = getUint32Memory();
        view[exnptr / 4] = 1;
        view[exnptr / 4 + 1] = addHeapObject(e);

    }
}

export function __wbg_new_3a83c5616033fe87() {
    return addHeapObject(new Object());
}

const __wbg_set_560a0f4bd944cd79_target = (typeof Reflect === 'undefined' ? null : Reflect.set || function() {
    throw new Error(`wasm-bindgen: Reflect.set does not exist`);
}).bind(Reflect);

export function __wbg_set_560a0f4bd944cd79(arg0, arg1, arg2, exnptr) {
    try {
        return __wbg_set_560a0f4bd944cd79_target(getObject(arg0), getObject(arg1), getObject(arg2)) ? 1 : 0;
    } catch (e) {
        const view = getUint32Memory();
        view[exnptr / 4] = 1;
        view[exnptr / 4 + 1] = addHeapObject(e);

    }
}

let cachedGlobalArgumentPtr = null;
function globalArgumentPtr() {
    if (cachedGlobalArgumentPtr === null) {
        cachedGlobalArgumentPtr = wasm.__wbindgen_global_argument_ptr();
    }
    return cachedGlobalArgumentPtr;
}

function getGlobalArgument(arg) {
    const idx = globalArgumentPtr() / 4 + arg;
    return getUint32Memory()[idx];
}

export function __wbg_new_5ec8c05dbefa35f1(arg0) {
    let cbarg0 = function(arg0, arg1) {
        let a = this.a;
        this.a = 0;
        try {
            return this.f(a, this.b, addHeapObject(arg0), addHeapObject(arg1));

        } finally {
            this.a = a;

        }

    };
    cbarg0.f = wasm.__wbg_function_table.get(arg0);
    cbarg0.a = getGlobalArgument(0);
    cbarg0.b = getGlobalArgument(0 + 1);
    try {
        return addHeapObject(new Promise(cbarg0.bind(cbarg0)));
    } finally {
        cbarg0.a = cbarg0.b = 0;

    }
}

const __wbg_resolve_8d5554a087221239_target = (typeof Promise === 'undefined' ? null : Promise.resolve || function() {
    throw new Error(`wasm-bindgen: Promise.resolve does not exist`);
}).bind(Promise);

export function __wbg_resolve_8d5554a087221239(arg0) {
    return addHeapObject(__wbg_resolve_8d5554a087221239_target(getObject(arg0)));
}

const __wbg_then_1aeefc3ac0acfca9_target = typeof Promise === 'undefined' ? null : Promise.prototype.then || function() {
    throw new Error(`wasm-bindgen: Promise.then does not exist`);
};

export function __wbg_then_1aeefc3ac0acfca9(arg0, arg1) {
    return addHeapObject(__wbg_then_1aeefc3ac0acfca9_target.call(getObject(arg0), getObject(arg1)));
}

const __wbg_then_2c5a1ba95df01fe2_target = typeof Promise === 'undefined' ? null : Promise.prototype.then || function() {
    throw new Error(`wasm-bindgen: Promise.then does not exist`);
};

export function __wbg_then_2c5a1ba95df01fe2(arg0, arg1, arg2) {
    return addHeapObject(__wbg_then_2c5a1ba95df01fe2_target.call(getObject(arg0), getObject(arg1), getObject(arg2)));
}

export function __wbindgen_object_clone_ref(idx) {
    // If this object is on the stack promote it to the heap.
    if ((idx & 1) === 1) return addHeapObject(getObject(idx));

    // Otherwise if the object is on the heap just bump the
    // refcount and move on
    const val = slab[idx >> 1];
    val.cnt += 1;
    return idx;
}

export function __wbindgen_object_drop_ref(i) {
    dropRef(i);
}

export function __wbindgen_string_new(p, l) {
    return addHeapObject(getStringFromWasm(p, l));
}

export function __wbindgen_number_get(n, invalid) {
    let obj = getObject(n);
    if (typeof(obj) === 'number') return obj;
    getUint8Memory()[invalid] = 1;
    return 0;
}

export function __wbindgen_is_null(idx) {
    return getObject(idx) === null ? 1 : 0;
}

export function __wbindgen_is_undefined(idx) {
    return getObject(idx) === undefined ? 1 : 0;
}

export function __wbindgen_boolean_get(i) {
    let v = getObject(i);
    if (typeof(v) === 'boolean') {
        return v ? 1 : 0;
    } else {
        return 2;
    }
}

export function __wbindgen_is_symbol(i) {
    return typeof(getObject(i)) === 'symbol' ? 1 : 0;
}

let cachedTextEncoder = new TextEncoder('utf-8');

function passStringToWasm(arg) {

    const buf = cachedTextEncoder.encode(arg);
    const ptr = wasm.__wbindgen_malloc(buf.length);
    getUint8Memory().set(buf, ptr);
    return [ptr, buf.length];
}

export function __wbindgen_string_get(i, len_ptr) {
    let obj = getObject(i);
    if (typeof(obj) !== 'string') return 0;
    const [ptr, len] = passStringToWasm(obj);
    getUint32Memory()[len_ptr / 4] = len;
    return ptr;
}

export function __wbindgen_cb_drop(i) {
    const obj = getObject(i).original;
    dropRef(i);
    if (obj.cnt-- == 1) {
        obj.a = 0;
        return 1;
    }
    return 0;
}

export function __wbindgen_closure_wrapper478(a, b, fi, di, _ignored) {
    const f = wasm.__wbg_function_table.get(fi);
    const d = wasm.__wbg_function_table.get(di);
    const cb = function(arg0) {
        this.cnt++;
        let a = this.a;
        this.a = 0;
        try {
            return f(a, b, addHeapObject(arg0));

        } finally {
            this.a = a;
            if (this.cnt-- == 1) d(this.a, b);

        }

    };
    cb.a = a;
    cb.cnt = 1;
    let real = cb.bind(cb);
    real.original = cb;
    return addHeapObject(real);
}

export function __wbindgen_throw(ptr, len) {
    throw new Error(getStringFromWasm(ptr, len));
}

