import React, { FC } from 'react';
import Button from '@material-ui/core/Button';
import AddBoxRoundedIcon from '@material-ui/icons/AddBoxRounded';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import {
 Content,
 Header,
 Page,
 pageTheme,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 return (
   
   <Page theme={pageTheme.tool}>
     <Header
       title="Repairing computer systems" type="group 18">
          <Button variant="contained" color="default" href="/" startIcon={<ExitToAppRoundedIcon />}> Logout
           </Button>
     </Header>

    <Content align="center">
     <Button variant="outlined" color="default" href="/recordusertable" startIcon={<AddBoxRoundedIcon />}> User information record
           </Button>
     </Content>



   </Page>
 );
};
export default WelcomePage;

