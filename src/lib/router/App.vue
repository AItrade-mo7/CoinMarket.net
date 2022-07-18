<script setup lang="ts">
import { RouterView } from 'vue-router';
import { NConfigProvider, zhCN, dateZhCN } from 'naive-ui';
import type { GlobalThemeOverrides } from 'naive-ui';
import { defineAsyncComponent } from 'vue';
import { Ping } from '@/api/Ping';
import { setToken, removeToken } from '@/utils/tools';
import LoadingView from './LoadingView.vue';

import { PingDataStore, LoadingStore } from '@/store';
const TopBar = defineAsyncComponent(() => import('@/lib/TopBar.vue'));

const fetchPing = async () => {
  const res = await Ping();
  if (res.Code > 0) {
    const Token = res.Data.Token;
    PingDataStore.update(res.Data);
    if (Token) {
      setToken(Token);
    }
  } else {
    removeToken();
  }
  LoadingStore.close();
};
fetchPing();

const themeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: '#9eb9fd',
  },
};
</script>

<template>
  <NConfigProvider :theme-overrides="themeOverrides" :locale="zhCN" :date-locale="dateZhCN">
    <n-message-provider>
      <LoadingView />
      <TopBar />
      <RouterView />
    </n-message-provider>
  </NConfigProvider>
</template>
