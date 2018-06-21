<template>
  <div id="wrapper" class="container">
    <div class="row">
      <div class="column">
        <h2>Neuen Account erstellen</h2>
        <label for="userNew">Dein neuer Account-Key:</label>
        <input type="text" id="userNew" readonly v-model.trim="accountNew">
        <label><input type="checkbox" v-model="rememberNew"> Auf diesem Computer merken</label>
        <button style="width: 100%" @click="login(accountNew)">Account anlegen</button>
        <p>Schreibe dir den Account-Key am besten auf - du benötigst ihn, um von einem anderen Gerät oder ohne Lesezeichen auf deinen Account zuzugreifen.</p>
      </div>
      <div class="column">
        <h2>Bestehenden Account verwenden</h2>
        <label for="user">Dein Account-Key:</label>
        <input type="text" id="user" v-model.trim="account" v-on:keyup.enter="login(account)">
        <label><input type="checkbox" v-model="remember"> Auf diesem Computer merken</label>
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
      if (localStorage.getItem("account")) this.$router.push("/" + encodeURIComponent(localStorage.getItem("account")));
      return {
        account: "",
        accountNew: generateUsername(),
        remember: true,
        rememberNew: true
      }
    },
    methods: {
      login(account, isNew) {
        account = account.replace(/ /g, "-");
        if ((isNew && this.rememberNew) || (!isNew && this.remember)) localStorage.setItem("account", account);
        this.$router.push("/" + encodeURIComponent(account));
      }
    }
  }
</script>
