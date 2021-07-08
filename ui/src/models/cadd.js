import { queryCADDList, addCADD, removeCADD, updateCADD } from '@/services/api';

export default {
  namespace: 'cadd',

  state: {
    cadd: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryCADDList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 2 ? removeCADD : updateCADD;
      } else {
        callback = addCADD;
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
        cadd: action.payload,
      };
    },
  },
};
