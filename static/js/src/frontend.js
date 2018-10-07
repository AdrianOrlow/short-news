function newsDescs() {
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
}

function checkForm() {
    var categories = ["polandCat", "worldCat", "economicsCat", "cultureCat", "sportsCat"];
    var media = ["tvn24Med", "rmf24Med", "polsatnewsMed"];
    var catBtn = document.getElementById("catBtn");
    var medBtn = document.getElementById("medBtn");
    var qInput = document.getElementById("quantity");

    var j = 0;
    for (i = 0; i < categories.length-1; i++) {
        category = document.getElementById(categories[i]);
        if (category.checked == true) {
            j++
        }
    }
    if (j == 0) {
        catBtn.classList.add("btn--red");
        return;
    } else if (catBtn.classList.contains("btn--red") == true) {
        catBtn.classList.remove("btn--red");
    }

    j = 0;
    for (i = 0; i < media.length; i++) {
        mediaX = document.getElementById(media[i]);
        if (mediaX.checked == true) {
            j++
        }
    }
    if (j == 0) {
        medBtn.classList.add("btn--red");
        return;
    } else if (medBtn.classList.contains("btn--red") == true) {
        medBtn.classList.remove("btn--red");
    }

    if (qInput.value < 40) {
        document.querySelector("#mainForm").submit();
    } else {
        qInput.classList.add("btn--red");
    } 
}