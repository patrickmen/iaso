import { queryBiotechCompanyList, addBiotechCompany, removeBiotechCompany, updateBiotechCompany } from '@/services/api';

export default {
  namespace: 'biotechCompany',

  state: {
    biotechCompany: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryBiotechCompanyList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 2 ? removeBiotechCompany : updateBiotechCompany;
      } else {
        callback = addBiotechCompany;
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
        biotechCompany: action.payload,
      };
    },
  },
};
