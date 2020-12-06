<template>
  <div class="upload-view">
    <input
      type="file"
      name="fakeinputfile"
      id="fakeinputfile"
      @change="onChangeFakeInput"
      ref="fakeInputFile"
    />
    <span class="view-title">Envoi de fichers</span>
    <el-divider></el-divider>
    <div class="upload-files">
      <div
        class="file-item-layout"
        v-for="(file, key) in files"
        :key="key"
        :title="file.file.name"
      >
        <div class="file-item">
          <el-col :span="22">
            <div class="item-name">{{ file.file.name }}</div>
          </el-col>
          <el-col :span="2">
            <div class="file-action" @click="onDeleteFile(file.localID)">
              <i class="el-icon-circle-close"></i>
            </div>
          </el-col>
        </div>
      </div>
      <div class="upload-button-add" v-show="!fileUpload">
        <el-button
          type="primary"
          icon="el-icon-plus"
          circle
          @click="selectFile"
        ></el-button>
      </div>
    </div>
    <div class="upload-options">
      <span>Options:</span>
      <div class="option-expire">
        <span class="expire-text">Durée</span>
        <span class="expire-action">
          <el-select v-model="expireValue" size="large">
            <el-option
              v-for="(item, key) in expireOptions"
              :key="key"
              :label="item"
              :value="key"
            >
            </el-option>
          </el-select>
        </span>
      </div>
      <div class="option-password">
        <span class="password-text">Password</span>
        <span class="password-action">
          <el-switch v-model="passwordActive"></el-switch>
          <div class="password-input">
            <transition name="el-zoom-in-top">
              <el-input
                v-show="passwordActive"
                placeholder="Entrez votre mot de passe"
                v-model="passwordText"
                show-password
              ></el-input>
            </transition>
          </div>
        </span>
      </div>
    </div>
    <div class="upload-actions">
      <el-button
        type="primary"
        style="width: 100%"
        size="medium"
        round
        :disabled="!fileUpload"
        @click="startUpload"
      >
        <font-awesome-icon icon="upload" style="font-size: 1.2rem" size="2x" />
        <span style="font-size: 18px; margin-left: 7px">Upload</span>
      </el-button>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Emit } from "vue-property-decorator";

// Services
import { UploadManager, FileList, HasFile } from "@/services/uploads-manager";

@Component
export default class UploadManagerComponent extends Vue {
  // Ref AutoCompletion
  $refs!: {
    fakeInputFile: HTMLFormElement;
    selectExpire: HTMLFormElement;
    inputPassword: HTMLFormElement;
    buttonSend: HTMLFormElement;
  };
  expireOptions = ["5min", "10min", "1h", "6h", "1J", "3J"];
  expireValue = 0;
  passwordActive = false;
  passwordText = "";
  fileUpload = false;
  files = FileList;
  uploadRunning = false
  progress = 0; // state upload

  selectFile(): void {
    this.$refs.fakeInputFile.click();
  }

  onChangeFakeInput(event: { target: HTMLInputElement }): void {
    const files = event.target.files;
    const file = files?.item(0);
    if (file) {
      UploadManager.register(file);
      this.fileUpload = true;
      this.$refs.fakeInputFile.value = ""
    }
  }

  onDeleteFile(id: string) {
    UploadManager.unregister(id);
    this.fileUpload = false;
  }
  startUpload() {
    const uploadOptions: TMPFiles.UploadOptions = {
      auth: this.passwordActive,
      password: this.passwordText,
      expire: this.expireValue,
    }
        this.uploadRunning = true;
        UploadManager.start(uploadOptions);
  }

  @Emit()
  uploadeRunning() {
    return
  }
}
</script>

<style lang="scss" scoped>
#fakeinputfile {
  display: none;
}
.upload-view {
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.view-title {
  font-size: 24px;
  font-weight: bold;
  text-align: left;
  margin-left: 10px;
}
.upload-button-add {
  text-align: center;
}
.upload-actions {
  box-sizing: border-box;
  text-align: center;
  padding: 0 60px;
  width: 100%;
  margin-top: 20px;
}
.upload-options {
  & > div {
    margin: 20px 20px 10px 20px;
    & [class$="-text"] {
      font-size: 15px;
    }
  }
  & > span {
    font-size: 15px;
    font-weight: bold;
  }

  .option-expire {
    display: flex;
    flex-direction: row;
    align-items: center;
    margin-top: 10px;
    .expire-text {
      margin-right: 5px;
    }
    .expire-action {
      width: auto;
      max-width: 100px;
    }
  }
  .option-password {
    height: 100px;
    .password-text {
      margin-right: 5px;
    }
    .password-input {
      box-sizing: border-box;
      // padding: 0 20px;
      margin: 20px 0 0 0;
    }
  }
}
.upload-files {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 50px;
  margin-bottom: 20px;
}
.file-item-layout {
  box-sizing: border-box;
  width: 100%;
  padding: 5px 10px;
  border-radius: 15px;
  background-color: rgba(0, 0, 0, 0.05);
  overflow: hidden;
  &:hover {
    background-color: rgba(0, 0, 0, 0.1);
  }
  .file-action {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    cursor: pointer;
    .el-icon-circle-close {
      font-size: 25px
    }
  }
  .file-item {
    margin: 2px;
    .item-name {
      max-width: 240px;
      margin: 2px 0;
      font-weight: 600;
      text-overflow: ellipsis;
      overflow: hidden;
      white-space: nowrap;
    }

    .item-size {
      font-size: 15px;
      opacity: 0.8;
    }
  }
}
</style>
