<template>
  <header>
    <div class="bg-rose-500 shadow-md rounded px-8 pt-6 pb-8 mb-4"></div>
  </header>
  <section id="app_container">
  <form
    class="bg-rose-500 shadow-md rounded px-8 pt-6 pb-8 mb-4"
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
        Buscar
      </button>
    </div>
  </form>

  <!-- content table -->
  <div class="flex flex-col md:flex-row px-5 min-w-full max-h-min h-96">
    <div class="md:h-auto md:max-h-50vh md:w-1/2 md:overflow-y-scroll">
      <table id="list_table" class="border-2 border-gray-300 bg-white rounded-lg table-fixed">
    <thead>
      
        <tr class="border-b-2 border-gray-300">
            <th class="w-1/3 px-4 py-2">Subject</th>
            <th class="w-1/3 px-4 py-2">From</th>
            <th class="w-1/3 px-4 py-2">To</th>
        </tr>
    </thead>
    <tbody>
      <tr v-for="(item, index) in result" :key="item.id" v-bind:class="[index % 2 !== 0 ? 'bg-white' : 'bg-gray-100']" @click="selectedContent = item.content">
            <td class="w-1/3 md:w-1/2 px-4 py-2">{{ item.subject }}</td>
            <td class="w-1/3 py-2">{{ item.from }}</td>
            <td class="w-1/3 px-4 py-2">{{ item.to }}</td>
        </tr>
    </tbody>
</table>
</div>
  <div class="content md:w-1/2 md:pl-5 pt-10 bg-gray-200 md:max-h-50vh  md:overflow-y-scroll">
    <p v-html="selectedContent" class="md:w-40%"></p>
  </div>
</div>
</section>
</template>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      key: '',
      result: [],
      selectedContent: '',
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
    handleClick(item) {
      this.selectedContent = item.content;
    },
  },
};
</script>
