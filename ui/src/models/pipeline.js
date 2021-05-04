import { queryPipelineList, addPipeline, removePipeline, updatePipeline } from '@/services/api';

export default {
  namespace: 'pipeline',

  state: {
    pipeline: [],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryPipelineList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response.data) ? response.data : [],
      });
    },
    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 2 ? removePipeline : updatePipeline;
      } else {
        callback = addPipeline;
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
        pipeline: action.payload,
      };
    },
  },
};
