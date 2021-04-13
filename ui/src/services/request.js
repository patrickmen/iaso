import axios from 'axios';
import Qs from 'qs';
import Cookie from 'js-cookie'

const xsrfHeaderName = 'Authorization'

function checkError(err) {
  return err;
}

function checkSuccess(res) {
  let msg = "";
  if ((res.status >= 200 && res.status < 400) && (res.data.code >= 200 && res.data.code < 400)) {
    switch (res.data.code) {
      case 201:
        msg = "创建成功";
        break;
      case 204:
        msg = "删除成功";
        break;
    }
    if (res.data.msg !== null && res.data.msg !== "") {
      msg = res.data.msg;
    }
    let auth = res.data.data
    if (auth.token !== null && auth.token !== "" && auth.token !== undefined) {
        Cookie.set(xsrfHeaderName, 'Bearer ' + auth.token, {expires: new Date(Date.parse(auth.expireAt))})
    }
  }
  let result = {
    data: {
      status: res.status,
      msg: msg,
    }
  };
  return result;
}

export default {
  async get(url, params) {
    console.log(url)
    return axios.get(
      url,
      params,
    ).then(checkSuccess).catch(checkError);
  },
  async post(url, data) {
    return axios.post(
      url,
      JSON.stringify(data),
    ).then(checkSuccess).catch(checkError);
  },
  async put(url, data) {
    return axios.put(
      url,
      Qs.stringify(data),
    ).then(checkSuccess).catch(checkError);
  },
  async delete(url, data) {
    return axios.delete(
      url,
      {data: Qs.stringify(data)},
    ).then(checkSuccess).catch(checkError);
  },
  async patch(url, data) {
    return axios.patch(
      url,
      data,
    ).then(checkSuccess).catch(checkError);
  },
  async jsonPost(url, data) {
    return axios.post(
      url,
      data,
    ).then((err) => {
    }).catch((err) => {
    });
  },
  async jsonPut(url, data) {
    return axios.put(
      url,
      data,
    ).then((err) => {
    }).catch((err) => {
    });
  },
  async jsonDelete(url, data) {
    return axios.delete(
      url,
      {data: data},
    ).then((err) => {
    }).catch((err) => {
    });
  },
};