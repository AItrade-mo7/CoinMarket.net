<script setup lang="ts">
import { routes } from '@/lib/router/index';
import { RouterLink } from 'vue-router';

let aboutListPage = [];

for (const item of routes) {
  if (item.path == '/about') {
    aboutListPage = item.children;
  }
}
</script>

<template>
  <div class="list">
    <template v-for="item in aboutListPage" :key="item.name">
      <div v-if="item.path">
        <RouterLink :to="'/about/' + item.path">
          <n-button class="a_link" text> {{ item.description }} </n-button>
        </RouterLink>
      </div>
    </template>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';
.list {
  margin-top: 20px;
}

.a_link {
  margin: 10px 6px;
  color: @mainColor;
}
</style>
