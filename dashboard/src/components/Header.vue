<script setup lang="ts">
import { useI18n } from "vue-i18n";

const { t } = useI18n();
</script>
<template>
  <nav class="navbar navbar-expand-lg bg-light">
    <div class="container-fluid">
      <router-link class="navbar-brand" :to="{ name: 'index' }"
        >Enduro</router-link
      >
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <router-link
              class="nav-link"
              active-class="active"
              :to="{ name: 'packages' }"
              >{{ t("Packages") }}</router-link
            >
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>
