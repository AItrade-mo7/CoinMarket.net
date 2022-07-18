<script setup lang="ts">
import { Logo } from '@/config/constant';
import XIcon from '@/lib/XIcon.vue';
import { logout } from '@/api/Account';

const UserInfo = {
  Avatar: '//file.mo7.cc/static/lxh_png/lxh.png',
  NickName: '墨七',
  Email: 'meichangliang@mo7.cc',
};
</script>

<template>
  <h1 class="PageTitle">Personal</h1>
  <div class="PageWrapper">
    <div className="Personal__info">
      <n-avatar round :size="70" :src="UserInfo.Avatar" :fallback-src="Logo" />
      <div className="Personal__NickName">{{ UserInfo.NickName }}</div>
      <div className="Personal__Email">{{ UserInfo.Email }}</div>
    </div>

    <div className="Personal__link">
      <RouterLink to="/edit_user_info">
        <n-button type="primary">
          <template #icon>
            <XIcon name="EditOutlined" />
          </template>
          编辑资料
        </n-button>
      </RouterLink>
      <br />
      <RouterLink to="/edit_password">
        <n-button>
          <template #icon>
            <XIcon name="KeyOutlined" />
          </template>
          修改密码
        </n-button>
      </RouterLink>
      <br />
      <RouterLink to="/about/duty">
        <n-button type="error">
          <template #icon>
            <XIcon name="DeleteOutlined" />
          </template>
          注销账户
        </n-button>
      </RouterLink>
      <br />
      <n-button type="warning" ghost @click="logout">
        <template #icon>
          <XIcon name="LoginOutlined" />
        </template>
        退出登录
      </n-button>
      <br />
    </div>
  </div>
</template>

<style lang="less" scoped>
.PageWrapper {
  text-align: center;
  .n-button {
    margin: 0 auto;
    margin-bottom: 16px;
    max-width: 317px;
    width: 100%;
    text-align: center;
  }
}

.Personal__info {
  margin-bottom: 24px;
  .n-avatar {
    margin-bottom: 6px;
    animation: Rotate infinite 36s linear;
  }
  .Personal__NickName {
    font-size: 16px;
    color: #333;
    margin-bottom: 4px;
  }
  .Personal__Email {
    font-size: 12px;
    color: #999;
  }
}
</style>
