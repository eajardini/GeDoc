import axios from 'axios';

export default class ProductService {

    getProductsSmall() {
		return axios.get('assets/layout/data/documentos.json').then(res => res.data.data);
	}
}
