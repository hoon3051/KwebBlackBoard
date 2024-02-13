import { RegisterParams, LoginParams } from "../../model/user";
import ApiManager from "../index";

export const postRegister = async (RegisterInfo: RegisterParams) => {
  const response = await ApiManager.post("/register", RegisterInfo);
  return response.data;
};

export const postLogin = async (LoginInfo: LoginParams) => {
  const response = await ApiManager.post("/login", LoginInfo);
  return response;
};