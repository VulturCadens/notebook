if ("serviceWorker" in navigator) {
    navigator.serviceWorker.register("/pwa/sw/service.js", { scope: "/pwa/sw/" })
        .then((registration) => {
            console.log("Registration succeeded. Scope is " + registration.scope)
        }).catch((error) => {
            console.log("Registration failed with " + error)
        })
}

window.onload = () => {
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
