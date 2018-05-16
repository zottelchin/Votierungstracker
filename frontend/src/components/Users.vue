<template>
  <div id="wrapper" class="container">
    <div class="row">
      <div class="column">
        <h2>Neuen Account erstellen</h2>
        <label for="userNew">Dein neuer Account-Key:</label>
        <input type="text" id="userNew" readonly v-model.trim="accountNew">
        <button style="width: 100%" @click="login(accountNew)">Account anlegen</button>
        <p>Schreibe dir den Account-Key am besten auf - du benötigst ihn, um von einem anderen Gerät oder ohne Lesezeichen auf deinen Account zuzugreifen.</p>
      </div>
      <div class="column">
        <h2>Bestehenden Account verwenden</h2>
        <label for="user">Dein Account-Key:</label>
        <input type="text" id="user" v-model.trim="account">
        <button style="width: 100%" @click="login(account)">Account verwenden</button>
      </div>
    </div>
  </div>
</template>
<script>
  import api from "../api"
  import { generateUsername } from "../username-generator"

  export default {
    data() {
      return {
        account: "",
        accountNew: generateUsername(),
      }
    },
    methods: {
      login(account) {
        this.$router.push("/" + encodeURIComponent(account.replace(/ /g, "-")));
      }
    }
  }
</script>
