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
})
  
const getters = {
    doneTodos: state => {
        return state.todos.filter(todos => todos.done)
    }
}
  
const mutations = {
   
    'USER_INFO' (state, obj) {
      state.user_id = obj.user.id
      state.user_name = obj.user.name
      state.user_role_ids = obj.user.role_ids
      localStorage.setItem('user_role_ids', state.user_role_ids)
      localStorage.setItem('user_name', state.user_name)
      localStorage.setItem('user_id', state.user_id)
    },
    'USER_AUTHS' (state, value) {
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
    },
    'AUTH_LIST_TREE' (state, value) {
      state.authListTree = value
    },

   
    
}
const actions = {
    'USER_INFO_ACTION' ({ commit }, obj) {
      commit('USER_INFO', obj)
    },
    'USER_TOKEN_ACTION' ({ commit }, obj) {
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