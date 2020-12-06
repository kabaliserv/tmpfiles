declare namespace TMPFiles {
  export interface UploadItem {
    localID: string;
    serverID: string;
    complete: boolean;
    pending: boolean;
    file: File;
  }

  interface UploadInfo {
      id: string
      size: number
      
  }

  export interface UploadOptions {
    auth: boolean;
    password: string;
    expire: number;
    filesid?: Array<string>;
  }

  export interface FileInfo {
    id: string
    name: string;
    type: string;
    size: number;
  }

  interface Meta {
    type: string;
  }

  export interface CompleteInfo {
    localID: string;
    serverID: string;
  }
}
