const state = () => ({
  isCollapse: false,
  tabList: Object.freeze([]),
  menuList: [
      {
          index: '/home',
          title: '首页',
          path: '/home'
      },
      {
          index: 'user_manage',
          title: '用户管理',
          icon: 'el-icon-user-solid',
          child: [
            {
                index: '/user/role',
                title: '角色',
                path: '/user/role'
            },{
                index: '/user/user',
                title: '用户',
                path: '/user/user'
            },{
                index: '/user/acl',
                title: '资源',
                path: '/user/acl'
            }
          ]
      },
      {
          index: 'provider_manage',
          title: '供应商管理',
          icon: 'el-icon-shopping-cart-full',
          child: [
            {
                index: '/provider/define',
                title: '商户管理',
                path: '/provider/define'
            },{
                index: '/provider/account',
                title: '账户管理',
                path: '/provider/account'
            }
          ]
      },
      {
          index: 'product',
          title: '产品管理2',
          path: '/product',
          icon: 'el-icon-s-promotion',
      },
  ]
})
  
const getters = {

}
  
const mutations = {
  updateisCollapse(state) {
    console.log("isCollapse=",state.isCollapse)
    state.isCollapse = !state.isCollapse
  },   
  
  updateTabList(state, item) {
    console.log("vuex updateTabListis", typeof item)
    state.tabList.push({
      title: item.title,
      path: item.path,
    })
    // const tmp = []
    // for (var i=0,l=state.tabList.length; i<l; i++){
    // 	console.log("for ",state.tabList[i].title);
    //   tmp.push({
    //     title: state.tabList[i].title,
    //     path: state.tabList[i].path
    //   })
    // }
    // console.log(tmp, typeof tmp)
    // state.tabList = Array.from(new Set(Array.from(tmp)));
    console.log("state.tabList 1=",typeof state.tabList, state.tabList)
    // state.tabList = this.$utils.unique(state.tabList);
    // set 类似数组，但是不能有重复值，Arry.form 转数据结构为数组
    state.tabList = Array.from(new Set(state.tabList));
  },

  initTabList(state, item) {
    console.log("vuex initTabList", item)
    state.tabList = item
  }   
}
const actions = {
  
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}