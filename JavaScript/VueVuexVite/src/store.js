import { createStore } from "vuex"

export default createStore({
    state: {
        count: 0,
        elements: []
    },

    getters: {},

    mutations: {
        increment(state) {
            state.elements.push(state.count)
            state.count++
        }
    },

    actions: {},

    modules: {}
})