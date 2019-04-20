import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './registerServiceWorker'
import VueApexCharts from 'vue-apexcharts'

Vue.use(VueApexCharts)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
