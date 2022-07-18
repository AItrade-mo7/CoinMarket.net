<script setup lang="ts">
import { Logo } from '@/config/constant';
import { TopBarStore } from '@/store';
import BackPage from '@/lib/BackPage.vue';
import XIcon from '@/lib/XIcon.vue';

import { useRouter } from 'vue-router';
const $router = useRouter();

const goHome = () => {
  $router.replace('/');
};

const goBack = () => {
  $router.go(-1);
};
</script>

<template>
  <h1 class="PageTitle"><BackPage /></h1>
  <div class="PageWrapper NotFundApp">
    <div className="NotFundApp__logo-box">
      <img className="NotFundApp__logo" :src="Logo" alt="" />
    </div>
    <h2>404 Not Found</h2>

    <n-button @click="goHome">
      <template #icon>
        <XIcon name="HomeOutlined" />
      </template>
      返回首页
    </n-button>

    <n-button type="primary" @click="TopBarStore.open()">
      <template #icon>
        <XIcon name="MenuFoldOutlined" />
      </template>
      浏览菜单
    </n-button>

    <n-button dashed type="primary" @click="goBack">
      <template #icon>
        <XIcon name="ArrowLeftOutlined" />
      </template>
      返回上一页
    </n-button>
  </div>
</template>

<style lang="less" scoped>
.NotFundApp {
  height: 80vh;
  text-align: center;
  .n-button {
    width: 70%;
    margin-bottom: 24px;
    text-align: center;
  }
}

.NotFundApp__logo-box {
  margin-bottom: 28px;
  max-width: 100%;
}

.NotFundApp__logo {
  width: 160px;
  height: 160px;
}
</style>
