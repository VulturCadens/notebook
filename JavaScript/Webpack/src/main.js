const CounterApp = {
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

Vue.createApp(CounterApp).mount("#counter")