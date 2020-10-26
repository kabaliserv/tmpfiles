<template>
    <div class="meta-view">
        <div class="meta-main">
            <div class="meta-icon"></div>
            <div class="meta-info">
                <div class="info-name">{{ file.name }}</div>
                <div class="info-size">{{ file.size }}</div>
            </div>
            <div class="meta-action">
                <div class="action-view">
                    <div class="action-option">
                        <div class="option-download">Download</div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue } from "vue-property-decorator";

@Component
export default class MetaFile extends Vue {
    @Prop({ required: true })
    readonly file: TMPFiles.FileInfo;

    @Emit()
    downloadQuery(id: string): string {
        return id
    }
}
</script>

<style lang="scss" scoped>
</style>