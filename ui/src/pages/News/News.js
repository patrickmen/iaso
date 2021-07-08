import React, { Component, Fragment } from 'react';
import { connect } from 'dva';
import moment from 'moment';
import { formatMessage, getLocale } from 'umi/locale';
import CssBaseline from '@material-ui/core/CssBaseline';
import Ellipsis from '@/components/Ellipsis';
import { Card, List, Button, Icon, Input, Modal, Form, message } from 'antd';
import HeadFeaturedPost from '@/components/Article/HeadFeaturedPost';
import CommonPictureLayout from '@/components/Form/CommonPictureLayout';
import styles from './News.less';

const { TextArea } = Input;

const headFeaturedPost = {
  title: 'MEET LOFLY BIO',
  description:
    "A Biopharmaceutical company, devoted to help the general public and investors better.",
  image: 'https://cdn.pharmcafe.com/news-banner-01.jpg',
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
      <Form.Item labelCol={{ span: 2 }} wrapperCol={{ span: 21, offset: 1 }} label={formatMessage({ id: 'app.form.title' })}>
        {form.getFieldDecorator('title', {
          rules: [{ required: true, message: formatMessage({ id: 'app.characters.limit' }), min: 5 }],
          initialValue: current.title,
        })(<Input />)}
      </Form.Item>
      <Form.Item labelCol={{ span: 2 }} wrapperCol={{ span: 21, offset: 1 }} label={formatMessage({ id: 'app.form.description' })}>
        {form.getFieldDecorator('description', {
          rules: [{ required: true, message: formatMessage({ id: 'app.characters.limit' }), min: 5 }],
          initialValue: current.description,
        })(<TextArea
          style={{ minHeight: 32 }}
          rows={3}
        />)}
      </Form.Item>
      <CommonPictureLayout pictureLayout={pictureLayout} form={form} current={current} onChange={handlePictureLayoutChange}/>
    </Modal>
  );
});

@connect(({ news, loading }) => ({
  news,
  loading: loading.models.news,
}))

@Form.create()
class NewsList extends Component {
  state = {
    modalVisible: false,
    pictureLayout: 'justify',
    currentLang: getLocale(),
  };

  componentDidMount() {
    const { dispatch } = this.props;
    dispatch({
      type: 'news/fetch',
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
      pictureLayout: 'justify',
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
      title: item.title,
      description: item.description,
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
      window.open(location.origin + `/#/preview/${fields.title}`)
  };

  handleAdd = (fields, align) => {
    const { dispatch } = this.props;
    dispatch({
      type: 'news/submit',
      payload: {
        title: fields.title,
        description: fields.description,
        content: JSON.stringify(fields.content),
        image: fields.image !== undefined ? JSON.stringify(fields.image) : "",
        align: align,
        lang: this.state.currentLang,
      },
    });
    message.success(formatMessage({ id: 'app.add.success' }));
    this.initModal();
  };

  handleUpdate = (fields, align) => {
    const { dispatch } = this.props;
    const { currentId } = this.state;
    dispatch({
      type: 'news/submit',
      payload: {
        id: currentId,
        title: fields.title,
        description: fields.description,
        content: JSON.stringify(fields.content),
        image: fields.image,
        align: align,
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
      type: 'news/submit',
      payload: {
        id: id,
        lang: this.state.currentLang,
      }
    });
    message.success(formatMessage({ id: 'app.add.success' }));
    this.initModal();
  };

  handlePictureLayoutChange = (e) => {
    this.setState({
      pictureLayout: e.target.value,
    })
  };

  render() {
    const {
      news: { news = [] },
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

    const headFeaturedPost = {
      title: 'Dynamic & News',
      description:
        "Welcome to browse and view.",
      image: 'https://cdn.pharmcafe.com/news-banner-01.jpg',
      imgText: 'head image description',
    };

    const ArticleListContent = ({ data: { description, title, createdAt, updatedAt} }) => (
      <div className={styles.listContent}>
        <div className={styles.title}>
          {title}
        </div>
        {/* <div className={styles.description}>
          {description.substring(0, 500)}...
        </div> */}
        {<Ellipsis lines={3}>{description}</Ellipsis>}
        <div className={styles.extra}>
          {formatMessage({ id: 'app.news.created-at' })}
          <em>{moment(createdAt).add(8, 'hours').format('YYYY-MM-DD HH:mm')}</em>
          &nbsp;&nbsp;&nbsp;{formatMessage({ id: 'app.news.updated-at' })}
          <em>{moment(updatedAt).add(8, 'hours').format('YYYY-MM-DD HH:mm')}</em>
          <Button type="link" onClick={() => this.props.history.push(`/news/${title}`)}>&nbsp;&nbsp;&nbsp;{formatMessage({ id: 'app.news.read-more' })}</Button>
        </div>
      </div>
    );

    return (
      <Fragment>
        <CssBaseline />
        <div>
          <HeadFeaturedPost post={headFeaturedPost} />
        </div>
        <Card
          bordered={false}
          bodyStyle={{ padding: '8px 8px 8px 68px', background: '#eff2f5' }}
        >
          <List
            size="large"
            key="news"
            loading={news.length === 0 ? loading : false}
            rowKey="id"
            itemLayout="vertical"
            dataSource={news}
            renderItem={item => (
              <List.Item
                key={item.id}
                extra={<div className={styles.listItemExtra} />}
                actions={[
                  <a onClick={() => this.showUpdateModal(item)}>
                    {formatMessage({ id: 'app.button.edit' })}
                  </a>,
                  <a onClick={() => this.deleteConfirm(item)}>
                    {formatMessage({ id: 'app.button.delete' })}
                </a>,
                ]}
              >
                <ArticleListContent data={item} />
              </List.Item>
            )}
          />
          <Button type="primary" icon="plus" onClick={this.showModal}>{formatMessage({ id: 'app.button.add' })}</Button> 
        </Card>
               
        <CreateForm {...parentMethods} modalVisible={modalVisible} pictureLayout={pictureLayout} current={current} />
      </Fragment>
    );
  }
}

export default NewsList;
