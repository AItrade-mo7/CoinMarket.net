import type { resType, resDataType } from './utils.d';
export const res_dispose = (response: resType): resDataType => {
  const data = response.data;

  if (data.Code < 0) {
    window.$message.error(data.Msg);
    return data;
  }

  return data;
};
