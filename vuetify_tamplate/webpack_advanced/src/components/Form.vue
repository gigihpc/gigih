<template>
  <v-form v-model="valid">
    <v-text-field
      label="Name"
      v-model="name"
      :rules="nameRules"
      :counter="10"
      required
    ></v-text-field>
    <v-text-field
      label="Address"
      v-model="address"
    ></v-text-field>
     <v-text-field
      label="Old"
      v-model="old"
    ></v-text-field>

    <v-btn @click="add">Add</v-btn>
    <v-btn @click="edit">Edit</v-btn>
    <v-btn @click="remove">Delete</v-btn>
  </v-form>
</template>

<script>
export default {
  $validates: true,
  data () {
    return {
      name: '',
      address: '',
      old: ''
    }
  },
  methods: {
    add () {},
    edit () {},
    remove () {}
  }
}
</script>
