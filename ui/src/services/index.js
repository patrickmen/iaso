import axios from 'axios';
import Cookie from 'js-cookie';
import { errorAnalysis } from './error';

// // 跨域认证信息 header 名
// const xsrfHeaderName = 'Authorization'

// axios.defaults.withCredentials = false;
// axios.defaults.xsrfHeaderName= xsrfHeaderName
// axios.defaults.xsrfCookieName= xsrfHeaderName

// axios.interceptors.request.use(
//   config => {
//     config.headers = {
//       'Content-Type': 'application/x-www-form-urlencoded'
//     }
//     if (checkLogin() && checkAuthorization()) {
//       config.withCredentials = true;
//       config.headers = {
//         'Authorization': getCookie(),
//         "Content-Type": "application/json",
//       };
//     }
//     return config
//   },
//   error => {
//     return Promise.reject(error)
//   }
// );

// export function setAuthorization(auth) {
//   Cookie.set(xsrfHeaderName, 'Bearer ' + auth.token, {expires: auth.expireAt})
// };

// function getCookie() {
//   return Cookie.get(xsrfHeaderName)
// };

// function checkAuthorization() {
//   if (Cookie.get(xsrfHeaderName)) {
//     return true
//   }
//   return false
// };

// function checkLogin() {
//   if (localStorage.getItem("Login")) {
//     return true
//   }
//   return false
// };

export * from './model';
export {errorAnalysis} from './error';