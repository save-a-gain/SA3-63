 
import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateUser from './components/Users';
import CreateReturninvoice from './components/RecordReturninvoice';
import ShowReturninvoice from './components/RecordReturninvoiceTable';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
	router.registerRoute('/RecordReturninvoice', CreateReturninvoice);
    router.registerRoute('/RecordReturninvoiceTable', ShowReturninvoice);
  },
});
 
