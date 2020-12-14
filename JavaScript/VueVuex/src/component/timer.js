export default {
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

    computed: {
        counter() {
            return this.$store.state.count
        }
    },

    mounted() {
        setInterval(() => {
            this.$store.commit("increment")
        }, 500)
    }
}
