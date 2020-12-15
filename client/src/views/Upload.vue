<template>
  <div class="upload-background">
    <transition name="el-fade-in">
      <div class="upload-layout" v-if="managerView">
        <UploadManagerComponent />
      </div>
    </transition>
    <transition name="el-fade-in">
      <div class="progress-upload" v-show="progessBarview">
        <el-progress
          type="circle"
          :percentage="progressUpload"
          v-show="uploadRunning"
        >
        </el-progress>
        <el-progress
          type="circle"
          :percentage="100"
          status="success"
          v-show="doneUpload"
        >
        </el-progress>
      </div>
    </transition>
    <transition name="el-fade-in">
      <div class="link-to-upload" v-show="linkView">
        <a :href="linkUpload">{{ linkUpload }}</a>
      </div>
    </transition>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { Upload } from "tus-js-client";
import UploadManagerComponent from "@/components/UplaodManager.vue";
import { UploadManager, FileList } from "@/services/uploads-manager";
import config from "@/config/config";

@Component({
  components: {
    UploadManagerComponent,
  },
})
export default class UploadView extends Vue {
  managerView = true;
  doneUpload = false;
  uploadRunning = false;
  progessBarview = false;
  progressUpload = 0;

  linkView = false;
  linkUpload = "";

  mounted() {
    UploadManager.getUploadID().subscribe({
      next: (id: string) => {
        UploadManager.unregister(FileList[0].localID)
        this.$nextTick(() => {
          this.uploadRunning = false;
          this.doneUpload = true;
        });
        setTimeout(() => {
          this.$nextTick(() => {
            this.uploadRunning = false;
            this.progessBarview = false;
          });
        }, 2000);
        this.generateLinkUpload(id);
        setTimeout(() => {
          this.linkView = true;
        }, 3500);
      },
    });
    UploadManager.getStart().subscribe((options: TMPFiles.UploadOptions) => {
      this.managerView = false;
      setTimeout(() => {
        this.uploadRunning = true;
        this.progessBarview = true;
      }, 500);
      setTimeout(() => this.StartUpload(options), 1000)
    });
  }

  generateLinkUpload(uploadID: string) {
    this.linkUpload = `${location.protocol}//${location.host}/file/${uploadID}`;
  }

  StartUpload(options: TMPFiles.UploadOptions) {
    const tusURL = config.API_URL + config.TUS_URL;
    const upload = new Upload(FileList[0].file, {
      endpoint: tusURL,
      retryDelays: [0, 3000, 5000, 10000, 20000],
      metadata: {
        filename: FileList[0].file.name,
        filetype: FileList[0].file.type,
        lastmodified: FileList[0].file.lastModified.toString(),
      },
      onError: (error) => {
        console.log("Failed because: " + error);
        this.$notify.error({
          title: "Upload",
          message: "Une erreur cest produit lors de l'upload veillez réessayer ulterieurment",
          duration: 0
        });
      },
      onProgress: (bytesUploaded, bytesTotal) => {
        this.progressUpload = Math.floor((bytesUploaded / bytesTotal) * 100);
      },
      onSuccess: () => {
        const serverID = upload.url?.split("/").pop() ?? "";
        options.filesid = [serverID];
        UploadManager.complete(options);
      },
    });

    // Start the upload
    upload.start();
  }
}
</script>

<style scoped>
.upload-background {
  position: absolute;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
  top: 0;
  left: 0;
  box-sizing: border-box;
}
.upload-layout {
  padding: 30px 50px 50px 50px;
  width: 350px;
  border-radius: 20px;
  border: 1px solid rgba(0, 0, 0, 0.1);
  box-shadow: 5px 5px 5px 0px rgba(0, 0, 255, 0.2);
}
</style>