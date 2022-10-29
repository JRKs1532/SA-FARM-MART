import React, { useEffect, useState } from "react";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Snackbar from "@mui/material/Snackbar";
import Divider from "@mui/material/Divider";
import InputLabel from '@mui/material/InputLabel';
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import Typography from "@mui/material/Typography";
import TextField from "@mui/material/TextField";
import FormControl from "@mui/material/FormControl";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import { Autocomplete, Button } from "@mui/material";

import { EmployeeInterface } from "../interfaces/IEmployee";
import { MemberInterface } from "../interfaces/IMember";
import { ProductInterface } from "../interfaces/IProduct";
import { OrderInterface } from "../interfaces/IOrder";
import { CartInterface } from "../interfaces/ICart";

import {
    GetEmployeeByEID,
    GetMembers,
    GetProducts,
    GetCarts,
    //GetOrders,
    Carts,
    Orders,

  } from "../services/HttpClientService";


const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
  ) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function Cart(){
    const [employee, setEmployee] = useState<EmployeeInterface>();
    const [members, setMembers] = useState<MemberInterface[]>([]);
    const [products, setProducts] = useState<ProductInterface[]>([]);
    const [Quantity, setQuantity] = React.useState<number | null>(null);
    const [carts, setCarts] = React.useState<Partial<CartInterface>>({});
    const [cartID, setcartID] = React.useState<number | null>(null);
    const [orders, setOrders] = React.useState<Partial<OrderInterface>>({});
     
    const [success, setSuccess] = React.useState(false);
    const [error, setError] = React.useState(false);
    const [lastcartID, setlastcartID] = React.useState<number | null>(null);
    

    const handleChangeProduct_Name = (event: any, value: any) => {
        setOrders({ ...orders, ProductID: value?.ID }); 
    };
    
    const handleChange = (event: SelectChangeEvent) => {
        const name = event.target.name as keyof typeof carts;
        setCarts({
          ...carts,
          [name]: event.target.value,
        });
    }
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
//============================================== START step 9 บันทึกใบสั่งยา() ==============================================

    /* Insert */
    async function submit() {
        
        const convertType = (data: string | number | undefined) => {
            let val = typeof data === "string" ? parseInt(data) : data;
            return val;
        };

        let data_cart = {
            Employee_ID: Number(localStorage.getItem("uid")),
            Telephone: carts.Telephone?? "",
        };
        let data_order = {
            ProducID:convertType(orders?.ProductID),
            Quantity: Quantity,
            //Quantity:typeof orders?.Product_quantity ==="string"?parseInt(orders.Product_quantity):0,
        };

        const apiUrl = "http://localhost:8080";

        const requestOptions_cart = {
            method: "POST",
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data_cart),
        };

        const requestOptions_order = {
            method: "POST",
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data_order),
        };

            // //ต้องใช้ await ไม่งั้นมันจะไปทำคำสั่งต่อไปเลยโดยไม่รอคำสั่งนี้ทำเสร็จ แล้วมันจะแจ้งว่าหา dispensationID ไม่เจอ */
            // if( lastcartID!= lastcartID){ // หากค่าเท่ากันจะไม่บันทึกซ้ำอีกรอบ
            //     // ตรวจสอบว่า Medicine ID และ Amount ได้ถูกกรอกร/เลือก หรือไม่ ถ้าไม่ถูกกรอกจะไม่ทำการ fetch และ
            //     // พอ fetch dsiepnsation medicine ก็จะแจ้ง error เพราะหา dispensation ไม่เจอ เนื่องจากมันไม่ถุก create จากตรงนี้
            //     if(data_order.Quantity && data_order.ProducID){ // หากเป็น null จะเป็นเท็จ
            //         await fetch(`${apiUrl}/carts`, requestOptions_cart)
            //             .then((response) => response.json())
            //             .then((res) => {
            //             if (res.data) {
            //                 setlastcartID(cartID)
            //                 setSuccess(true)
            //             } else {
            //                 setError(true)
            //             }
            //         });
            //     }
            // }
        

        fetch(`${apiUrl}/orders`, requestOptions_order)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) { 
                    setSuccess(true)
                } else {
                    setError(true)
                }
            console.log("orders", orders)   
            });
    }
    //** 5: ดึงข้อมูลทั้งหมด() */
    const getCarts = async () => {
        const apiUrl = "http://localhost:8080/carts";
        const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        };
       
        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setcartID((res.data.at(-1).ID)+1); // ตรวจสอบเอาเฉพาะ cartID ล่าสุด
                }else{
                    setcartID(1);
                }
            });
    };
    
    // Get Employee
    const getEmployee = async () => {
        let res = await GetEmployeeByEID();
            carts.EmployeeID = res.ID;
        if (res) {
        setEmployee(res);
        }
    };

    const getMembers = async () => {
      let res = await GetMembers();
          carts.MemberID = res.ID;
      if (res) {
      setMembers(res);
      }
  };

    /*ดึงข้อมูลทั้งหมด() */
    const getProducts = async () => {
        const apiUrl = "http://localhost:8080/products";
        const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        };
       
        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setProducts(res.data);
                }
            });
    };
    
    React.useEffect(() => {
        getCarts();
        getMembers();
        getEmployee();
        getProducts();
    }, []);
    var productUnitArray = products.map((item: ProductInterface) => (item.Product_name));


return (
    <Container maxWidth="md">
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
            Success!
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          Failed!
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
              Shopping Cart
            </Typography>
          </Box>
        </Box>
        <Divider />

        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={12}>
            <FormControl fullWidth >
            <InputLabel id="InputLabeltel">Telephone number</InputLabel>      
                <Select
                  required
                  labelId="labelIdtel"
                  id="idtel"
                  label="Telephone Number"
                  native
                  value={carts.MemberID + ""}
                  onChange={handleChange}
                  inputProps={{
                    name: "MemberID",
                  }}                
                >
                    <option aria-label="None" value=""></option>
                    {members.map((item: MemberInterface) => (
                        <option value={item.ID} key={item.ID}>
                        {item.Telephone}
                        </option>
                    ))}
                </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
            <Autocomplete
                id="product"
                options={products}
          
                sx={{ width: 420,paddingY: 2 }}
                style={{ float: "right" }}
                size="medium"
                onChange={handleChangeProduct_Name}
                getOptionLabel={(option: any) => 
                    `${option.Product_name}`
                } //filter value
                renderInput={(params) => <TextField {...params} label="Product " />}
                renderOption={(props: any, option: any) => {
                console.log("props :",props)
                console.log("option :",option);
                return (
                    <li
                    {...props}
                    value={`${option.ID}`}
                    key={`${option.ID}`}
                    >{`${option.Product_name}`}</li>
                ); 
                }}
            />
            </FormControl>
          </Grid>

          <Grid item xs={6} >
            <FormControl fullWidth variant="outlined" >
            <TextField
                margin="normal"
                required
                // sx={{ paddingY: 4 }}
                id="Product_quantity"
                variant="outlined"
                type="number"
                size="medium"  
                value={orders?.Product_quantity }
                onChange={(event) => setQuantity(Number(event.target.value))}
                label="Quantity"
            /> 
            </FormControl>
          </Grid>
          {/* <Grid item xs={2}>
              <h4> x {productUnitArray[Number(orders.ProductID) - 1]} </h4>
          </Grid> */}
    
          <Grid item xs={12}>
            <FormControl fullWidth variant="outlined">
            <InputLabel id="demo-simple-select-label">Employee</InputLabel>      
              <Select
                native
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                label="Employee"                
                value={carts.EmployeeID + ""}
                onChange={handleChange}
                disabled
                inputProps={{
                  name: "EmployeeID",
                }}
              >
                <option value={employee?.ID} key={employee?.ID}>
                  {employee?.FirstName} {employee?.LastName}
                </option>    
              </Select>
            </FormControl>
          </Grid>
        
          <Grid item xs={12}>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              Check Out
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}
export default Cart;