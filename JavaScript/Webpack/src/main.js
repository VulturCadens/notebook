const timercomponent = {
    template: '<div id="counter">Counter: {{ counter }}</div>',

    data() {
        return {
            counter: 0
        }
    },

    mounted() {
        setInterval(() => {
            this.counter++
        }, 200)
    }
}

const app = {
    components: {
        "timer-component": timercomponent
    }
}

Vue.createApp(app).mount("#application")