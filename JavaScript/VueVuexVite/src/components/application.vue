<template>
    <img alt="Image" src="../assets/image.png" />

    <section-component :message="title" />
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
    img {
        max-width: 512px;
        width: 98%;
        height: auto;
    }

    #application {
        font-family: "Oswald";
        color: #2c3e50;
        text-align: center;
        margin-top: 5em;
    }
</style>