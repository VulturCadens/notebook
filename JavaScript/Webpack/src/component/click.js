export const clickcomponent = {
    name: "Click",

    template: `
        <div>
            <h1>Click counter is {{ counter }}</h1>
            <button v-on:click="clicked">Click</button>
        </div>
    `,

    data() {
        return {
            counter: 0
        }
    },

    methods: {
        clicked() {
            this.counter++
        }
    }
}
