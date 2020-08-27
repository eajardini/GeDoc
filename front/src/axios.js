import Vue from 'vue'
import axios from 'axios'


Vue.use({
  install(Vue) {    
    Vue.prototype.$http = axios.create({
      baseURL: 'http://localhost:3000',
      // timeout:	10000,
    })   
    Vue.prototype.$http.interceptors.request.use(config => {
      
      return config
    }, error => Promise.reject(error))

    Vue.prototype.$http.interceptors.response.use(resp => {
      
      return resp
    }, error => Promise.reject(error))



  }
})


// this.$acl.interceptors.request.use((config) => {
//   const token = this.$cookies.get("token");

//   if (token) {
//     config.headers.Authorization = `Bearer ${token}`
//   }

//   return config
// }, (err) => {
//   return Promise.reject(err)
// })
