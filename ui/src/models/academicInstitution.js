import { queryAcademicInstitutionList, addAcademicInstitution, removeAcademicInstitution, updateAcademicInstitution } from '@/services/api';

export default {
  namespace: 'academicInstitution',

  state: {
    academicInstitution: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryAcademicInstitutionList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 2 ? removeAcademicInstitution : updateAcademicInstitution;
      } else {
        callback = addAcademicInstitution;
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
        academicInstitution: action.payload,
      };
    },
  },
};
