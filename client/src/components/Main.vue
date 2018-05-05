<template>
<div class="login-page">
  <h1>{{ msg }}</h1>
  <v-form ref="form" v-model="valid">
    <v-text-field v-model="username" label="username" :rules="req" required></v-text-field>
    <v-text-field v-model="password" label="password" :rules="req" type="password" required></v-text-field>
    <v-btn :disabled="!valid" @click="submit">
      submit
    </v-btn>
  </v-form>
</div>
</template>

<script>
export default {
  name: 'Main',
  data() {
    return {
      msg: 'File Storage System',
      valid: true,
      req: [
        v => !!v || 'field is required',
      ],
      username: '',
      password: '',
    };
  },
  methods: {
    submit() {
      if (this.$refs.form.validate()) {
        // Native form submission is not yet supported
        this.axios.post('login', {
          username: this.username,
          password: this.password,
        });
      }
    },
  },
};
</script>

<style scoped>
form {
  width: 300px;
  margin-top: 30px;
}

.login-page {
  display: flex;
  flex-direction: column;
  align-items: center;
}

h1,
h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
