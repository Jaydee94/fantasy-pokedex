<template>
  <v-row>
    <v-col
      v-for="pokemon in pokemons"
      :key="pokemon.id"
      cols="12"
      md="4"
      sm="6"
    >
      <PokemonCard :pokemon="pokemon" />
    </v-col>
  </v-row>
</template>

<script setup>
  import { onMounted, ref } from 'vue'
  import api from '../services/api'
  import PokemonCard from './PokemonCard.vue'

  const pokemons = ref([])

  onMounted(async () => {
    const res = await api.get('/pokemon')
    pokemons.value = res.data
  })
</script>
