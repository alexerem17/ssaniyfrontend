<template>
    <section id="app">
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNavAltMarkup" aria-controls="navbarNavAltMarkup"
        aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNavAltMarkup">
        <div class="container-fluid bg-dark">
          <div class="navbar">
            <div class="row col-md-4" id="navbarSearch">
              <form class="form-inline">
                <input class="form-control mr-sm-2" type="search" placeholder="Поиск..." aria-label="Search">
                <button class="btn btn-outline-success my-2 my-sm-0" @click="onClick()">Найти</button>
              </form>
            </div>
            <div class="col text-white text-center col-md-4">
              <p>{{ date }}</p>
            </div>
            <div class="col text-white text-center col-md-4">
              <span>$ {{ this.username }}</span>
            </div>
          </div>
        </div>
      </div>
    </nav>
  </section>
</template>


<script>
    export default {
        name: 'myHeader',
        data: () => ({
            x: 0,
            date: null,
            username: 'Ну ты и пидор'
        }),
        methods: {
            onClick() { alert("анус себе ткни") },
            startTime() {
                var tm = new Date();
                var h = tm.getHours();
                var m = tm.getMinutes();
                m = this.checkTime(m);
                this.date = `${h}:${m}`;
            },
            checkTime(i) {
                if (i < 10) {
                    i = "0" + i;
                }
                return i;
            }
        },
        mounted() {
            setInterval(this.startTime, 1000)
        }
    }
</script>