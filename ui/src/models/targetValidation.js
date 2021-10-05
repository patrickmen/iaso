import { queryTargetValidationList, addTargetValidation, removeTargetValidation, updateTargetValidation } from '@/services/api';

export default {
  namespace: 'targetValidation',

  state: {
    targetValidation: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryTargetValidationList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 2 ? removeTargetValidation : updateTargetValidation;
      } else {
        callback = addTargetValidation;
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
        targetValidation: action.payload,
      };
    },
  },
};
