<template>
<div class='main-page'>
  <v-chip v-if=status.text label color='red' text-color='white'>{{ status.text }}</v-chip>
  <div>
    {{ files }}
  </div>
</div>
</template>

<script>
export default {
  name: 'Main',
  data() {
    return {
      files: [],
      status: {
        text: '',
      },
    };
  },
  // eslint-disable-next-line
  beforeCreate: function() {
    this.axios.get('items').then((response) => {
      this.files = response.data;
    });
  },
  methods: {
  },
};
</script>

<style scoped>
</style>
