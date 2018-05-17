<template>
<div class='main-page'>
  <v-chip v-if=status.text label color='red' text-color='white'>{{ status.text }}</v-chip>
  <b>Current Directory: {{dir ? dir : '/'}}</b>
  <div class="buttons">
    <v-btn flat small icon color="gray" @click="moveBack">←</v-btn>
    <v-btn-toggle v-model="file_view">
      <v-btn flat>
        <v-icon>icons</v-icon>
      </v-btn>
      <v-btn flat>
        <v-icon>details</v-icon>
      </v-btn>
    </v-btn-toggle>
  </div>
  <br/>
  <v-layout column wrap v-if="file_view == 1">
    <v-flex xs2 class="file-details" v-for="file in files" :key="file.name">
      <a class="file-name" v-bind:class="{file: !file.isDirectory}" v-on:click.once="moveTo(file)">{{file.name}}</a>
    </v-flex>
  </v-layout>
  <v-layout row wrap v-else>
    <v-flex v-for="file in files" :key="file.name" md2 sm4 xs6 class="file-icon-group">
      <img src="" class="file-icon" />
      <a class="file-name" v-if="file.isDirectory" v-on:click.once="moveTo(file.name)">{{file.name}}</a>
      <div class="file-name" v-else>{{ file.name }}</div>
    </v-flex>
  </v-layout>
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
      file_view: 1,
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
    moveTo(file) {
      if (file.isDirectory) {
        this.directories.push(file.name);
        this.updateDirectory();
        this.getItems();
      } else {
        this.getItem(file.name);
      }
    },
    getItem(name) {
      this.axios.get('item', { params: { name: this.dir.concat(`/${name}`) } }).then((response) => {
        console.log(response);
      }).catch((error) => {
        if (error.response.status === 401) {
          this.$router.replace('/login');
        } else {
          this.status.text = error.response.data.message;
        }
      });
    },
    getItems() {
      this.axios.get('items', { params: { dir: this.dir } }).then((response) => {
        this.files = response.data;
      }).catch((error) => {
        if (error.response.status === 401) {
          this.$router.replace('/login');
        } else {
          this.status.text = error.response.data.message;
        }
      });
    },
  },
};
</script>

<style scoped>
  a:hover {
    color: #84baf0;
  }
  a.file {
    color: #000;
  }
  a.file:hover {
    color: #aaa;
  }
  .file-icon-group {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
  .file-name {
    width: 100%;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    display: block;
  }
  .file-icon {
    height: 150px;
    width: 150px;
  }
  .file-details {
    text-align: left;
  }
  .buttons {
    display: flex;
    justify-content: space-between;
    width: 100%;
  }
</style>
