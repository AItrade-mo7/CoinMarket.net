<script setup lang="ts">
import { isPwa } from '@/utils/tools';
import { useRouter } from 'vue-router';
const $router = useRouter();

const installPwa = async () => {
  console.info('window.deferredPrompt', window.deferredPrompt);
  if (window.deferredPrompt) {
    window.deferredPrompt.prompt();
    const { outcome } = await window.deferredPrompt.userChoice;
    if (outcome === 'accepted') {
      window.deferredPrompt = null;
    }
  } else {
    $router.push('/about/pwa');
  }
};
</script>

<template>
  <div v-if="!isPwa()" @click="installPwa">
    <slot></slot>
  </div>
</template>

<style lang="less" scoped></style>
