import {Grid} from "@mui/material";
import TextField from "@mui/material/TextField";
import MenuItem from "@mui/material/MenuItem";
import { AdapterMoment } from '@mui/x-date-pickers/AdapterMoment';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DesktopDatePicker } from '@mui/x-date-pickers/DesktopDatePicker';
import * as React from "react";
import {useEffect, useState} from "react";
import axios from "axios";
import Autocomplete from "@mui/material/Autocomplete";

export default function AdvancedSearchEmployees(props){
    const [cityAndPostcodes, setCityAndPostcodes] = useState([]);
    const [postcodes, setPostcodes] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/api/v1/postcodes_by_settlement')
            .then(
                (response) => {
                    setCityAndPostcodes(response.data);
                }
            )
    }, []);

    const handleChangeCity = (event, value) => {
        let my_event = {};
        const target = {name: "settlement", value: value};
        my_event.target = target;
        props.onChange(my_event);
        const postcodes = cityAndPostcodes[value];
        if (typeof postcodes === 'undefined') return;
        setPostcodes(postcodes);
    }

    const handleDateStartChange = value =>{
        let event = {};
        const target = {name: "birth_date_start", value: value.format('YYYY-M-DD')};
        event.target = target;
        props.onChange(event);
    }

    const handleDateEndChange = value =>{
        let event = {};
        const target = {name: "birth_date_end", value: value.format('YYYY-M-DD')};
        event.target = target;
        props.onChange(event);
    }

    return (
        <div>
            <Grid container spacing={2}>
                <Grid item xs={4.5}>
                    <Autocomplete
                        fullWidth
                        name="settlement"
                        renderInput={(params => <TextField
                            {...params}
                            label={"???????????????????? ??????????"}
                        />)}
                        onChange={handleChangeCity}
                        options={Object.keys(cityAndPostcodes)}
                    />
                </Grid>
                <Grid item xs={3}>
                    <TextField
                        name="postcode"
                        select
                        label="????????????"
                        fullWidth
                        value={props.values.postcode}
                        onChange={props.onChange}
                    >
                        {
                            postcodes.map((postcode) => (
                                <MenuItem key={postcode} value={postcode}>{postcode}</MenuItem>
                            ))
                        }
                    </TextField>
                </Grid>
                <Grid item xs={4.5}>
                    <TextField
                        fullWidth
                        select
                        label="??????????????????"
                        id="position"
                        name="position"
                        value={props.values.position}
                        onChange={props.onChange}
                    >
                        <MenuItem value="">???? ??????????????</MenuItem>
                        <MenuItem value={'??????????????????'}>??????????????????</MenuItem>
                        <MenuItem value={'????????????????'}>????????????????</MenuItem>
                        <MenuItem value={'?????????????????? ?????????????????? ??????????'}>?????????????????? ?????????????????? ??????????</MenuItem>
                        <MenuItem value={'?????????????????? ???????????????????????????? ????????????'}>?????????????????? ???????????????????????????? ????????????</MenuItem>
                    </TextField>
                </Grid>
                <Grid item xs={3}>
                    <LocalizationProvider dateAdapter={AdapterMoment}>
                        <DesktopDatePicker
                            name="date_start"
                            onChange={handleDateStartChange}
                            value={props.values.birth_date_start}
                            label={"???????? ???????????????? ??"}
                            inputFormat="DD/MM/YYYY"
                            renderInput={(params) => <TextField fullWidth {...params} />}
                        />
                    </LocalizationProvider>
                </Grid>
                <Grid item xs={3}>
                    <LocalizationProvider dateAdapter={AdapterMoment}>
                        <DesktopDatePicker
                            name="date_end"
                            onChange={handleDateEndChange}
                            value={props.values.birth_date_end}
                            label={"???????? ???????????????? ????"}
                            inputFormat="DD/MM/YYYY"
                            renderInput={(params) => <TextField fullWidth {...params} />}
                        />
                    </LocalizationProvider>
                </Grid>
                <Grid item xs={2}>
                    <TextField
                        fullWidth
                        select
                        label="??????"
                        name="gender"
                        value={props.values.gender}
                        onChange={props.onChange}
                    >
                        <MenuItem value="">???? ??????????????</MenuItem>
                        <MenuItem value={'??'}>??????????????</MenuItem>
                        <MenuItem value={'??'}>??????????????</MenuItem>
                    </TextField>
                </Grid>
                <Grid item xs={4}>
                    <TextField
                        fullWidth
                        label="??????????????"
                        name="phone_number"
                        value={props.values.phone_number}
                        onChange={props.onChange}
                        />
                </Grid>
            </Grid>
        </div>
    );
}