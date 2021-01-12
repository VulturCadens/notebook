import { createApp } from "vue"

import application from "./components/application.vue"

import store from "./store"

const app = createApp(application)

app.use(store)

app.mount("#application")
