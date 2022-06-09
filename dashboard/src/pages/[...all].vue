<script setup lang="ts">
import { useI18n } from "vue-i18n";

const { t } = useI18n();
</script>

<template>
  <h1>{{ t("Not found!") }}</h1>
</template>

<style scoped></style>
