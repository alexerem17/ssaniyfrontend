<template>
 <h1>{{id}}</h1>
</template>

<script>
export default {
    asyncData ({ params, error }) {
 return {
  id: params.name
}
}
}
</script>