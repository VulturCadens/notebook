<template>
    <h1>{{ message }}</h1>

    <button @click="increment">Count is: {{ count }}</button>

    <template v-for="element in elements" :key="element">
        <p>{{ element }}. Lorem ipsum dolores.</p>
    </template>
</template>

<script>
    import { mapMutations, mapState } from "vuex";

    export default {
        name: "Section Component",

        props: ["message"],

        methods: {
            ...mapMutations([
                "increment" // -> this.$store.commit("increment")
            ])
        },

        computed: {
            ...mapState([
                "count", // -> return this.$store.state.count
                "elements"
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