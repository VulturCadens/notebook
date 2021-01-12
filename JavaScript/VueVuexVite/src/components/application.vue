<template>
    <img alt="Image" src="../assets/image.png" />

    <section-component :message=title />
</template>

<script>
    import SectionComponent from "./section.vue";

    export default {
        components: {
            SectionComponent
        },

        data() {
            return {
                title: "Hello World"
            };
        }
    };
</script>

<style>
    #application {
        font-family: "Oswald";
        color: #2c3e50;
        text-align: center;
        margin-top: 5em;
    }
</style>