<template>
  <div class='main-page'>
    <v-dialog v-model="videoDialog">
      <video-player
      ref="videoPlayer"
      class="vjs-custom-skin"
      :options="playerOptions"
      :playsinline="true">
      </video-player>
    </v-dialog>
    <v-chip v-if=status.text label color='red' text-color='white'>{{ status.text }}</v-chip>
    <h2>Current Directory: {{dir ? dir : '/'}}</h2>
    <div class="buttons">
      <v-btn flat small icon color="gray" @click="moveBack">←</v-btn>
      <v-btn-toggle v-model="fileView">
        <v-btn flat>
          <v-icon>icons</v-icon>
        </v-btn>
        <v-btn flat>
          <v-icon>details</v-icon>
        </v-btn>
      </v-btn-toggle>
    </div>
    <br/>
    <v-layout column wrap v-if="fileView == 1">
      <v-flex xs2 class="file-details" v-for="file in files" :key="file.name">
        <a class="file-name" v-bind:class="{file: !file.isDirectory}"
        v-on:click.once="moveTo(file)">{{file.name}}</a>
      </v-flex>
    </v-layout>
    <v-layout row wrap v-else>
      <v-flex v-for="file in files" :key="file.name" md2 sm4 xs6 class="file-icon-group">
        <img src="" class="file-icon" />
        <a class="file-name" v-if="file.isDirectory"
        v-on:click.once="moveTo(file.name)">{{file.name}}</a>
        <div class="file-name" v-else>{{ file.name }}</div>
      </v-flex>
    </v-layout>
  </div>
</template>

<script>
import 'video.js/dist/video-js.css';
import 'vue-video-player/src/custom-theme.css';
import { videoPlayer } from 'vue-video-player';

export default {
  name: 'Main',
  components: { videoPlayer },
  data() {
    return {
      playerOptions: {
        height: '600',
        autoplay: true,
        language: 'en',
        playbackRates: [0.7, 1.0, 1.5, 2.0],
        sources: [],
      },
      videoDialog: false,
      files: [],
      directories: ['Users', 'yaojie', 'Documents', 'Music'],
      dir: '',
      fileView: 1,
      status: {
        text: '',
      },
    };
  },
  // eslint-disable-next-line
  created: function() {
    this.updateDirectory();
    this.getItems();
  },
  watch: {
    videoDialog: function() {
      if (!this.videoDialog) {
         this.playerOptions.sources = [];
      }
    }
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
    isVideo(name) {
      const items = name.split('.');
      const ext = items[items.length - 1];
      return [['mp4', 'flv', 'webm', 'avi', 'mov', 'wmv'].includes(ext), ext];
    },
    getItem(name) {
      const fullName = this.dir.concat(`/${name}`);
      const [isVideo, ext] = this.isVideo(name);
      if (isVideo) {
        if (['mp4', 'flv', 'webm'].includes(ext)) {
          this.videoDialog = true;
          const src = `${this.axios.defaults.baseURL}/item?name=${fullName}`;
          this.playerOptions.sources = [{
            type: `video/${ext}`,
            src,
          }];
        } else {
          this.axios.get('item', {
            params: { name: fullName }})
          .then((response) => {
            console.log(response);
          }).catch((error) => {
            console.err(error);
          })
        }
      } else {
        this.axios.get('item', {
          params: { name: fullName }, responseType: 'blob' })
          .then((response) => {
            this.status.text = '';
            const link = document.createElement('a');
            link.href = window.URL.createObjectURL(new Blob([response.data]));
            link.download = name;
            link.click();
          }).catch((error) => {
            if (error.response.status === 401) {
              this.$router.replace('/login');
            } else {
              this.status.text = error.response.data;
            }
          });
      }
    },
    getItems() {
      this.axios.get('items', { params: { dir: this.dir } }).then((response) => {
        this.status.text = '';
        this.files = response.data;
      }).catch((error) => {
        if (error.response.status === 401) {
          this.$router.replace('/login');
        } else {
          this.status.text = error.response.data;
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
