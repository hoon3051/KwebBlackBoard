import ApiManager from "..";

export const applyCourse = async (courseId: number) => {
    const response = await ApiManager.post(`/apply/${courseId}`);
    return response.data;
}

export const getAppliedStudents = async (courseId: string) => {
    const response = await ApiManager.get(`/apply/${courseId}`);
    return response.data;
}

export const deleteStudent = async (courseId: string, studentId: number) => {
    const response = await ApiManager.delete(`/apply/${courseId}`, {
      data: { Studentid: studentId }
    });
    return response.data;
  };