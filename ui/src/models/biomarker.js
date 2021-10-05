import { queryBiomarkerList, addBiomarker, removeBiomarker, updateBiomarker } from '@/services/api';

export default {
  namespace: 'biomarker',

  state: {
    biomarker: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryBiomarkerList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 2 ? removeBiomarker : updateBiomarker;
      } else {
        callback = addBiomarker;
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
        biomarker: action.payload,
      };
    },
  },
};
