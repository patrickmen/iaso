import { queryDELList, addDEL, removeDEL, updateDEL } from '@/services/api';

export default {
  namespace: 'del',

  state: {
    del: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryDELList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 2 ? removeDEL : updateDEL;
      } else {
        callback = addDEL;
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
        del: action.payload,
      };
    },
  },
};
