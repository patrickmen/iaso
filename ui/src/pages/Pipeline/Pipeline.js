import React, { Component } from 'react';
import { Button, Icon, Input, Modal, Form, Row, message, Col } from 'antd';
import { connect } from 'dva';
import { formatMessage, getLocale } from 'umi/locale';
import 'github-markdown-css';
import ReactMarkdown from 'react-markdown';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';
import HeadFeaturedPost from '@/components/Article/HeadFeaturedPost';
import Exception404 from '@/pages/ExceptionBeta/E404';

const { TextArea } = Input;

const CreateForm = Form.create()(props => {
  const { modalVisible, form, handleAdd, handleUpdate, handlePreview, current, initModal } = props;
  const okHandle = () => {
    form.validateFields((err, fieldsValue) => {
      if (err) return;
      form.resetFields();
      if (current.id !== null && current.id !== undefined) {
        handleUpdate(fieldsValue);
      } else {
        handleAdd(fieldsValue);
      }
    });
  };
  const onPreview =() => {
    form.validateFields((err, fieldsValue) => {
      if(err) return;
      handlePreview(fieldsValue);
    });
  }
  return (
    <Modal
      destroyOnClose
      title={formatMessage({ id: 'app.dialog.title' })}
      visible={modalVisible}
      footer={[
          <Button onClick={() => initModal()}>{formatMessage({ id: 'app.dialog.cancel' })}</Button>,
          <Button onClick={onPreview}>{formatMessage({ id: 'app.dialog.preview' })}</Button>,
          <Button onClick={okHandle}>{formatMessage({ id: 'app.dialog.submit' })}</Button>,
      ]}
      onOk={okHandle}
      onCancel={() => initModal()}
    >
      <Form.Item labelCol={{ span: 5 }} wrapperCol={{ span: 18 }} label="markdown">
        {form.getFieldDecorator('content', {
          rules: [{ required: true, message: formatMessage({ id: 'app.characters.limit' }), min: 5 }],
          initialValue: current.content,
        })(<TextArea rows={10} />)}
      </Form.Item>
    </Modal>
  );
});

@connect(({ pipeline, loading }) => ({
  pipeline,
  loading: loading.models.pipeline,
}))

@Form.create()
export default class Pipeline extends Component {
  // state = {
  //   markdown: [],
  // }
  // componentDidMount() {
  //   data.map((item) => {
  //     fetch(item)
  //       .then(res => res.text())
  //       .then(text => this.setState({markdown: [...this.state.markdown, text]}));
  //   })    
  // };
  state = {
    modalVisible: false,
    currentLang: getLocale(),
  };

  componentDidMount() {
    const { dispatch } = this.props;
    dispatch({
      type: 'pipeline/fetch',
      payload: {
        lang: this.state.currentLang,
      },
    });
  };

  initModal = () => {
    this.setState({
      modalVisible: false,
      current: {},
      currentId: null,
    });
  };

  showModal = () => {
    this.setState({
      modalVisible: true,
      current: undefined,
    });
  };

  showUpdateModal = item => {
    let currentValue = {
      id: item.id,
      content: JSON.parse(item.content)
    }
    this.setState({
      modalVisible: true,
      currentId: item.id,
      current: currentValue,
    });
  };

  handlePreview = fields => {
    let data = {
        content: JSON.stringify(fields.content),
    }
    window["data"] = data;
    window.open(location.origin + `/#/preview/${new Date().getTime()}`)
  };

  handleAdd = fields => {
    const { dispatch } = this.props;
    dispatch({
      type: 'pipeline/submit',
      payload: {
        content: JSON.stringify(fields.content),
        lang: this.state.currentLang,
      },
    });
    message.success(formatMessage({ id: 'app.add.success' }));
    this.initModal();
  };

  handleUpdate = fields => {
    const { dispatch } = this.props;
    const { currentId } = this.state;
    dispatch({
      type: 'pipeline/submit',
      payload: {
        id: currentId,
        content: JSON.stringify(fields.content),
        lang: this.state.currentLang,
      }
    });
    message.success(formatMessage({ id: 'app.update.success' }));
    this.initModal();
  };

  deleteConfirm = item => {
    Modal.confirm({
      title: formatMessage({ id: 'app.delete-confirm-title' }),
      content: formatMessage({ id: 'app.delete-confirm-content' }),
      okText: formatMessage({ id: 'app.dialog.ok' }),
      cancelText: formatMessage({ id: 'app.dialog.cancel' }),
      onOk: () => this.handleDelete(item.id),
    })
  };

  handleDelete = id => {
    const { dispatch } = this.props;
    dispatch({
      type: 'pipeline/submit',
      payload: {
        id: id,
        lang: this.state.currentLang,
      }
    });
    message.success(formatMessage({ id: 'app.update.success' }));
    this.initModal();
  };

  render() {
    // const { markdown } = this.state;
    const {
      pipeline: { pipeline = [] },
      loading,
    } = this.props;

    const { modalVisible, current = {}, currentId = null } = this.state;
    const parentMethods = {
      initModal: this.initModal,
      handleAdd: this.handleAdd,
      handleUpdate: this.handleUpdate,
      handlePreview: this.handlePreview,
    };
    const headFeaturedPost = {
      title: 'MEET LOFLY BIO',
      description:
        "A Biopharmaceutical company, devoted to help the general public and investors better.",
      image: 'https://source.unsplash.com/random',
      imgText: 'head image description',
    };
 
    return (
      <React.Fragment>
        <CssBaseline />
        <div>
          <HeadFeaturedPost post={headFeaturedPost} />
        </div>
        <Container maxWidth="lg">
          <main>
            { pipeline.length ? 
              <Grid container>
                { pipeline.map((post) => (
                  <div key={JSON.parse(post.content).substring(0, 40)}>
                    <ReactMarkdown
                      className="markdown-body"
                      source={JSON.parse(post.content)}
                      // key={JSON.parse(post.content).substring(0, 40)}
                      escapeHtml={true}
                    />
                    <br />
                    <br />
                    <Row type="flex" justify="start">
                      <Col span={3}>
                        <Button 
                          type="primary" 
                          icon="plus" 
                          onClick={this.showModal}
                        >
                          {formatMessage({ id: 'app.button.add' })}
                        </Button>
                      </Col>
                      <Col span={3}>
                        <Button 
                          type="primary" 
                          icon="edit" 
                          onClick={() => this.showUpdateModal(post)}
                        >
                          {formatMessage({ id: 'app.button.edit' })}
                        </Button> 
                      </Col>
                      <Col span={3}>
                        <Button 
                          type="primary" 
                          icon="close-circle" 
                          onClick={() => this.deleteConfirm(post)}
                        >
                          {formatMessage({ id: 'app.button.delete' })}
                        </Button> 
                      </Col>
                    </Row>
                  </div> 
                  
                ))}
              </Grid> : 
              <div>
                <br/>
                <Button 
                  type="primary" 
                  icon="add" 
                  size="large"
                  onClick={this.showModal}
                >
                  {formatMessage({ id: 'app.button.add-new' })}
                </Button> 
              </div> 
            }
          </main>
        </Container>
        <CreateForm {...parentMethods} modalVisible={modalVisible} current={current} /> 
      </React.Fragment>
    );
  }
}