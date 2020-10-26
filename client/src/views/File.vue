<template>
    <div>
        <div class="file-container" v-if="existUpload && authValid">
            <Tree :files="treeFiles" @download-query="downloadFile" @meta-query="getMetaFile" />
            <Meta :file="meta" />
        </div>
        <NotFound v-if="!existUpload" />
        <Auth v-if="requireAuth && !authValid" @sucess-auth="authSucess" />
    </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Axios, { AxiosError, AxiosInstance, AxiosRequestConfig } from "axios";
import jwt_decode from "jwt-decode";

//Components
import Tree from "@/components/Files/Tree.vue";
import Meta from "@/components/Files/Meta.vue";
import NotFound from "@/components/NotFound.vue";
import Auth from "@/components/Auth.vue";


@Component({
    components: {
        Tree,
        Meta,
        NotFound,
        Auth,
    },
})
export default class File extends Vue {
    // Status Upload
    existUpload: boolean = true;
    requireAuth: boolean = false;

    // Auth
    authValid: boolean = false;
    uploadToken: string = "";
    authError: object = {
        password: false,
    };

    // Upload
    uploadID: string = this.$route.params.id;

    // Tree
    treeFiles: Array<TMPFiles.FileInfo> = [];
    // Meta
    meta: TMPFiles.FileInfo

    caca: string


    created() {
        Axios.interceptors.request.use(
            (config) => {
                if (this.uploadToken) {
                    config.headers.Authorization = `Bearer ${this.uploadToken}`;
                }
                return config;
            },
            (err) => {
                return Promise.reject(err);
            }
        );
    }

    async beforeMount(): Promise<any> {
        try {
            console.log("toto");
            this.treeFiles = await this.getMetaFiles();
            this.authValid = true;
            console.log(this.treeFiles);
        } catch (err) {
            if (err.response.status === 404) {
                this.existUpload = false;
            }
            if (err.response.status === 401) {
                this.requireAuth = true;
                this.authValid = false;
                console.log(err.response.status);
            }
        }
    }
    async getMetaFiles(): Promise<any> {
        const response = await Axios.get("../meta/" + this.uploadID);
        this.getMetaFile(response.data.files[0].id)
        return response.data.files;
    }

    async authSucess(token: string): Promise<any> {
        this.uploadToken = token;
        this.treeFiles = await this.getMetaFiles();
        this.authValid = true;
    }

    validateToken(): boolean {
        if (this.uploadToken) {
            let decode: any = jwt_decode(this.uploadToken);
            if (decode.exp < decode.iat) return false;
        }
        return true;
    }

    downloadFile(id: string) {
        let link = document.createElement("a");
        let url = `../d/${this.uploadID}?f=${id}`;
        if (this.uploadToken) {
            if (!this.validateToken()) {
                this.authValid = false;
                this.requireAuth = true;
                return;
            }
            url += `&t=${this.uploadToken}`;
        }
        link.href = url;
        link.click();
    }

    getMetaFile(id: string): void {
        this.meta = this.treeFiles.filter((file: TMPFiles.FileInfo) => file.id === id)[0]
    }
}
</script>

<style lang="scss" scoped>
</style>