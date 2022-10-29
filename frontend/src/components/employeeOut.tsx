import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { EmployeeInterface } from "../interfaces/IEmployee";
import { GetEmployees } from "../services/HttpClientService";

function WatchVideos() {
  const [watchVideos, setWatchVideos] = useState<EmployeeInterface[]>([]);

  useEffect(() => {
    getWatchVideos();
  }, []);

  const getWatchVideos = async () => {
    let res = await GetEmployees();
    if (res) {
      setWatchVideos(res);
    } 
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ลำดับ", width: 50 },
    {
      field: "FirstName",
      headerName: "First name",
      width: 100,
      valueFormatter: (params) => params.value.FirstName,
    },
    {
      field: "LastName",
      headerName: "Last name",
      width: 100,
      valueFormatter: (params) => params.value.LastName,
    },
    {
      field: "Email",
      headerName: "Email",
      width: 150,
      valueFormatter: (params) => params.value.Email,
    },
    {
      field: "Telephone",
      headerName: "Telephone",
      width: 150,
      valueFormatter: (params) => params.value.Telephone,
    },
    {
      field: "Slary",
      headerName: "Slary",
      width: 100,
      valueFormatter: (params) => params.value.Slary,
    },

    {
      field: "Gender",
      headerName: "Gender",
      width: 100,
      valueFormatter: (params) => params.value.Gender,
    },
    {
      field: "Education",
      headerName: "Education",
      width: 150,
      valueFormatter: (params) => params.value.Education,
    },
    {
      field: "Position",
      headerName: "Position",
      width: 150,
      valueFormatter: (params) => params.value.Position,
    },
    
  ];

  return (
    <div>
      <Container maxWidth="md">
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              Information
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/employee/create"
              variant="contained"
              color="primary"
            >
              Create Information
            </Button>
          </Box>
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={watchVideos}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
          />
        </div>
      </Container>
    </div>
  );
}

export default WatchVideos;