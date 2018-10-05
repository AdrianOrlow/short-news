var newsDescs = document.getElementsByClassName("news--desc");

for (var i = 0; i < newsDescs.length; i++) {
    newsDescs[i].addEventListener("click", function () {
        this.childNodes[3].classList.toggle("fa-angle-up");
        var content = this.parentNode.childNodes[3];
        if (content.style.display === "block") {
            content.style.display = "none";
        } else {
            content.style.display = "block";
        }
    });
}