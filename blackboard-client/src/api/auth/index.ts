import { SignUpParams, SignInParams } from "../../model/auth";
import ApiManager from "../index";

export const postSignUp = async (signUpInfo: SignUpParams) => {
  const response = await ApiManager.post("/sign_up", signUpInfo);
  return response.data;
};

export const postSignIn = async (signInInfo: SignInParams) => {
  const response = await ApiManager.post("/sign_in", signInInfo);
  return response;
};