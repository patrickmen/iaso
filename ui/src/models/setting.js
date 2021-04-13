import { message } from 'antd';
import defaultSettings from '../defaultSettings';

export default {
  namespace: 'setting',
  state: defaultSettings,
  reducers: {
    getSetting(state) {
      const setting = {};
      const urlParams = new URL(window.location.href);
      Object.keys(state).forEach(key => {
        if (urlParams.searchParams.has(key)) {
          const value = urlParams.searchParams.get(key);
          setting[key] = value === '1' ? true : value;
        }
      });
      return {
        ...state,
        ...setting,
      };
    },
  },
};
