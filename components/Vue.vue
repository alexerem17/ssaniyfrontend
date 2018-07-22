<template>
<section>
  <input v-model='Date' placeholder="Введите дату">
</section>
</template>
 
 <script>
 export default {
     name: 'myVue',
     props: 'Date'
 }
 </script>