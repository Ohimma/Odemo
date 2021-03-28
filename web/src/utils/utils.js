'use strict'
// import moment from 'moment'
import JSEncrypt from 'jsencrypt';


const Utils = {}

/**
 * @param word（要加密的字符串）
 * @param keyStr（加密串）
 * @return 返回加密后字符串
 */
Utils.encrypt = function (word, keyStr) {
    console.log("word, keyStr = ", word, keyStr)
    // let pubKey = keyStr
    let encryptStr = new JSEncrypt()
    encryptStr.setPublicKey(keyStr)    // 设置 keyStr 公钥
    let  data = encryptStr.encrypt(word.toString())  // 对 word 进行加密
    return data
}




// 存值 存储在用户localstorage 中
Utils.setItem = ((key, value) =>{
  if(typeof value === 'string'){
      localStorage.setItem(key, value);
  }else{
      localStorage.setItem(key, JSON.stringify(value));
  }
})

// 取值 在用户localstorage 中
Utils.getItem = ((key) =>{
  return key&&JSON.parse(localStorage.getItem(key)) || null;
})

// 取值 在用户localstorage 中
Utils.removeItem = ((key) =>{
  localStorage.removeItem(key);
})


// 普通json 转 树形组件要求 json


export default Utils