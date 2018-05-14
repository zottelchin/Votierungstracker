<template>
  <div id="app">
    <h1><img src="/static/logo.png"> Votierungstracker</h1>
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
