import { PureComponent } from 'react';
import { Icon } from 'antd';
import SelectLang from '../SelectLang';
import styles from './index.less';

export default class GlobalHeaderRight extends PureComponent {
  render() {
    const {
      theme,
    } = this.props;
    
    const loadMoreProps = {
    };
    let className = styles.right;
    if (theme === 'dark') {
      className = `${styles.right}  ${styles.dark}`;
    }
    return (
      <div className={className}>
        {/* <Spin size="small" style={{ marginLeft: 8, marginRight: 8 }} /> */}
        <Icon type="user" style={{color: "#61bcca"}} />
        <SelectLang className={styles.action} />
      </div>
    );
  }
}
