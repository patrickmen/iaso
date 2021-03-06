import dialog from './en-US/dialog';
import exception from './en-US/exception';
import menu from './en-US/menu';
import news from './en-US/news';

export default {
  'navBar.lang': 'Languages',
  'app.footer.description': 'IASO BM System',
  'app.add.success': 'Add Success!',
  'app.update.success': 'Update Success!',
  'app.delete.success': 'Delete Success!',
  'app.delete-confirm-title': 'Delete Card',
  'app.delete-confirm-content': 'Sure to delete the record?',
  'app.button.add': 'Add',
  'app.button.edit': 'Edit',
  'app.button.delete': 'Delete',
  'app.button.add-new': 'Add New',
  'app.characters.limit': 'Cannot less than 5 characters!',
  'app.characters.empty': 'Cannot be empty!',
  ...dialog,
  ...exception,
  ...menu,
  ...news,
};
