import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import router from "./router";
import { Quasar } from "quasar";
import "quasar/dist/quasar.css";

// require('dotenv').config();
// const WS_BASE_URL = import.meta.env.VITE_WS_BASE_URL;
// const WS_BASE_URL = import.meta.env.VITE_WS_BASE_URL;

const app = createApp(App);

// app.config.globalProperties.$WS_BASE_URL = WS_BASE_URL;
// Use Quasar
app.use(Quasar);

// Use the router
app.use(router);

// Mount the app
app.mount("#app");
