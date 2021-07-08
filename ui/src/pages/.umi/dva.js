import dva from 'dva';
import { Component } from 'react';
import createLoading from 'dva-loading';
import history from '@tmp/history';

let app = null;

export function _onCreate() {
  const plugins = require('umi/_runtimePlugin');
  const runtimeDva = plugins.mergeConfig('dva');
  app = dva({
    history,
    
    ...(runtimeDva.config || {}),
    ...(window.g_useSSR ? { initialState: window.g_initialData } : {}),
  });
  
  app.use(createLoading());
  (runtimeDva.plugins || []).forEach(plugin => {
    app.use(plugin);
  });
  
  app.model({ namespace: 'aboutus', ...(require('/Users/jenson/go/src/iaso/ui/src/models/aboutus.js').default) });
app.model({ namespace: 'cadd', ...(require('/Users/jenson/go/src/iaso/ui/src/models/cadd.js').default) });
app.model({ namespace: 'careers', ...(require('/Users/jenson/go/src/iaso/ui/src/models/careers.js').default) });
app.model({ namespace: 'del', ...(require('/Users/jenson/go/src/iaso/ui/src/models/del.js').default) });
app.model({ namespace: 'global', ...(require('/Users/jenson/go/src/iaso/ui/src/models/global.js').default) });
app.model({ namespace: 'menu', ...(require('/Users/jenson/go/src/iaso/ui/src/models/menu.js').default) });
app.model({ namespace: 'news', ...(require('/Users/jenson/go/src/iaso/ui/src/models/news.js').default) });
app.model({ namespace: 'partnering', ...(require('/Users/jenson/go/src/iaso/ui/src/models/partnering.js').default) });
app.model({ namespace: 'pipeline', ...(require('/Users/jenson/go/src/iaso/ui/src/models/pipeline.js').default) });
app.model({ namespace: 'products', ...(require('/Users/jenson/go/src/iaso/ui/src/models/products.js').default) });
app.model({ namespace: 'sbdd', ...(require('/Users/jenson/go/src/iaso/ui/src/models/sbdd.js').default) });
app.model({ namespace: 'setting', ...(require('/Users/jenson/go/src/iaso/ui/src/models/setting.js').default) });
app.model({ namespace: 'targetProtein', ...(require('/Users/jenson/go/src/iaso/ui/src/models/targetProtein.js').default) });
  return app;
}

export function getApp() {
  return app;
}

export class _DvaContainer extends Component {
  render() {
    const app = getApp();
    app.router(() => this.props.children);
    return app.start()();
  }
}
