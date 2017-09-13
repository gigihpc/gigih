<template>
  <v-form>
      <title>VueJs</title>
	<div class="container">
		<h1>VueJs Tutorial</h1>

		<table class="table">
			<thead>
				<tr>
					<th>Name</th>
					<th>Address</th>
					<th>Old</th>
					<th>Action</th>
				</tr>
				<tr>
                    <td><v-text-field v-model="name"></v-text-field></td>
					          <td><v-text-field v-model="address"></v-text-field></td>
					          <td><v-text-field v-model="old"></v-text-field></td>
					          <td><button class="btn btn-primary add" v-on:click.prevent="add">Add</button>
                    <button class="btn btn-primary" v-on:click.prevent="show_all">SHOW ALL</button> </td>
				</tr>
                <tr></tr>
                <tr>
                    <template v-for="(mhsw,index) in mhsws">
                       <tr :key="mhsw.id">
                         <td>{{mhsw.name}}</td>
                         <td>{{mhsw.address}}</td>
                         <td>{{mhsw.old}}</td>
                          <td><!--<button class="btn btn-warning edit" v:on:click.native=edit>Edit</button>
                             <button class="btn btn-primary remove" v:on:click.prevent=remove>Remove</button>
                             <button class="btn btn-success update" v:on:click="update">Update</button> 
                             <button class="btn btn-danger cancel" v:on:click="cancel">Cancel</button> -->
                             <v-btn v-if="edit_mhsw" @click="edit($event,index)">Edit</v-btn>
                             <v-btn v-if="delete_mhsw" @click="remove($event,mhsw.id)">Delete</v-btn>
                             <v-btn v-if="update_mhsw" @click="update($event,mhsw.id)">Update</v-btn>
                             <v-btn v-if="cancel_mhsw" @click="cancel">Cancel</v-btn></td>
                       </tr>
                    </template>
                    <template>
                       <tr v-if="update_mhsw" v-on="edit"><text-field-edit></text-field-edit></tr>
                    </template>
                </tr>
			</thead>
		</table>
	</div>
  </v-form>
</template>

<script>
import {HTTP} from '@/router/index'
import Vue from 'vue'
// import axios from 'axios'
export default {
  data () {
    return {
      name: '',
      address: '',
      old: '',
      name_in: '',
      address_in: '',
      old_in: '',
      mhsw: [],
      mhsws: [],
      update_mhsw: false,
      cancel_mhsw: false,
      edit_mhsw: false,
      delete_mhsw: false
    }
  },
  ready: function () {
    this.fetchmhsws()
  },
  methods: {
    show_all: function () {
      this.fetchmhsws()
      this.edit_mhsw = true
      this.delete_mhsw = true
    },
    fetchmhsws: function () {
      HTTP.get('/api/mhsws').then(response => {
        this.mhsws = response.data.data
      })
      .catch(err => {
        console.log(err)
      })
    },
    add: function () {
      var name = this.name
      var address = this.address
      var old = this.old
      var mhswJson = {
        data: {
          name,
          address,
          old
        }
      }
      console.log(mhswJson)
      HTTP.post('/api/mhsws', mhswJson)
      .then(function (response) {
        console.log('succesfully create mahasiswa')
        this.fatchmhsws()
      })
      .catch(function (err) {
        console.log(err.response)
      })
      // axios.post('http://192.168.1.8:3003/api/mhsws', {name: 'suci', address: 'galek', old: '21'})
      // .then(res => {
      //   console.log('success')
      // }).catch(err => {
      //   console.log(err.res)
      // })
    },
    edit: function (event, index) {
      this.update_mhsw = true
      this.cancel_mhsw = true
      this.edit_mhsw = !this.edit_mhsw
      this.delete_mhsw = false
      this.name_in = this.mhsws[index].name
      this.address_in = this.mhsws[index].address
      this.old_in = this.mhsws[index].old
    },
    remove: function (event, _id) {
      HTTP.delete('/api/mhsws/' + _id).then(response => {
        console.log('succesfully delete _id: ' + _id)
        this.fetchmhsws()
      })
      .catch(err => {
        console.log(err.response)
      })
    },
    update: function (event, _id) {
      HTTP.get('/api/mhsws/' + _id).then(response => {
        this.mhsw = response.data
      })
      .catch(err => {
        console.log(err.response)
      })
      HTTP.put('/api/mhsws/' + _id, this.mhsw).then(response => {
        console.log('success update _id:' + _id)
      })
      .catch(err => {
        console.log(err.response)
      })
    },
    valid () {},
    cancel: function () {
      this.fetchmhsws()
      this.update_mhsw = !this.update_mhsw
      this.cancel_mhsw = !this.cancel_mhsw
      this.edit_mhsw = true
      this.delete_mhsw = true
    }
  },
  event: {
    click () {}
  }

}
Vue.component('text-field-edit', {
  template: '<tr><td><v-text-field v-model="name_in">this.name_in</v-text-field></td><td><v-text-field v-model="address_in">this.address_in</v-text-field></td><td><v-text-field v-model="old_in">this.old_in</v-text-field></td></tr>',
  data: function () {
    return {
      name_in: '',
      address_in: '',
      old_in: ''
    }
  },
  methods: {
    edit: function (event, index, mhsw) {
      this.fetchmhsws()
      this.name_in = this.mhsws[index].name
      this.address_in = this.mhsws[index].address
      this.old_in = this.mhsws[index].old
      console.log(this.name_in, this.address_in, this.old_in)
    }
  }
})
</script>