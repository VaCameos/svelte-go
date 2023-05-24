import axios from "axios";

const instance = axios.create({
  baseURL: "/api",
  timeout: 5000,
});

export const http = {
  get(url, params) {
    return instance({
      method: "get",
      url,
      params,
    });
  },
  post(url, data) {
    return instance({
      method: "post",
      url,
      data,
    });
  },
};
