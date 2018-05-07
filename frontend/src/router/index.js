import Vue from "vue"
import Router from "vue-router"

import Courses from "@/components/Courses"
import Tasks from "@/components/Tasks"

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: "/",
      name: "Courses",
      component: Courses
    },
    {
      path: "/class",
      name: "Tasks",
      component: Tasks
    }
  ]
})
