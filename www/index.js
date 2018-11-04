import * as wasm from "colada-lottery";

wasm.init()

document.getElementById("lotteryBtn").addEventListener("click", () => {
    document.getElementById("lotteryResultsContainer").style.display = "block";
});