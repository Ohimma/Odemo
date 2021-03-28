'use strict'
// import moment from 'moment'
// import JSEncrypt from 'jsencrypt';


const Conver = {}


// 权限列表 数据库原生数据，递归转成树形结构
Conver.convTotree = function (list, id) {
  var child = []
  var i 
  for (i in list) {
      if (list[i].pid == id ) {
          var obj = {
              id: list[i].id,
              label: list[i].name,
              name: list[i].name,
              pid: list[i].pid,
              level: list[i].level,
              url: list[i].url,
              method: list[i].method,
              icon: list[i].icon,
              detail: list[i].detail,
              status: list[i].status,
          }
          obj.children = this.convTotree(list , obj.id)
          child.push(obj)
      }
  }
  return child
}

// 去重操作
Conver.unique = (arr) => { 
  const res = new Map()
  let temp = (arr) => !res.has(arr.index) && res.set(arr.index, 1)
  return arr.filter(temp)
}

export default Conver