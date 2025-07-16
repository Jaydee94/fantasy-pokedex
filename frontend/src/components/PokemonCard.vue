<template>
  <div style="position: relative;">
    <v-card class="mx-auto my-4" max-width="500">
    <img
      alt="Pokemon image"
      :src="`data:image/png;base64,${pokemon.ImageBase64}`"
      class="pokemon-image-preview"
      @click="showImageModal = true"
    />

    <v-card-title class="d-flex flex-column">
      <span class="text-h6" style="text-align: center; width: 100%; display: block;">{{ pokemon.Name }}</span>
      <span class="text-subtitle-2 text-grey">{{ $t('card.category') }}: {{ pokemon.Category }}</span>
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

    <v-card-actions style="justify-content: center;">
      <v-btn
        icon
        variant="text"
        @click="expanded = !expanded"
        style="margin: 0 8px;"
      >
        <v-icon>{{ expanded ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
      </v-btn>
      <v-btn
        icon
        color="primary"
        variant="text"
        @click="showStatsModal = true"
        :title="$t('card.showStats')"
      >
        <v-icon>mdi-radar</v-icon>
      </v-btn>
      <v-btn
        v-if="isAdmin"
        icon
        color="red"
        variant="text"
        @click="onDelete"
        :title="$t('card.deleteTitle')"
        style="margin: 0 8px;"
      >
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </v-card-actions>

    <v-expand-transition>
      <template v-if="expanded">
        <div>
          <v-divider class="my-2" />
          <v-card-text>
            <v-list dense>
              <v-list-item>
                <v-list-item-title class="font-weight-bold">{{ $t('card.ability') }}:</v-list-item-title>
                <v-list-item-subtitle class="no-ellipsis">{{ pokemon.Ability }}</v-list-item-subtitle>
              </v-list-item>
              <v-list-item>
                <v-list-item-title class="font-weight-bold">{{ $t('card.height') }}:</v-list-item-title>
                <v-list-item-subtitle class="no-ellipsis">{{ pokemon.HeightM }} m</v-list-item-subtitle>
              </v-list-item>
              <v-list-item>
                <v-list-item-title class="font-weight-bold">{{ $t('card.weight') }}:</v-list-item-title>
                <v-list-item-subtitle class="no-ellipsis">{{ pokemon.WeightKg }} kg</v-list-item-subtitle>
              </v-list-item>
              <v-list-item>
                <v-list-item-title class="font-weight-bold">{{ $t('card.pokedexEntry') }}:</v-list-item-title>
                <v-list-item-subtitle class="no-ellipsis">{{ pokemon.PokedexEntry }}</v-list-item-subtitle>
              </v-list-item>
              <v-list-item v-if="pokemon.Appearance">
                <v-list-item-title class="font-weight-bold">{{ $t('card.appearance') }}:</v-list-item-title>
                <v-list-item-subtitle class="no-ellipsis">{{ pokemon.Appearance }}</v-list-item-subtitle>
              </v-list-item>
              <v-list-item v-if="pokemon.Attacks?.length">
                <v-list-item-title class="font-weight-bold">{{ $t('card.attacks') }}:</v-list-item-title>
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
      </template>
    </v-expand-transition>
    </v-card>
    <v-dialog v-model="showImageModal" max-width="600" persistent>
      <v-card>
        <v-card-title class="text-h6 font-weight-bold" style="text-align:center;">{{ pokemon.Name }}</v-card-title>
        <v-card-text style="display:flex; justify-content:center; align-items:center;">
          <img
            :src="`data:image/png;base64,${pokemon.ImageBase64}`"
            alt="Pokemon image large"
            style="max-width:100%; max-height:400px; object-fit:contain;"
          />
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="primary" @click="showImageModal = false">Close</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog v-model="showStatsModal" max-width="600" persistent>
      <v-card>
        <v-card-title class="text-h6 font-weight-bold" style="text-align:center;">Stats for {{ pokemon.Name }}</v-card-title>
        <v-card-text style="display:flex; justify-content:center; align-items:center; min-height:400px;">
          <div style="width:350px; height:350px; display:flex; justify-content:center; align-items:center;">
            <PokemonStatsChart :stats="extractStats(pokemon) || defaultStats" />
          </div>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="primary" @click="showStatsModal = false">Close</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup>
// Fallback stats to prevent undefined errors in chart
const defaultStats = {
  HP: 0,
  Attack: 0,
  Defense: 0,
  Speed: 0,
  SpAttack: 0,
  SpDefense: 0,
}
import { ref, computed } from 'vue'
import PokemonStatsChart from './PokemonStatsChart.vue'
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
const emit = defineEmits(['delete', 'preview-image'])
const expanded = ref(false)
const showStatsModal = ref(false)
const showImageModal = ref(false)
const appStore = useAppStore()
const isAdmin = computed(() => appStore.isAdmin)


// Helper to extract stats from the pokemon object
function extractStats(p) {
  return {
    HP: p.HP ?? 0,
    Attack: p.Attack ?? 0,
    Defense: p.Defense ?? 0,
    Speed: p.Speed ?? 0,
    SpAttack: p.SpAttack ?? 0,
    SpDefense: p.SpDefense ?? 0,
  }
}

function onDelete() {
  emit('delete', props.pokemon)
}
function previewImage() {
  emit('preview-image', `data:image/png;base64,${props.pokemon.ImageBase64}`)
  showImageModal.value = true
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
/* ...existing code... */
.pokemon-image-preview {
  display: block;
  margin: 0 auto;
  max-width: 100%;
  max-height: 250px;
  object-fit: contain;
  cursor: pointer;
  transition: box-shadow 0.2s, transform 0.2s;
}
.pokemon-image-preview:hover {
  box-shadow: 0 0 0 4px #356abc55, 0 2px 16px rgba(0,0,0,0.15);
  transform: scale(1.04);
  filter: brightness(1.08);
}
</style>
