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

const __wbg_log_420772b956a38cc5_target = console.log;

export function __wbg_log_420772b956a38cc5(arg0, arg1) {
    let varg0 = getStringFromWasm(arg0, arg1);
    __wbg_log_420772b956a38cc5_target(varg0);
}
/**
* @returns {void}
*/
export function greet() {
    return wasm.greet();
}

/**
* @returns {void}
*/
export function init() {
    return wasm.init();
}

/**
* @returns {void}
*/
export function draw_barista_and_cleaner() {
    return wasm.draw_barista_and_cleaner();
}

