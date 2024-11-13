import { createRouter, createWebHistory } from "vue-router";
import routes from "./config";


const router = createRouter({
    history: createWebHistory(),
    routes,
    strict: true, // applies to all routes
  })


export default router