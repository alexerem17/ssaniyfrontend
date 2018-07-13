<template><section>
 <h1>{{id}}</h1>
  <ul>
  <li v-for="(n, index) in parseInt(id)" :key="index">
    <p v-if="id/2==n"><strong>{{n}}:{{index}}</strong></p><p v-else>{{n}}:{{index}}</p>
    </li>
</ul>
</section>
</template>

<script>
export default {
  asyncData({ params, error }) {
    return {
      id: params.id
    };
  },
};
</script>