import React from 'react';
import {
  Router as DefaultRouter,
  Route,
  Switch,
  StaticRouter,
} from 'react-router-dom';
import dynamic from 'umi/dynamic';
import renderRoutes from 'umi/lib/renderRoutes';
import history from '@@/history';
import RendererWrapper0 from '/Users/jenson/go/src/iaso/ui/src/pages/.umi/LocaleWrapper.jsx';
import _dvaDynamic from 'dva/dynamic';

const Router = require('dva/router').routerRedux.ConnectedRouter;

const routes = [
  {
    path: '/login',
    component: __IS_BROWSER
      ? _dvaDynamic({
          component: () =>
            import(/* webpackChunkName: "p__Login__Login" */ '../Login/Login'),
          LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
            .default,
        })
      : require('../Login/Login').default,
    exact: true,
  },
  {
    path: '/',
    component: __IS_BROWSER
      ? _dvaDynamic({
          component: () =>
            import(/* webpackChunkName: "layouts__BasicLayout" */ '../../layouts/BasicLayout'),
          LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
            .default,
        })
      : require('../../layouts/BasicLayout').default,
    routes: [
      {
        path: '/',
        redirect: '/about-us',
        exact: true,
      },
      {
        path: '/about-us',
        name: 'ABOUT-US',
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__AboutUs__About" */ '../AboutUs/About'),
              LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                .default,
            })
          : require('../AboutUs/About').default,
        exact: true,
      },
      {
        path: '/technology',
        name: 'TECHNOLOGY',
        routes: [
          {
            path: '/technology/gene-to-protein-platform',
            name: 'GENE-TO-PROTEIN-PLATFORM',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  component: () =>
                    import(/* webpackChunkName: "layouts__BasicLayout" */ '../Technology/TargetProtein'),
                  LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                    .default,
                })
              : require('../Technology/TargetProtein').default,
            exact: true,
          },
          {
            path: '/technology/cadd-platform',
            name: 'CADD-PLATFORM',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  component: () =>
                    import(/* webpackChunkName: "layouts__BasicLayout" */ '../Technology/CADD'),
                  LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                    .default,
                })
              : require('../Technology/CADD').default,
            exact: true,
          },
          {
            path: '/technology/sbdd-platform',
            name: 'SBDD-PLATFORM',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  component: () =>
                    import(/* webpackChunkName: "layouts__BasicLayout" */ '../Technology/SBDD'),
                  LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                    .default,
                })
              : require('../Technology/SBDD').default,
            exact: true,
          },
          {
            path: '/technology/del-platform',
            name: 'DEL-PLATFORM',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  component: () =>
                    import(/* webpackChunkName: "layouts__BasicLayout" */ '../Technology/DEL'),
                  LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                    .default,
                })
              : require('../Technology/DEL').default,
            exact: true,
          },
          {
            component: __IS_BROWSER
              ? _dvaDynamic({
                  component: () =>
                    import(/* webpackChunkName: "layouts__BasicLayout" */ '../404'),
                  LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                    .default,
                })
              : require('../404').default,
            exact: true,
          },
          {
            component: () =>
              React.createElement(
                require('/Users/jenson/go/src/iaso/ui/node_modules/umi-build-dev/lib/plugins/404/NotFound.js')
                  .default,
                { pagesPath: 'src/pages', hasRoutesInConfig: true },
              ),
          },
        ],
      },
      {
        path: '/pipeline',
        name: 'PIPELINE',
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__Pipeline__Pipeline" */ '../Pipeline/Pipeline'),
              LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                .default,
            })
          : require('../Pipeline/Pipeline').default,
        exact: true,
      },
      {
        path: '/partnering',
        name: 'PARTNERING',
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__Partnering__Partnering" */ '../Partnering/Partnering'),
              LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                .default,
            })
          : require('../Partnering/Partnering').default,
        exact: true,
      },
      {
        path: '/news',
        name: 'NEWS',
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__News__News" */ '../News/News'),
              LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                .default,
            })
          : require('../News/News').default,
        exact: true,
      },
      {
        path: '/careers',
        name: 'CAREERS',
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__Careers__Careers" */ '../Careers/Careers'),
              LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                .default,
            })
          : require('../Careers/Careers').default,
        exact: true,
      },
      {
        path: '/exception',
        name: 'EXCEPTION',
        hideInMenu: true,
        routes: [
          {
            path: '/exception/403',
            name: 'NOT-PERMISSION',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  app: require('@tmp/dva').getApp(),
                  models: () => [
                    import(/* webpackChunkName: 'p__Exception__models__error.js' */ '/Users/jenson/go/src/iaso/ui/src/pages/Exception/models/error.js').then(
                      m => {
                        return { namespace: 'error', ...m.default };
                      },
                    ),
                  ],
                  component: () =>
                    import(/* webpackChunkName: "layouts__BasicLayout" */ '../Exception/403'),
                  LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                    .default,
                })
              : require('../Exception/403').default,
            exact: true,
          },
          {
            path: '/exception/404',
            name: 'NOT-FOUND',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  app: require('@tmp/dva').getApp(),
                  models: () => [
                    import(/* webpackChunkName: 'p__Exception__models__error.js' */ '/Users/jenson/go/src/iaso/ui/src/pages/Exception/models/error.js').then(
                      m => {
                        return { namespace: 'error', ...m.default };
                      },
                    ),
                  ],
                  component: () =>
                    import(/* webpackChunkName: "layouts__BasicLayout" */ '../Exception/404'),
                  LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                    .default,
                })
              : require('../Exception/404').default,
            exact: true,
          },
          {
            path: '/exception/500',
            name: 'SERVER-ERROR',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  app: require('@tmp/dva').getApp(),
                  models: () => [
                    import(/* webpackChunkName: 'p__Exception__models__error.js' */ '/Users/jenson/go/src/iaso/ui/src/pages/Exception/models/error.js').then(
                      m => {
                        return { namespace: 'error', ...m.default };
                      },
                    ),
                  ],
                  component: () =>
                    import(/* webpackChunkName: "layouts__BasicLayout" */ '../Exception/500'),
                  LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                    .default,
                })
              : require('../Exception/500').default,
            exact: true,
          },
          {
            component: () =>
              React.createElement(
                require('/Users/jenson/go/src/iaso/ui/node_modules/umi-build-dev/lib/plugins/404/NotFound.js')
                  .default,
                { pagesPath: 'src/pages', hasRoutesInConfig: true },
              ),
          },
        ],
      },
      {
        path: '/preview/:title',
        name: 'PREVIEW',
        hideInMenu: true,
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__Preview" */ '../Preview'),
              LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                .default,
            })
          : require('../Preview').default,
        exact: true,
      },
      {
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__404" */ '../404'),
              LoadingComponent: require('/Users/jenson/go/src/iaso/ui/src/components/PageLoading/index')
                .default,
            })
          : require('../404').default,
        exact: true,
      },
      {
        component: () =>
          React.createElement(
            require('/Users/jenson/go/src/iaso/ui/node_modules/umi-build-dev/lib/plugins/404/NotFound.js')
              .default,
            { pagesPath: 'src/pages', hasRoutesInConfig: true },
          ),
      },
    ],
  },
  {
    component: () =>
      React.createElement(
        require('/Users/jenson/go/src/iaso/ui/node_modules/umi-build-dev/lib/plugins/404/NotFound.js')
          .default,
        { pagesPath: 'src/pages', hasRoutesInConfig: true },
      ),
  },
];
window.g_routes = routes;
const plugins = require('umi/_runtimePlugin');
plugins.applyForEach('patchRoutes', { initialValue: routes });

export { routes };

export default class RouterWrapper extends React.Component {
  unListen() {}

  constructor(props) {
    super(props);

    // route change handler
    function routeChangeHandler(location, action) {
      plugins.applyForEach('onRouteChange', {
        initialValue: {
          routes,
          location,
          action,
        },
      });
    }
    this.unListen = history.listen(routeChangeHandler);
    // dva 中 history.listen 会初始执行一次
    // 这里排除掉 dva 的场景，可以避免 onRouteChange 在启用 dva 后的初始加载时被多执行一次
    const isDva =
      history.listen
        .toString()
        .indexOf('callback(history.location, history.action)') > -1;
    if (!isDva) {
      routeChangeHandler(history.location);
    }
  }

  componentWillUnmount() {
    this.unListen();
  }

  render() {
    const props = this.props || {};
    return (
      <RendererWrapper0>
        <Router history={history}>{renderRoutes(routes, props)}</Router>
      </RendererWrapper0>
    );
  }
}
