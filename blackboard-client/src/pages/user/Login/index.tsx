import React, { useState, useEffect } from "react"
import { Box, Button, TextField, Typography, Snackbar } from "@mui/material"
import MuiAlert, { AlertProps } from '@mui/material/Alert';
import { LoginParams } from "../../../model/user"
import { postLogin } from "../../../api/user"
import { useNavigate } from "react-router-dom"
import axios from "axios"
import Header from "../../../components/header"


const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref,
  ) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });



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
      
const LoginPage: React.FC = () =>{

    const [loginData, setLoginData] = useState<LoginParams>({
        username: "",
        password: "",
    })

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) =>{
        const {name, value} = e.target
        setLoginData({...loginData, [name]: value})
    }

    const [user, setUser] = useState<{ isProfessor?: boolean } | null>(null);


    useEffect(() => {
        // 세션에서 사용자 정보를 가져오는 가상의 함수, 실제 구현 필요
        const userInfo = getUserInfoFromSession(); // getUserInfoFromSession 함수는 구현해야 함
        setUser(userInfo);
    }, []);


    const navigate = useNavigate()

    const [message, setMessage] = useState(""); // 메시지 상태
    const [messageType, setMessageType] = useState<"success" | "error">("success"); // 메시지 타입
    const [openSnackbar, setOpenSnackbar] = useState(false); // Snackbar의 열림 상태를 관리

    const handleLogin = async () => {
        
        try{
            const response = await postLogin(loginData)
            setMessage("로그인이 성공적으로 완료되었습니다!");
            setMessageType("success");
            setOpenSnackbar(true);

            const user = response.data.user;

            sessionStorage.setItem("user", JSON.stringify(user));

            // 스낵바가 충분히 표시된 후 페이지를 이동시킵니다.
            setTimeout(() => {
                if(!user.Isprofessor){
                    navigate("/student");
                }
                else{
                    navigate("/professor");
                }
            }, 1000); // 1000ms = 1초 후에 페이지 이동

        } catch (error) {
            let errorMessage = "알 수 없는 에러가 발생했습니다";
            if (axios.isAxiosError(error)) {
                // Axios 오류 응답인 경우
                if (error.response && error.response.status === 400) {
                    // 서버로부터 받은 구체적인 오류 메시지 사용
                    errorMessage = error.response.data.error || "잘못된 요청입니다.";
                }
            } else if (error instanceof Error) {
                // 일반 JavaScript 오류인 경우
                errorMessage = error.message;
            }
            setMessage(errorMessage);
            setMessageType("error");
            setOpenSnackbar(true);
        }
    }

    const handleCloseSnackbar = (event: React.SyntheticEvent | Event, reason?: string) => {
        if (reason === 'clickaway') {
            return;
        }
        setOpenSnackbar(false);
    };   

    return(
        <>
        <Header />
        <Box style={{
            width: "100vw",
            height: "100vh",
            display: "flex",
            flexDirection: "column",
            justifyContent: "center",
            alignItems: "center",
            }}>
            <Typography variant="h2" component={"div"}>로그인 페이지</Typography>
            <TextField 
            name="username"
            required
            autoFocus
            value={loginData.username} 
            label="아이디" 
            onChange={handleChange}
            placeholder="your_ID@korea.ac.kr"
            />
            <TextField
            name="password"
            required
            value={loginData.password} 
            label="비밀번호" 
            onChange={handleChange}
            placeholder="your_PassWord"
            type="password"
            />
            <Button onClick={handleLogin}>
                로그인
            </Button>
            <Snackbar open={openSnackbar} autoHideDuration={6000} onClose={handleCloseSnackbar}>
                <Alert onClose={handleCloseSnackbar} severity={messageType} sx={{ width: '100%' }}>
                    {message}
                </Alert>
            </Snackbar>
        </Box>
        </>
    )
}

export default LoginPage