<template>
  <v-card class="mx-auto my-4" max-width="500">
    <img
      alt="Pokemon image"
      :src="`data:image/png;base64,${pokemon.ImageBase64}`"
      style="display: block; margin: 0 auto; max-width: 100%; max-height: 250px; object-fit: contain;"
    />

    <v-card-title class="d-flex flex-column">
      <span class="text-h6">{{ pokemon.Name }}</span>
      <span class="text-subtitle-2 text-grey">{{ pokemon.Category }}</span>
    </v-card-title>

    <v-card-subtitle>
      <v-chip
        v-for="type in pokemon.Types"
        :key="type"
        class="ma-1"
        :color="typeColor(type)"
        small
        text-color="white"
      >
        <v-icon start size="small">{{ typeIcon(type) }}</v-icon>
        {{ type }}
      </v-chip>
    </v-card-subtitle>

    <v-card-actions>
      <v-spacer />
      <v-btn
        icon
        variant="text"
        @click="expanded = !expanded"
      >
        <v-icon>{{ expanded ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
      </v-btn>
      <v-btn
        v-if="isAdmin"
        icon
        color="red"
        variant="text"
        @click="onDelete"
        title="Delete Pokémon"
      >
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </v-card-actions>

    <v-expand-transition>
      <div v-show="expanded">
        <v-divider class="my-2" />

        <v-card-text>
          <v-list dense>
            <v-list-item>
              <v-list-item-title class="font-weight-bold">Ability:</v-list-item-title>
              <v-list-item-subtitle class="no-ellipsis">{{ pokemon.Ability }}</v-list-item-subtitle>
            </v-list-item>

            <v-list-item>
              <v-list-item-title class="font-weight-bold">Height:</v-list-item-title>
              <v-list-item-subtitle class="no-ellipsis">{{ pokemon.HeightM }} m</v-list-item-subtitle>
            </v-list-item>

            <v-list-item>
              <v-list-item-title class="font-weight-bold">Weight:</v-list-item-title>
              <v-list-item-subtitle class="no-ellipsis">{{ pokemon.WeightKg }} kg</v-list-item-subtitle>
            </v-list-item>

            <v-list-item>
              <v-list-item-title class="font-weight-bold">Pokédex Entry:</v-list-item-title>
              <v-list-item-subtitle class="no-ellipsis">{{ pokemon.PokedexEntry }}</v-list-item-subtitle>
            </v-list-item>

            <v-list-item v-if="pokemon.Appearance">
              <v-list-item-title class="font-weight-bold">Appearance:</v-list-item-title>
              <v-list-item-subtitle class="no-ellipsis">{{ pokemon.Appearance }}</v-list-item-subtitle>
            </v-list-item>

            <v-list-item v-if="pokemon.Attacks?.length">
              <v-list-item-title class="font-weight-bold">Attacks:</v-list-item-title>
              <v-list-item-subtitle class="no-ellipsis">
                <v-chip
                  v-for="attack in pokemon.Attacks"
                  :key="attack"
                  class="ma-1 no-ellipsis"
                  :color="attackColor(attack)"
                  small
                  text-color="white"
                  style="white-space: normal; overflow: visible; text-overflow: unset;"
                >
                  <v-icon start size="small">{{ attackIcon(attack) }}</v-icon>
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
import { ref, computed } from 'vue'
import { useAppStore } from '../stores/app'

const typeColorMap = {
  Fire: 'deep-orange darken-2',
  Water: 'blue lighten-2',
  Grass: 'green lighten-1',
  Electric: 'yellow darken-2',
  Ice: 'cyan lighten-3',
  Fighting: 'red darken-3',
  Poison: 'purple lighten-1',
  Ground: 'brown lighten-1',
  Flying: 'indigo lighten-3',
  Psychic: 'pink lighten-2',
  Bug: 'light-green darken-1',
  Rock: 'amber darken-2',
  Ghost: 'deep-purple darken-2',
  Dragon: 'indigo darken-3',
  Dark: 'grey darken-3',
  Steel: 'blue-grey lighten-2',
  Fairy: 'pink lighten-4',
  Normal: 'grey lighten-1',
}

const typeIconMap = {
  Fire: 'mdi-fire',
  Water: 'mdi-water',
  Grass: 'mdi-leaf',
  Electric: 'mdi-flash',
  Ice: 'mdi-snowflake',
  Fighting: 'mdi-sword',
  Poison: 'mdi-skull',
  Ground: 'mdi-shovel',
  Flying: 'mdi-feather',
  Psychic: 'mdi-eye',
  Bug: 'mdi-ladybug',
  Rock: 'mdi-diamond-stone',
  Ghost: 'mdi-ghost',
  Dragon: 'mdi-dragon',
  Dark: 'mdi-weather-night',
  Steel: 'mdi-hammer-wrench',
  Fairy: 'mdi-magic-staff',
  Normal: 'mdi-pokeball',
}

function typeColor(type) {
  return typeColorMap[type] || 'grey lighten-1';
}
function typeIcon(type) {
  return typeIconMap[type] || 'mdi-pokeball';
}

// Example for some attacks, add more as needed
const attackColorMap = {
  Thunderbolt: 'yellow darken-2',
  Flamethrower: 'deep-orange darken-2',
  Surf: 'blue lighten-2',
  SolarBeam: 'green lighten-1',
  IceBeam: 'cyan lighten-3',
  // ...add more attacks as needed
}
const attackIconMap = {
  Thunderbolt: 'mdi-flash',
  Flamethrower: 'mdi-fire',
  Surf: 'mdi-water',
  SolarBeam: 'mdi-leaf',
  IceBeam: 'mdi-snowflake',
  // ...add more attacks as needed
}
function attackColor(attack) {
  return attackColorMap[attack] || 'purple lighten-2';
}
function attackIcon(attack) {
  return attackIconMap[attack] || 'mdi-sword';
}

const props = defineProps(['pokemon'])
const emit = defineEmits(['delete'])
const expanded = ref(false)
const appStore = useAppStore()
const isAdmin = computed(() => appStore.isAdmin)

function onDelete() {
  emit('delete', props.pokemon)
}
</script>

<style scoped>
.no-ellipsis {
  white-space: normal !important;
  overflow: visible !important;
  text-overflow: unset !important;
  max-height: none !important;
  display: block !important;
  word-break: break-word !important;
}
</style>
