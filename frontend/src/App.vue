<template>
  <v-app>
    <v-app-bar app color="indigo" dark>
      <v-toolbar-title>Fantasy Pokédex</v-toolbar-title>
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
