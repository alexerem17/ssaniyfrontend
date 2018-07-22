<template>
  <footer class="footer">
    <div class="navbar navbar-dark bg-dark">
      <span class="text-muted text-center text-white">{{ textFooter }}</span>
    </div>
  </footer>
</template>

<script>
export default {
  name: "myFooter",
  data: () => ({
    textFooter: "Проверка текста"
  })
};
</script>