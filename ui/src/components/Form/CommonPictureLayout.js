import PropTypes from 'prop-types';
import { Row, message, Col, Input, Form } from 'antd';
import { formatMessage } from 'umi/locale';
import RadioAlign from './RadioAlign';

const { TextArea } = Input;

export default function CommonPictureLayout(props) {
  const { pictureLayout, form, current, onChange } = props
  const pictureItemLayout = pictureLayout === "left" ? {
      span: 6,
      rows: 4,
    } : pictureLayout === "right" ? {
      span: 6,
      push: 18,
      rows: 4,
    } : null;
  const contentItemLayout = pictureLayout === "left" ? {
    span: 16,
    offset: 2,
    rows: 10,
  } : pictureLayout === "right"  ? {
    span: 16,
    pull: 6,
    rows: 10,
  } : null;

  return (
    <div>
      <RadioAlign onChange={onChange} layout={pictureLayout}/>
      { pictureLayout != 'justify' ? 
        <Form layout={pictureLayout}>
          <Row >
            <Col {...pictureItemLayout}>
              <Form.Item label={formatMessage({ id: 'app.dialog.image' })} >
                {form.getFieldDecorator('image', {
                  initialValue: current.image,
                })(<TextArea rows={pictureItemLayout.rows} />)}
              </Form.Item>
            </Col>
            <Col {...contentItemLayout}>
              <Form.Item label={formatMessage({ id: 'app.dialog.content' })}>
                {form.getFieldDecorator('content', {
                  rules: [{ required: true, message: formatMessage({ id: 'app.characters.limit' }), min: 5 }],
                  initialValue: current.content,
                })(<TextArea rows={contentItemLayout.rows} />)}
              </Form.Item>
            </Col>
          </Row>
        </Form> :
        <Form>
          <Form.Item label={formatMessage({ id: 'app.dialog.content' })}>
            {form.getFieldDecorator('content', {
              rules: [{ required: true, message: formatMessage({ id: 'app.characters.limit' }), min: 5 }],
              initialValue: current.content,
            })(<TextArea rows={10} />)}
          </Form.Item>
        </Form>
      }  
    </div>
  )      
}

CommonPictureLayout.propTypes = {
    post: PropTypes.object,
  };