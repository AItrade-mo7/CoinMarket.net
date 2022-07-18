<script setup lang="ts">
import { GetBingUrl } from '@/api/bing';

const imgString1 = `<img src="http://file.mo7.cc/api/public/bz" />`;
const imgString2 = `<img src="http://file.mo7.cc/api/public/bz?idx=3" />`;

let bingRes = $ref({});

const getBing = () => {
  GetBingUrl().then((res) => {
    bingRes = res;
  });
};
</script>

<template>
  <h1 class="PageTitle">DailyCover</h1>
  <div class="PageWrapper DailyCover">
    <h2>每日一图</h2>
    <div>golang 实现获取 每日一图，每 8 小时自动更新一次 , 跟 bing 每日一图保持一致</div>
    <h3>使用方式：</h3>
    <p>{{ imgString1 }}</p>
    <p>{{ imgString2 }}</p>

    <br />

    <div>
      参数： idx=0 为随即返回一张图片,1-8则返回对应的壁纸,例如：
      <a href="https://file.mo7.cc/api/public/bz?idx=1" target="_blank"> https://file.mo7.cc/api/public/bz?idx=1 </a>
      固定返回当日的 cn.bing.com 壁纸
    </div>
    <br />
    <div>
      效果图：
      <img src="//file.mo7.cc/api/public/bz?idx=1" />
    </div>

    <h3>接口请求示例: https://file.mo7.cc/api/public/url</h3>
    <n-button @click="getBing" type="primary"> 发起 GET 请求 </n-button>
    <p>
      {{ JSON.stringify(bingRes) }}
    </p>
  </div>
</template>

<style lang="less" scoped>
.DailyCover {
  word-break: break-all;
  font-size: 14px;
  color: #333;
  img {
    display: block;
    max-width: 92%;
    margin: 0 auto;
  }
}
</style>
