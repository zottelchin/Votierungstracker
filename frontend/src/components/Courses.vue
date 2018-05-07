<template>
  <div id="wrapper" class="container">
    <h1>Votierungstracker</h1>

    <label for="user"></label>
    <input type="text" id="user" v-model="username">
    <button @click="searchUser">Suchen</button>

    <ul>
      <li v-for="entry in userClasses">
        <a :href=getLink(entry)>{{ entry.class }}</a>
      </li>
    </ul>
  </div>
</template>

<script>
  import api from "../api"
  export default {
    data() {
      return {
        username: "",
        userClasses: []
      }
    },
    methods: {
      searchUser() {
        api.get("/getUser/" + this.username).then(res => {
          this.userClasses = res.body.items ? res.body.items : []
        })
        console.log("yaaa")
      },
      getLink(entry) {
        return "/class?user=" + entry.user + "&class=" + entry.class
      }
    }
  }
</script>
