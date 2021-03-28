import { createStore } from 'vuex'

import layout from './modules/layout'

const store = createStore({
  strict: process.env.NODE_ENV !== 'production',
  modules: {
    layout: layout,
  },
  plugins: []
})

export default store