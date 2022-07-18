import { ajax_json } from '@/utils/http';
import { Md5, removeToken } from '@/utils/tools';
import type { loginType } from './api.d';

export const login = (data: loginType) => {
  const param = {
    ...data,
    Password: Md5(data.Password),
  };

  return ajax_json({
    url: '/api/public/login',
    data: param,
    method: 'post',
  });
};

export const logout = () => {
  removeToken();
  window.location.replace('/login');
};

export const getUserInfo = () => {
  return ajax_json({
    url: '/api/private/get_user_info',
    data: null,
    method: 'get',
  });
};
