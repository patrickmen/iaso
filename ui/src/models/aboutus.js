import { queryAboutUsList, addAboutUs, removeAboutUs, updateAboutUs } from '@/services/api';

export default {
  namespace: 'aboutus',

  state: {
    aboutus: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryAboutUsList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 2 ? removeAboutUs : updateAboutUs;
      } else {
        callback = addAboutUs;
      }
      const response = yield call(callback, payload); // post
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
  },

  reducers: {
    queryList(state, action) {
      return {
        ...state,
        aboutus: action.payload,
      };
    },
  },
};
