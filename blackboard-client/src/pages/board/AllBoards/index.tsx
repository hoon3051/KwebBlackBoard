import React, {useState, useEffect} from 'react';
import Header from '../../../components/header';
import { getAllBoards } from '../../../api/board';
import { TableContainer, Paper, Table, TableHead, TableRow, TableCell, TableBody, Typography } from '@mui/material';
import axios from 'axios';

const AllBoardsPage: React.FC = () => {
    const [boards, setBoards] = useState([]);
  
    useEffect(() => {
      // API 호출을 통해 게시물 데이터를 가져옵니다.
    const fetchBoards = async () => {
        try {
            const response = await getAllBoards();
            setBoards(response.boards);
        } catch (error) {
          if (axios.isAxiosError(error)) {
            // 서버에서 반환된 에러 메시지를 가져와 alert에 표시합니다.
            const message = error.response?.data?.error;
            alert(message);
          } else {
            // Axios 에러가 아닌 경우 일반 에러 메시지를 출력합니다.
            alert('최근 게시물들을 불러오는 중 에러가 발생했습니다.1');
          }
          console.error('최근 게시물들을 불러오는 중 에러가 발생했습니다.2', error);
        }
    };

    fetchBoards();
    }, []);

    function DisplayFormattedContent({ htmlContent }) {
      return <div dangerouslySetInnerHTML={{ __html: htmlContent }} />;
    }
  


  return (
    <>
    <Header />
    <TableContainer component={Paper} sx={{ maxWidth: 800, margin: 'auto', marginTop: 10 }}>
      <Typography variant="h4" sx={{ textAlign: 'center', marginY: 3 }}>
        최근 게시물 목록
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
                    <TableCell align="center" component="th" scope="row">
                        {board.Title}
                    </TableCell>
                    <TableCell align="center">
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

export default AllBoardsPage;
