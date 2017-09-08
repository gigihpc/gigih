import Vue from 'vue'
import Router from 'vue-router'
import axios from 'axios'
import Hello from '@/components/Hello'
import Form from '@/components/Form'
import Mhsw from '@/components/mhsw'

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
    },
    {
      path: '/mhsw',
      name: 'MHSW',
      component: Mhsw
    }
  ],
  mode: 'history'
})

export const HTTP = axios.create({
  baseURL: 'http://192.168.1.8:8001',
  timeout: 1000,
  withCredentials: true
})
