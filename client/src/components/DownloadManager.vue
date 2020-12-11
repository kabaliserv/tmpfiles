<template>
  <div class="download-view">
    <div class="view-item">
      <div class="item-info" :title="meta.name">
        <span class="file-name">{{ meta.name }}</span>
        <span class="file-size">{{ getSize(meta.size) }}</span>
      </div>
      <div class="item-actions">
        <el-button type="primary" round plain @click="download">
          Download
        </el-button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Emit, Prop } from "vue-property-decorator";

@Component
export default class DownloadManagerComponent extends Vue {
  @Prop({ required: false })
  readonly meta!: TMPFiles.FileInfo;

  @Emit()
  download() {
    return;
  }

  getSize(value: number): string {
    const Size = {
      TB: 2 ** 40,
      GB: 2 ** 30,
      MB: 2 ** 20,
      KB: 2 ** 10,
    };
    let result = "";
    switch (true) {
      case value > Size.TB:
        result = (value / Size.TB).toFixed(1) + " To";
        break;
      case value > Size.GB:
        result = (value / Size.GB).toFixed(1) + " Go";
        break;
      case value > Size.MB:
        result = (value / Size.MB).toFixed(1) + " Mo";
        break;
      case value > Size.KB:
        result = (value / Size.KB).toFixed(1) + " Ko";
        break;
      default:
        result = value + " Octet";
    }
    return result;
  }
}
</script>

<style lang="scss" scoped>
.download-view {
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.view-item {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}
.item-info {
  display: flex;
  flex-direction: column;
}
.file-name {
  font-size: 20px;
  font-weight: bold;
  margin-bottom: 5px;
}
.file-size {
  color: rgba(0, 0, 0, 0.7);
}
.item-actions,
.item-info {
  margin: 0 20px;
}
</style>