import { queryPartneringList, addPartnering, removePartnering, updatePartnering } from '@/services/api';

export default {
  namespace: 'partnering',

  state: {
    partnering: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryPartneringList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 1 ? removePartnering : updatePartnering;
      } else {
        callback = addPartnering;
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
        partnering: action.payload,
      };
    },
  },
};
