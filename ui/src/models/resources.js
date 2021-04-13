import { queryResourcesList, addResource, removeResource, updateResource } from '@/services/api';

export default {
  namespace: 'resources',

  state: {
    resources: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryResourcesList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 1 ? removeResource : updateResource;
      } else {
        callback = addResource;
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
        resources: action.payload,
      };
    },
  },
};
