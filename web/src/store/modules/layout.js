const state = () => ({
    token: localStorage.getItem('token') || '',
    name: localStorage.getItem('name') || '',
    id: localStorage.getItem('id') || 0,
    
    user_id: localStorage.getItem('user_id') || 0,
    user_name: localStorage.getItem('user_name') || '',
    user_role_ids: localStorage.getItem('user_role_ids') || '',

    user_auths: localStorage.getItem('user_auths') || [],
    authListInit:[],
    authListTree:[],
    tabList:[],
})
  
const getters = {
    doneTodos: state => {
        return state.todos.filter(todos => todos.done)
    }
}
  
const mutations = {
   
    'USER_INFO' (state, obj) {
      console.log("enter store USER_INFO = ", obj)
      state.user_id = obj.user.id
      state.user_name = obj.user.name
      state.user_role_ids = obj.user.role_ids
      localStorage.setItem('user_role_ids', state.user_role_ids)
      localStorage.setItem('user_name', state.user_name)
      localStorage.setItem('user_id', state.user_id)
    },
    'USER_AUTHS' (state, value) {
      console.log("enter store layout USER_INFO_ACTION")
      state.user_auths = value
      localStorage.setItem('user_auths', state.user_auths)
    },
    'USER_TOKEN' (state, token) {
      state.token = token
      localStorage.setItem('token', state.token)
    },
    'USER_LOGOUT' (state) {
      state.token = ''
      localStorage.removeItem('token')
    },
    'AUTH_LIST' (state, value) {
      state.authListInit = value
      console.log("enter store AUTH_LIST = ", state.authListInit )
    },
    'AUTH_LIST_TREE' (state, value) {
      state.authListTree = value
      console.log("enter store AUTH_LIST_TREE = ", state.authListTree )
    },

    // updateisCollapse(state) {
    //   state.isCollapse = !state.isCollapse
    // },   
    initTabList(state, item) {
      state.tabList = item
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
}
const actions = {
    'USER_INFO_ACTION' ({ commit }, obj) {
      console.log("enter store layout USER_INFO_ACTION")
      commit('USER_INFO', obj)
    },
    'USER_TOKEN_ACTION' ({ commit }, obj) {
      console.log("enter store layout USER_TOKEN_ACTION")
      commit('USER_TOKEN', obj)
    },

    'USER_LOGOUT_ACTION' ({ commit }) {
      commit('USER_LOGOUT')
    },

    'USER_AUTHS_ACTION' ({ commit }, value) {
      commit('USER_AUTHS', value)
    },
    'AUTH_LIST_ACTION' ({ commit }, value) {
      commit('AUTH_LIST', value)
    },
    'AUTH_LIST_TREE_ACTION' ({ commit }, value) {
      commit('AUTH_LIST_TREE', value)
    },

}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}