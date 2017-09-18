<script>
import Vue from 'vue'

Vue.component('text-input', {
  template: `<td><input v-bind:value="value" @input="updateValue($event.target.value)"></input></td>`,
  // template: `<td><v-text-field v-bind:value="value" @input="updateValue($event.target.value)"></v-text-field></td>`,
  props: {
    value: {
      type: String,
      default: ''
    }
  },
  methods: {
    updateValue: function (value) {
      // console.log(value)
      this.$emit('input', value)
    }
  }
})
</script>