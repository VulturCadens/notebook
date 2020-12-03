import { timercomponent } from "./component/timer"

const app = {
    components: {
        "timer-component": timercomponent
    }
}

Vue.createApp(app).mount("#application")