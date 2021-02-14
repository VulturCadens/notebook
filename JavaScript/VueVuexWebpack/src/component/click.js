export default {
    name: "Click",

    template: `
        <div>
            <div class="container">
                <transition name="animation">
                    <div class="box" v-show="show"></div>
                </transition>
            </div>
            
            <h1>Click counter is {{ counter }}</h1>
            <button v-on:click="clicked">Click</button>
        </div>
    `,

    data() {
        return {
            counter: 0,
            show: true
        }
    },

    methods: {
        clicked() {
            this.show = !this.show
            this.counter++
        }
    }
}
