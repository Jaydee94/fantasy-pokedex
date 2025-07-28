<template>
  <v-app>
    <AdminSetPassword v-if="showSetPassword" />
    <v-app-bar app color="#356abc" dark height="80" style="box-shadow: 0 2px 8px rgba(0,0,0,0.15); background: linear-gradient(90deg, #356abc 0%, #3d2376 100%);">
      <div style="position: absolute; left: 0; right: 0; top: 0; bottom: 0; display: flex; align-items: center; justify-content: center; pointer-events: none;">
        <v-img
          max-height="130"
          src="@/assets/fantasy_pokedex.png"
          style="max-width: 200px; margin: 0 auto; pointer-events: auto;"
        />
      </div>
      <div style="margin-left: auto; display: flex; align-items: center; z-index: 1; gap: 1rem;">
        <v-btn text to="/">{{ $t('app.home') }}</v-btn>
        <v-btn text to="/create">{{ $t('app.create') }}</v-btn>
        <v-btn v-if="isAdmin" text @click="logout">{{ $t('app.logout') }}</v-btn>
        <AdminLogin v-else />
        <v-select
          v-model="currentLocale"
          dense
          hide-details
          item-title="label"
          item-value="code"
          :items="locales"
          style="width: 120px; min-width: 100px;"
          variant="outlined"
        />
      </div>
    </v-app-bar>

    <v-main>
      <v-container class="py-5" fluid>
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import AdminLogin from './components/AdminLogin.vue'
import AdminSetPassword from './components/AdminSetPassword.vue'
import { useAppStore } from './stores/app'
import { computed, ref, onMounted, watch } from 'vue'
import { isAdminPasswordSet } from './services/admin'
import { useI18n } from 'vue-i18n'

const appStore = useAppStore()
const isAdmin = computed(() => appStore.isAdmin)
const showSetPassword = ref(false)

const { locale } = useI18n()
const locales = [
  { code: 'en', label: 'English' },
  { code: 'de', label: 'Deutsch' }
]
const currentLocale = ref(locale.value)

watch(currentLocale, (val) => {
  locale.value = val
})

onMounted(async () => {
  try {
    const set = await isAdminPasswordSet()
    showSetPassword.value = !set
  } catch (e) {
    showSetPassword.value = false
  }
})

function logout() {
  appStore.logout()
}
</script>
