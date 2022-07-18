import { ajax_json } from '@/utils/http';

export const Ping = (data?: any) => {
  return ajax_json({
    url: '/api/ping',
    data,
    method: 'post',
  });
};
