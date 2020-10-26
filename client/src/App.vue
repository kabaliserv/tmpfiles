<template>
    <v-app>
        <v-app-bar flat absolute color="white">
            <div class="d-flex align-center">
                <span>TMPFiles</span>
            </div>

            <v-spacer></v-spacer>
        </v-app-bar>

        <v-main>
            <router-view></router-view>
        </v-main>
    </v-app>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
    name: "App",
    data: () => ({}),
});
</script>
