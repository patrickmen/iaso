import { queryTargetProteinList, addTargetProtein, removeTargetProtein, updateTargetProtein } from '@/services/api';

export default {
  namespace: 'targetProtein',

  state: {
    targetProtein: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryTargetProteinList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 1 ? removeTargetProtein : updateTargetProtein;
      } else {
        callback = addTargetProtein;
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
        targetProtein: action.payload,
      };
    },
  },
};
