<template>
  <div id="wrapper" class="container">
    <label>Meine Kurse:</label>
    <ul class="course-list">
      <small v-if="!courses || !courses.length">Du hast noch keine Kurse angelegt.</small>
      <li v-else v-for="course in courses">
        <router-link class="button button-outline" :to="getLink(course)">{{ course }}</router-link>
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
      setTimeout(() => this.list());
      return {
        account: this.$route.params.account,
        courses: [],
        coursename: "",
      }
    },
    methods: {
      list() {
        api.GET("/courses/" + encodeURIComponent(this.account)).then(res => {
          this.courses = res.body || []
        })
      },
      getLink(course) {
        return "/" + encodeURIComponent(this.account) + "/" + encodeURIComponent(course);
      },
      addCourse() {
        this.$router.push(this.getLink(this.coursename));
      }
    }
  }
</script>
