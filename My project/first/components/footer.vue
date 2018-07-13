<template>
  <footer class="footer">
      <div class="containder">
          <span class="text-muted">$$$ {{ text }}</span>
      </div>
  </footer>
</template>

<script>
export default {
  name: 'myFooter',
  props: ['text'] 
}
</script>