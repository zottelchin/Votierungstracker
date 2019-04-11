import Vue from "vue"
import App from "./App"

import router from "./router"

Vue.config.productionTip = false

window.router = router;

window.app = new Vue({
  el: "#app",
  router,
  render: h => h(App)
})
