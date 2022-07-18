<script setup lang="ts">
import { Icon } from '@vicons/utils';
import { defineAsyncComponent } from 'vue';

const props = defineProps({
  name: String,
});

function MyIconLoad() {
  switch (props.name) {
    case 'HomeOutlined':
      return defineAsyncComponent(() => import('@vicons/antd/HomeOutlined'));
    case 'LoginOutlined':
      return defineAsyncComponent(() => import('@vicons/antd/LoginOutlined'));
    case 'InfoCircleOutlined':
      return defineAsyncComponent(() => import('@vicons/antd/InfoCircleOutlined'));
    case 'MenuUnfoldOutlined':
      return defineAsyncComponent(() => import('@vicons/antd/MenuUnfoldOutlined'));
    case 'DownloadOutlined':
      return defineAsyncComponent(() => import('@vicons/antd/DownloadOutlined'));
    case 'EditOutlined':
      return defineAsyncComponent(() => import('@vicons/antd/EditOutlined'));
    case 'KeyOutlined':
      return defineAsyncComponent(() => import('@vicons/antd/KeyOutlined'));
    case 'DeleteOutlined':
      return defineAsyncComponent(() => import('@vicons/antd/DeleteOutlined'));
    case 'LogoutOutlined':
      return defineAsyncComponent(() => import('@vicons/antd/LogoutOutlined'));
    case 'MailOutlined':
      return defineAsyncComponent(() => import('@vicons/antd/MailOutlined'));
    case 'PictureOutlined':
      return defineAsyncComponent(() => import('@vicons/antd/PictureOutlined'));
    case 'ArrowLeftOutlined':
      return defineAsyncComponent(() => import('@vicons/antd/ArrowLeftOutlined'));
    case 'MenuFoldOutlined':
      return defineAsyncComponent(() => import('@vicons/antd/MenuFoldOutlined'));

    default:
      return defineAsyncComponent(() => import('@vicons/antd/Html5Filled'));
  }
}

const MyIcon = MyIconLoad();
</script>

<template>
  <Icon>
    <MyIcon />
  </Icon>
</template>

<style lang="less"></style>
