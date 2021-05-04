import { querySBDDList, addSBDD, removeSBDD, updateSBDD } from '@/services/api';

export default {
  namespace: 'sbdd',

  state: {
    sbdd: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(querySBDDList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 2 ? removeSBDD : updateSBDD;
      } else {
        callback = addSBDD;
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
        sbdd: action.payload,
      };
    },
  },
};
