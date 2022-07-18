<script setup lang="ts">
import { TopBarStore, UserInfoStore, PingDataStore } from '@/store';
import { RouterLink } from 'vue-router';
import XIcon from '@/lib/XIcon.vue';
import PWAInstall from '@/lib/PWAInstall.vue';

const ViteConst = window.ViteConst;

const closeDrawer = () => {
  TopBarStore.close();
};
const showDrawer = () => {
  TopBarStore.open();
};
</script>

<template>
  <div class="TopBar">
    <n-avatar class="ShowBtn" @click="showDrawer" round :size="36" :src="UserInfoStore.value.Avatar" />
    <n-drawer :auto-focus="false" width="38%" :show="TopBarStore.status" :on-mask-click="closeDrawer">
      <n-drawer-content class="TopBarDrawer">
        <template #header>
          <RouterLink to="/personal" class="UserInfo">
            <n-avatar round :size="50" :src="UserInfoStore.value.Avatar" />
            <div className="TopBar__NickName">{{ UserInfoStore.value.NickName }}</div>
          </RouterLink>
        </template>
        <div className="TopBar__ul">
          <div className="TopBar__item">
            <RouterLink to="/" class="item_btn">
              <n-button text size="tiny">
                <template #icon>
                  <XIcon name="HomeOutlined" />
                </template>
                Home
              </n-button>
            </RouterLink>
          </div>
        </div>

        <template #footer>
          <div class="TopBar__pendant">
            <PWAInstall className="TopBar__instPWA">
              <n-button type="primary" size="small">
                <template #icon>
                  <XIcon name="DownloadOutlined" />
                </template>
                安装此程序
              </n-button>
            </PWAInstall>
          </div>

          <div class="TopBar__footer">
            <n-button text size="large" @click="closeDrawer">
              <template #icon>
                <XIcon name="MenuUnfoldOutlined" />
              </template>
            </n-button>
            <div class="TopBar__version">
              <span> 当前版本: {{ ViteConst.AppVersion }} </span>
              <span v-if="ViteConst.AppVersion !== PingDataStore.value.AppInfo.version">
                最新版本: {{ PingDataStore.value.AppInfo.version }}
              </span>
            </div>
          </div>
        </template>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<style lang="less" scoped>
.TopBar {
  position: fixed;
  top: 0;
  right: 0;
  padding: 10px;
  z-index: 50;
}
.ShowBtn {
  cursor: pointer;
  display: block;
  animation: Rotate infinite 7s linear;
}
</style>

<style lang="less">
@import '@/config/constant.less';

.TopBarDrawer {
  .n-drawer-header__main {
    width: 100%;
  }

  .TopBar__footer {
    display: flex;
    width: 100%;
    align-items: center;
    justify-content: space-between;
    .n-button {
      color: #f0b90b;
      &:hover {
        color: @mainColor;
      }
    }
  }

  .TopBar__pendant {
    position: absolute;
    left: 0;
    bottom: 100%;
    width: 100%;
    margin-bottom: 16px;
    text-align: center;
  }

  .TopBar__version {
    color: #d9d9d9;
    font-size: 12px;
    cursor: pointer;
    span {
      display: block;
    }
  }

  [aria-current='page'].item_btn {
    .n-button {
      color: @mainColor;
    }
  }

  .item_btn {
    display: block;
    font-size: 14px;
    .n-button {
      width: 100%;
      justify-content: flex-start;
      padding: 4px 15px;
    }
  }

  [aria-current='page'].UserInfo {
    .TopBar__NickName {
      color: @mainColor;
    }
  }

  .UserInfo {
    text-align: center;
    display: block;
    .TopBar__NickName {
      margin-top: 8px;
      font-size: 14px;
    }
  }
}

.n-drawer .TopBarDrawer {
  .n-drawer-body-content-wrapper {
    padding: 1px;
  }
  .n-drawer-header {
    padding-top: 20px;
    padding-left: 0;
    padding-right: 0;
  }
  .n-drawer-footer {
    padding: 16px 15px;
    position: relative;
  }
}
</style>
