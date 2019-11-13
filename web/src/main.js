import Vue from 'vue'
import ECharts from 'vue-echarts'
import 'echarts/lib/chart/bar'
import 'echarts/lib/component/visualMap'
import 'echarts-gl'

import App from './App.vue'
import router from './router'
import './registerServiceWorker'

Vue.component('v-chart', ECharts)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
