import request from '@/utils/request';

export default async function queryError(code) {
  return request(`/api/${code}`);
}

export function errorAnalysis(err) {
  let msg = "";
  switch (err.status) {
    case 400:
      msg = "请求错误";
      break;
    case 401:
      msg = "您还没有登陆";
      break;
    case 403:
      msg = "您没有该项权限";
      break;
    case 404:
      msg = "资源不存在";
      break;
    case 409:
      msg = "您已重复提交";
      break;
    case 500:
      msg = "服务器发生了点意外";
      break;
  }
  if (err.data !== null && err.data !== "") {
    msg = err.data.msg;
  };
  let result = {
    status: err.status,
    msg: msg,
  };
  return result
};
