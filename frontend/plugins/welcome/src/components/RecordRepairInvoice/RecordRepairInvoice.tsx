import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import {
  Select,
  MenuItem,
} from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert, AlertTitle } from '@material-ui/lab';
import InputLabel from '@material-ui/core/InputLabel';
import SaveRoundedIcon from '@material-ui/icons/SaveRounded';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import InputAdornment from '@material-ui/core/InputAdornment';
import PersonIcon from '@material-ui/icons/Person';

import { DefaultApi } from '../../api/apis';
import { EntDevice } from '../../api/models/EntDevice'; // import interface Device
import { EntUser } from '../../api/models/EntUser'; // import interface User
import { EntSymptom } from '../../api/models/EntSymptom'; // import interface Symptom
import { EntStatusR } from '../../api/models/EntStatusR'; // import interface StatusR
import { EntRepairInvoice } from '../../api';

// css style 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
      margin: theme.spacing(2),
   },
   insideLabel: {
    margin: theme.spacing(1),
  },
   button: {
    marginLeft: '40px',
  },
   textField: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
   },
    select: {
      width: 500 ,
      marginLeft:7,
      marginRight:-7,
      //marginTop:10,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
  }),
);

interface recordRepairInvoice {
  Rename: string;
  Device: number;
  StatusR: number;
  Symptom: number;
  User: number;
}

export default function recordRepairInvoice() {
 const classes = useStyles();
 const http = new DefaultApi();
 
 const [repiarinvoices, setRepairInvoices] = React.useState<EntRepairInvoice[]>([]);


 const [devices, setDevices] = React.useState<EntDevice[]>([]);
 const [statusrs, setStatusRs] = React.useState<EntStatusR[]>([]);
 const [symptoms, setSymptoms] = React.useState<EntSymptom[]>([]);
 const [users, setUsers] = React.useState<EntUser[]>([]);

 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);

 const [loading, setLoading] = useState(true);

 const [rename, setRename] = useState(String);
 const [device, setDevice] = useState(Number);
 const [statusr, setStatusR] = useState(Number);
 const [symptom, setSymptom] = useState(Number);
 const [user, setUser] = useState(Number);

 useEffect(() => {
  const getUsers = async () => {
    const res = await http.listUser({ limit: 10, offset: 0 });
    setLoading(false);
    setUsers(res);
    console.log(res);
  };
  getUsers();

  const getDevices = async () => {
    const res = await http.listDevice({ limit: 10, offset: 0 });
    setLoading(false);
    setDevices(res);
    console.log(res);
  };
  getDevices();

  const getStatusRs = async () => {
    const res = await http.listStatusr({ limit: 10, offset: 0 });
    setLoading(false);
    setStatusRs(res);
    console.log(res);
  };
  getStatusRs();

  const getSymptoms = async () => {
    const res = await http.listSymptom({ limit: 10, offset: 0 });
    setLoading(false);
    setSymptoms(res);
    console.log(res);
  };
  getSymptoms();

}, [loading]);


const getRepairInvoice = async () => {
  const res = await http.listRepairinvoice({ limit: 10, offset: 0 });
  setRepairInvoices(res);
};

const handleRenameChange = (event: any) => {
  setRename(event.target.value as string);
};

const UserhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setUser(event.target.value as number);
};

const StatusRhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setStatusR(event.target.value as number);
};

const DevicehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setDevice(event.target.value as number);
};

const SymptomhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setSymptom(event.target.value as number);
};

// create repairinvoice
const CreateRepairInvoice = async () => {
  const repairinvoice = {
    Rename: rename,
    device: device,
    user: user,
    statusr: statusr,
    symptom: symptom,
  };
  console.log(repairinvoice);
  const res: any = await http.createRepairInvoice({ repairInvoice: repairinvoice });
  console.log("bruhhhhhhhhh");
  setStatus(true);
  
  if (res.id != '') {
    setAlert(true);
  } else {
    setAlert(false);
  }
  const timer = setTimeout(() => {
    setStatus(false);
  }, 3000);
};

  // clear input form
  function clear() {
    setRepairInvoices([]);
  }

  return (
    <Page theme={pageTheme.tool}>

      <Header
        title={`Repaired Invoice record`}
        type="Repairing computer systems"> 
      </Header>

      <Content>
        <ContentHeader title="Repaired Invoice information"> 
              <Button onClick={() => {CreateRepairInvoice();}} variant="contained"  color="primary" startIcon={<SaveRoundedIcon/>}> Create new repaired invoice </Button>
              <Button style={{ marginLeft: 20 }} component={RouterLink} to="/recordrepairinvoicetable" variant="contained" startIcon={<CancelRoundedIcon/>}>  Dismiss </Button>
        </ContentHeader>  
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl
              //fullWidth
              //className={classes.margin}
              variant="outlined"
             
            >
               <div className={classes.paper}><strong>รหัสใบส่งซ่อม</strong></div>
              <TextField className={classes.textField}
    //          style={{ width: 500 ,marginLeft:7,marginRight:-7}}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                    <PersonIcon />
                  </InputAdornment>
                ),
              }}
                id="personalID"
                label=""
                variant="standard"
                color="secondary"
                type="string"
                size="medium"
                value={rename}
                onChange={handleRenameChange}
              />

              <div className={classes.paper}><strong>อุปกรณ์</strong></div>
              <Select className={classes.select}
                //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
                color="secondary"
                labelId="device-label"
                id="device"
                value={device}
                onChange={DevicehandleChange}
              >
                <InputLabel className={classes.insideLabel} id="device-label">เลือกอุปกรณ์(Device)</InputLabel>

                {devices.map((item: EntDevice) => (
                  <MenuItem value={item.id}>{item.dname}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>ชื่อผู้ใช้บริการ</strong></div>
              <Select className={classes.select}
                //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
                color="secondary"
                id="user"
                value={user}
                onChange={UserhandleChange}
              >
                <InputLabel className={classes.insideLabel}>เลือกชื่อผู้ใช้บริการ(User)</InputLabel>

                {users.map((item: EntUser) => (
                  <MenuItem value={item.id}>{item.uname}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>รูปแบบอาการ</strong></div>
              <Select className={classes.select}
                //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
                color="secondary"
                id="symptom"
                value={symptom}
                onChange={SymptomhandleChange}
              >
                <InputLabel className={classes.insideLabel}>เลือกรูปแบบอาการ(Symptom)</InputLabel>

                {symptoms.map((item: EntSymptom) => (
                  <MenuItem value={item.id}>{item.syname}</MenuItem>
                ))}
              </Select>

              <div className={classes.paper}><strong>สถานะการซ่อม</strong></div>
              <Select className={classes.select}
               // style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
                color="secondary"
                id="statusr"
                value={statusr}
                onChange={StatusRhandleChange}
              >
                <InputLabel className={classes.insideLabel}>เลือกสถานะการซ่อม(Status)</InputLabel>

                {statusrs.map((item: EntStatusR) => (
                  <MenuItem value={item.id}>{item.sname}</MenuItem>
                ))}
              </Select>

              {status ? ( 
                      <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
              {alert ? ( 
                      <Alert severity="success"> <AlertTitle>Success</AlertTitle> Complete data — check it out! </Alert>) 
              : (     <Alert severity="warning"> <AlertTitle>Warining</AlertTitle> Incomplete data — please try again!</Alert>)} </div>
            ) : null}
            
            </FormControl>

          </form>
        </div>
      </Content>
    </Page>
  );
 }
