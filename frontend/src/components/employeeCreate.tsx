import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import TextField from "@mui/material/TextField";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";

import { EducationsInterface } from "../interfaces/IEducation";
import { GendersInterface } from "../interfaces/IGender";
import { PositionsInterface } from "../interfaces/IPosition";
import { EmployeeInterface } from "../interfaces/IEmployee";

import {
  GetGender,
  GetEducation,
  GetPosition,
  CreateEmployee
} from "../services/HttpClientService";
const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function WatchVideoCreate() {
  const [videos, setVideos] = useState<GendersInterface[]>([]);
  const [resolutions, setResolutions] = useState<EducationsInterface[]>([]);
  const [positions, setPositions] = useState<PositionsInterface[]>([]);
  const [watchVideo, setWatchVideo] = useState<EmployeeInterface>({
    StartJob: new Date(),
  });

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof WatchVideoCreate;
    const { value } = event.target;
    setWatchVideo({ ...watchVideo, [id]: value });
  };

  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof watchVideo;
    setWatchVideo({
      ...watchVideo,
      [name]: event.target.value,
    });
  };

  const getVideos = async () => {
    let res = await GetGender();
    if (res) {
      setVideos(res);
    }
  };

  const getResolution = async () => {
    let res = await GetEducation();
    if (res) {
      setResolutions(res);
    }
  };

  const getPositon = async () => {
    let res = await GetPosition();
    if (res) {
      setPositions(res);
    }
  };


  useEffect(() => {
    getVideos();
    getResolution();
    getPositon();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      FirstName: watchVideo.FirstName ?? "",
      LastName: watchVideo.LastName ?? "",
      Email: watchVideo.Email ?? "",
      Password: watchVideo.Password ?? "",
      Telephone: watchVideo.Telephone ?? "",
      Slary: watchVideo.Slary ?? "",
      EducationID: convertType(watchVideo.EducationID),
      PositionID: convertType(watchVideo.PositionID),
      GenderID: convertType(watchVideo.GenderID),
      StartJob: watchVideo.StartJob,
    };

    let res = await CreateEmployee(data);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
  }

  return (
    <Container maxWidth="md">
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
        Successed!!
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
        Unsuccess!!
        </Alert>
      </Snackbar>
      <Paper>
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 1 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              Information Employee
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={6}>
            <FormControl fullWidth >
            <p>First name</p>
              <TextField
                margin="normal"
                required
                id="FirstName"
                type="string"
                size="medium"
                autoFocus
                value={watchVideo.FirstName || ""}
                onChange={handleInputChange}
                label="First Name"
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth >
            <p>Last name</p>
              <TextField
                margin="normal"
                required
                id="LastName"
                type="string"
                size="medium"
                autoFocus
                value={watchVideo.LastName || ""}
                onChange={handleInputChange}
                label="Last Name"
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Gender</p>
              <Select
                native
                value={watchVideo.GenderID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "GenderID",
                }}
              >
                <option aria-label="None" value="">
                choose gender
                </option>
                {videos.map((item: GendersInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Gender}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth >
            <p>Email</p>
              <TextField
                margin="normal"
                required
                id="Email"
                type="string"
                size="medium"
                autoFocus
                value={watchVideo.Email || ""}
                onChange={handleInputChange}
                label="Email"
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth >
            <p>Password</p>
              <TextField
                margin="normal"
                required
                id="Password"
                type="string"
                size="medium"
                autoFocus
                value={watchVideo.Password || ""}
                onChange={handleInputChange}
                label="Password"
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth >
            <p>Telephone</p>
              <TextField
                margin="normal"
                required
                id="Telephone"
                type="string"
                size="medium"
                autoFocus
                value={watchVideo.Telephone || ""}
                onChange={handleInputChange}
                label="Telephone"
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth >
            <p>Salary</p>
              <TextField
                margin="normal"
                required
                id="Slary"
                type="number"
                size="medium"
                autoFocus
                value={watchVideo.Slary || ""}
                onChange={handleInputChange}
                label="Salary"
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Education</p>
              <Select
                native
                value={watchVideo.EducationID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "EducationID",
                }}
              >
                <option aria-label="None" value="">
                choose education
                </option>
                {resolutions.map((item: EducationsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Education}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Position</p>
              <Select
                native
                value={watchVideo.PositionID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "PositionID",
                }}
              >
                <option aria-label="None" value="">
                choose position
                </option>
                {positions.map((item: PositionsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Position}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

        

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Start Job</p>
              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                  value={watchVideo.StartJob}
                  onChange={(newValue) => {
                    setWatchVideo({
                      ...watchVideo,
                      StartJob: newValue,
                    });
                  }}
                  renderInput={(params) => <TextField {...params} />}
                />
              </LocalizationProvider>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/employee"
              variant="contained"
              color="inherit"
            >
              Back
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              Save
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default WatchVideoCreate;