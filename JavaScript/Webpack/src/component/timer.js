export const timercomponent = {
    render() {
        return Vue.h("div", {}, this.counter)
    },

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