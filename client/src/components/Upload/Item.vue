<template>
    <v-card class="file-item pa-0 elevation-0" elevation="0">
        <v-col class="pa-0">
            <v-progress-linear
                :value="progressStatus"
                background-color="white"
            ></v-progress-linear>
            <v-row class="full-with" no-gutters>
                <div class="item-info">
                    <span>{{ file.name }}</span>
                    <span>{{ file.size | size }}</span>
                </div>
                <div class="space" />
                <div class="item-buttons">
                    <button class="remove-button" @click="removeItem">X</button>
                </div>
            </v-row>
        </v-col>
    </v-card>
</template>

<script lang="ts">
import { Subscription } from "rxjs";
import { Upload } from "tus-js-client";
import { Component, Vue, Prop, Emit } from "vue-property-decorator";

import { Utils } from "./types";

// Upload Manager
import { UploadManager } from "@/services/upload";

@Component
export default class File extends Vue {
    @Prop({ required: true })
    readonly fileItem!: TMPFiles.UploadItem;

    @Prop({ required: true })
    readonly url!: string;

    id = this.fileItem.localID;
    file = this.fileItem.file;

    progressStatus: number = 0;
    showProgressBar: boolean = false;

    uploadManagerStart$!: Subscription;

    created() {
        this.uploadManagerStart$ = UploadManager.getStart().subscribe(
            (id: string) => {
                if (id === this.id) {
                    this.startUpload();
                }
            }
        );
    }

    beforeDestroy() {
        this.uploadManagerStart$.unsubscribe();
        // UploadRegister.unregister(this.fileItem.localID)
    }

    removeItem(): void {
        UploadManager.unregister(this.fileItem.localID);
    }

    complete() {}

    startUpload() {
        var upload = new Upload(this.file, {
            endpoint: this.url,
            retryDelays: [0, 3000, 5000, 10000, 20000],
            metadata: {
                filename: this.file.name,
                filetype: this.file.type,
                lastmodified: this.file.lastModified.toString(),
            },
            onError: (error) => {
                console.log("Failed because: " + error);
            },
            onProgress: (bytesUploaded, bytesTotal) => {
                this.progressStatus = (bytesUploaded / bytesTotal) * 100;
                //console.log(this.progressStatus);
            },
            onSuccess: () => {
                console.log(
                    "Download %s from %s",
                    // upload.file.name,
                    upload.url
                );
                let info: TMPFiles.CompleteInfo = {
                    localID: this.id,
                    serverID: upload.url?.split("/").pop() ?? "",
                };
                UploadManager.complete(info);
            },
        });

        // Start the upload
        upload.start();
    }
}
</script>

<style lang="scss" scoped>
.file-item {
    display: flex;
    flex-direction: row;
    border: solid 1px rgb(85, 84, 84);
    border-radius: 10px;
    margin: 5px 0;
    padding: 5px 10px;

    .space {
        flex-grow: 1;
    }
}
.item-info {
    display: flex;
    flex-direction: column;
}
.item-buttons {
    align-items: center;
}
</style>
