<template>
  <div class="hello">
    <h3> Helloword component </h3>

      <!-- <HelloWorld msg="Testes com upload e download de arquivos"/>
      <form @submit.prevent="greet">
        <InputText type="text" v-model="text"/>
        <Button type="submit" label="Submit"/>
        <h3>{{message}}</h3>
      </form> -->

    <h1>{{ msg }}</h1>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  methods: {
        greet() {
           
            this.message = 'Olá ' + this.text;
        }
    },
}
</script>