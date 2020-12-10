import { timercomponent } from "./component/timer"
import { clickcomponent } from "./component/click"

const app = {
    components: {
        "timer-component": timercomponent,
        "click-component": clickcomponent
    }
}

Vue.createApp(app).mount("#application")