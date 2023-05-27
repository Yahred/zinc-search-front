import axios from "axios";


const baseURL = 'http://localhost:8080/api';

const axiosInstance = axios.create({
    baseURL,
    timeout: 1000,
});

axiosInstance.interceptors.response.use(
    response => response.data,
    error => {
        console.log(error);
        return Promise.reject(error);
    }
);

export default axiosInstance;