<script setup lang="ts">
import { RouterView } from 'vue-router';
import BackPage from '@/lib/BackPage.vue';
</script>

<template>
  <h1 class="PageTitle">
    <BackPage />
  </h1>
  <div class="PageWrapper">
    <RouterView />
  </div>
</template>
<style lang="less" scoped>
.PageWrapper {
  padding-top: 36px;
}
</style>
