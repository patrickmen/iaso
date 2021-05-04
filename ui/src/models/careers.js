import { queryCareersList, addCareers, removeCareers, updateCareers } from '@/services/api';

export default {
  namespace: 'careers',

  state: {
    careers: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryCareersList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 2 ? removeCareers : updateCareers;
      } else {
        callback = addCareers;
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
        careers: action.payload,
      };
    },
  },
};
