import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { VitePWA } from 'vite-plugin-pwa';
import path from 'path';

import eslintPlugin from 'vite-plugin-eslint';
import Inspect from 'vite-plugin-inspect';

import Components from 'unplugin-vue-components/vite';
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers';

const PwaConfig = {
  workbox: {
    sourcemap: true,
  },
  manifest: {
    name: 'VueDemo',
    short_name: 'VueDemo',
    theme_color: '#F0B90B',
    description: 'VueDemo, golang serve',
    lang: 'zh',
    icons: [
      {
        src: '//file.mo7.cc/static/lxh_png/2.png',
        sizes: '192x192',
        type: 'image/png',
      },
      {
        src: '//file.mo7.cc/static/lxh_png/2.png',
        sizes: '512x512',
        type: 'image/png',
      },
      {
        src: '//file.mo7.cc/static/lxh_png/2.png',
        sizes: '192x192',
        type: 'image/png',
        purpose: 'maskable',
      },
      {
        src: '//file.mo7.cc/static/lxh_png/2.png',
        sizes: '512x512',
        type: 'image/png',
        purpose: 'maskable',
      },
    ],
    start_url: './?mode=pwa',
    display: 'standalone',
    background_color: '#333333',
  },
};

import AppPackage from './package.json';

// const ProxyUrl = 'https://file.mo7.cc';
const ProxyUrl = `http://127.0.0.1:${AppPackage.Port}`;

// https://vitejs.dev/config/
const pathSrc = path.resolve(__dirname, 'src');
export default defineConfig({
  resolve: {
    alias: {
      '@': pathSrc,
    },
  },
  plugins: [
    vue({
      reactivityTransform: true,
    }),
    Components({
      resolvers: [NaiveUiResolver()],
      dts: path.resolve(pathSrc, 'components.d.ts'),
    }),
    VitePWA(PwaConfig),
    eslintPlugin(),
    Inspect(),
  ],
  define: {
    ViteConst: JSON.stringify({
      AppVersion: AppPackage.version,
      AppName: AppPackage.name,
      ProxyUrl,
    }),
  },
  server: {
    host: true,
    port: AppPackage.Port + 1,
    strictPort: true, // 端口已被占用则会直接退出
    proxy: {
      '/api': {
        // 设置你调用的接口域名和端口号 别忘了加http
        target: ProxyUrl,
        changeOrigin: true, // 允许跨域
      },
    },
  },
});
