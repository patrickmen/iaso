import fetch from 'dva/fetch';

export const dva = {
  config: {
    onError(err) {
      err.preventDefault();
    },
  },
};

export function render(oldRender) {
  oldRender();
}
