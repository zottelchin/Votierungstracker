<template>
  <div id="app">
    <h1><img src="/static/logo.png"> Votierungstracker</h1>
    <router-view/>
    <footer class="footer">
      <div class="container">
        <div class="row">
          <div class="column">
            <router-link class="" :to='"/"+"impressum"'>Impressum &amp; Datenschutz</router-link>
          </div>
          <div class="column">
            <p class="footext">Logo by <a href="https://icons8.com">Icons8</a></p>
          </div>
          <div class="column">
            <p class="footext"><i class="fas fa-copyright"></i> <a href="https://mo-mar.de">Moritz</a> und <a href="https://github.com/zottelchin">Phillipp</a></p>
          </div>
          <div class="column">
            <p class="footext">
              <a class=" link" href="https://github.com/zottelchin/Votierungstracker"><i class="fab footext fa-github"></i> Repository</a>
            </p>
          </div>
          <div class="column">
            
            <a class=" link" href="https://github.com/zottelchin/Votierungstracker/issues"><i class="fas footext fa-exclamation-circle"></i> Issue erstellen</a>
            
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
import api from "./api";
if (
  window.location.host.match(/^(?:localhost|127\.\d+\.\d+\.\d+|\[::1?\]):8999/)
)
  api.setEndpoint("http://" + window.location.hostname + ":8900/api");
export default {
  name: "app"
};
</script>

<style scoped>
@media screen and (max-width: 40rem) {
  .footer {
    position: relative;
    margin-top: 2rem;  
  }
}
@media screen and (min-width: 40.001rem) {
  .footer {
position: fixed;
  }
}
.footer {
  left: 0;
  bottom: 0;
  width: 100%;
  background-color: #eee8e8;
  color: white;
  text-align: center;
  border-top: 2px solid #9b4dca;
}
.row {
  margin-top: 3px;
  margin-bottom: 3px;
}
.footext {
  color: black;
}
.column p {
  margin-bottom: 0px;
}
</style>
