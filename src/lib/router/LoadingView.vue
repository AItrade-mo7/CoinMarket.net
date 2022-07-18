<script setup lang="ts">
import { useMessage } from 'naive-ui';
import { LoadingStore } from '@/store';
window.$message = useMessage();
</script>

<template>
  <Transition>
    <div v-if="LoadingStore.status" class="Loading"></div>
  </Transition>
</template>

<style lang="less" scoped>
.Loading {
  position: fixed;
  left: 0;
  top: 0;
  width: 100vw;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: rgba(3, 3, 3, 0.4);
  z-index: 1999;
}

.Loading::after {
  content: '';
  display: block;
  width: 36px;
  height: 36px;
  animation: AppLoader 1s infinite linear;
  background-color: rgb(64, 158, 255.7);
}
@keyframes AppLoader {
  0% {
    transform: rotate(0deg);
    border-radius: 50%;
  }

  50% {
    transform: rotate(90deg);
    border-radius: 0%;
  }

  100% {
    transform: rotate(180deg);
    border-radius: 50%;
  }
}
</style>
