<script setup lang="ts">
function UpLoadFileEncrypt() {
  const stamp = new Date().getTime();
  const time = stamp / 1000 / 30;
  const timeStamp = parseInt(`${time}`);
  const EncStr = `mo7${timeStamp}`; // mo7 + 分钟时间戳
  return window.mo7Md5(EncStr); // 然后融合MD5
}
</script>

<template>
  <h1 class="PageTitle">UpLoadFile</h1>
  <div class="PageWrapper">
    <n-upload
      action="/api/public/upload"
      :headers="{
        'File-Path': 'disk',
        'Auth-File': UpLoadFileEncrypt(),
      }"
    >
      <n-button>上传文件</n-button>
    </n-upload>
  </div>
</template>

<style lang="less" scoped></style>
