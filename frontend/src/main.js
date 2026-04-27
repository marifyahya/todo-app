import { createApp } from "vue";
import { createPinia } from "pinia";
import router from "./router";
import VCalendar from "v-calendar";
import "v-calendar/style.css";
import "./style.css";
import App from "./App.vue";

const app = createApp(App);
const pinia = createPinia();

app.use(pinia);
app.use(router);
app.use(VCalendar, {});
app.mount("#app");
