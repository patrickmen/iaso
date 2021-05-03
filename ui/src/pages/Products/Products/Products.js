// import React, { PureComponent, Fragment } from 'react';
// import CssBaseline from '@material-ui/core/CssBaseline';
// import { formatMessage } from 'umi/locale';
// import { connect } from 'dva';
// import { Card, List, Button, Icon, Input, Modal, Form, message } from 'antd';

// import Ellipsis from '@/components/Ellipsis';
// import HeadFeaturedPost from '@/components/Article/HeadFeaturedPost';
// import styles from './Products.less';

// const { TextArea } = Input;

// const CreateForm = Form.create()(props => {
//     const { modalVisible, form, handleAdd, handleUpdate, handlePreview, current, initModal } = props;
//     const okHandle = () => {
//       form.validateFields((err, fieldsValue) => {
//         if (err) return;
//         form.resetFields();
//         if (current.id !== null && current.id !== undefined) {
//           handleUpdate(fieldsValue);
//         } else {
//           handleAdd(fieldsValue);
//         }
//       });
//     };
//     const onPreview =() => {
//       form.validateFields((err, fieldsValue) => {
//         if(err) return;
//         handlePreview(fieldsValue);
//       });
//     }
//     return (
//       <Modal
//         destroyOnClose
//         title={formatMessage({ id: 'app.dialog.title' })}
//         visible={modalVisible}
//         footer={[
//             <Button onClick={() => initModal()}>{formatMessage({ id: 'app.dialog.cancel' })}</Button>,
//             <Button onClick={onPreview}>{formatMessage({ id: 'app.dialog.preview' })}</Button>,
//             <Button onClick={okHandle}>{formatMessage({ id: 'app.dialog.submit' })}</Button>,
//         ]}
//         onOk={okHandle}
//         onCancel={() => initModal()}
//       >
//         <Form.Item labelCol={{ span: 5 }} wrapperCol={{ span: 18 }} label={formatMessage({ id: 'app.form.title' })}>
//           {form.getFieldDecorator('title', {
//             rules: [{ required: true, message: formatMessage({ id: 'app.characters.limit' }), min: 5 }],
//             initialValue: current.title,
//           })(<Input />)}
//         </Form.Item>
//         <Form.Item labelCol={{ span: 5 }} wrapperCol={{ span: 18 }} label={formatMessage({ id: 'app.form.cover' })}>
//           {form.getFieldDecorator('cover', {
//             rules: [{ required: true, message: formatMessage({ id: 'app.characters.empty' })}],
//             initialValue: current.cover,
//           })(<Input />)}
//         </Form.Item>
//         <Form.Item labelCol={{ span: 5 }} wrapperCol={{ span: 18 }} label={formatMessage({ id: 'app.form.description' })}>
//           {form.getFieldDecorator('description', {
//             rules: [{ required: true, message: formatMessage({ id: 'app.characters.limit' }), min: 5 }],
//             initialValue: current.description,
//           })(<TextArea
//             style={{ minHeight: 32 }}
//             rows={3}
//           />)}
//         </Form.Item>
//         <Form.Item labelCol={{ span: 5 }} wrapperCol={{ span: 18 }} label="markdown">
//           {form.getFieldDecorator('content', {
//             rules: [{ required: true, message: formatMessage({ id: 'app.characters.limit' }), min: 5 }],
//             initialValue: current.content,
//           })(<TextArea rows={10} />)}
//         </Form.Item>
//       </Modal>
//     );
//   });

// // @connect(({ products, loading }) => ({
// //     products,
// //     loading: loading.models.products,
// // }))

// @Form.create()
// class ProductsList extends PureComponent {
//   state = {
//     modalVisible: false,
//   };

//   componentDidMount() {
//     const { dispatch } = this.props;
//     dispatch({
//       type: 'products/fetch',
//       payload: {
//       },
//     });
//   };

//   initModal = () => {
//     this.setState({
//       modalVisible: false,
//       current: {},
//       currentId: null,
//     });
//   };

//   showModal = () => {
//     this.setState({
//       modalVisible: true,
//       current: undefined,
//     });
//   };

//   showUpdateModal = item => {
//     let currentValue = {
//       id: item.id,
//       title: item.title,
//       cover: item.cover,
//       description: item.description,
//       content: JSON.parse(item.content)
//     }
//     this.setState({
//       modalVisible: true,
//       currentId: item.id,
//       current: currentValue,
//     });
//   };

//   handlePreview = fields => {
//       let data = {
//           content: JSON.stringify(fields.content),
//       }
//       window["data"] = data;
//       window.open(location.origin + `/#/preview/${fields.title}`)
//   };

//   handleAdd = fields => {
//     const { dispatch } = this.props;
//     dispatch({
//       type: 'products/submit',
//       payload: {
//         title: fields.title,
//         cover: fields.cover,
//         description: fields.description,
//         content: JSON.stringify(fields.content),
//       },
//     });
//     message.success(formatMessage({ id: 'app.add.success' }));
//     this.initModal();
//   };

//   handleUpdate = fields => {
//     const { dispatch } = this.props;
//     const { currentId } = this.state;
//     dispatch({
//       type: 'products/submit',
//       payload: {
//         id: currentId,
//         title: fields.title,
//         cover: fields.cover,
//         description: fields.description,
//         content: JSON.stringify(fields.content),
//       }
//     });
//     message.success(formatMessage({ id: 'app.update.success' }));
//     this.initModal();
//   };

//   deleteConfirm = item => {
//     Modal.confirm({
//       title: formatMessage({ id: 'app.delete-confirm-title' }),
//       content: formatMessage({ id: 'app.delete-confirm-content' }),
//       okText: formatMessage({ id: 'app.dialog.ok' }),
//       cancelText: formatMessage({ id: 'app.dialog.cancel' }),
//       onOk: () => this.handleDelete(item.id),
//     })
//   };

//   handleDelete = id => {
//     const { dispatch } = this.props;
//     dispatch({
//       type: 'products/submit',
//       payload: {
//         id: id,
//       }
//     });
//     message.success(formatMessage({ id: 'app.delete.success' }));
//     this.initModal();
//   };

//   render() {
//     const {
//       products: { products = [] },
//       loading,
//     } = this.props;

//     const { modalVisible, current = {}, currentId = null } = this.state;

//     const parentMethods = {
//       initModal: this.initModal,
//       handleAdd: this.handleAdd,
//       handleUpdate: this.handleUpdate,
//       handlePreview: this.handlePreview,
//     };

//     const headFeaturedPost = {
//       title: 'Products',
//       description:
//         "Welcome to browse and view.",
//       image: 'https://source.unsplash.com/random',
//       imgText: 'head image description',
//     };

//     const cardList = products ? (
//       <List
//         rowKey="id"
//         loading={loading}
//         grid={{ gutter: 24, xl: 4, lg: 3, md: 3, sm: 2, xs: 1 }}
//         dataSource={['', ...products]}
//         renderItem={item => item ? (
//           <List.Item
//             actions={[
//                 <a onClick={() => this.showUpdateModal(item)}>
//                 {formatMessage({ id: 'app.button.edit' })}
//                 </a>,
//                 <a onClick={() => this.deleteConfirm(item)}>
//                 {formatMessage({ id: 'app.button.delete' })}
//             </a>,
//             ]}
//           >
//             <Card
//               className={styles.card}
//               hoverable
//               cover={<img alt={item.title} src={item.cover} />}
//             >
//               <Card.Meta
//                 title={<a>{item.title}</a>}
//                 description={<Ellipsis lines={2}>{item.subDescription}</Ellipsis>}
//               />
//             </Card>
//           </List.Item>
//         ): (
//           <List.Item>
//             <Button type="dashed" className={styles.newButton} onClick={this.showModal}>
//             <Icon type="plus" /> {formatMessage({ id: 'app.button.add-new' })}
//             </Button>
//           </List.Item>
//         )}
//       />
//     ) : null;

//     return (
//       <Fragment>
//         <CssBaseline />
//         <div>
//           <HeadFeaturedPost post={headFeaturedPost} />
//         </div>
//         <Card
//           bordered={false}
//           bodyStyle={{ padding: '8px 8px 30px 42px' }}
//         >
//           <div className={styles.coverCardList}>  
//             <div className={styles.cardList}>{cardList}</div>
//           </div>
//         </Card> 
//         <CreateForm {...parentMethods} modalVisible={modalVisible} current={current} />   
//       </Fragment>
//     );
//   }
// }

// export default ProductsList;
