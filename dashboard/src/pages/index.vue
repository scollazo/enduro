<script setup lang="ts">
import Footer from "../components/Footer.vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
</script>

<template>
  {{ t("Welcome to Enduro!") }}

  <ul>
    <li>
      <router-link :to="{ name: 'packages' }">{{ t("Packages") }}</router-link>
    </li>
  </ul>

  <Footer />
</template>
