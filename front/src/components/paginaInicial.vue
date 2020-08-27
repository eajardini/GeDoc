<template>
  <div class="hello">
    <h3> Página Inicial </h3>
    <p>Status do servidor: {{statusServer}} </p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      name: "paginaInicial",
      statusServer: "Servidor não alcançado!",
    };
  },
  props: {
    msg: String
  },
   mounted() {
    this.$http.get("/GetStatus")
    .then(response => {
      console.log(response);
      this.statusServer = response.data.Resposta;
      })
  },
  methods: {},
}
</script>