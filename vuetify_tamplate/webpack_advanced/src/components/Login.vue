<template>
  <v-form ref="form">
      <v-layout justify-center>
    <v-flex xs12 sm10 md8 lg6>
      <li>{{ isfail }}</li>
    <v-text-field
      label="Email"
      v-model="email"
    ></v-text-field>
    <v-text-field
      label="Password"
      v-model="password"
      :append-icon="isShow ? 'visibility' : 'visibility_off'"
              :append-icon-cb="() => (isShow = !isShow)"
              :type="isShow ? 'text' : 'password'"
              counter
    ></v-text-field>

    <v-btn @click="Login">LOGIN</v-btn>
    <td><router-link to="regist">Create User</router-link></td>
    <router-view></router-view>
    </v-flex>
    </v-layout>
  </v-form>
</template>

<script>
import {HTTP} from '@/router/index'
export default{
  data () {
    return {
      email: null,
      password: null,
      user_list: [],
      isShow: false,
      isfail: ''
    }
  },
  methods: {
    Login: function () {
      var email = this.email
      var pass = this.password
    //   console.log('var: ', email, pass)
      HTTP.get('/api/user').then(response => {
        this.user_list = response.data.data
        this.user_list.forEach(function (element) {
        //   console.log('from db: ', element.email, element.password)
          if (email === element.email && pass === element.password) {
            console.log('Login Success!')
            this.$router.push('mhsw')
          } else {
            this.isfail = 'Login Failed !'
            console.log('Login Failed!!')
          }
        }, this)
      })
      .catch(err => {
        console.log(err.response)
      })
    }
  }
}
</script>