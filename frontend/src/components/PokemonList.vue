<template>
  <div>
    <v-alert
      v-if="alertMsg"
      :type="alertType"
      dense
      class="mb-2"
      border="start"
      transition="fade-transition"
    >
      {{ alertMsg }}
    </v-alert>
    <v-row>
      <v-col
        v-for="pokemon in pokemons"
        :key="pokemon.Name"
        cols="12"
        md="4"
        sm="6"
      >
        <PokemonCard :pokemon="pokemon" @delete="deletePokemon" />
      </v-col>
    </v-row>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import api from '../services/api'
import PokemonCard from './PokemonCard.vue'

const pokemons = ref([])

async function fetchPokemons() {
  const res = await api.get('/pokemon')
  pokemons.value = res.data
}

onMounted(fetchPokemons)

const alertMsg = ref('')
const alertType = ref('success')

function showAlert(msg, type = 'success') {
  alertMsg.value = msg
  alertType.value = type
  setTimeout(() => {
    alertMsg.value = ''
  }, 2000)
}

async function deletePokemon(pokemon) {
  if (!confirm(`Delete ${pokemon.Name}?`)) return
  try {
    await api.delete(`/pokemon/${encodeURIComponent(pokemon.Name)}`)
    pokemons.value = pokemons.value.filter(p => p.Name !== pokemon.Name)
    showAlert('Pokémon erfolgreich gelöscht', 'success')
  } catch (e) {
    showAlert('Failed to delete Pokemon', 'error')
  }
}
</script>
