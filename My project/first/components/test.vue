<template>
    <section>
        <nav class="navbar-expand-x1 navbar-dark">
            <div class="collapse" id="navbarToggleExternalContent">
                <div class="bg-dark p-4">
                    <div class="container-fluid bg-dark">
                        <nav class="navbar navbar-dark bg-dark">
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
                        </nav>
                    </div>
                </div>
            </div>
        </nav>
        <nav class="navbar navbar-dark bg-dark">
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarToggleExternalContent" aria-controls="navbarToggleExternalContent"
                aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
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