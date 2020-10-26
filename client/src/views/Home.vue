<template>
    <div class="home">
        <v-overlay light :dark="false" :value="true">
            <MainUploader />
        </v-overlay>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import MainUploader from "@/components/Upload/Main.vue";

@Component({
    components: {
        MainUploader,
    },
})
export default class Home extends Vue {}
</script>

<style lang="scss" scoped>
.home {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: rgb(255, 255, 255);
}
</style>
