import { queryIndustrialInstitutionList, addIndustrialInstitution, removeIndustrialInstitution, updateIndustrialInstitution } from '@/services/api';

export default {
  namespace: 'industrialInstitution',

  state: {
    industrialInstitution: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryIndustrialInstitutionList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 2 ? removeIndustrialInstitution : updateIndustrialInstitution;
      } else {
        callback = addIndustrialInstitution;
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
        industrialInstitution: action.payload,
      };
    },
  },
};
