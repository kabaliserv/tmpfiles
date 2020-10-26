<template>
    <v-card rounded="lg" elevation="1" class="px-5 py-6 my-12">
            <Manager />
    </v-card>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { VueClass } from "vue-class-component/lib/declarations";

// Components
import Manager from "@/components/Upload/Manager.vue";

@Component({
    components: {
        Manager,
    },
})
export default class MainUploader extends Vue {}
</script>

<style lang="scss" scoped>
.main-uploader {
    height: 80vh;
    width: 50vw;
    max-width: 600px;
    padding: 30px 20px;
}
</style>
