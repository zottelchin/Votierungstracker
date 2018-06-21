<template>
  <div id="wrapper" class="container">
    <label>Meine Kurse:</label>
    <ul class="course-list">
      <small v-if="!courses || !courses.length">Du hast noch keine Kurse angelegt.</small>
      <li v-else v-for="course in courses">
        <router-link class="button button-outline courselink" :to="getLink(course.course)">
          <strong>{{ course.course }}</strong>
          <span style="font-weight: normal; margin-left: auto;" v-show="course.minimumType != 'none'">{{ course.actual.toFixed(course.minimumType == "percent" ? 1 : 0).replace(/\.0$/, "") + (course.minimumType == "percent" ? "%" : "") }} von {{ course.minimum + (course.minimumType == "percent" ? "%" : " Punkten") }}</span>
        </router-link>
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
<style>
  .courselink {
    position: relative; display: flex !important; font: inherit; font-size: inherit; text-transform: none; padding: 0.6rem 1rem !important; align-items: center;
    background: linear-gradient(180deg, rgba(253,251,255,1) 0%, rgba(248,246,249,1) 100%);
    letter-spacing: 0;
  }
  .courselink>strong {
    flex-shrink: 1; overflow: hidden; text-overflow: ellipsis;
    font-size: 1.15em;
    margin-right: 0.6rem;
  }
  .courselink>span { flex: none; font-size: 1.05em; }

  .course-list>li:not(:first-child)>.courselink { border-top-left-radius: 0; border-top-right-radius: 0; margin-top: -1px; }
  .course-list>li:not(:last-child)>.courselink { border-bottom-left-radius: 0; border-bottom-right-radius: 0; }
  .courselink:hover, .courselink:active { z-index: 1000; }
  .course-list>li, .courselink { margin-bottom: 0; }

  /* TODO: mobile view: word wrap with long course names */
</style>
