<template>
  <v-card class="mx-auto my-4" max-width="500">
    <img
      alt="Pokemon Bild"
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
        title="Delete Pokemon"
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
              <v-list-item-title class="font-weight-bold">Fähigkeit:</v-list-item-title>
              <v-list-item-subtitle class="no-ellipsis">{{ pokemon.Ability }}</v-list-item-subtitle>
            </v-list-item>

            <v-list-item>
              <v-list-item-title class="font-weight-bold">Größe:</v-list-item-title>
              <v-list-item-subtitle class="no-ellipsis">{{ pokemon.HeightM }} m</v-list-item-subtitle>
            </v-list-item>

            <v-list-item>
              <v-list-item-title class="font-weight-bold">Gewicht:</v-list-item-title>
              <v-list-item-subtitle class="no-ellipsis">{{ pokemon.WeightKg }} kg</v-list-item-subtitle>
            </v-list-item>

            <v-list-item>
              <v-list-item-title class="font-weight-bold">Pokédex-Eintrag:</v-list-item-title>
              <v-list-item-subtitle class="no-ellipsis">{{ pokemon.PokedexEntry }}</v-list-item-subtitle>
            </v-list-item>

            <v-list-item v-if="pokemon.Appearance">
              <v-list-item-title class="font-weight-bold">Aussehen:</v-list-item-title>
              <v-list-item-subtitle class="no-ellipsis">{{ pokemon.Appearance }}</v-list-item-subtitle>
            </v-list-item>

            <v-list-item v-if="pokemon.Attacks?.length">
              <v-list-item-title class="font-weight-bold">Attacken:</v-list-item-title>
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
  Feuer: 'deep-orange darken-2',
  Wasser: 'blue lighten-2',
  Pflanze: 'green lighten-1',
  Elektro: 'yellow darken-2',
  Eis: 'cyan lighten-3',
  Kampf: 'red darken-3',
  Gift: 'purple lighten-1',
  Boden: 'brown lighten-1',
  Flug: 'indigo lighten-3',
  Psycho: 'pink lighten-2',
  Käfer: 'light-green darken-1',
  Gestein: 'amber darken-2',
  Geist: 'deep-purple darken-2',
  Drache: 'indigo darken-3',
  Unlicht: 'grey darken-3',
  Stahl: 'blue-grey lighten-2',
  Fee: 'pink lighten-4',
  Normal: 'grey lighten-1',
}

const typeIconMap = {
  Feuer: 'mdi-fire',
  Wasser: 'mdi-water',
  Pflanze: 'mdi-leaf',
  Elektro: 'mdi-flash',
  Eis: 'mdi-snowflake',
  Kampf: 'mdi-sword',
  Gift: 'mdi-skull',
  Boden: 'mdi-shovel',
  Flug: 'mdi-feather',
  Psycho: 'mdi-eye',
  Käfer: 'mdi-ladybug',
  Gestein: 'mdi-diamond-stone',
  Geist: 'mdi-ghost',
  Drache: 'mdi-dragon',
  Unlicht: 'mdi-weather-night',
  Stahl: 'mdi-hammer-wrench',
  Fee: 'mdi-magic-staff',
  Normal: 'mdi-pokeball',
}

function typeColor(type) {
  return typeColorMap[type] || 'grey lighten-1';
}
function typeIcon(type) {
  return typeIconMap[type] || 'mdi-pokeball';
}

// Attacken können nach Typ oder Name gemappt werden, hier ein Beispiel für einige Attacken
const attackColorMap = {
  Donnerblitz: 'yellow darken-2',
  Flammenwurf: 'deep-orange darken-2',
  Surfer: 'blue lighten-2',
  Solarstrahl: 'green lighten-1',
  Eisstrahl: 'cyan lighten-3',
  // ...weitere Attacken nach Bedarf
}
const attackIconMap = {
  Donnerblitz: 'mdi-flash',
  Flammenwurf: 'mdi-fire',
  Surfer: 'mdi-water',
  Solarstrahl: 'mdi-leaf',
  Eisstrahl: 'mdi-snowflake',
  // ...weitere Attacken nach Bedarf
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
