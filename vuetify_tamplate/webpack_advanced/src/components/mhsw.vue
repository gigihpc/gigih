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
                    <td><v-text-field v-model="name" id="namein"></v-text-field></td>
					          <td><v-text-field v-model="address" id="addressin"></v-text-field></td>
					          <td><v-text-field v-model="old" id="oldin"></v-text-field></td>
					          <td><button class="btn btn-primary add" v-on:click.prevent="add">Add</button>
                    <button class="btn btn-primary" v-on:click.prevent="fetchmhsws">SHOW ALL</button></td>
				</tr>
                <tr></tr>
                <tr>
                    <template v-for="mhsw in mhsws">
                       <tr :key="mhsw.id">
                         <td>{{mhsw.id}}</td>
                         <td colspan="100%">{{mhsw.name}}</td>
                         <td colspan="100%">{{mhsw.address}}</td>
                         <td colspan="100%">{{mhsw.old}}</td>
                          <td><!--<button class="btn btn-warning edit" v:on:click.native=edit>Edit</button>
                             <button class="btn btn-primary remove" v:on:click.prevent=remove>Remove</button>
                             <button class="btn btn-success update" v:on:click="update">Update</button> 
                             <button class="btn btn-danger cancel" v:on:click="cancel">Cancel</button> -->
                             <v-btn @click="edit">Edit</v-btn>
                             <v-btn @click="remove">Remove</v-btn>
                             <v-btn @click="update">Update</v-btn>
                             <v-btn @click="cancel">Cancel</v-btn></td>
                       </tr>
                    </template>
                </tr>
			</thead>
		</table>
	</div>
  </v-form>
</template>

<script>
import {HTTP} from '@/router/index'
// import axios from 'axios'
export default {
  data () {
    return {
      name: '',
      address: '',
      old: '',
      mhsw: [],
      mhsws: []
    }
  },
  ready: function () {
    this.fetchmhsws()
  },
  methods: {
    showAll () {
      HTTP.get('/api/mhsws').then(response => {
        this.result = response.data
        this.result.data.forEach(function (element) {
        }, this)
      })
      .catch(e => {
        console.log(e)
      })
    },
    fetchmhsws: function () {
      HTTP.get('/api/mhsws').then(response => {
        this.mhsws = response.data.data
      })
      .catch(err => {
        console.log(err)
      })
    },
    add () {
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
        console.log('succesfully')
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
    edit: function () {
      console.log('edit')
    },
    remove: function () {
      console.log('test')
      HTTP.delete('/api/mhsws/:id').then(response => {
        console.log('succesfully delete: ' + response.data.data)
      })
      .catch(err => {
        console.log(err.response)
      })
    },
    update () {
      console.log('update')
    },
    valid () {},
    cancel () {}
  },
  event: {
    click () {}
  }

}
</script>