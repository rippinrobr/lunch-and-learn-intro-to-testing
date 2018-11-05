import * as wasm from "colada-lottery";

wasm.init().then((data) => {
    if (data) {
        if (data.error) {
            document.getElementById("error-msg").innerText = "No previous drawing found";
            document.getElementById("error-msg").style.display = "block";
            return;
        }

        displayBaristaAndCleaner(data.result);
    }
})

document.getElementById("error-msg").addEventListener('click', (evt) => {
    evt.target.style.display="none";
});

document.getElementById("lotteryBtn").addEventListener("click", () => {
    wasm.draw_barista_and_cleaner().then((data) => {
        document.getElementById("error-msg").style.display = "none";
        displayBaristaAndCleaner(data.result);
        
    });
});

function displayBaristaAndCleaner(res) {
    document.getElementById("baristaName").innerText =  "Barista: " + res.barista;
    document.getElementById("baristaDrawnAt").innerText = res.drawnAt;
    document.getElementById("baristaHeadshot").setAttribute("src", res.baristaImg);

    document.getElementById("cleanerName").innerText = "Cleaner: " + res.cleaner;
    document.getElementById("cleanerDrawnAt").innerText = res.drawnAt;
    document.getElementById("cleanerHeadshot").setAttribute("src", "http://localhost:8080"+res.cleanerImg);
    
    document.getElementById("lotteryResultsContainer").style.display = "block";
}