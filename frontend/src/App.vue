<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
  import api from "./api"
  if (window.location.host.match(/^(?:localhost|127\.\d+\.\d+\.\d+|\[::1?\]):8999/)) api.setEndpoint("http://" + window.location.hostname + ":8900/api")
  export default {
    name: "app"
  }
</script>
