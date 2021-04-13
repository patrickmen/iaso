import request from './request';

let origin = window.location.origin;

const urlPrefix = '/api/v1';
export const submitLogin = data => request.post(`${urlPrefix}/users`, data);
 