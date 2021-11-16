const searchWrapper = document.querySelector(".search-input");
const inputBox = searchWrapper.querySelector("input");
const suggBox = searchWrapper.querySelector(".autocomplete")

let userInput
let d
inputBox.onkeyup = async (e) => {
    userInput = e.target.value;
    if (userInput) {
        await fetch('/search?a=' + userInput)
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
    let tabIndex = 0

    let arr;

    data.forEach((data) => {

        if (!Array.isArray(data)) {
            for (const [key, value] of Object.entries(data)) {
                d = data
                let listItem = document.createElement('li');
                listItem.setAttribute("onclick", "result(" + key + ")");
                listItem.setAttribute("tabindex", tabIndex);
                listItem.innerHTML = value;
                suggBox.appendChild(listItem);
            }
        } else {
            // Did you mean: text

            if (userInput !== data[0]) {
                let p = document.createElement('p')
                p.innerText = "Did you mean: "
                p.className = "didMean"
                suggBox.appendChild(p)

                arr = []
                data.forEach((value, i) => {
                    let listItem = document.createElement('li');
                    listItem.setAttribute("onclick", `suggestion("` + value + `")`);
                    listItem.setAttribute("tabindex", tabIndex);
                    listItem.innerHTML = value;

                    // check if suggBox don't have same strings
                    if (!arr.includes(value)) {
                        arr.push(value)
                        suggBox.appendChild(listItem);
                    }
                })
            }

        }

        searchWrapper.classList.add("active")
    })


}

function result(id) {
    window.location.href = "/artist?id=" + id;
}

function suggestion(x) {
    inputBox.value = x;

    inputBox.dispatchEvent(new KeyboardEvent('keyup', {
        'key': 'a'
    }));
}

document.addEventListener("keydown", function (evt) {
    if (evt.ctrlKey && evt.key === 's') {
        inputBox.focus()
    }
    getSelectedElement(evt)
});


const getSelectedElement = (evt) => {
    if (suggBox.contains(document.activeElement)) {
        if (evt.key === 'Enter') {
            Function("return " + document.activeElement.getAttribute("onclick"))()
        }
    }
}


inputBox.addEventListener("focusin", function (evt) {
    el = document.querySelector("i")
    if (inputBox === document.activeElement || inputBox.value.length !== 0) {
        el.classList.add("disable")
    }
})
inputBox.addEventListener("focusout", function (evt) {
    el = document.querySelector("i")
    if (inputBox !== document.activeElement && inputBox.value.length === 0) {
        el.classList.remove("disable")
    }
})


// if press key enter and focus in input
inputBox.addEventListener("keydown", function (evt) {
    if (evt.key === 'Enter') {
        evt.preventDefault();

        for (const [key, value] of Object.entries(d)) {
            if (userInput.toLowerCase() === value.toLowerCase()) {
                result(key)
            }
        }
    }

}, false)

