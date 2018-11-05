import * as wasm from "colada-lottery";

wasm.init().then((results) => {
    console.log("results", results);
    if (results) {
        if (results.error) {
            alert(results.error);
        }
    }
})

document.getElementById("lotteryBtn").addEventListener("click", () => {
    wasm.draw_barista_and_cleaner().then((results) => {
        console.log(results);
        var logEntry = results.result;
        console.log(logEntry);

        document.getElementById("baristaName").innerText = logEntry.barista;
        document.getElementById("baristaDrawnAt").innerText = logEntry.drawnAt;
        document.getElementById("baristaHeadshot").setAttribute("src", logEntry.baristaImg);
        document.getElementById("cleanerName").innerText = logEntry.cleaner;
        document.getElementById("cleanerDrawnAt").innerText = logEntry.drawnAt;
        document.getElementById("lotteryResultsContainer").style.display = "block";
    });
});