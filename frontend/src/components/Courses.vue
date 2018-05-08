<template>
  <div id="wrapper" class="container">
    <h1>Votierungstracker</h1>

    <label for="user">Wähle einen User:</label>
    <input type="text" id="user" v-model="username">
    <button @click="searchUser">Suchen</button>

    <ul>
      <li v-for="entry in userClasses">
        <a :href=getLink(entry)>{{ entry.class }}</a>
      </li>
    </ul>
    <label for="newCourse">Erstelle einen neuen Kurs</label>
    <input type="text" id="newCourse" v-model="coursename">
    <button @click="addCourse">Erstellen</button>
  </div>
</template>

<script>
  import api from "../api"
  export default {
    data() {
      return {
        username: "",
        userClasses: [],
        coursename: "",
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
      },
      addCourse() {
        window.location.href = "/class?user=" + this.username + "&class=" + this.coursename;
      }
    }
  }
</script>
