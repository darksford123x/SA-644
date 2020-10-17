 
import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateUser from './components/RecordUser';
import ShowUser from './components/RecordUserTable';
import SignIn from './components/SignIn'
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/Home', WelcomePage);
    router.registerRoute('/RecordUser', CreateUser);
    router.registerRoute('/RecordUserTable', ShowUser);
    router.registerRoute('/', SignIn);
  },
});
