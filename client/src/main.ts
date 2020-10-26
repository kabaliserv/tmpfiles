import Vue from 'vue'
import App from './App.vue'
import router from './router'
import VueRx from 'vue-rx'

// Filters
import './filters/size'

// Plugins
import vuetify from './plugins/vuetify';

// Style


Vue.use(VueRx)
Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
