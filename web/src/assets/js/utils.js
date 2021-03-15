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

// 对列表进行去重
Utils.unique = (arr) => { 
  console.log("entry unique")
  // 根据唯一标识 index 来对数组进行过滤
  // 定义一个Map对象实例，has index存在时则不继续执行，不存在则set
  // 将map对象用数组的filter过滤器过去去重
  const res = new Map()
  let temp = (arr) => !res.has(arr.index) && res.set(arr.index, 1)
  return arr.filter(temp)

  // 简单去重的方式 https://www.cnblogs.com/zsp-1064239893/p/11196501.html
  // indexOf / filter / ES6 Set / sort排序 相邻比较 /includes
  // return Array.from(new Set(arr))
  // var arr1 = [];       // 新建一个数组来存放arr中的值
  //     for(var i=0,len=arr.length;i<len;i++){
  //         if(arr1.indexOf(arr[i]) === -1){
  //             arr1.push(arr[i]);
  //         }
  //     }   
  // return arr1;
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

export default Utils