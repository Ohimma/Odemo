'use strict'
// 请求拦截、响应拦截、错误统一处理

import axios from 'axios';
// import qs from 'qs';
// qs.parse()是将URL解析成对象的形式
// qs.stringify()是将对象 序列化成URL的形式，以&进行拼接
import { ElMessage } from 'element-plus'
import store from '@/store/index.js'
import router from '@/router/index.js'
 
// 环境的切换
if (process.env.NODE_ENV == 'development') {    
    axios.defaults.baseURL = '/api';
} else if (process.env.NODE_ENV == 'debug') {    
    axios.defaults.baseURL = '';
} else if (process.env.NODE_ENV == 'production') {    
    axios.defaults.baseURL = 'http://api.123dailu.com/';
}

// 请求超时时间
axios.defaults.timeout = 10000;
 
// post请求头
// axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8';
 
// 请求拦截器
axios.interceptors.request.use(    
    config => {
        // 每次发送请求之前判断是否存在token，如果存在，则统一在http请求的header都加上token，不用每次请求都手动添加了
        // 即使本地存在token，也有可能token是过期的，所以在响应拦截器中要对返回状态进行判断
        // const token = store.state.token;        
        // token && (config.headers.Authorization = token);        
        console.log("进入请求拦截器 request =", config)
        return config;    
    },    
    error => {   
        console.log("进入请求拦截器 err.request =", config)     
        return Promise.error(error);    
    })

// 响应拦截器
axios.interceptors.response.use(   
    response => {    
        console.log("进入响应拦截器 response =", response)     
        if (response.status === 200) {            
            return Promise.resolve(response);        
        } else {            
            return Promise.reject(response);        
        }    
    },
    // 服务器状态码不是200的情况    
    error => {    
        console.log("进入响应拦截器 error.response = ", error.response)    
        if (error.response.status) {            
            switch (error.response.status) {                
                case 401:                    
                    router.replace({                        
                        path: '/login',                        
                        query: { redirect: router.currentRoute.fullPath } 
                    });
                    break;
                // 403 token过期                               
                case 403:     
                    ElMessage.warning('code: 403 token过期喽, 重新登录');                
                                    
                    // 清除token                    
                    localStorage.removeItem('token');                    
                    // store.commit('loginSuccess', null);                    
                    // 跳转登录页面，并将要浏览的页面fullPath传过去，登录成功后跳转需要访问的页面
                    setTimeout(() => {                        
                        router.replace({                            
                            path: '/login',                            
                            query: { 
                                redirect: router.currentRoute.fullPath 
                            }                        
                        });                    
                    }, 1000);                    
                    break; 
                case 404:          
                    ElMessage.warning('code: 404 没找到页面');          
                    break;  
                case 400: 
                    ElMessage.warning('code: 400 提交格式错误');                  
                    break;              
                default:   
                    ElMessage.warning(error.response.data.message); 
                    // this.$message({
                    //     message: error.response.data.message,
                    //     type: 'error'
                    // });
                           
            }            
            return Promise.reject(error.response);        
        }       
    }
);


const http = {}

// get方法，对应get请求 
// @param {String} url [请求的url地址] 
// @param {Object} params [请求时携带的参数] 
http.get = function (url, params, needToken = true) {
    let config = {
      method: 'GET',
      url: url,
      params: params, // params : { key : value}
      headers: {
        'Content-Type': 'application/xxxx-form; charset=UTF-8'
      },
      responseType: 'json'
    }
    if (needToken) {
        config.headers['token'] = store.state.layout.token
        config.headers['user_id'] = store.state.layout.user_id
    }
    return axios(config)
}

http.post = function (url, obj, needToken = true) {
    let config = {
      method: 'POST',
      url: url,
    //   data: qs.stringify(obj), // change obj to url
      data: obj, 
      headers: {
        'Content-Type': 'application/json;charset=UTF-8'
      },
      responseType: 'json'
    }
    if (needToken) {
        // config.headers['Authorization'] = `Bearer ${store.state.layout.token}`
        config.headers['token'] = store.state.layout.token
        config.headers['user_id'] = store.state.layout.user_id
    }
    return new Promise((resolve, reject) => {
        axios(config)
            .then(response => {
                // 判断id或者token是否存在，不存在则报错
                if (response.status == 200 || response.data.code == 200) {
                    resolve(response)
                    this.$message({
                        message: 'post 请求创建成功',
                        type: 'success'
                    });
                } else {
                    reject(response)
                }
            })
            .catch(e => {
                reject(e)
            })
    })
}

http.put = function (url, obj, needToken = true) {
    let config = {
      method: 'PUT',
      url: url,
    //   data: qs.stringify(obj), // change obj to url
      data: obj, // change obj to url
      headers: {
        'Content-Type': 'application/json;charset=UTF-8',
      },
      responseType: 'json'
    }
    if (needToken) {
      config.headers['token'] = store.state.layout.token
      config.headers['user_id'] = store.state.layout.user_id
    }

    return new Promise((resolve, reject) => {
        axios(config)
            .then(response => {
                if (response.data.code === 200 || response.status === 200 ) {
                    resolve(response)
                    this.$message({
                        message: 'put 更新成功',
                        type: 'success'
                    });
                } else {
                    this.$message({
                        type: 'error',
                        message: 'put 更新失败',
                    });
                    reject(response)
                }
            })
            .catch(e => {
                reject(e)
            })
        })
}


http.delete = function (url, params, needToken = true) {
    let config = {
      method: 'DELETE',
      url: url,
      params: params,
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
      },
      responseType: 'json'
    }
   
    if (needToken) {
    //   config.headers['Authorization'] = `Bearer ${store.state.layout.token}`
      config.headers['token'] = store.state.layout.token
      config.headers['user_id'] = store.state.layout.user_id
    }
    return new Promise((resolve, reject) => {
        axios(config)
        .then(response => {
            if (response.data.code === 200 || response.status === 200 ) {
                resolve(response)
                this.$message({
                    message: 'delete 删除成功',
                    type: 'success'
                });
            } else {
                reject(response)
            }
        })
        .catch(e => {
            reject(e)
        })
    })
}

export default http