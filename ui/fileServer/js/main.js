const searchWrapper = document.querySelector(".search-input");
const inputBox = searchWrapper.querySelector("input");
const suggBox = searchWrapper.querySelector(".autocomplete");

inputBox.onkeyup = async (e) => {
    let userData = e.target.value;
    if (userData) {
        await fetch('/search?a=' + userData)
            .then(function (response) {
                // The response is a Response instance.
                // You parse the data into a useable format using `.json()`
                return response.json();
            }).then(function (data) {
                suggestions(data)
            });


    } else {
        searchWrapper.classList.remove("active"); //hide autocomplete box
    }
}

function suggestions(data) {

    suggBox.innerHTML = ""


    for (const [key, value] of Object.entries(data)) {
        let listItem = document.createElement('li');
        listItem.id = key
        listItem.setAttribute("onclick", "select(" + key + ")");
        listItem.innerHTML = value;
        suggBox.appendChild(listItem);
    }
    searchWrapper.classList.add("active")
    let allList = suggBox.querySelectorAll("li");

}

function select(id) {
    window.location.href = "http://localhost:8080/artist?id=" + id;
}

