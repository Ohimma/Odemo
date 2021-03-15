import { createStore } from 'vuex'

import login from './modules/login'
import layout from './modules/layout'

const store = createStore({
  strict: process.env.NODE_ENV !== 'production',
  modules: {
    login: login,
    layout: layout,
  },
  plugins: []
})

export default store