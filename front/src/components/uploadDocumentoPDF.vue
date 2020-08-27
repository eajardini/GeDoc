<template>
  <div class="hello">
    <h3>{{name}}</h3>

    <p>Forneça uma data e escolha um arquivo PDF para fazer Upload</p>

    <div class="p-field">
      <Calendar v-model="ValorDaData" :manualInput="false" :monthNavigator="true" :yearNavigator="true" :value="ValorDaData" dateFormat="dd/mm/yy" :showOnFocus="true" yearRange="1930:2040" :showIcon="true"  />
    </div>

    <div>
      <label>
        <input accept="application/pdf" class="buttonUpload" type="file" id="file" ref="file" v-on:change="myUploader()" />
      </label>     
    </div>
    <Button  style="padding-left: 0px;"  class="p-button-rounded p-button-help" label="Enviar Arquivo" id="myFile" type="file" name="documento" v-on:click="submitFile()" />       
  </div>
</template>

<script>
export default {
  data() {
    return {
      mensagem:null,
      name: "Upload Documento PDF",
      nomeArquivo: "Sem nome",
      ValorDaData: null,
      ValorDaDataString:"",
    };
  },
  props: {
    msg: String,
  },
  mounted:function(){
        this.setHora();
  },
  methods: {
    myUploader: function () {
      this.nomeArquivo = this.$refs.file.files[0];
    },
    setHora(){
      this.ValorDaData = new Date();
    },
    submitFile(){
      let formData = new FormData();
      //Converter a Date para format DD/MM/YYYY
      var dataLocal =  new Date(this.ValorDaData);
      var dia = String(dataLocal.getDate()).padStart(2, '0');
      var mes = String(dataLocal.getMonth() + 1).padStart(2, '0'); 
      var ano = dataLocal.getFullYear();
      dataLocal = dia+"/"+mes+"/"+ano;     
      formData.append('ValorDaData', dataLocal)
      formData.append('documento', this.nomeArquivo);
    
      this.$http.post( '/SetDocs',
      formData,
      {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
      }
      ).then(function(){
        // this.mensagem = "Deu certo";
        alert("Deu certo");
      })
      .catch(function(error){
        alert("Falha:" + error.response.data.Resposta);       
      });
    }
  },
};
</script>


<style scoped lang="scss">
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 60px;
  margin-left: 250px;
}

.buttonUpload {
  font-size: 20px;
  padding-top: 10px;
  padding-bottom: 10px;
  
}
</style>