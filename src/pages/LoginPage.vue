<script setup lang="ts">
import { cloneDeep, setToken } from '@/utils/tools';
import { useRouter } from 'vue-router';
import { Logo } from '@/config/constant';
import XIcon from '@/lib/XIcon.vue';
import { login } from '@/api/Account';

const $router = useRouter();

let SubmitStatus: boolean = $ref(false);

const formValue = $ref({
  Email: '',
  Password: '',
});

const Submit = async () => {
  SubmitStatus = true;
  const res = await login({
    ...cloneDeep(formValue),
  });
  SubmitStatus = false;

  if (res.Code > 0) {
    await setToken(res.Data.Token);
    $router.replace('/');
  }
};
</script>

<template>
  <h1 class="PageTitle">Login</h1>
  <div class="PageWrapper">
    <div className="Login__logo-box">
      <img className="Login__logo" :src="Logo" alt="" />
    </div>
    <h2 className="Login__title">
      <div>登录</div>
    </h2>

    <n-form ref="loginForm" :model="formValue" size="small" class="myForm">
      <n-form-item path="age" class="myForm__item">
        <n-input
          name="Email"
          v-model:value="formValue.Email"
          :inputProps="{ autocomplete: 'password' }"
          placeholder="请输入邮箱地址"
        >
          <template #prefix> <XIcon name="MailOutlined" /> </template>
        </n-input>
      </n-form-item>
      <n-form-item path="password" class="myForm__item">
        <n-input
          v-model:value="formValue.Password"
          type="password"
          name="Password"
          show-password-on="mousedown"
          placeholder="请输入密码"
          :inputProps="{ autocomplete: 'password' }"
        ></n-input>
      </n-form-item>

      <n-form-item path="password" class="myForm__item">
        <n-button :disabled="SubmitStatus" type="primary" @click="Submit"> 登录 </n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<style lang="less" scoped>
.Login__logo-box {
  padding-bottom: 36px;
}

.Login__logo {
  display: block;
  border-radius: 100px;
  overflow: hidden;
  width: 100px;
  height: 100px;
  margin: 0 auto;
}

.Login__title {
  margin: 0;
  text-align: center;
}

.Login__title-str {
  display: block;
}
.Login__forget {
  text-align: center;
}
</style>
