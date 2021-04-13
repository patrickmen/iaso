import React, { Fragment } from 'react';
import { Layout, Icon } from 'antd';
import { formatMessage } from 'umi/locale';
import GlobalFooter from '@/components/GlobalFooter';

const { Footer } = Layout;
const FooterView = () => (
  <Footer style={{ padding: 0 }}>
    <GlobalFooter
      links={[
        {
          key: 'app.footer.description',
          title: formatMessage({ id: 'app.footer.description' }),
          href: 'https://pro.ant.design',
          blankTarget: true,
        },
        {
          key: 'github',
          title: <Icon type="github" />,
          href: 'https://github.com/patrickmen/iaso',
          blankTarget: true,
        },
        {
          key: 'IASO',
          title: 'IASO',
          href: 'https://github.com/patrickmen/iaso',
          blankTarget: true,
        },
      ]}
      copyright={
        <Fragment>
          Copyright <Icon type="copyright" /> 2021 Patrick
        </Fragment>
      }
    />
  </Footer>
);
export default FooterView;
