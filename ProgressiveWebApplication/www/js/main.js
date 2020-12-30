if ("serviceWorker" in navigator) {
    navigator.serviceWorker.register("/service.js")
        .then((registration) => {
            console.log("Registration succeeded. Scope is " + registration.scope)
        }).catch((error) => {
            console.log("Registration failed with " + error)
        })
}

window.onload = () => {
    setTimeout(fetchImage, 3000)
}

function fetchImage() {
    const imageElement = document.createElement("img")
    const parentElement = document.getElementById("parent")

    fetch("/image/picture.jpg")
        .then(response => response.blob())

        .then(blob => {
            const imageURL = window.URL.createObjectURL(blob)

            imageElement.src = imageURL
            parentElement.appendChild(imageElement)
        })

        .catch((error) => {
            console.log("Image failed with " + error)
        })
}
