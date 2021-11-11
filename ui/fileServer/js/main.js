const searchWrapper = document.querySelector(".search-input");
const inputBox = searchWrapper.querySelector("input");
const suggBox = searchWrapper.querySelector(".autocomplete");


inputBox.onkeyup = async (e) => {
    let userData = e.target.value;
    if (userData) {
        await fetch('/search?a=' + userData)
            .then(function (response) {
                // The response is a Response instance.
                // You parse the data into a usable format using `.json()`
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


    let arr;
    if (!Array.isArray(data)) {
        for (const [key, value] of Object.entries(data)) {
            let listItem = document.createElement('li');
            listItem.setAttribute("onclick", "result(" + key + ")");
            listItem.innerHTML = value;
            suggBox.appendChild(listItem);
        }
    } else {
        // Did you mean: text
        let p = document.createElement('p')
        p.innerText = "Did you mean: "
        p.className = "didMean"
        suggBox.appendChild(p)

        arr = []
        data.forEach((value, i) => {
            let listItem = document.createElement('li');
            listItem.setAttribute("onclick", `suggestion("` + value + `")`);
            listItem.innerHTML = value;

            // check if suggBox don't have same strings
            if (!arr.includes(value)) {
                arr.push(value)
                suggBox.appendChild(listItem);
            }
        })
    }

    searchWrapper.classList.add("active")
}

function result(id) {
    window.location.href = "http://localhost:8080/artist?id=" + id;
}

function suggestion(x) {
    inputBox.value = x;

    inputBox.dispatchEvent(new KeyboardEvent('keyup', {
        'key': 'a'
    }));
}



