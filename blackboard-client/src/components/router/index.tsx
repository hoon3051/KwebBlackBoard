import { Route, Routes } from "react-router-dom";
import MainPage from "../../pages/main";
import StudentPage from "../../pages/studnet";
import RegisterPage from "../../pages/user/Register";
import LoginPage from "../../pages/user/Login";
import AllCoursesPage from "../../pages/course/AllCourses";
import MyAppliedCoursesPage from "../../pages/course/MyAppliedCourses";
import MyTeachedCoursesPage from "../../pages/course/MyTeachedCourse";
import CourseBoardsPage from "../../pages/board/CourseBoards";
import ProfessorPage from "../../pages/professor";
import CreateCoursePage from "../../pages/course/CreateCourse";
import CreateBoardPage  from "../../pages/board/CreateBoard";
import AppliedStudentsPage from "../../pages/apply/getAppliedStudent";
import AllBoardsPage from "../../pages/board/AllBoards";


/**
 * 어느 url에 어떤 페이지를 보여줄지 정해주는 컴포넌트입니다.
 * Routes 안에 Route 컴포넌트를 넣어서 사용합니다.
 */

const RouteComponent = () => {
  return (
    <Routes>
      <Route path="/" element={<MainPage />} />
      <Route path="/student" element={<StudentPage />} />
      <Route path="/professor" element={<ProfessorPage />} />
      <Route path="/user/register" element={<RegisterPage />} />
      <Route path="/user/login" element={<LoginPage />} />
      <Route path="/course" element={<AllCoursesPage />} />
      <Route path="/course/student" element={<MyAppliedCoursesPage />} />
      <Route path="/course/professor" element={<MyTeachedCoursesPage />} />
      <Route path="/board/:courseId" element={< CourseBoardsPage />} />
      <Route path="/course/create" element={<CreateCoursePage />} />
      <Route path="/board/:courseId/create" element={<CreateBoardPage />} />
      <Route path="/course/:courseId/students" element={<AppliedStudentsPage />} />
      <Route path="/board" element={<AllBoardsPage />} />
    </Routes>
  );
};

export default RouteComponent;