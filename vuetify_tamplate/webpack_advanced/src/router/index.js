import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import Form from '@/components/Form'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/hello',
      name: 'Hello',
      component: Hello
    },
    {
      path: '/form',
      name: 'Form',
      component: Form
    }
  ],
  mode: 'history'
})
