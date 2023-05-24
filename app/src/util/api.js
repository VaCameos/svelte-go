import { http } from "./axios";

// 创建用户
export const createUserAjax = (data) => http.post("/reg", data);

// 登录
export const loginAjax = (data) => http.post("/login", data);

// 获取所有用户信息
export const getUserList = (data) => http.post("/getUserList", data);

// 删除用户
export const deleteUser = (userid) => http.get("/deleteUser", { userid });

// 获取用户信息
export const getUser = (userid) => http.get("/getUser", { userid });

// 修改用户信息
export const updateUserInfo = (data) => http.post("/updateUser", data);
