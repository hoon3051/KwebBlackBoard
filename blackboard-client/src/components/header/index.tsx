import {
    AppBar,
    Toolbar,
    IconButton,
    Typography,
    Box,
  } from "@mui/material";

import { Link } from "react-router-dom";
  

  const Header: React.FC = () => {
    return (
      <AppBar component="nav" color = {"transparent"}>
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            edge="start"
            //   onClick={handleDrawerToggle}
            sx={{ mr: 2, display: { sm: "none" } }}
          >
            {/* <MenuIcon /> */}
          </IconButton>
          <Typography
            variant="h6"
            component="div"
            sx={{ flexGrow: 1, display: { xs: "none", sm: "block" } }}
          >
            BlackBoard
          </Typography>
          <Box sx={{ display: { xs: "none", sm: "block" } }}>
            <Link to={"/auth/sign-in"}>로그인&nbsp;</Link>
            <Link to={"/auth/sign-up"}>&nbsp;회원가입</Link>
          </Box>
        </Toolbar>
      </AppBar>
    );
  };
  
  export default Header;