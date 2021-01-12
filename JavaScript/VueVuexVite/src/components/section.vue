<template>
    <h1>{{ message }}</h1>

    <button @click="increment">Count is: {{ count }}</button>

    <p>Lorem ipsum dolores.</p>
</template>

<script>
    import { mapMutations, mapState } from "vuex";

    export default {
        name: "Component",

        props: ["message"],

        methods: {
            ...mapMutations([
                "increment" // -> this.$store.commit("increment")
            ])
        },

        computed: {
            ...mapState([
                "count" // -> return this.$store.state.count
            ])
        }
    };
</script>

<style scoped>
    button {
        font-size: 1.5em;
        padding: 0.5em;
        margin: 1em;
        border: 1px black solid;
        border-radius: 0.2em;
    }
</style>