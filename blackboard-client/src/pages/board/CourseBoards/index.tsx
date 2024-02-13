import React, {useState, useEffect} from 'react';
import Header from '../../../components/header';
import { getCourseBoards } from '../../../api/board';
import { TableContainer, Paper, Table, TableHead, TableRow, TableCell, TableBody, Typography } from '@mui/material';
import axios from 'axios';
import { useParams } from 'react-router-dom';

const CourseBoardsPage: React.FC = () => {
    const [boards, setBoards] = useState([]);

    type RouteParams = {
      courseId: string; // 또는 courseId가 숫자 타입이라면 number로 지정합니다.
    };
    
    const { courseId } = useParams<RouteParams>();
  
    useEffect(() => {
      // API 호출을 통해 강의 데이터를 가져옵니다.
    const fetchBoards = async () => {
        try {
          const response = await getCourseBoards(courseId!);
          setBoards(response.boards);
        } catch (error) {
          if (axios.isAxiosError(error)) {
            // 서버에서 반환된 에러 메시지를 가져와 alert에 표시합니다.
            const message = error.response?.data?.error;
            alert(message);
          } else {
            // Axios 에러가 아닌 경우 일반 에러 메시지를 출력합니다.
            alert('강의 게시물을 불러오는 중 에러가 발생했습니다.');
          }
          console.error('강의 게시물을 불러오는 중 에러가 발생했습니다.', error);
        }
    };

    if(courseId) fetchBoards();
    }, [courseId]);

    function DisplayFormattedContent({ htmlContent }) {
      return <div dangerouslySetInnerHTML={{ __html: htmlContent }} />;
    }
  


  return (
    <>
    <Header />
    <TableContainer component={Paper} sx={{ maxWidth: 1000, margin: 'auto', marginTop: 10 }}>
      <Typography variant="h4" sx={{ textAlign: 'center', marginY: 3 }}>
        강의 게시물 목록
      </Typography>
      <Table aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">제목</TableCell>
            <TableCell align="center">내용</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
            {boards.map((board: { ID: number, Title: string, Content: string }) => (
                <TableRow
                    key={board.ID.toString()}
                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                    <TableCell component="th" scope="row">
                        {board.Title}
                    </TableCell>
                    <TableCell align="right">
                      <DisplayFormattedContent htmlContent={board.Content} />
                    </TableCell>
                </TableRow>
            ))}
        </TableBody>
      </Table>
    </TableContainer>
    </>
  );
}

export default CourseBoardsPage;
