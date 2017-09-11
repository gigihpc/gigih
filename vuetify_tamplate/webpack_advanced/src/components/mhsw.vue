<template>
  <v-form v-model="valid">
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
                       <tr>
                         <td colspan="100%">{{mhsw.name}}</td>
                         <td colspan="100%">{{mhsw.address}}</td>
                         <td colspan="100%">{{mhsw.old}}</td>
                         <td><button class="btn btn-warning edit" v:on:click="edit">Edit</button>
                             <button class="btn btn-primary remove" v:on:click.prevent="remove">Remove</button>
                             <button class="btn btn-success update" v:on:click="update">Update</button> 
                             <button class="btn btn-danger cancel" v:on:click="cancel">Cancel</button></td>
                       </tr>
                    </template>
                </tr>
			</thead>
			<tbody class="mhsw-list"></tbody>
		</table>
	</div>
  </v-form>
</template>

<script>
import {HTTP} from '@/router/index'
export default {
  $validates: true,
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
      this.mhsw = [this.name, this.address, this.old]
      var name = this.name
      var address = this.address
      var old = this.old
      var mhswJson = {name, address, old}
      console.log(JSON.stringify(mhswJson))
      HTTP.post('/api/mhsws', JSON.stringify(mhswJson))
      .then(function (response) {
        console.log(response.data.data)
      })
      .catch(function (err) {
        console.log(err)
      })
    },
    edit () {},
    remove () {
      console.log('test')
      HTTP.delete('/api/mhsws/:id').then(response => {
        console.log('succesfully delete: ' + response.data.data)
      })
      .catch(err => {
        console.log(err)
      })
    },
    update () {},
    valid () {},
    cancel () {}
  },
  event: {
    click () {}
  }

}
</script>