const state = () => ({
    count: 0,
    todos: [
      { id: 2, done: false },
      { id: 3, done: true },
    ],
    token: localStorage.getItem('token') || '',
    userName: localStorage.getItem('userName') || '',
})
  
const getters = {
    doneTodos: state => {
        return state.todos.filter(todos => todos.done)
    }
}
  
const mutations = {
    increment1 (state, payload) {
        console.log("mutations", payload)
        state.count += payload.amount
    },
    'USER_NAME' (state, userName) {
      state.userName = userName
      localStorage.setItem('userName', state.userName)
    },
    'USER_TOKEN' (state, token) {
      state.token = token
      localStorage.setItem('token', state.token)
    },
    'USER_LOGOUT' (state) {
      state.token = ''
      localStorage.removeItem('token')
    }
}
const actions = {
    increment2 (context, payload) {
        console.log("actions",payload)
        context.commit('increment1', payload)
    },
    'USER_NAME_ACTION' ({ commit }, obj) {
      commit('USER_NAME', obj)
    },
    'USER_TOKEN_ACTION' ({ commit }, obj) {
      commit('USER_TOKEN', obj)
    },
    'USER_LOGOUT_ACTION' ({ commit }) {
      commit('USER_LOGOUT')
    }
}

export default {
//   namespaced: true,
  state,
  getters,
  actions,
  mutations
}