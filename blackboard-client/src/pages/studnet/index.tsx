import Header from "../../components/header"
import { useNavigate } from "react-router-dom"

const StudentPage: React.FC = () =>{
    const navigate = useNavigate(); // useNavigate 훅을 사용해 navigate 함수 초기화

    // 전체 강의 조회 페이지로 이동하는 함수
    const goToAllCourses = () => {
        navigate('/course');
    };

    // 내 수강 강의 페이지로 이동하는 함수
    const goToMyCourses = () => {
        navigate('/course/student');
    };

    const goToAllBoards = () => {
        navigate('/board');
    }

    return (
    <>
        <Header />
        <div style={{
                display: "flex",
                justifyContent: "center", // 수평 중앙 정렬
                alignItems: "center", // 수직 중앙 정렬
                height: "calc(100vh - 64px)", // 헤더의 높이를 뺀 나머지 높이
                width: "100%", // 너비 100%
                textAlign: "center", // 텍스트 중앙 정렬
                fontSize: "4rem", // 폰트 크기를 크게
        }}>
            {/* 전체 강의 조회 및 수강신청 버튼 */}
            <button onClick={goToAllCourses} style={{ fontSize: "2rem", margin: "20px" }}>
                전체 강의 조회 및 수강신청
            </button>
            {/* 내 수강 강의 조회 버튼 */}
            <button onClick={goToMyCourses} style={{ fontSize: "2rem", margin: "20px" }}>
                나의 수강 강의 목록
            </button>
             {/* 최근 게시물 목록 버튼 */}
             <button onClick={goToAllBoards} style={{ fontSize: "2rem", margin: "20px" }}>
                최근 게시물 목록
            </button>
        </div>
    </>
    )
}

export default StudentPage