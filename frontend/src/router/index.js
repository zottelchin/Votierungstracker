import Vue from "vue"
import Router from "vue-router"

import Users from "@/components/Users"
import Courses from "@/components/Courses"
import Tasks from "@/components/Tasks"
import Impressum from "@/components/Impressum"

Vue.use(Router)

export default new Router({
  mode: "history",
  routes: [
    {
      path: "/",
      name: "Users",
      component: Users
    },
    {
      path: "/impressum",
      name: "Impressum",
      component: Impressum
    },
    {
      path: "/:user",
      name: "Courses",
      component: Courses
    },
    {
      path: "/:user/:class",
      name: "Tasks",
      component: Tasks
    }
  ]
})
