export default [
  {
    path: '/login',
    component: './Login/Login',
  },

  {
    path: '/',
    component: '../layouts/BasicLayout',
    routes: [
      // dashboard
      { path: '/', redirect: '/about-us' },
      {
        path: '/about-us',
        name: 'ABOUT-US',
        component: './AboutUs/About',
      },
      {
        path: '/technology',
        name: 'TECHNOLOGY',
        routes: [
          {
            path: '/technology/gene-to-protein-platform',
            name: 'GENE-TO-PROTEIN-PLATFORM',
            component: './Technology/TargetProtein',
          },
          {
            path: '/technology/cadd-platform',
            name: 'CADD-PLATFORM',
            component: './Technology/CADD',
          },
          {
            path: '/technology/sbdd-platform',
            name: 'SBDD-PLATFORM',
            component: './Technology/SBDD',
          },
          {
            path: '/technology/del-platform',
            name: 'DEL-PLATFORM',
            component: './Technology/DEL',
          },
          {
            component: '404',
          },
        ],
      },
      {
        path: '/pipeline',
        name: 'PIPELINE',
        component: './Pipeline/Pipeline',
      },
      // {
      //   path: '/products',
      //   name: 'PRODUCTS',
      //   routes: [
      //     {
      //       path: '/products/summary',
      //       name: 'SUMMARY',
      //       component: './Products/Summary/Summary',
      //     },
      //     {
      //       path: '/products/products',
      //       name: 'PRODUCTS',
      //       component: './Products/Products/Products',
      //     },
      //     {
      //       path: '/products/services',
      //       name: 'SERVICES',
      //       component: './Products/Services/Services',
      //     },
      //   ],
      // },
      {
        path: '/partnering',
        name: 'PARTNERING',
        component: './Partnering/Partnering',
      },
      {
        path: '/news',
        name: 'NEWS',
        component: './News/News',
      },
      {
        path: '/careers',
        name: 'CAREERS',
        component: './Careers/Careers',
      },
      {
        path: '/exception',
        name: 'EXCEPTION',
        hideInMenu: true,
        routes: [
          // exception
          {
            path: '/exception/403',
            name: 'NOT-PERMISSION',
            component: './Exception/403',
          },
          {
            path: '/exception/404',
            name: 'NOT-FOUND',
            component: './Exception/404',
          },
          {
            path: '/exception/500',
            name: 'SERVER-ERROR',
            component: './Exception/500',
          },
        ],
      },
      {
        path: '/preview/:title',
        name: 'PREVIEW',
        hideInMenu: true,
        component: 'Preview',
      },
      {
        component: '404',
      },
    ],
  }, 
];
