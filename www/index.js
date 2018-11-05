import * as wasm from "colada-lottery";

wasm.init().then((results) => {
    console.log("results", results);
    if (results) {
        if (results.error) {
            document.getElementById("error-msg").innerText = results.error;
            document.getElementById("error-msg").style.display = "block";
        }
    }
})

document.getElementById("error-msg").addEventListener('click', (evt) => {
    evt.target.style.display="none";
});

document.getElementById("lotteryBtn").addEventListener("click", () => {
    wasm.draw_barista_and_cleaner().then((results) => {
        var DrawingResult = results.result;
        
        document.getElementById("baristaName").innerText =  "Barista: " + DrawingResult.barista;
        document.getElementById("baristaDrawnAt").innerText = DrawingResult.drawnAt;
        document.getElementById("baristaHeadshot").setAttribute("src", DrawingResult.baristaImg);

        document.getElementById("cleanerName").innerText = "Cleaner: " + DrawingResult.cleaner;
        document.getElementById("cleanerDrawnAt").innerText = DrawingResult.drawnAt;
        document.getElementById("cleanerHeadshot").setAttribute("src", "http://localhost:8080"+DrawingResult.cleanerImg);
        
        document.getElementById("lotteryResultsContainer").style.display = "block";
    });
});