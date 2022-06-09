import App from "./App.vue";
import { createClient, clientProviderKey, api } from "./client";
import "./styles/main.scss";
import messages from "@intlify/vite-plugin-vue-i18n/messages";
import humanizeDuration from "humanize-duration";
import moment from "moment";
import { createPinia } from "pinia";
import { createApp } from "vue";
import { createI18n } from "vue-i18n";
import { createRouter, createWebHistory } from "vue-router";
import routes from "~pages";

const getNavigatorLanguage = () => {
  if (navigator.languages && navigator.languages.length) {
    return navigator.languages[0];
  } else {
    return navigator.language;
  }
};

const i18n = createI18n({
  legacy: false,
  locale: getNavigatorLanguage(),
  fallbackLocale: {
    default: ["en"],
  },
  formatFallbackMessages: true,
  //missingWarn: false,
  //fallbackWarn: false,
  messages,
});

const router = createRouter({
  history: createWebHistory("/"),
  routes,
  strict: false,
});

const client = createClient();
const pinia = createPinia();

const app = createApp(App);
app.use(router);
app.use(pinia);
app.use(i18n);
app.mount("#app");
app.provide(clientProviderKey, client);

interface Filters {
  [key: string]: (...value: any[]) => string;
}

declare module "@vue/runtime-core" {
  interface ComponentCustomProperties {
    $filters: Filters;
  }
}

app.config.globalProperties.$filters = {
  formatDateTimeString(value: string) {
    const date = new Date(value);
    return date.toLocaleString();
  },
  formatDateTime(value: Date | undefined) {
    if (!value) {
      return "";
    }
    return value.toLocaleString();
  },
  formatDuration(from: Date, to: Date) {
    const diff = moment(to).diff(from);
    return humanizeDuration(moment.duration(diff).asMilliseconds());
  },
  formatPreservationActionStatus(
    value: api.EnduroPackagePreservationActionsActionResponseBodyStatusEnum
  ) {
    switch (value) {
      case api.EnduroPackagePreservationActionsActionResponseBodyStatusEnum
        .Complete:
        return "bg-success";
      case api.EnduroPackagePreservationActionsActionResponseBodyStatusEnum
        .Failed:
        return "bg-danger";
      case api.EnduroPackagePreservationActionsActionResponseBodyStatusEnum
        .Processing:
        return "bg-warning";
      default:
        return "bg-secondary";
    }
  },
};
