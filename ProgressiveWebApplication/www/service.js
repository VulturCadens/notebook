self.addEventListener("install", (event) => {
    console.log("Install event.")

    event.waitUntil(
        caches.open("V1").then((cache) => {
            return cache.addAll([
                "./image/picture.jpg"
            ])
        })
    )
})

self.addEventListener("fetch", (event) => {
    console.log("Fetch event.")

    event.respondWith(
        caches.match(event.request).then((response) => {
            return response || fetch(event.request);
        })
    )
})

self.addEventListener("activate", (event) => {
    console.log("Activate event.")
})
