export const timercomponent = {
    name: "Timer",

    render() {
        return Vue.h(
            "h1", {}, "Counter is " + this.counter
        )
    },

    /*
    template: `
        <div>
            Counter is {{ counter }}
        </div>
    `,
    */

    data() {
        return {
            counter: 0
        }
    },

    mounted() {
        setInterval(() => {
            this.counter++
        }, 500)
    }
}
