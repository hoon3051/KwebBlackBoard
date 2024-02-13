import React, { useState } from "react"
import { useNavigate } from "react-router-dom"
import { createCourse } from "../../../api/course";
import { CourseParams } from "../../../model/course";
import { Box, Button, TextField, Typography, Snackbar } from "@mui/material"
import MuiAlert, { AlertProps } from '@mui/material/Alert';
import axios from "axios"
import Header from "../../../components/header"

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref,
  ) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });


const CreateCouresPage: React.FC = () =>{
    const [courseData, setCourseData] = useState<CourseParams>({
        coursenumber: "",
        coursename: "",
    })

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) =>{
        const {name, value} = e.target
        setCourseData({...courseData, [name]: value})
    }

    const navigate = useNavigate()

    const [message, setMessage] = useState(""); // 메시지 상태
    const [messageType, setMessageType] = useState<"success" | "error">("success"); // 메시지 타입
    const [openSnackbar, setOpenSnackbar] = useState(false); // Snackbar의 열림 상태를 관리


    const handleCreateCourse = async () => {
        
        
        try {
            await createCourse(courseData);
            setMessage("강의 등록이 성공적으로 완료되었습니다!");
            setMessageType("success");
            setOpenSnackbar(true);

            // 스낵바가 충분히 표시된 후 페이지를 이동시킵니다.
            setTimeout(() => {
                navigate("/professor");
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
            <Typography variant="h2" component={"div"}>강의 등록</Typography>
            <TextField
            name="coursenumber"
            required
            autoFocus
            value={courseData.coursenumber}
            label="강의번호"
            onChange={handleChange}
            placeholder="CCCCNNN"
            />
            <TextField 
            name="coursename"
            required
            value={courseData.coursename} 
            label="강의명" 
            onChange={handleChange}
            placeholder="강의명"
            />
            <Button onClick={handleCreateCourse}>
                강의 등록
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

export default CreateCouresPage