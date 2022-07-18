/* eslint-env node */
require('@rushstack/eslint-patch/modern-module-resolution');

module.exports = {
  root: true,
  extends: ['plugin:vue/vue3-essential', 'eslint:recommended', '@vue/eslint-config-typescript/recommended'],
  env: {
    'vue/setup-compiler-macros': true,
  },
  globals: {
    // 你的全局变量（设置为 false 表示它不允许被重新赋值）
    //
    window: false,
    $ref: false,
  },
  rules: {
    '@typescript-eslint/no-explicit-any': 0,
    '@typescript-eslint/no-empty-function': 0,
    'no-console': [
      'warn',
      {
        allow: ['warn', 'error', 'info', 'group', 'groupCollapsed', 'groupEnd', 'table'],
      },
    ],
    // 禁止使用嵌套的三元表达式
    'no-nested-ternary': 'error',
    // 调用构造函数必须带括号
    'new-parens': 'error',
    // this别名
    'consistent-this': ['error', '_this'],
    // 对象中的属性和方法使用简写
    'object-shorthand': 'error',
    // 不要省括号
    curly: 'error',
    // switch
    'default-case': 'error',
    // const
    'prefer-const': 'error',
    // 模板字符串
    'prefer-template': 'error',
  },
};
