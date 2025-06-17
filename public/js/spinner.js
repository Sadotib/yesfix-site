
document.body.addEventListener("htmx:request", function (evt) {
    document.getElementById("submit-text").classList.add("hidden");
    document.getElementById("submit-loader").classList.remove("hidden");
});

document.body.addEventListener("htmx:afterSwap", function (evt) {
    document.getElementById("submit-text").classList.remove("hidden");
    document.getElementById("submit-loader").classList.add("hidden");
});

document.body.addEventListener("htmx:responseError", function (evt) {
    document.getElementById("submit-text").classList.remove("hidden");
    document.getElementById("submit-loader").classList.add("hidden");
});

