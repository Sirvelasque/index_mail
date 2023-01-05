

<template>
  <form
    class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4"
    @submit="handleSubmit"
  >
    <div class="mb-4">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="key">
        Ingrese su búsqueda:
      </label>
      <input
        class="
          shadow
          appearance-none
          border
          rounded
          w-full
          py-2
          px-3
          text-gray-700
          leading-tight
          focus:outline-none focus:shadow-outline
        "
        type="text"
        id="key"
        v-model="key"
      />
    </div>
    <div class="flex items-center justify-between">
      <button
        class="
          bg-blue-500
          hover:bg-blue-700
          text-white
          font-bold
          py-2
          px-4
          rounded
          focus:outline-none focus:shadow-outline
        "
        type="submit"
      >
        Enviar
      </button>
    </div>
  </form>
  <!-- <ul>
      <li v-for="item in result" :key="item.id">
        {{ item.name }} - {{ item.email }}
      </li>
  </ul> -->
</template>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      key: '',
      result: [],
    };
  },
  methods: {
    handleSubmit(event) {
      // Prevenir la acción por defecto del formulario (recargar la página)
      event.preventDefault();
      // Enviar la solicitud GET
      axios.get(`http://localhost:8080/search?key=${this.key}`)
  .then((response) => {
    // Almacenar la respuesta en una variable
    this.result = response.data;
    console.log(response)
    console.log(this.result)
  })
  .catch((error) => {
    // Manejar el error
  });
    },
  },
};
</script>
