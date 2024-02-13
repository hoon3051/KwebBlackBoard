import React, {useState, useEffect} from 'react';
import Header from '../../../components/header';
import { getMyTeachedCourses } from '../../../api/course';
import { TableContainer, Paper, Table, TableHead, TableRow, TableCell, TableBody, Typography } from '@mui/material';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const MyTeachedCoursesPage: React.FC = () => {
    const [courses, setCourses] = useState([]);
    const navigate = useNavigate();
  
    useEffect(() => {
      // API 호출을 통해 강의 데이터를 가져옵니다.
    const fetchCourses = async () => {
        try {
            const data = await getMyTeachedCourses();
            // 서버로부터 받은 데이터를 Coursenumber 기준으로 오름차순 정렬
            const sortedCourses = data.courses.sort((a: { Coursenumber: string }, b: { Coursenumber: string }) => {
                return a.Coursenumber.localeCompare(b.Coursenumber);
            });
            setCourses(sortedCourses);
        } catch (error) {
          if (axios.isAxiosError(error)) {
            // 서버에서 반환된 에러 메시지를 가져와 alert에 표시합니다.
            const message = error.response?.data?.error;
            alert(message);
          } else {
            // Axios 에러가 아닌 경우 일반 에러 메시지를 출력합니다.
            alert('내 등록 강의를 불러오는 중 에러가 발생했습니다.');
          }
          console.error('내 등록 강의를 불러오는 중 에러가 발생했습니다.', error);
        }
    };

    fetchCourses();
    }, []);

    const handleBoardsClick = (courseId: number) => {
      navigate(`/board/${courseId}`);
  };

  const handleCreateBoardClick = (courseId: number) => {
    navigate(`/board/${courseId}/create`);
};

    const handleCourseStudentClick = (courseId: number) => {
      navigate(`/course/${courseId}/students`);
    };


  return (
    <>
    <Header />
    <TableContainer component={Paper} sx={{ maxWidth: 800, margin: 'auto', marginTop: 10 }}>
      <Typography variant="h4" sx={{ textAlign: 'center', marginY: 3 }}>
        나의 등록 강의 목록
      </Typography>
      <Table aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">강의 번호</TableCell>
            <TableCell align="center">강의명</TableCell>
            <TableCell align="center"></TableCell>
            <TableCell align="center"></TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
            {courses.map((course: { ID: number, Coursenumber: string, Coursename: string }) => (
                <TableRow
                    key={course.ID.toString()}
                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                    hover
                    onClick={() => handleBoardsClick(course.ID)}
                    style={{ cursor: 'pointer' }}
                >
                    <TableCell component="th" scope="row">{course.Coursenumber}</TableCell>
                    <TableCell align="center">{course.Coursename}</TableCell>
                    <TableCell align="right">
                        <button onClick={(event) => {
                          event.stopPropagation(); // TableRow의 onClick 이벤트가 실행되지 않도록 합니다.
                          handleCreateBoardClick(course.ID)
                        }}>게시물 등록</button>
                    </TableCell>
                    <TableCell align="right">
                        <button onClick={(event) => {
                          event.stopPropagation(); // TableRow의 onClick 이벤트가 실행되지 않도록 합니다.
                          handleCourseStudentClick(course.ID)
                        }}>학생 관리</button>
                    </TableCell>
                </TableRow>
            ))}
        </TableBody>
      </Table>
    </TableContainer>
    </>
  );
}

export default MyTeachedCoursesPage;
