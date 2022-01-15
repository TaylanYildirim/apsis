import React, {forwardRef, useCallback, useEffect, useState} from "react";
import MaterialTable from "material-table";
import axios from 'axios'
import CircularProgress from '@mui/material/CircularProgress';
import AddBox from '@material-ui/icons/AddBox';
import ArrowDownward from '@material-ui/icons/ArrowDownward';
import Check from '@material-ui/icons/Check';
import ChevronLeft from '@material-ui/icons/ChevronLeft';
import ChevronRight from '@material-ui/icons/ChevronRight';
import Clear from '@material-ui/icons/Clear';
import DeleteOutline from '@material-ui/icons/DeleteOutline';
import Edit from '@material-ui/icons/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import FilterList from '@material-ui/icons/FilterList';
import FirstPage from '@material-ui/icons/FirstPage';
import LastPage from '@material-ui/icons/LastPage';
import Remove from '@material-ui/icons/Remove';
import SaveAlt from '@material-ui/icons/SaveAlt';
import Search from '@material-ui/icons/Search';
import ViewColumn from '@material-ui/icons/ViewColumn';
import {Fab, Grid, Tooltip} from "@material-ui/core";

import moment from 'moment';
import * as _filefy from "filefy";
import Alert from '@mui/material/Alert';
import Stack from '@mui/material/Stack';
import Zoom from "@material-ui/core/Zoom";
import swal from "sweetalert";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import TextField from "@mui/material/TextField";
import DialogActions from "@mui/material/DialogActions";
import Button from "@mui/material/Button";
import AddIcon from "@mui/icons-material/Add";
import {Slider} from "@mui/material";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import {string} from "prop-types";

const ApsisTable = () => {
    const [blocks, setBlocks] = useState([]);
    const [fetchErr, setFetchErr] = useState(false);
    const [open, setOpen] = React.useState(false);
    const [isEdit, setIsEdit] = useState(true);
    const [employeeID, setEmployeeID] = useState(0);
    const [teamID, setTeamID] = useState(0);
    const [teamScore, setTeamScore] = useState(0);
    const [score, setScore] = useState(0);
    const [teamIdTextFieldErr, setTeamIdTextFieldErr] = useState(false);
    let timeout;

    const tableIcons = {
        Add: forwardRef((props, ref) => <AddBox {...props} ref={ref}/>),
        Check: forwardRef((props, ref) => <Check {...props} ref={ref}/>),
        Clear: forwardRef((props, ref) => <Clear {...props} ref={ref}/>),
        Delete: forwardRef((props, ref) => <DeleteOutline {...props} ref={ref}/>),
        DetailPanel: forwardRef((props, ref) => <ChevronRight {...props} ref={ref}/>),
        Edit: forwardRef((props, ref) => <Edit {...props} ref={ref}/>),
        Export: forwardRef((props, ref) => <SaveAlt {...props} ref={ref}/>),
        Filter: forwardRef((props, ref) => <FilterList {...props} ref={ref}/>),
        FirstPage: forwardRef((props, ref) => <FirstPage {...props} ref={ref}/>),
        LastPage: forwardRef((props, ref) => <LastPage {...props} ref={ref}/>),
        NextPage: forwardRef((props, ref) => <ChevronRight {...props} ref={ref}/>),
        PreviousPage: forwardRef((props, ref) => <ChevronLeft {...props} ref={ref}/>),
        ResetSearch: forwardRef((props, ref) => <Clear {...props} ref={ref}/>),
        Search: forwardRef((props, ref) => <Search {...props} ref={ref}/>),
        SortArrow: forwardRef((props, ref) => <ArrowDownward {...props} ref={ref}/>),
        ThirdStateCheck: forwardRef((props, ref) => <Remove {...props} ref={ref}/>),
        ViewColumn: forwardRef((props, ref) => <ViewColumn {...props} ref={ref}/>)
    };

    const columns = [
        {title: 'Team ID', field: 'TeamID', defaultGroupOrder: 0},
        {title: 'Employee ID', field: 'EmployeeID'},
        {title: 'Score', field: 'Score'},
        {title: 'Team Score', field: 'TeamScore'},
    ]
    const tableOptions = () => (
        {
            headerStyle: {
                backgroundColor: '#1876d2',
                color: '#ffffff',
                fontWeight: 'bold',
            },

            pageSize: 10,
            pageSizeOptions: [10, 20, 30],
            exportButton: {csv: true, pdf: true},
            exportCsv: (columns, data) => handleExportCsv(columns, data),
            loadingType: 'linear',
            actionsColumnIndex: -1,
            minBodyHeight: 200,
            maxBodyHeight: 600,
            toolbar: true,
            paging: true,
            search: true,
        }
    )
    function split(array, n) {
        let [...arr]  = array;
        var res = [];
        while (arr.length) {
            res.push(arr.splice(0, n));
        }
        return res;
    }
    const handleExportCsv = (allColumns, allData) => {
        let arr = [], resultArr=[];
        const columns = allColumns.filter(columnDef => columnDef["export"] !== false);
        for (let i in allData) {
           arr.push(...allData[i].data.map(rowData => columns.map(columnDef => rowData[columnDef.field])));
        }

        new _filefy.CsvBuilder('teams_' + moment().format('YYYY-MM-DDTHH_mm') + '.csv')
            .setDelimeter(';')
            .setColumns(columns.map(columnDef => columnDef.title))
            .addRows(arr)
            .exportFile();
    }

    useEffect(async () => {
        const domain = process.env.NODE_ENV === 'production' ?
            'https://apsis-code.herokuapp.com/' :
            `http://127.0.0.1:10000/`;
        const baseURL = `${domain}employees`;
        let response = await axios.get(baseURL)
            .then((response) => {
                setBlocks(response?.data);
                setFetchErr(false);
            })
            .catch((err) => {
                setFetchErr(true);
                console.error(err);
            })
    }, []);

    useEffect(async () => {
        const domain = process.env.NODE_ENV === 'production' ?
            'https://apsis-code.herokuapp.com/' :
            `http://127.0.0.1:10000/`;
        const baseURL = `${domain}employees`
        let response = await axios.get(baseURL)
            .then((response) => {
                setBlocks(response?.data);
                setFetchErr(false);
            })
            .catch((err) => {
                setFetchErr(true);
                console.error(err);
            })
    }, [score]);


    const onClickAdd = async (rowData) => {
        setOpen(true)
        const {EmployeeID, Score, TeamID, TeamScore} = rowData;
        setIsEdit(false)
        setEmployeeID(EmployeeID)
        setScore(Score);
        setTeamID(TeamID);
        setTeamScore(TeamScore);
    }

    const onClickEdit = async (rowData) => {
        setOpen(true)
        const {EmployeeID, Score, TeamID, TeamScore} = rowData;
        setIsEdit(true)
        setEmployeeID(EmployeeID)
        setScore(Score);
        setTeamID(TeamID);
        setTeamScore(TeamScore);
    }

    const onClickDelete = async (rowData) => {
        const domain = process.env.NODE_ENV === 'production' ?
            'https://apsis-code.herokuapp.com/' :
            `http://127.0.0.1:10000/`;
        const {EmployeeID, TeamID} = rowData;
        const baseURL = `${domain}teams/${TeamID}/employees/${EmployeeID}`;
        let response = await axios.delete(baseURL)
            .then((response) => {
                showSuccessPopup("deleted")
            })
            .catch((err) => {
                showFailPopup("deleted").then();
                setFetchErr(true);
            })
        setScore(score + 1)
        setScore(score - 1)
    };

    const handleClickOpen = () => {
        setIsEdit(false);
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const handleAddEmployee = async () => {
        const domain = process.env.NODE_ENV === 'production' ?
            'https://apsis-code.herokuapp.com/' :
            `http://127.0.0.1:10000/`;
        const newEmployee = {score: score, teamScore: teamScore, teamID: teamID};
        if (isEdit) {
            const baseURL = `${domain}employees/${employeeID}`
            let response = await axios.put(baseURL, newEmployee)
                .then((response) => {
                    showSuccessPopup("edited");
                })
                .catch((err) => {
                    showFailPopup("edited");
                    console.error(err);
                })

        } else {
            const baseURL = `${domain}teams/${teamID}/employee`
            let response = await axios.post(baseURL, newEmployee)
                .then((response) => {
                    showSuccessPopup("added");
                })
                .catch((err) => {
                    showFailPopup("added");
                    console.error(err);
                })
        }
        setOpen(false);
        setScore(score + 1)
        setScore(score - 1)
    };

    const handleSliderChange = (e) => {
        timeout && clearTimeout(timeout);
        timeout = setTimeout(() => {
            setScore(e.target.value);
        }, 400);
    }
    const handleTeamIdTextFieldChange = (e) => {
        Number(e.target.value) >= 0 && Number(e.target.value) <= 20 ?
            setTeamIdTextFieldErr(false) :
            setTeamIdTextFieldErr(true);
        setTeamID(e.target.value);
    }

    return (
        <Grid container justifyContent="center">
            <div>
                <Dialog open={open} onClose={handleClose}>
                    <DialogTitle>Employee</DialogTitle>
                    <DialogContent>
                        <DialogContentText>
                            Plese fill employee step count:
                        </DialogContentText>
                        <TextField
                            autoFocus
                            defaultValue="0"
                            type="number"
                            margin="dense"
                            id="teamID"
                            label="Team Id"
                            fullWidth
                            variant="standard"
                            value={teamID}
                            onChange={handleTeamIdTextFieldChange}
                            error={teamIdTextFieldErr}
                            helperText={teamIdTextFieldErr ? "Please enter less than 20 and positive" : ""}
                            required={true}
                            disabled={isEdit}
                        />
                        <TextField
                            autoFocus
                            margin="dense"
                            id="employeeId"
                            label="Employee Id"
                            type="email"
                            fullWidth
                            variant="standard"
                            value={employeeID}
                            disabled
                        />
                        <Box width={400}>
                            <Typography style={{marginTop: "20px"}} id="input-slider" gutterBottom>
                                Step count
                            </Typography>
                            <Slider defaultValue={score} aria-label="Default" valueLabelDisplay="auto"
                                    max={5000}
                                    onChange={handleSliderChange}/>
                        </Box>
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={handleClose}>Cancel</Button>
                        <Button disabled={teamIdTextFieldErr} onClick={handleAddEmployee}>Add</Button>
                    </DialogActions>
                </Dialog>
            </div>
            <Grid item xs={11} style={{display: 'table-cell', width: '100%', height: '100%', verticalAlign: 'middle', textAlign: 'center', minHeight: '50px'}}>
                {fetchErr ?
                    <Stack sx={{width: '100%'}} spacing={2}>
                        <Alert severity="error">Sorry, error occurred during fetching API!</Alert>
                    </Stack>
                    :
                    !blocks?.length < 0 ?
                        <CircularProgress style={{marginTop: '20px'}}/>
                        :
                        <MaterialTable
                            title="TEAMS"
                            columns={columns}
                            data={blocks== null ?[] : blocks}
                            actions={[
                                {
                                    icon: () => <Fab color="primary" style={{backgroundColor: '#1876d2'}}
                                                     aria-label="add" onClick={handleClickOpen}>
                                        <AddIcon/>
                                    </Fab>,
                                    tooltip: 'Add team',
                                    isFreeAction: true,
                                    onClick: (event, rowData) => onClickAdd(rowData)
                                },
                                {
                                    icon: deleteIcon,
                                    onClick: (event, rowData) => onClickDelete(rowData)
                                },
                                {
                                    icon: editIcon,
                                    onClick: (event, rowData) => onClickEdit(rowData)
                                },
                            ]}
                            icons={tableIcons}
                            options={{
                                headerStyle: {
                                    backgroundColor: '#1876d2',
                                    color: '#ffffff',
                                    fontWeight: 'bold',
                                },

                                pageSize: 10,
                                pageSizeOptions: [10, 20, 30],
                                exportButton: {csv: true, pdf: false},
                                exportCsv: (columns, data) => handleExportCsv(columns, data),
                                loadingType: 'linear',
                                actionsColumnIndex: -1,
                                minBodyHeight: 200,
                                maxBodyHeight: 600,
                                toolbar: true,
                                paging: true,
                                search: true,
                            }}
                        />
                }
            </Grid>
        </Grid>
    )
}

const editIcon = () => (
    <Tooltip placement={"top"} TransitionComponent={Zoom} title={'Edit'} arrow
             enterDelay={200}>
        <div style={{padding: '3px'}}>
            <EditIcon fontSize="default" style={{fill: 'gray'}}/>
        </div>
    </Tooltip>
)

const deleteIcon = () => (
    <Tooltip placement={"top"} TransitionComponent={Zoom} title={'Delete'} arrow
             enterDelay={200}>
        <div style={{padding: '3px'}}>
            <DeleteIcon fontSize="default" style={{fill: 'gray'}}/>
        </div>
    </Tooltip>
)

const showSuccessPopup = (message) => swal({
    title: "Operation success!",
    text: `Successfully ${message}!`,
    icon: "success",
    timer: "1200",
});

const showFailPopup = (message) => swal({
    title: "Operation failed!",
    text: `Couldn't be ${message}`,
    icon: "error",
    timer: "1200",
});

export default ApsisTable;