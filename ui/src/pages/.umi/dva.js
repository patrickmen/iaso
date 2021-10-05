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
  
  app.model({ namespace: 'aboutus', ...(require('/Users/patrick/go/src/iaso/ui/src/models/aboutus.js').default) });
app.model({ namespace: 'academicInstitution', ...(require('/Users/patrick/go/src/iaso/ui/src/models/academicInstitution.js').default) });
app.model({ namespace: 'biomarker', ...(require('/Users/patrick/go/src/iaso/ui/src/models/biomarker.js').default) });
app.model({ namespace: 'biotechCompany', ...(require('/Users/patrick/go/src/iaso/ui/src/models/biotechCompany.js').default) });
app.model({ namespace: 'careers', ...(require('/Users/patrick/go/src/iaso/ui/src/models/careers.js').default) });
app.model({ namespace: 'global', ...(require('/Users/patrick/go/src/iaso/ui/src/models/global.js').default) });
app.model({ namespace: 'menu', ...(require('/Users/patrick/go/src/iaso/ui/src/models/menu.js').default) });
app.model({ namespace: 'news', ...(require('/Users/patrick/go/src/iaso/ui/src/models/news.js').default) });
app.model({ namespace: 'pipeline', ...(require('/Users/patrick/go/src/iaso/ui/src/models/pipeline.js').default) });
app.model({ namespace: 'products', ...(require('/Users/patrick/go/src/iaso/ui/src/models/products.js').default) });
app.model({ namespace: 'sbdd', ...(require('/Users/patrick/go/src/iaso/ui/src/models/sbdd.js').default) });
app.model({ namespace: 'setting', ...(require('/Users/patrick/go/src/iaso/ui/src/models/setting.js').default) });
app.model({ namespace: 'targetValidation', ...(require('/Users/patrick/go/src/iaso/ui/src/models/targetValidation.js').default) });
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
