import { Box, Button, TextField, Typography } from "@mui/material"
import { useState } from "react"
import { SignInParams } from "../../../model/auth"
import { postSignIn } from "../../../api/auth"
import { useNavigate } from "react-router-dom"
import { AxiosError } from "axios"

const SignInPage: React.FC = () =>{

    const [signInData, setSignInData] = useState<SignInParams>({
        username: "",
        password: "",
    })

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) =>{
        const {name, value} = e.target
        setSignInData({...signInData, [name]: value})
    }

    const navigate = useNavigate()

    const handleSignIn = async () => {
        console.log(signInData)
        
        try{
            await postSignIn(signInData)
            navigate("/")
        } catch(error){
            if(error instanceof AxiosError){
                alert(error.response?.data.message ?? "서버에서 에러 발생")
                return
            }
            alert("알 수 없는 에러가 발생했습니다")
        }
    }

    return(
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
            value={signInData.username} 
            label="아이디" 
            onChange={handleChange}
            placeholder="hoon_3051@korea.ac.kr"
            />
            <TextField
            name="password"
            value={signInData.password} 
            label="비밀번호" 
            onChange={handleChange}
            placeholder="string"
            type="password"
            />
            <Button onClick={handleSignIn}>로그인</Button>
        </Box>
    )
}

export default SignInPage