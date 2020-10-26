<template>
    <div>
        <!-- FAKE INPUT FILE -->
        <input
            type="file"
            name="fakeinputfile"
            id="fakeinputfile"
            @change="onChangeFakeInput"
            ref="fakeInputFile"
            multiple
        />
        <v-card-title class="text-center">Fichier Temporaire</v-card-title>

        <v-divider class="mb-5"></v-divider>

        <v-card-text>
            <v-card class="tmpfiles--upload-card-items">
                <file
                    v-for="file in files"
                    :key="file.localID"
                    :file-item="file"
                    :url="urlTUS"
                />
            </v-card>
            <div class="d-flex justify-center">
                <v-btn
                    v-show="!uploadRunning"
                    class="mx-6"
                    fab
                    dark
                    small
                    :depressed="true"
                    color="indigo"
                    :disabled="uploadRunning"
                    @click="insertNewFiles"
                >
                    <v-icon dark> mdi-plus </v-icon>
                </v-btn>
            </div>
        </v-card-text>

        <v-card-text>
            <div class="tmpfiles--upload-card-options">
                <div class="upload-option">
                    <v-chip-group
                        v-model="expire"
                        active-class="deep-purple accent-4 white--text"
                        column
                    >
                        <v-chip
                            v-for="(value, index) in expireOptions"
                            :value="index"
                            :key="index"
                            :disabled="uploadRunning"
                            >{{ value }}</v-chip
                        >
                    </v-chip-group>

                    <v-container class="option-password">
                        <v-row>
                            <v-text-field
                                type="password"
                                label="Password"
                                :disabled="!showInputPassword || uploadRunning"
                            >
                            </v-text-field>

                            <v-spacer></v-spacer>

                            <v-switch
                                v-model="showInputPassword"
                                dense
                                inset
                                :disabled="uploadRunning"
                            >
                            </v-switch>
                        </v-row>
                    </v-container>
                </div>
            </div>
        </v-card-text>

        <v-spacer></v-spacer>
        <v-card-actions class="d-flex flex-row-reverse">
            <v-btn
                :loading="uploadRunning"
                :disabled="uploadRunning || files.length < 1"
                color="blue-grey"
                class="ma-2 white--text"
                fab
                small
                @click="startUpload()"
            >
                <v-icon dark> mdi-cloud-upload </v-icon>
            </v-btn>
        </v-card-actions>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import { nanoid } from "nanoid";
import { Observable, Subscription } from "rxjs";
import Axios, { AxiosResponse } from "axios";

// Components
import File from "@/components/Upload/Item.vue";

// Upload Manager
import { UploadManager, FileList, ExpireOptions } from "@/services/upload";

//Vue.use(Button)

@Component({
    components: {
        File,
    },
})
export default class Manager extends Vue {
    // Ref AutoCompletion
    $refs!: {
        fakeInputFile: HTMLFormElement;
        selectExpire: HTMLFormElement;
        inputPassword: HTMLFormElement;
        buttonSend: HTMLFormElement;
    };

    // Upload option handle
    readonly parallelUpload: number = 2;

    // URL send data to api
    urlTUS = "./upload/cache/";
    urlUpload = "./upload";

    // Array Files Item from send
    files: Array<TMPFiles.UploadItem> = FileList;

    // Config
    expireOptions: Array<string> = ExpireOptions;

    // Layout
    showInputPassword: boolean = false;
    uploadRunning: boolean = false;

    // Form
    password: string = "";
    expire: number = 0;

    // RXJS handles
    completeNotify$!: Subscription;

    created() {
        this.completeNotify$ = UploadManager.getComplete().subscribe({
            complete: () => setTimeout(this.sendFormUpload, 0),
        });
    }

    beforeDestroy() {
        this.completeNotify$.unsubscribe();
    }

    insertNewFiles(): void {
        this.$refs.fakeInputFile.click();
    }

    onChangeFakeInput(event: Event): void {
        const files = (<HTMLInputElement>event.target)!.files;
        if (files != (null || undefined)) {
            let t = files!.length;
            for (let i = 0; i < t; i++) {
                const file = files[i];
                console.log(file);
                UploadManager.register(file);
            }
        }
    }

    startUpload() {
        this.uploadRunning = true;
        UploadManager.start(this.parallelUpload);
    }

    async sendFormUpload(): Promise<void> {
        const expire = this.expire;
        const password = this.password;
        const filesid: Array<string> = [];
        this.files.forEach((item: TMPFiles.UploadItem) => {
            console.log(item.serverID);
            filesid.push(item.serverID);
        });

        const UploadID = await Axios.post(
            this.urlUpload, // URL
            {
                // Data
                expire,
                password,
                filesid,
            },
            {
                // Config
                headers: {
                    "Content-Type": "application/json",
                },
            }
        ).then((response: AxiosResponse) => response.data);
        console.log(UploadID);
    }
}
</script>

<style lang="scss" scoped>
#fakeinputfile {
    display: none;
}
</style>