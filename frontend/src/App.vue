<template>
  <v-app>
    <v-app-bar app color="blue-grey-darken-2" dark>
      <v-container class="d-flex justify-center align-center" fluid>
        <v-img
          src="@/assets/fantasy-pokedex-logo.png"
          max-height="150"
          style="max-width: 200px"
        />
      </v-container>

      <v-spacer />

      <v-btn text to="/">Home</v-btn>
      <v-btn text to="/create">Create</v-btn>
    </v-app-bar>

    <v-main>
      <v-container class="py-5" fluid>
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>
