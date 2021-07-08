import dialog from './zh-CN/dialog';
import exception from './zh-CN/exception';
import menu from './zh-CN/menu';
import news from './zh-CN/news';

export default {
  'navBar.lang': '语言',
  'app.footer.description': 'IASO后台管理系统',
  'app.add.success': '新增成功！',
  'app.update.success': '更新成功!',
  'app.delete.success': '删除成功!',
  'app.delete-confirm-title': '删除卡片',
  'app.delete-confirm-content': '确定删除该条记录吗?',
  'app.button.add': '新增',
  'app.button.edit': '编辑',
  'app.button.delete': '删除',
  'app.button.add-new': '新增',
  'app.characters.limit': '不能少于5个字符哦!',
  'app.characters.empty': '不能为空哦!',
  ...dialog,
  ...exception,
  ...menu,
  ...news,
};
