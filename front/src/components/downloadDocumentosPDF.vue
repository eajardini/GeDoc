<template>
  <div class="hello">
    <h3> Download de Documentos Armazenados </h3>
      <div style="background-color: black"> 
          <DataTable  :value="documentos" :rowHover="true" :selection.sync="selectedProduct1" dataKey="id" >
            <Column field="id" header="Código"></Column>
            <Column field="titulo" header="Título"></Column>
            <Column field="datadocumento" header="Data Criação"></Column>   
             <Column header="Download Documento" headerStyle=" text-align: center" bodyStyle="text-align: center; ">
                  <template #body="slotProps">
                      <Button type="button"  label="Download" icon="pi pi-file-pdf" class="p-button-help" @click="GetDocumentoParaDownload(slotProps.data)" ></Button>
                      <!-- <img       
                          src= "../assets/images/pdfDownload.png"
                          width="70px"
                          /> -->
                  </template>
             </Column>
          </DataTable>
      </div>
  </div>
</template>

<script>
import documentosService from '../service/ProductService';

export default {
  data() {
        return {
            documentos: null,
            selectedProduct1: null,
             
        }
    },
  documentosService: null,
  created() {
        this.documentosService = new documentosService();
    },
  mounted() {
        this.$http.get("/GetAllDocs")
          .then(response => {
            console.log(response);
            this.documentos = response.data;
            })    
  },
  methods: {
    GetDocumentoParaDownload(documento){        
        console.log("Valor do ID:" + documento.id);
        let formData = new FormData();   
        formData.append('IDDocumento', documento.id)     

        this.$http.post("/GetDocs", formData, {responseType: 'blob'})
          .then(function(response) {
            var blob = new Blob([response.data], {
            type: 'application/pdf'
    });
          var url = window.URL.createObjectURL(blob)
          window.open(url);
  })
    }
  },
}
</script>