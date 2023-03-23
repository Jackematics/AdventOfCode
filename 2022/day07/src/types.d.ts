import IDirectory from './directory/idirectory.interface';

export type File = {
  name: string;
  size: number;
};

export type Contents = {
  directories: IDirectory[];
  files: File[];
};
