<template>
    <div>
        <div v-if="error.password">Fail Auth</div><br>
        <input type="password" name="pwd" id="pwd" @keyup.enter="sendPassword" />
    </div>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue } from "vue-property-decorator";
import Axios from "axios";

@Component
export default class Auth extends Vue {

    error = {
        password: false
    }

    @Emit()
    sendToken(token: string): string {
        return token;
    }

    async sendPassword(event: Event): Promise<void> {
        let password = (<HTMLInputElement>event.target).value;
        try {
            const response = await Axios.post("../auth", {
                id: this.$route.params.id,
                password
            }, {
                headers: {
                    'content-type': 'application/json'
                }
            })
            console.log(response)
            this.$emit('sucess-auth', response.data)
        } catch (error) {
            if (error.response.status = 401){
                this.error.password = true
            }
        }
    }
}
</script>

<style scoped>
</style>