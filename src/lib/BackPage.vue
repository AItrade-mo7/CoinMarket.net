<script setup lang="ts">
import XIcon from '@/lib/XIcon.vue';
import { useRouter } from 'vue-router';
const $router = useRouter();

const back = () => {
  $router.go(-1);
};
</script>

<template>
  <n-button text size="tiny" class="BackPage" @click="back">
    <template #icon>
      <XIcon name="ArrowLeftOutlined" />
    </template>
    Back
  </n-button>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';

.BackPage {
  color: @mainColor;
}
</style>
