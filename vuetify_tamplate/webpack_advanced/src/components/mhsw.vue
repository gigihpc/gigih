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
                    <td><v-text-field v-model="name"></v-text-field></td>
                    <td><v-text-field v-model="address"></v-text-field></td>
                    <td><v-text-field v-model="old"></v-text-field></td>
					<td><button class="btn btn-primary add" @click="add">Add</button>
                    <button class="btn btn-warning edit" @click="edit">Edit</button><button class="btn btn-danger remove" @click="remove">Remove</button>
                    <button class="btn btn-primary"@click="fetchmhsws">SHOW ALL</button>
                    <button class="btn btn-success update" @click="update" style="display:none">Update</button> 
                    <button class="btn btn-danger cancel" @click="cancel" style="display:none">Cancel</button></td>
				</tr>
                <tr></tr>
                <tr>
                    <template v-for="mhsw in mhsws">
                       <tr> 
                         <td>{{mhsw.name}}</td>
                         <td>{{mhsw.address}}</td>
                         <td>{{mhsw.old}}</td>
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
      result: [],
      error: [],
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
          console.log(element.id + ';' + element.name + ';' + element.address + ';' + element.old)
        }, this)
      })
      .catch(e => {
        console.log(e)
      })
    },
    fetchmhsws: function () {
      console.log('test')
      HTTP.get('/api/mhsws').then(response => {
        this.mhsws = response.data.data
      })
      .catch(err => {
        console.log(err)
      })
    },
    add () {},
    edit () {},
    remove () {},
    update () {},
    valid () {},
    cancel () {}
  },
  event: {
    click () {}
  }

}
</script>