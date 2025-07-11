<template>
  <v-card class="mx-auto my-4" max-width="500">
    <!-- Bild -->
    <img
      alt="Pokemon Bild"
      :src="`data:image/png;base64,${pokemon.ImageBase64}`"
      style="display: block; margin: 0 auto; max-width: 100%; max-height: 250px; object-fit: contain;"
    />

    <!-- Titel -->
    <v-card-title class="d-flex flex-column">
      <span class="text-h6">{{ pokemon.Name }}</span>
      <span class="text-subtitle-2 text-grey">{{ pokemon.Category }}</span>
    </v-card-title>

    <!-- Typen -->
    <v-card-subtitle>
      <v-chip
        v-for="type in pokemon.Types"
        :key="type"
        class="ma-1"
        color="deep-orange lighten-2"
        small
        text-color="white"
      >
        {{ type }}
      </v-chip>
    </v-card-subtitle>

    <!-- Toggle Button -->
    <v-card-actions>
      <v-spacer />
      <v-btn
        icon
        variant="text"
        @click="expanded = !expanded"
      >
        <v-icon>{{ expanded ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
      </v-btn>
    </v-card-actions>

    <!-- Collapsible Inhalt -->
    <v-expand-transition>
      <div v-show="expanded">
        <v-divider class="my-2" />

        <v-card-text>
          <v-list dense>
            <v-list-item>
              <v-list-item-title class="font-weight-bold">Fähigkeit:</v-list-item-title>
              <v-list-item-subtitle>{{ pokemon.Ability }}</v-list-item-subtitle>
            </v-list-item>

            <v-list-item>
              <v-list-item-title class="font-weight-bold">Größe:</v-list-item-title>
              <v-list-item-subtitle>{{ pokemon.HeightM }} m</v-list-item-subtitle>
            </v-list-item>

            <v-list-item>
              <v-list-item-title class="font-weight-bold">Gewicht:</v-list-item-title>
              <v-list-item-subtitle>{{ pokemon.WeightKg }} kg</v-list-item-subtitle>
            </v-list-item>

            <v-list-item>
              <v-list-item-title class="font-weight-bold">Pokédex-Eintrag:</v-list-item-title>
              <v-list-item-subtitle>{{ pokemon.PokedexEntry }}</v-list-item-subtitle>
            </v-list-item>

            <v-list-item v-if="pokemon.Appearance">
              <v-list-item-title class="font-weight-bold">Aussehen:</v-list-item-title>
              <v-list-item-subtitle>{{ pokemon.Appearance }}</v-list-item-subtitle>
            </v-list-item>

            <v-list-item v-if="pokemon.Attacks?.length">
              <v-list-item-title class="font-weight-bold">Attacken:</v-list-item-title>
              <v-list-item-subtitle>
                <v-chip
                  v-for="attack in pokemon.Attacks"
                  :key="attack"
                  class="ma-1"
                  color="purple lighten-2"
                  small
                  text-color="white"
                >
                  {{ attack }}
                </v-chip>
              </v-list-item-subtitle>
            </v-list-item>
          </v-list>
        </v-card-text>
      </div>
    </v-expand-transition>
  </v-card>
</template>

<script setup>
  import { ref } from 'vue'

  defineProps(['pokemon'])

  const expanded = ref(false)
</script>
