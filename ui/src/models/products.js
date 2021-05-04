import { queryProductsList, addProduct, removeProduct, updateProduct } from '@/services/api';

export default {
  namespace: 'products',

  state: {
    products: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryProductsList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 2 ? removeProduct : updateProduct;
      } else {
        callback = addProduct;
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
        products: action.payload,
      };
    },
  },
};
