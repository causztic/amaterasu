<template>
<div class='main-page'>
  <v-chip v-if=status.text label color='red' text-color='white'>{{ status.text }}</v-chip>
  <div class='file-container'>
    <b>Current Directory: {{dir ? dir : '/'}}</b>
    <div class="buttons">
      <v-btn flat small icon color="gray" @click="moveBack">←</v-btn>
    </div>
    <br/>
    <div v-for="file in files" :key="file.name">
      <a v-if="file.isDirectory" v-on:click.once="moveTo(file.name)">{{file.name}}</a>
      <div v-else>{{ file.name }}</div>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: 'Main',
  data() {
    return {
      files: [],
      directories: [],
      dir: '',
      status: {
        text: '',
      },
    };
  },
  // eslint-disable-next-line
  created: function() {
    this.getItems();
  },
  methods: {
    updateDirectory() {
      this.dir = this.directories.reduce((a, value) => `${a}/${value}`, '');
    },
    moveBack() {
      this.directories.pop();
      this.updateDirectory();
      this.getItems();
    },
    moveTo(name) {
      this.directories.push(name);
      this.updateDirectory();
      this.getItems();
    },
    getItems() {
      this.axios.get('items', { params: { dir: this.dir } }).then((response) => {
        this.files = response.data;
      });
    },
  },
};
</script>

<style scoped>
  .file-container {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    padding: 0 50px;
  }
</style>
