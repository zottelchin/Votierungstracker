<template>
  <div id="wrapper" class="container">
    <label>Meine Kurse:</label>
    <ul class="course-list">
      <small v-if="!userClasses || !userClasses.length">Du hast noch keine Kurse angelegt.</small>
      <li v-else v-for="entry in userClasses">
        <router-link class="button button-outline" :to="getLink(entry)">{{ entry.class }}</router-link>
      </li>
    </ul>
    <label for="newCourse">Erstelle einen neuen Kurs:</label>
    <div class="manual-row">
      <div style="flex-grow: 1; margin-right: 5px">
        <input type="text" id="newCourse" v-model="coursename" v-on:keyup.enter="addCourse">
      </div>
      <div>
        <button class="slim-button" @click="addCourse">Erstellen</button>
      </div>
    </div>
    <p><strong>Tipp:</strong> Setze dir ein Lesezeichen auf diese Seite, um nicht jedes Mal deinen Account-Key eingeben zu müssen.</p>
  </div>
</template>

<script>
  import api from "../api"
  export default {
    data() {
      setTimeout(() => this.searchUser());
      return {
        username: this.$route.params.user,
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
        return "/" + encodeURIComponent(entry.user) + "/" + encodeURIComponent(entry.class);
      },
      addCourse() {
        console.log({ user: this.username, class: this.coursename });
        this.$router.push(this.getLink({ user: this.username, class: this.coursename }));
      }
    }
  }
</script>
