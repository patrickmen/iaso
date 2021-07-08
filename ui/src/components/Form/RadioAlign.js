import PropTypes from 'prop-types';
import { Radio } from 'antd';
import { formatMessage } from 'umi/locale';
import FormatAlignJustifyIcon from '@material-ui/icons/FormatAlignJustify';
import FormatAlignLeftIcon from '@material-ui/icons/FormatAlignLeft';
import FormatAlignRightIcon from '@material-ui/icons/FormatAlignRight';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles(() => ({
    root: {
      overflow: 'hidden',
      verticalAlign: 'middle',
    }
  }));

export default function RadioAlign(props) {
  const { layout, onChange } = props;
  const classes = useStyles();
  return (
    <div>
      <span>{formatMessage({ id: 'app.dialog.pictureLayout' })}:&nbsp;&nbsp;</span>
      <Radio.Group defaultValue={layout} size="large" buttonStyle="solid" onChange={onChange}>
        <Radio.Button value="justify">
          <FormatAlignJustifyIcon className={classes.root} />
        </Radio.Button>
        <Radio.Button value="left">
          <FormatAlignLeftIcon className={classes.root} />
        </Radio.Button>
        <Radio.Button value="right">
          <FormatAlignRightIcon className={classes.root} />
        </Radio.Button>
      </Radio.Group>
    </div>
  )
}

RadioAlign.propTypes = {
    post: PropTypes.object,
  };