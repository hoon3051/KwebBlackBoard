import React, {useState, useEffect} from 'react';
import Header from '../../../components/header';
import { getAppliedStudents } from '../../../api/apply';
import { TableContainer, Paper, Table, TableHead, TableRow, TableCell, TableBody, Typography } from '@mui/material';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import { deleteStudent } from '../../../api/apply';

const AppliedStudentsPage: React.FC = () => {
    const [students, setStudents] = useState([]);

    type RouteParams = {
      courseId: string; // 또는 courseId가 숫자 타입이라면 number로 지정합니다.
    };
    
    const { courseId } = useParams<RouteParams>();
  
    useEffect(() => {
      // API 호출을 통해 수강 학생  데이터를 가져옵니다.
    const fetchStudents = async () => {
        try {
          const response = await getAppliedStudents(courseId!);
          setStudents(response.users);
        } catch (error) {
          if (axios.isAxiosError(error)) {
            // 서버에서 반환된 에러 메시지를 가져와 alert에 표시합니다.
            const message = error.response?.data?.error;
            alert(message);
          } else {
            // Axios 에러가 아닌 경우 일반 에러 메시지를 출력합니다.
            alert('수강 학생들을 불러오는 중 에러가 발생했습니다.');
          }
          console.error('수강 학생들을 불러오는 중 에러가 발생했습니다.', error);
        }
    };

    if(courseId) fetchStudents();
    }, [courseId]);

    const handleDeleteStudent = async (courseId: string, studentId: number) => {
      try { 
        await deleteStudent(courseId, studentId);
        alert('수강 취소가 완료되었습니다.');
        window.location.reload();
      } catch (error) {
        if (axios.isAxiosError(error)) {
          // 서버에서 반환된 에러 메시지를 가져와 alert에 표시합니다.
          const message = error.response?.data?.error;
          alert(message);
        } else {
          // Axios 에러가 아닌 경우 일반 에러 메시지를 출력합니다.
          alert('수강 취소 중 에러가 발생했습니다.');
        }
        console.error('수강 취소 중 에러가 발생했습니다.', error);
      }
    }


  return (
    <>
    <Header />
    <TableContainer component={Paper} sx={{ maxWidth: 1000, margin: 'auto', marginTop: 10 }}>
      <Typography variant="h4" sx={{ textAlign: 'center', marginY: 3 }}>
        수강 학생 목록
      </Typography>
      <Table aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">이름</TableCell>
            <TableCell align="center">학번</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
            {students.map((users: { ID: number, Displayname: string, Studentnumber: string }) => (
                <TableRow
                    key={users.ID.toString()}
                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                    <TableCell align="center" component="th" scope="row">
                        {users.Displayname}
                    </TableCell>
                    <TableCell align="center">{users.Studentnumber}</TableCell>
                    <TableCell align="center">
                        <button onClick={() => handleDeleteStudent(courseId!, users.ID)}>수강 취소</button>
                    </TableCell>
                </TableRow>
            ))}
        </TableBody>
      </Table>
    </TableContainer>
    </>
  );
}

export default AppliedStudentsPage;
