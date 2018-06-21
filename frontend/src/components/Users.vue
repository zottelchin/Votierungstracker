<template>
  <div id="wrapper" class="container">
    <div class="row">
      <div class="column">
        <h2>Neuen Account erstellen</h2>
        <label for="userNew">Dein neuer Account-Key:</label>
        <input type="text" id="userNew" readonly v-model.trim="accountNew">
        <button style="width: 100%" @click="login(accountNew)">Account anlegen</button>
        <p>Schreibe dir den Account-Key am besten auf - du benötigst ihn, um von einem anderen Gerät auf deinen Account zuzugreifen.</p>
      </div>
      <div class="column">
        <h2>Bestehenden Account verwenden</h2>
        <label for="user">Dein Account-Key:</label>
        <input type="text" id="user" v-model.trim="account" v-on:keyup.enter="login(account)">
        <button style="width: 100%" @click="login(account)">Account verwenden</button>
      </div>
    </div>
    <div class="row"><div class="column">
      <hr>
      <label>
        <input type="checkbox" v-model="remember"> Login-Daten auf diesem Computer merken<br>
        <span style="font-weight: normal">Wenn du dies deaktivierst oder wenn dein Browser regelmäßig Cookies löscht, kommst du jedes Mal nur mit deinem Account-Key oder per Lesezeichen auf deine Votierungen.</span>
      </label>
    </div></div>
  </div>
</template>
<script>
  import api from "../api"
  import { generateUsername } from "../username-generator"

  export default {
    data() {
      if (localStorage.getItem("account")) this.$router.push("/" + encodeURIComponent(localStorage.getItem("account")));
      return {
        account: "",
        accountNew: generateUsername(),
        remember: true
      }
    },
    methods: {
      login(account) {
        account = account.replace(/ /g, "-");
        if (this.remember) localStorage.setItem("account", account);
        this.$router.push("/" + encodeURIComponent(account));
      }
    }
  }
</script>
