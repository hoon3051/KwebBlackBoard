import {
    AppBar,
    Toolbar,
    IconButton,
    Typography,
    Box,
  } from "@mui/material";

import { useEffect, useState } from "react";

import { Link } from "react-router-dom";

const getUserInfoFromSession = () => {
  const userInfo = sessionStorage.getItem('user');
  if (!userInfo) {
      return null;
  }

  try {
      // 세션 스토리지에는 문자열로 저장되어 있으므로, JSON 형식으로 파싱합니다.
      const parsedUserInfo = JSON.parse(userInfo);
      return parsedUserInfo;
  } catch (error) {
      // 파싱에 실패한 경우, 오류를 콘솔에 기록하고 null을 반환합니다.
      console.error('Failed to parse user info from session storage:', error);
      return null;
  }
}
  

  const Header: React.FC = () => {

    const [user, setUser] = useState<{ isprofessor?: boolean } | null>(null);


    useEffect(() => {
        // 세션에서 사용자 정보를 가져오는 가상의 함수, 실제 구현 필요
        const userInfo = getUserInfoFromSession(); // getUserInfoFromSession 함수는 구현해야 함
        setUser(userInfo);
    }, []);

    // 사용자 상태에 따라 리디렉션 경로 설정
    const getPath = () => {
        
        if (!user) return "/"; // 사용자 정보가 없으면 초기화면 메인페이지로
        if (user.isprofessor) return "/professor"; // 사용자가 교수면 교수 페이지로
        return "/student"; // 그 외에는 학생 페이지로
    };

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
            <Link to={getPath()} style={{ color: 'black', textDecoration: 'none' }}>BlackBoard</Link>
          </Typography>
          <Box sx={{ display: { xs: "none", sm: "block" } }}>
            <Link to={"/user/login"}>로그인&nbsp;</Link>
            <Link to={"/user/register"}>&nbsp;회원가입</Link>
          </Box>
        </Toolbar>
      </AppBar>
    );
  };
  
  export default Header;