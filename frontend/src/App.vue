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
          <div class="column right">
            <p class="footext">Logo by <a href="https://icons8.com">Icons8</a></p>
          </div>
          <div class="column copyright">
            <p class="footext"><i class="fas fa-copyright"></i> <a href="https://mo-mar.de">Moritz</a> und <a href="https://github.com/zottelchin">Phillipp</a></p>
          </div>
          <div class="column">
            <p class="footext">
              <a class=" link" href="https://github.com/zottelchin/Votierungstracker"><i class="fab footext fa-github"></i> Repository</a>
            </p>
          </div>
          <div class="column right">
            <a class=" link" href="https://github.com/zottelchin/Votierungstracker/issues"><i class="fas footext fa-exclamation-circle"></i> Fehler melden</a>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
  import api from "./api";
  if (window.location.host.match(/^(?:localhost|127\.\d+\.\d+\.\d+|\[::1?\]):8999/)) {
    api.setEndpoint("http://" + window.location.hostname + ":8900/api");
  }

  export default {
    name: "app"
  };
</script>

<style scoped>
.footer {
  left: 0;
  bottom: 0;
  width: 100%;
  position: fixed;
  background-color: #f2f2f2;
  max-width: 100%;
  margin: 0 auto;
  color: white;
  text-align: center;
  border-top: 2px solid #9b4dca;
  margin-top: 2rem;
  padding: 0.8rem 0;
  white-space: nowrap;
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
@media screen and (max-width: 800px) {
  .footer {
    position: relative;
    max-width: 800px;
    background-color: transparent;
    padding: 1.5rem 0;
    white-space: normal;

  }
  .footer .row {
    display: grid;
    grid-template-columns: 3fr 2fr;
    grid-gap: 1rem;
    text-align: left;
    margin: 0 !important; width: auto;
  }
  .footer .column {
    padding: 0 0 0.5rem;
  }
  .copyright {
    grid-column: span 2;
  }
  .footer .right { text-align: right; }
}
</style>
