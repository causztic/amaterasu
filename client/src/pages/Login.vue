<template>
<div class='login-page'>
  <v-chip v-if=status.text label color='red' text-color='white'>{{ status.text }}</v-chip>
  <h1>{{ msg }}</h1>
  <v-form ref='form' v-model='valid' @submit="whatever">
    <v-text-field v-model='username' label='username' :rules='req' required />
    <v-text-field v-model='password' label='password' :rules='req' type='password' required />
    <br/>
    <v-btn type='submit' :disabled='!valid' @click='submit'>
      submit
    </v-btn>
  </v-form>
</div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      msg: 'File Storage System',
      valid: true,
      req: [
        v => !!v || 'field is required',
      ],
      username: '',
      password: '',
      status: {
        text: ''
      },
    };
  },
  methods: {
    submit() {
      if (this.$refs.form.validate()) {
        this.axios.post('auth/login', {
          username: this.username,
          password: this.password,
        }).then((response) => {
          localStorage.setItem('amaterasu_token', response.data.token);
          localStorage.setItem('amaterasu_token_expire', response.data.expire);
          this.$router.replace('/');
        }).catch((error) => {
          if (error.response.status === 401) {
            this.status.text = error.response.data.message;
          } else {
            this.status.text = 'Server error.';
          }
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
