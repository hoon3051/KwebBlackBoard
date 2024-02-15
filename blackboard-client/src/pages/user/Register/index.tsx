import React, { useState } from "react"
import { useNavigate } from "react-router-dom"
import { postRegister } from "../../../api/user"
import { RegisterParams } from "../../../model/user"
import { Box, Button, Checkbox, FormControlLabel, TextField, Typography, Snackbar } from "@mui/material"
import MuiAlert, { AlertProps } from '@mui/material/Alert';
import axios from "axios"
import Header from "../../../components/header"

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref,
  ) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });


const RegisterPage: React.FC = () =>{
    const [registerData, setRegisterData] = useState<RegisterParams>({
        username: "",
        password: "",
        displayname: "",
        studentnumber: "",
        isprofessor: false,
    })

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) =>{
        const {name, value} = e.target
        setRegisterData({...registerData, [name]: value})
    }

    const navigate = useNavigate()

    const [message, setMessage] = useState(""); // 메시지 상태
    const [messageType, setMessageType] = useState<"success" | "error">("success"); // 메시지 타입
    const [openSnackbar, setOpenSnackbar] = useState(false); // Snackbar의 열림 상태를 관리


    const handleRegister = async () => {
        
        try {
            await postRegister(registerData);
            setMessage("회원가입이 성공적으로 완료되었습니다!");
            setMessageType("success");
            setOpenSnackbar(true);

            // 스낵바가 충분히 표시된 후 페이지를 이동시킵니다.
            setTimeout(() => {
                navigate("/user/login");
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
            <Typography variant="h2" component={"div"}>회원가입 페이지</Typography>
            <TextField 
            name="username"
            required
            autoFocus
            value={registerData.username} 
            label="아이디" 
            onChange={handleChange}
            placeholder="your_ID@korea.ac.kr"
            />
            <TextField
            name="password"
            required
            value={registerData.password}
            label="비밀번호"
            onChange={handleChange}
            placeholder="your_PassWord"
            type="password"
            />
            <TextField
            name="displayname"
            required
            value={registerData.displayname}
            label="이름"
            onChange={handleChange}
            placeholder="your_Displayname"
            />
            <TextField
            name="studentnumber"
            required
            value={registerData.studentnumber}
            label="학번"
            onChange={handleChange}
            placeholder="YYYYCCDNNN"
            />
            <FormControlLabel
            control={
                <Checkbox
                checked={registerData.isprofessor}
                onChange={(e) => setRegisterData({...registerData, isprofessor: e.target.checked})}
                />
            }
            label="교수 여부"
            />
            <Button onClick={handleRegister}>
                회원가입
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

export default RegisterPage