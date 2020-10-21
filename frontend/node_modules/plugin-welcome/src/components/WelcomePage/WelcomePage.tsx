import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponentsRecordReturnInvoiceTable from '../RecordReturnInvoiceTable';
import ComponanceDemo from '../Demo';
import Button from '@material-ui/core/Button';
import Autocomplete from '@material-ui/lab/Autocomplete';

import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
 
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'Return the repaired device' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`Welcome ${profile.givenName || 'to Backstage'}`}
       subtitle="Some quick intro and links."
     ></Header>
     <Content>
       <ContentHeader title=" ">
         <Link component={RouterLink} to="/RecordReturnInvoice">
           <Button variant="contained" color="primary">
             Add User
           </Button>
         </Link>
       </ContentHeader>
      
     </Content>
   </Page>
 );
};
 
export default WelcomePage;
