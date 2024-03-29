import axios from "axios";

const API_URL = "http://localhost:8080"; // TODO: 환경변수로 변경

const ApiManager = axios.create({
  baseURL: API_URL,
  responseType: "json",
  withCredentials: true,
  timeout: 10000,
  headers: { "Content-Type": "application/json" },
});

ApiManager.interceptors.request.use(async (config) => {
    //TO DO: token 받아오기
  const token = localStorage.getItem("_auth");

  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default ApiManager;