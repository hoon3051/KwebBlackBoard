import ApiManager from "..";
import { BoardParams } from "../../model/board";

export const getCourseBoards = async (courseId: string) => {
  const response = await ApiManager.get(`/board/${courseId}`);
  return response.data;
};

export const createBoard = async (courseId: string, board: BoardParams) => {
  const response = await ApiManager.post(`/board/${courseId}`, board);
  return response.data;
}

export const getAllBoards = async () => {
  const response = await ApiManager.get("/board");
  return response.data;
}