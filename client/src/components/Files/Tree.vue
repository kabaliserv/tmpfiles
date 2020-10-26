<template>
    <div>
        <div v-for="file in files" :key="file.id">
            <div
                class="file-view"
                @click.prevent="metaQuery(file.id)"
                @contextmenu.prevent="$refs.menu.open($event, { id: file.id })"
            >
                <div class="file-icon"></div>
                <div class="file-info">
                    <span>{{ file.name }}</span>
                    <span>{{ file.size }}</span>
                </div>
                <div class="file-option">
                    <span class="option-download"></span>
                </div>
            </div>
        </div>
        <vue-context ref="menu" v-slot="{ data }">
            <template v-if="data">
                <li>
                    <a @click.prevent="metaQuery(data.id)">
                        Metadata
                    </a>
                </li>
                <li>
                    <a @click.prevent="downloadQuery(data.id)">
                        Download
                    </a>
                </li>
            </template>
        </vue-context>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Emit, Vue } from "vue-property-decorator";
import VueContext from "vue-context";

// Style
import "vue-context/src/sass/vue-context.scss";

// Interface
// import { FileInfo } from "./types";

@Component({
    components: {
        VueContext,
    },
})
export default class TreeFiles extends Vue {
    @Prop()
    readonly files!: Array<TMPFiles.FileInfo>;

    @Emit()
    downloadQuery(id: string): string {
        return id
    }

    @Emit()
    metaQuery(id: string): string {
        return id
    }
}
</script>

<style lang="scss" scoped>
.file-view {
    display: flex;
    flex-direction: row;
    border-radius: 10px;
    border: solid 1px rgb(54, 54, 54);

    &:hover {
        background-color: rgb(114, 112, 112);
    }
}
.file-info {
    margin: 10px 0;
    display: flex;
    flex-direction: column;
}
</style>