window.onload = () => {
    if ("serviceWorker" in navigator) {
        navigator.serviceWorker.register("/pwa/sw/service.js",{ scope: "/pwa/sw/" })
            .then((registration) => {
                console.log("Registration succeeded. Scope is " + registration.scope)
            }).catch((error) => {
                console.log("Registration failed with " + error)
            })
    }
}