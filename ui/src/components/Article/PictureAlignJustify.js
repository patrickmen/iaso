import PropTypes from 'prop-types';
import { Row, Col } from 'antd';
import classNames from 'classnames';
import ReactMarkdown from 'react-markdown';
import Grid from '@material-ui/core/Grid';
import styles from './index.less';

export default function PictureAlignJustify(props) {
  const { post } = props;
  return (
    <Grid container justify="space-between" alignItems="center">
      <ReactMarkdown
        className={classNames(styles.markdown, "markdown-body")}
        source={JSON.parse(post.content)}
        escapeHtml={true}
      />
    </Grid>
  )
}

PictureAlignJustify.propTypes = {
    post: PropTypes.object,
  };