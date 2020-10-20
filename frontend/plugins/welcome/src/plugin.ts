 
import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateRepairInvoice from './components/RecordRepairInvoice';
import ShowRepairInvoice from './components/RecordRepairInvoiceTable';
import SignIn from './components/SignIn'
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/Home', WelcomePage);
    router.registerRoute('/RecordRepairInvoice', CreateRepairInvoice);
    router.registerRoute('/RecordRepairInvoiceTable', ShowRepairInvoice);
    router.registerRoute('/', SignIn);
  },
});
