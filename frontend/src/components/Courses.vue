<template>
  <div id="wrapper" class="container">
    <label>Meine Kurse:</label>
    <ul class="course-list">
      <small v-if="!courses || !courses.length">Du hast noch keine Kurse angelegt.</small>
      <li v-else v-for="course in courses">
        <router-link class="button button-outline" :to="getLink(course.course)">{{ course.course }} <small style="padding-left: 4em;" v-show="course.show">{{ course.actualPercent }}% von {{ course.Percent }}%</small></router-link>
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
    <hr>
    <p style="margin-bottom: 0"><strong>Dein Account-Key:</strong> <code>{{account}}</code></p>
    <p v-if="remembered">
      Du bist auf diesem Computer dauerhaft angemeldet.<br>
      <a href="/" onclick="localStorage.removeItem('account')"><span class="square-icon"><i class="fa fa-sign-out-alt"></i></span> <strong>Abmelden</strong></a>
    </p>
    <p v-else>
      <strong>Tipp:</strong> Setze dir ein Lesezeichen auf diese Seite, um nicht jedes Mal deinen Account-Key eingeben zu müssen.<br>
      <a href="#" onclick="localStorage.setItem('account', app.account); return false" @click="remembered = true"><span class="square-icon"><i class="fa fa-sign-in-alt"></i></span> Auf diesem Computer dauerhaft angemeldet bleiben</a><br>
      <a href="/"><span class="square-icon"><i class="fa fa-users"></i></span> Account wechseln</a>
    </p>
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
        remembered: !!localStorage.getItem("account")
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
