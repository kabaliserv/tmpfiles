import Axios, { AxiosResponse } from "axios";
import config from "@/config/config";

export let Token = "";

const instance = Axios.create({
  baseURL: config.API_URL,
  params: {},
});
instance.interceptors.request.use(
  (config) => {
    config.headers.Authorization = `Bearer ${Token}`;
    return config;
  },
  (err) => {
    return Promise.reject(err);
  }
);

export async function GetMeta(id: string): Promise<any> {
  return await instance.get("/meta/" + id);
}

export async function GetAuth(id: string, password: string): Promise<boolean> {
  let err = false;
  try {
    const response = await instance.post(
      "/auth",
      {
        id: id,
        password: password,
      },
      {
        headers: {
          "content-type": "application/json",
        },
      }
    );
    
    Token = response.data;
  } catch (error) {
    if (error.response.status == 401) {
      err = true;
    }
  }
  return !err;
}
