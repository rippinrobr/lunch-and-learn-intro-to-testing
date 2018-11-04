extern crate cfg_if;
extern crate wasm_bindgen;
extern crate futures;
extern crate js_sys;
extern crate wasm_bindgen_futures;
extern crate web_sys;
#[macro_use]
extern crate serde_derive;
extern crate serde;

use futures::{future, Future};
use js_sys::Promise;
use wasm_bindgen::prelude::*;
use wasm_bindgen::JsCast;
use wasm_bindgen_futures::future_to_promise;
use wasm_bindgen_futures::JsFuture;
use web_sys::{Request, RequestInit, RequestMode, Response};

mod utils;

use cfg_if::cfg_if;
use wasm_bindgen::prelude::*;

cfg_if! {
    // When the `wee_alloc` feature is enabled, use `wee_alloc` as the global
    // allocator.
    if #[cfg(feature = "wee_alloc")] {
        extern crate wee_alloc;
        #[global_allocator]
        static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;
    }
}

#[wasm_bindgen]
extern {
    fn alert(s: &str);
    #[wasm_bindgen(js_namespace = console)]
    fn log(msg: &str);
}

#[wasm_bindgen]
pub fn greet() {
    alert("Hello, colada-lottery!");
}

#[wasm_bindgen]
pub fn init() {
    log("initializing");
    log("I will fetch the history");
    get_previous_results();
}

#[wasm_bindgen]
pub fn draw_barista_and_cleaner() {
    log("pretend draw")
}

// #[derive(Debug,Serialize)]
// struct LogEntry {
// 	id:      i32,
// 	barista: String,
// 	cleaner: String,
// 	drawnAt: String,
// }

fn get_previous_results()  -> Promise {
    log("[get_previous_results] Top");
    let mut opts = RequestInit::new();
    opts.method("GET");
    opts.mode(RequestMode::Cors);

    log("[get_previous_results] about to create the request");
    let request = Request::new_with_str_and_init(
        "http://localhost:9999/v1/history/latest",
        &opts,
    ).unwrap();

    log("[get_previous_results] setting up headers");
    request
        .headers()
        .set("Accept", "application/json")
        .unwrap();

    let window = web_sys::window().unwrap();
    let request_promise = window.fetch_with_request(&request);
    log("[get_previous_results] mad the fetch_with_request call");

    let future = JsFuture::from(request_promise)
        .and_then(|resp_value| {
            log("IN the first and_then)")
            // `resp_value` is a `Response` object.
            assert!(resp_value.is_instance_of::<Response>());
            let resp: Response = resp_value.dyn_into().unwrap();
            resp.json()
        }).and_then(|json_value: Promise| {
            log("In the second and_then");
            // Convert this other `Promise` into a rust `Future`.
            JsFuture::from(json_value)
        }).and_then(|json| {
            log("In the last and_then");
            // Use serde to parse the JSON into a struct.
            //let latest_history: LogEntry = json.into_serde().unwrap();

            // Send the `Branch` struct back to JS as an `Object`.
            log(&format!("{:#?}", json));
            future::ok(json)
        });

    // Convert this Rust `Future` back into a JS `Promise`.
    future_to_promise(future)

}   
