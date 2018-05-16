<template>
  <div id="wrapper" class="container">
    <div class="row">
      <div class="column">
        <h2>Neuen Account erstellen</h2>
        <label for="userNew">Dein neuer Account-Key:</label>
        <input type="text" id="userNew" readonly v-model="usernameNew">
        <button style="width: 100%" @click="login(usernameNew)">Account anlegen</button>
        <p>Schreibe dir den Account-Key am besten auf - du benötigst ihn, um von einem anderen Gerät oder ohne Lesezeichen auf deinen Account zuzugreifen.</p>
      </div>
      <div class="column">
        <h2>Bestehenden Account verwenden</h2>
        <label for="user">Dein Account-Key:</label>
        <input type="text" id="user" v-model="username" v-on:keyup.enter="login(username)">
        <button style="width: 100%" @click="login(username)">Account verwenden</button>
      </div>
    </div>
  </div>
</template>
<script>
  import api from "../api"
  import { generateUsername } from "../username-generator"
  console.log(generateUsername);

  export default {
    data() {
      return {
        username: "",
        usernameNew: generateUsername(),
      }
    },
    methods: {
      login(username) {
        this.$router.push("/" + encodeURIComponent(username.replace(/ /g, "-")));
      }
    }
  }
</script>
