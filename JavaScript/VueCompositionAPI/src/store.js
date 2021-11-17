import { reactive } from "vue"

const state = reactive({
    text: {
        style: "foo",
        value: "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
    },

    menu: {
        style: "foo",
        title: "This is menu title",
        entries: [
            {
                text: "Text ON",
                args: "ON"
            },
            {
                text: "Text OFF",
                args: "OFF"
            }
        ]
    }
})

const methods = {
    func(args) {
        if (args === "ON") this.textOn()
        if (args === "OFF") this.textOff()
    },

    textOff() {
        state.text.value = ""
    },

    textOn() {
        state.text.value = "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
    }
}

export default {
    state,
    methods
}