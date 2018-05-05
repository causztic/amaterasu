<template>
  <v-app>
    <v-content>
      <div v-if="$auth.ready()">
        <router-view/>
      </div>

      <div v-if="!$auth.ready()">
        Loading ...
      </div>
    </v-content>
  </v-app>
</template>

<script>
export default {
  name: 'amaterasu',
};
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  padding-top: 60px;
}
</style>
