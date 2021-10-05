import React, { Component } from 'react';
import { Button, Icon, Modal, Form, Row, message, Col } from 'antd';
import { connect } from 'dva';
import { formatMessage, getLocale } from 'umi/locale';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';
import HeadFeaturedPost from '@/components/Article/HeadFeaturedPost';
import PictureAlignRight from '@/components/Article/PictureAlignRight';
import PictureAlignLeft from '@/components/Article/PictureAlignLeft';
import PictureAlignJustify from '@/components/Article/PictureAlignJustify';
import CommonPictureLayout from '@/components/Form/CommonPictureLayout';
import Exception404 from '@/pages/ExceptionBeta/E404';


const headFeaturedPost = {
  title: "",
  description:
    "",
  image: 'https://cdn.pharmcafe.com/partnering-banner-01.jpg',
  imgText: 'head image description',
};

const CreateForm = Form.create()(props => {
  const { modalVisible, pictureLayout, form, handleAdd, handleUpdate, handlePreview, current, initModal, handlePictureLayoutChange } = props;
  const okHandle = () => {
    form.validateFields((err, fieldsValue) => {
      if (err) return;
      form.resetFields();
      if (current.id !== null && current.id !== undefined) {
        handleUpdate(fieldsValue, pictureLayout);
      } else {
        handleAdd(fieldsValue, pictureLayout);
      }
    });
  };
  const onPreview =() => {
    form.validateFields((err, fieldsValue) => {
      if(err) return;
      handlePreview(fieldsValue, pictureLayout);
    });
  };

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
      <CommonPictureLayout pictureLayout={pictureLayout} form={form} current={current} onChange={handlePictureLayoutChange}/>
    </Modal>
  );
});

@connect(({ academicInstitution, loading }) => ({
  academicInstitution,
  loading: loading.models.academicInstitution,
}))

@Form.create()
export default class AcademicInstitution extends Component {
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
    pictureLayout: 'justify',
    currentLang: getLocale(),
  };

  componentDidMount() {
    const { dispatch } = this.props;
    dispatch({
      type: 'academicInstitution/fetch',
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
      pictureLayout: "justify",
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
      content: JSON.parse(item.content),
      image: item.image,
    }
    this.setState({
      modalVisible: true,
      currentId: item.id,
      current: currentValue,
      pictureLayout: item.align,
    });
  };

  handlePreview = (fields, align) => {
    let data = {
        headPost: headFeaturedPost,
        content: JSON.stringify(fields.content),
        image: fields.image !== undefined ? fields.image : "",
        align: align,
    }
    window["data"] = data;
    window.open(location.origin + `/#/preview/${new Date().getTime()}`)
  };

  handleAdd = (fields, align) => {
    const { dispatch } = this.props;
    dispatch({
      type: 'academicInstitution/submit',
      payload: {
        lang: this.state.currentLang,
        content: JSON.stringify(fields.content),
        image: fields.image !== undefined ? JSON.stringify(fields.image) : "",
        align: align,
      },
    });
    message.success(formatMessage({ id: 'app.add.success' }));
    this.initModal();
  };

  handleUpdate = (fields, align) => {
    const { dispatch } = this.props;
    const { currentId } = this.state;
    dispatch({
      type: 'academicInstitution/submit',
      payload: {
        id: currentId,
        lang: this.state.currentLang,
        content: JSON.stringify(fields.content),
        image: fields.image,
        align: align,
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
      type: 'academicInstitution/submit',
      payload: {
        id: id,
        lang: this.state.currentLang,
      }
    });
    message.success(formatMessage({ id: 'app.delete.success' }));
    this.initModal();
  };

  handlePictureLayoutChange = (e) => {
    this.setState({
      pictureLayout: e.target.value,
    })
  };

  render() {
    const {
      academicInstitution: { academicInstitution = [] },
      loading,
    } = this.props;

    const { modalVisible, pictureLayout, current = {}, currentId = null } = this.state;
    const parentMethods = {
      initModal: this.initModal,
      handleAdd: this.handleAdd,
      handleUpdate: this.handleUpdate,
      handlePreview: this.handlePreview,
      handlePictureLayoutChange: this.handlePictureLayoutChange,
    };
 
    return (
      <React.Fragment>
        <CssBaseline />
        <div>
          <HeadFeaturedPost post={headFeaturedPost} />
        </div>
        <Container maxWidth="lg">
          <main>
            { academicInstitution.length ? 
              <Grid container>
                { academicInstitution.map((post) => (
                  <div key={JSON.parse(post.content).substring(0, 40)}>
                    <div>
                      {post.align == "right" ? <PictureAlignRight post={post} /> : post.align == "left" ? <PictureAlignLeft post={post} /> : <PictureAlignJustify post={post} />}
                    </div>
                    <div>
                    <Row justify="space-between">
                      <Col xs={8} sm={8} md={8} lg={8}>
                        <Button 
                          type="primary" 
                          icon="plus" 
                          onClick={this.showModal}
                        >
                          {formatMessage({ id: 'app.button.add' })}
                        </Button>
                      </Col>
                      <Col xs={8} sm={8} md={8} lg={8}>
                        <Button 
                          type="primary" 
                          icon="edit" 
                          onClick={() => this.showUpdateModal(post)}
                        >
                          {formatMessage({ id: 'app.button.edit' })}
                        </Button> 
                      </Col>
                      <Col xs={8} sm={8} md={8} lg={8}>
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
        <CreateForm {...parentMethods} modalVisible={modalVisible} pictureLayout={pictureLayout} current={current} /> 
      </React.Fragment>
    );
  }
}