import { Subject } from 'rxjs'
import { nanoid } from 'nanoid'

const starFilesID = new Subject<string>()
const completeFilesID = new Subject<TMPFiles.CompleteInfo>()

export const FileList: Array<TMPFiles.UploadItem> = []

export const UploadManager = {
    register: (file: File) => {
        let fileUpload: TMPFiles.UploadItem = {
            localID: nanoid(10),
            serverID: "",
            pending: false,
            complete: false,
            file
        }
        FileList.push(fileUpload)
    },
    unregister: (id: string) => {
        let fileindex = -1
        FileList.every((item: TMPFiles.UploadItem, index: number) => {
            if (item.localID === id) {
                fileindex = index
                return false
            }
            return true
        })
        if (fileindex < 0) return;
        FileList.splice(fileindex, 1);
    },
    start: (num: number) => {
        for (let i = 0; i < num; i++) {
            if(FileList[i] !== undefined) {
                sendStart(FileList[i].localID)
            }
        }
    },
    complete: (item: TMPFiles.CompleteInfo) => completeFilesID.next(item),
    getStart: () => starFilesID.asObservable(),
    getComplete: () => completeFilesID.asObservable(),
}

function sendStart(id :string) {
    FileList.every((item: TMPFiles.UploadItem, index: number) => {
        if (item.localID === id) {
            FileList[index].pending = true
            return false
        }
        return true
    })
    starFilesID.next(id)
}

function sendCompleteUpload() {
    completeFilesID.complete()
}

UploadManager.getComplete().subscribe({
    next: (completeItem: TMPFiles.CompleteInfo) => {
        FileList.every((item: TMPFiles.UploadItem, index: number) => {
            if (item.localID === completeItem.localID) {
                FileList[index].serverID = completeItem.serverID
                FileList[index].pending = false
                FileList[index].complete = true
                return false
            }
            return true
        })
        const nextUpload = FileList.filter((item: TMPFiles.UploadItem) => !item.pending && !item.complete)[0]
        if (nextUpload !== undefined) {
            sendStart(nextUpload.localID)
        } else {
            const pendingUpload = FileList.filter((item: TMPFiles.UploadItem) => item.pending)
            if (pendingUpload.length < 1) {
                sendCompleteUpload()
            }
        }
    },
})

export const ExpireOptions = ['10 Min', '30 Min', '1H', '6H', '1J', '3J']