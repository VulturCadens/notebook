import { createStore } from "vuex"

export default createStore({
    state: {
        count: 0,
        elements: []
    },

    getters: {},

    mutations: {
        increment(state) {
            state.count++
            state.elements.push(state.count)
        }
    },

    actions: {},

    modules: {}
})