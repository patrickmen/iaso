import React, { Component, Fragment } from 'react';
import { connect } from 'dva';
import moment from 'moment';
import { formatMessage } from 'umi/locale';
import CssBaseline from '@material-ui/core/CssBaseline';
import Ellipsis from '@/components/Ellipsis';
import { Card, List, Button, Icon, Input, Modal, Form, message } from 'antd';
import HeadFeaturedPost from '@/components/Article/HeadFeaturedPost';
import styles from './News.less';

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
      <Form.Item labelCol={{ span: 5 }} wrapperCol={{ span: 18 }} label={formatMessage({ id: 'app.form.title' })}>
        {form.getFieldDecorator('title', {
          rules: [{ required: true, message: formatMessage({ id: 'app.characters.limit' }), min: 5 }],
          initialValue: current.title,
        })(<Input />)}
      </Form.Item>
      <Form.Item labelCol={{ span: 5 }} wrapperCol={{ span: 18 }} label={formatMessage({ id: 'app.form.description' })}>
        {form.getFieldDecorator('description', {
          rules: [{ required: true, message: formatMessage({ id: 'app.characters.limit' }), min: 5 }],
          initialValue: current.description,
        })(<TextArea
          style={{ minHeight: 32 }}
          rows={3}
        />)}
      </Form.Item>
      <Form.Item labelCol={{ span: 5 }} wrapperCol={{ span: 18 }} label="markdown">
        {form.getFieldDecorator('content', {
          rules: [{ required: true, message: formatMessage({ id: 'app.characters.limit' }), min: 5 }],
          initialValue: current.content,
        })(<TextArea rows={10} />)}
      </Form.Item>
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
  };

  componentDidMount() {
    const { dispatch } = this.props;
    dispatch({
      type: 'news/fetch',
      payload: {
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
      title: item.title,
      description: item.description,
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
      window.open(location.origin + `/preview/${fields.title}`)
  };

  handleAdd = fields => {
    const { dispatch } = this.props;
    dispatch({
      type: 'news/submit',
      payload: {
        title: fields.title,
        description: fields.description,
        content: JSON.stringify(fields.content),
      },
    });
    message.success(formatMessage({ id: 'app.add.success' }));
    this.initModal();
  };

  handleUpdate = fields => {
    const { dispatch } = this.props;
    const { currentId } = this.state;
    dispatch({
      type: 'news/submit',
      payload: {
        id: currentId,
        title: fields.title,
        description: fields.description,
        content: JSON.stringify(fields.content),
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
      }
    });
    message.success(formatMessage({ id: 'app.add.success' }));
    this.initModal();
  };

  render() {
    const {
      news: { news = [] },
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
      title: 'Dynamic & News',
      description:
        "Welcome to browse and view.",
      image: 'https://source.unsplash.com/random',
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
          bodyStyle={{ padding: '8px 8px 8px 42px' }}
        >
          <List
            size="large"
            key="test"
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
               
        <CreateForm {...parentMethods} modalVisible={modalVisible} current={current} />
      </Fragment>
    );
  }
}

export default NewsList;
