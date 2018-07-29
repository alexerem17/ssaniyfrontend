<template>
    <div id="login">
        <h1>Войти</h1>
        <hr>
        <div class="col-xs-3">
        <input type="text" class="form-control" name="username" v-model="input.username" placeholder="Логин" />
        </div>
        <div class="col-xs-3">
        <input type="password" class="form-control" name="password" v-model="input.password" placeholder="Пароль" />
        </div>
        <button type="button" class="btn btn-default" v-on:click="login()">Войти</button>
    </div>
</template>

<script>
    export default {
        name: 'Login',
        data() {
            return {
                input: {
                    username: "",
                    password: ""
                }
            }
        },
        methods: {
            login() {
                if(this.input.username != "" && this.input.password != "") {
                    if(this.input.username == this.$parent.mockAccount.username && this.input.password == this.$parent.mockAccount.password) {
                        this.$emit("authenticated", true);
                        this.$router.replace({ name: "secure" });
                    } else {
                        console.log("Пользователь авторизован");
                    }
                } else {
                    console.log("A username and password must be present");
                }
            }
        }
    }
</script>
