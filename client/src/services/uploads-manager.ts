import { Subject } from "rxjs";
import { nanoid } from "nanoid";
import Axios, { AxiosResponse } from "axios";
import config from "@/config/config";

const starFilesID = new Subject<TMPFiles.UploadOptions>();
const completeFilesID = new Subject<TMPFiles.UploadOptions>();
const uploadID = new Subject<string>();

export const FileList: Array<TMPFiles.UploadItem> = [];
export let HasFile = false;

export const UploadManager = {
  register: (file: File) => {
    const fileUpload: TMPFiles.UploadItem = {
      localID: nanoid(10),
      serverID: "",
      pending: false,
      complete: false,
      file,
    };
    FileList.push(fileUpload);
    HasFile = true;
  },
  unregister: (id: string) => {
    let fileindex = -1;
    FileList.every((item: TMPFiles.UploadItem, index: number) => {
      if (item.localID === id) {
        fileindex = index;
        return false;
      }
      return true;
    });
    if (fileindex < 0) return;
    FileList.splice(fileindex, 1);
    HasFile = false;
  },
  start: (options: TMPFiles.UploadOptions) => starFilesID.next(options),
  complete: (options: TMPFiles.UploadOptions) => completeFilesID.next(options),
  sendUploadID: (id: string) => uploadID.next(id),
  getStart: () => starFilesID.asObservable(),
  getComplete: () => completeFilesID.asObservable(),
  getUploadID: () => uploadID.asObservable(),
};

function sendCompleteUpload() {
  completeFilesID.complete();
}

const uploadPath = config.API_URL + "upload";

UploadManager.getComplete().subscribe({
  next: async (options: TMPFiles.UploadOptions) => {
    const UploadID = await Axios.post(
      uploadPath, // URL
      options,
      {
        // Config
        headers: {
          "Content-Type": "application/json",
        },
      }
    ).then((response: AxiosResponse) => response.data);
    UploadManager.sendUploadID(UploadID)
  },
});



/* UploadManager.getStart().subscribe(() => {
    startUpload();
  }); */
