import { ajax_json } from '@/utils/http';

export const GetBingUrl = () => {
  return ajax_json({
    url: '//file.mo7.cc/api/public/url',
    data: {},
    method: 'get',
  });
};
