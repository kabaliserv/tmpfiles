<template>
  <div class="download-backgound" v-loading="processAuth">
    <div class="download-layout" v-if="initialize && !err404">
      <div v-if="!auth">
        <DownloadManagerComponent :meta="meta" @download="startDownload" />
      </div>
      <div class="auth-layout" v-show="auth">
        <el-input
          placeholder="Mot de passe"
          v-model="password"
          show-password
          autofocus
          ref="inputpassword"
          @keyup.enter.native="getAuth"
        >
        </el-input>
        <div class="auth-action">
          <el-button
            type="primary"
            icon="el-icon-arrow-right"
            round
            :disabled="password === '' || processAuth"
            @click="getAuth"
          ></el-button>
        </div>
      </div>
    </div>
    <NotFound v-if="err404"/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import DownloadManagerComponent from "@/components/DownloadManager.vue";
import { GetMeta, GetAuth, Token } from "@/services/download-manager";
import NotFound from '@/views/Errors/NotFound.vue';

@Component({
  components: {
    DownloadManagerComponent,
    NotFound,
  },
})
export default class FileView extends Vue {
  $refs!: {
    inputpassword: HTMLFormElement;
  };
  initialize = false;
  auth = false;
  uploadAuth = false
  password = "";
  errPassword = false;
  processAuth = false;
  meta!: TMPFiles.FileInfo;
  uploadID = this.$route.params.id;

  err404 = false

  async mounted() {
    //InitInterceptor()
    try {
      const response = await GetMeta(this.uploadID);
      this.auth = false;
      this.meta = response.data.files[0];
    } catch (err) {
      if (err.response.status === 404) {
        this.err404 = true
      }
      if (err.response.status === 401) {
        this.auth = true;
        this.uploadAuth = true
        this.$nextTick(() => {
          this.$refs.inputpassword.focus();
        });
      }
    }
    this.initialize = true;
  }

  async getAuth() {
    if (this.password === "") return
    const auth = await GetAuth(this.uploadID, this.password);
    if (auth) {
      const response = await GetMeta(this.uploadID);
      this.auth = false;
      this.meta = response.data.files[0];
      return;
    } else {
      this.$notify.error({
        title: "Error",
        message: "Mot de passe incorect",
      });
    }
  }

  startDownload() {
    const link = document.createElement("a")
    link.href = `https://${location.host}/d/${this.uploadID}?f=${this.meta.id}`
    if (this.uploadAuth) {
      link.href += `&t=${Token}`
    }
    link.click()
    /* let link = `https://${location.host}/d/${this.uploadID}?f=${this.meta.id}`
    if (this.uploadAuth) {
      link += `&t=${Token}`
    }
    window.open(link, '_blank'); */
  }
}
</script>

<style lang="scss" scoped>
.download-backgound {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
  top: 0;
  left: 0;
  box-sizing: border-box;
}

.download-layout {
  padding: 20px;
  max-width: 700px;
  border-radius: 20px;
  border: 1px solid rgba(0, 0, 0, 0.1);
  box-shadow: 5px 5px 5px 0px rgba(0, 0, 255, 0.2);
}

.auth-layout {
  display: flex;
  flex-direction: row;
  & > .auth-action {
    margin: 0 0 0 25px;
  }
}
</style>