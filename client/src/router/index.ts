import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import UploadView from "../views/Upload.vue"
import DownloadView from "../views/Download.vue"

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "Home",
    component: UploadView
  },
  {
    path: "/download/:id",
    name: "Download",
    component: DownloadView
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
