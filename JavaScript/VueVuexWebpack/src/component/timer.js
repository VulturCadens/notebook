export default {
    name: "Timer",

    render() {
        return Vue.h(
            "h1", {}, "Counter is " + this.count
        )
    },

    /*
    template: `
        <h1>
            Counter is {{ this.count }}
        </h1>
    `,
    */

    methods: {
        ...Vuex.mapMutations([
            "increment"     // -> this.$store.commit("increment")
        ])
    },

    computed: {
        ...Vuex.mapState([
            "count"         // -> return this.$store.state.count
        ])
    },

    mounted() {
        setInterval(() => {
            this.increment()
        }, 500)
    }
}
