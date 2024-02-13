import React, { useState } from "react"
import { useNavigate, useParams } from "react-router-dom"
import { createBoard } from "../../../api/board";
import { BoardParams } from "../../../model/board";
import { Box, Button, TextField, Typography, Snackbar } from "@mui/material"
import MuiAlert, { AlertProps } from '@mui/material/Alert';
import axios from "axios"
import Header from "../../../components/header"
import ReactQuill from 'react-quill';
import 'react-quill/dist/quill.snow.css'; // Quill editor의 스타일을 import합니다.



const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref,
  ) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });


const CreateBoardPage: React.FC = () =>{
    const [boardData, setBoardData] = useState<BoardParams>({
        title: "",
        content: "",
    })

    type RouteParams = {
        courseId: string; // 또는 courseId가 숫자 타입이라면 number로 지정합니다.
      };
      
    const { courseId } = useParams<RouteParams>();

    const handleChange = (value: string, path: string) => {
        setBoardData({ ...boardData, [path]: value });
      };

    const navigate = useNavigate()

    const [message, setMessage] = useState(""); // 메시지 상태
    const [messageType, setMessageType] = useState<"success" | "error">("success"); // 메시지 타입
    const [openSnackbar, setOpenSnackbar] = useState(false); // Snackbar의 열림 상태를 관리


    const handleCreateBoard = async () => {
        
        try {

            await createBoard(courseId!, boardData);
            setMessage("게시물 등록이 성공적으로 완료되었습니다!");
            setMessageType("success");
            setOpenSnackbar(true);

            // 스낵바가 충분히 표시된 후 페이지를 이동시킵니다.
            setTimeout(() => {
                navigate(`/board/${courseId}`);
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
            <Typography variant="h2">게시물 등록</Typography>
        <TextField name="title" required autoFocus value={boardData.title} label="제목" onChange={(e) => handleChange(e.target.value, "title")} placeholder="제목" />
        <ReactQuill theme="snow" value={boardData.content} onChange={(content) => handleChange(content, "content")} />
        <Button onClick={handleCreateBoard}>게시물 등록</Button>
            <Snackbar open={openSnackbar} autoHideDuration={6000} onClose={handleCloseSnackbar}>
                <Alert onClose={handleCloseSnackbar} severity={messageType} sx={{ width: '100%' }}>
                    {message}
                </Alert>
            </Snackbar>
        </Box>
        </>
    )

}

export default CreateBoardPage