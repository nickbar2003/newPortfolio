
<script setup>
import Header from './Header.vue'
import Projects from './Projects.vue'
import Education from './Education.vue'
import Work from './Work.vue'
</script>

<template>
    <div id="appwrapper">
        <Header />
        <Projects />
        <Education />
        <Work />
    </div>
</template>

<style scoped>
#appwrapper {
    display: flex;
    flex-direction: column;
    /* justify-content: center; */
    align-items: center;

}
</style>