<template>
  <div id="app">
   
    <div class="sidenav">
        <!-- <PanelMenu :model="items" style="font-size: 15px; text-align:left" /> -->
         <PanelMenu  :model="items" style="text-align:left"/>
       
    </div>

    <div class="app-container">
      <router-view />
      
       
      </div>
    
    <Toast />
  </div>
</template>

<script>
export default {
  data() {
    return {
      visibleLeft: false,
      message: null,
      text: null,
      items: [
        { label: "Início",
          icon: "pi pi-fw pi-home",
          items:[
            {
              label: "Ir para página inicial",
              icon: "",
              to:"/"
            }
        ]
        },
        { label: "Documentos",
          icon: "pi pi-fw pi-file",
          items: [
            {
              label: "Tipo PDF",
              icon: "pi pi-fw pi-caret-right",
              items: [
                {
                  label: "Fazer Upload de Documentos",
                  icon: "",
                  to:"/uploadPDF"
                },
                {
                  label: "Realizar Download de Documentos",
                  icon: "",
                  to:"/downloadPDF"
                },
              ],
            },
          ],
        },        
      ],
    };
  },
  methods: {
    computed() {
      this.$menuitemTextColor = "green";
    },
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

.p-panelmenu {
  font-size: 16px;

}

.p-panelmenu-header{
  font-size: 15px;
}

.p-panelmenu-list{
  font-size: 05px;
}


.app-container {
  text-align: Left;
}

body #app .p-button {
  margin-left: 200px; // 0.2em;
}

form {
  margin-top: 2em;
}

* {
  box-sizing: border-box;
}

body {
  margin: 0;
  font-family: Arial, Helvetica, sans-serif;
   margin-left: 200px;
}
/* Peguei do w3school. Exemplos->Templates
/* Style the side navigation */
.sidenav {  
  height: 100%;
  width: 250px;
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  background-color: #111;
  overflow-x: hidden;
  
}


/* Side navigation links */

/* Style the content */
.content {
  margin-left: 200px;
  padding-left: 20px;
}
</style>
