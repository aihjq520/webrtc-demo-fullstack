<script setup lang="ts">
import { onMounted } from 'vue';

onMounted(()=>{
  console.log('mount')
})

</script>

<template> 
<div style="display: flex; justify-content: space-between; flex-direction:column;">
  <router-view ></router-view>
</div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
