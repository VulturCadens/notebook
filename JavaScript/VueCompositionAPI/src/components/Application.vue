<script setup>
    import Menu from "./Menu.vue"
    import Text from "./Text.vue"

    import { provide } from "vue"

    import store from "../store"

    provide("store", store)
</script>

<template>
    <img id="logo" src="../assets/logo.png" />

    <Menu />
    <Text />
</template>

<style>
    @font-face {
        font-family: "Oswald";
        font-style: normal;
        font-weight: 400;
        src: local("Oswald Regular"), local("Oswald-Regular"), url(/font/Oswald-Regular.woff) format("woff");
        unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
    }

    * {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }

    html {
        font-size: 16px;
    }

    body {
        display: grid;
        place-items: center;
        min-height: 100vh;
        background-color: #AAA;
    }

    #application {
        position: relative;

        width: 1000px;
        height: 500px;

        border: solid 1px #222;
        background-color: #EEE;

        font-family: "Oswald";
        font-weight: normal;

        text-align: left;

        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
    }

    #logo {
        position: absolute;
        top: 10px;
        left: 20px;
        filter: drop-shadow(0 0 1rem #FFD993);
    }
</style>
