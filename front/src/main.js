import Vue from 'vue';
import App from './App.vue';
import router from './router';
import "./axios"

import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import Calendar from 'primevue/calendar';
import Column from 'primevue/column';
import DataTable from 'primevue/datatable';
import Toast from 'primevue/toast';
import ToastService from 'primevue/toastservice';
import PanelMenu from 'primevue/panelmenu';
import Sidebar from 'primevue/sidebar';
import Panel from 'primevue/panel';
import Menu from 'primevue/menu';
import Message from 'primevue/message';
import FileUpload from 'primevue/fileupload';

Vue.use(ToastService);

Vue.component('InputText', InputText);
Vue.component('Button', Button);
Vue.component('Calendar', Calendar);
Vue.component('Column', Column);
Vue.component('DataTable', DataTable);
Vue.component('Toast', Toast);
Vue.component('PanelMenu', PanelMenu);
Vue.component('Sidebar', Sidebar);
Vue.component('Panel', Panel);
Vue.component('Menu', Menu);
Vue.component('Message', Message);
Vue.component('FileUpload', FileUpload);


// import 'primevue/resources/themes/nova-light/theme.css';
import "primevue/resources/themes/bootstrap4-dark-blue/theme.css"
// import "primevue/resources/themes/luna-blue/theme.css"
// import "primevue/resources/themes/md-dark-indigo/theme.css"
// import "primevue/resources/themes/mdc-dark-indigo/theme.css"
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';

Vue.config.productionTip = false;

new Vue({
  router,  
  render: h => h(App),
}).$mount('#app')
