import ApiManager from "..";
import { CourseParams } from "../../model/course";

export const getAllCourses = async () => {
  const response = await ApiManager.get("/course");
  return response.data;
}


export const getMyAppliedCourses = async () => {
  const response = await ApiManager.get("/course/student");
  return response.data;
}

export const getMyTeachedCourses = async () => {
  const response = await ApiManager.get("/course/professor");
  return response.data;
}

export const createCourse = async (course: CourseParams) => {
  const response = await ApiManager.post("/course", course);
  return response.data;
}

