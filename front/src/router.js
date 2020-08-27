import Vue from 'vue';
import Router from 'vue-router';
import paginaInicial from './components/paginaInicial';
import HelloWorld from './components/HelloWorld.vue'
import uploadDocumentoPDF from './components/uploadDocumentoPDF.vue'
import downloadDocumentosPDF from './components/downloadDocumentosPDF.vue'

Vue.use(Router);

export default new Router({
	mode: "history",
	routes: [
		{
			path: '/',
			name: 'paginaInicial',
			component: paginaInicial
		},		
		{
			path: '/hello',
			name: 'HelloWorld',
			component: HelloWorld
		},
		{
			path: '/uploadPDF',
			name: 'uploadDocumentoPDF',
			component: uploadDocumentoPDF
		},
		{
			path: '/downloadPDF',
			name: 'downloadDocumentosPDF',
			component: downloadDocumentosPDF
		},
	
	],
	scrollBehavior() {
		return {x: 0, y: 0};
	}
});