<script setup>
    import { inject } from "vue"

    const store = inject("store")
</script>

<template>
    <div :class="store.state.menu.style"
        v-if="Object.keys(store.state.menu).length > 0"
    >

        <div class="title">
            {{ store.state.menu.title }}
        </div>

        <div class="entry"
            v-for="(entry, index) in store.state.menu.entries"
            :key="index"
            @click="store.methods.func(entry.args)"
        >
            {{ entry.text }}
        </div>

    </div>
</template>

<style scoped>
    .foo {
        display: inline-block;
        position: absolute;
        top: 50px;
        left: 450px;
        background-color: #EEF;
        border: solid 1px #222;
    }

    .title {
        font-size: 1.5rem;     
        padding: .1rem .3rem .5rem .5rem;
    }

    .entry {
        font-size: 1.2rem;
        border-top: solid 1px #222;
        padding: .2rem .2rem .5rem .5rem;
    }

    .entry:hover {
        background-color: #DDE;
    }
</style>
