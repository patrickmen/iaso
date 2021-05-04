import { queryNewsList, addNews, removeNews, updateNews } from '@/services/api';

export default {
  namespace: 'news',

  state: {
    news: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryNewsList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 2 ? removeNews : updateNews;
      } else {
        callback = addNews;
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
        news: action.payload,
      };
    },
  },
};
