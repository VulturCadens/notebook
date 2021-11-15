<script setup>
    import { inject } from "vue"

    const store = inject("store")
</script>

<template>
    <div :class="store.state.text.style" v-if="store.state.text.value.length > 0">
        {{ store.state.text.value }}
    </div>
</template>

<style scoped>
    .foo {
        position: absolute;  
        bottom: 20px;
        left: 50px;
        font-size: 2.5rem;
        color: #464;
        text-shadow: 3px 3px 2px #111;
    }
</style>
