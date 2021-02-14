import timercomponent from "./component/timer"
import clickcomponent from "./component/click"

import store from "./store"

const app = Vue.createApp({
    components: {
        "timer-component": timercomponent,
        "click-component": clickcomponent
    }
})

app.use(store)

app.mount("#application")