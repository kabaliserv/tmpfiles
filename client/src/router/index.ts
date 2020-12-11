import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import UploadView from "../views/Upload.vue"
import FileView from "../views/File.vue"

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "Home",
    component: UploadView
  },
  {
    path: "/file/:id",
    name: "File",
    component: FileView
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
