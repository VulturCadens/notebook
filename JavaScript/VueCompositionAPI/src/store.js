import { reactive } from "vue"

const state = reactive({
    text: {
        style: "foo",
        value: "This is text."
    },

    menu: {
        style: "foo",
        title: "This is menu title",
        entries: ["First entry", "Second entry"]
    }
})

const methods = {
    textOff() {
        state.text.value = ""
    },
    
    textOn() {
        state.text.value = "This is text."
    }      
}
export default {
    state,
    methods
}