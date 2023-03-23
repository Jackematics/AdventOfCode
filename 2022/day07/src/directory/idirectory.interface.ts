import { Contents, File } from '../types';

export default interface IDirectory {
  getChildDirectories(): IDirectory[];
  getName(): string;
  getParent(): IDirectory;
  addChildDirectory(targetDirName: string): IDirectory;
  addFile(file: File): File;
  cd(rawCdArg: string): IDirectory;
  ls(): Contents;
  toJson(): any;
  getDirectoriesBelowSizeCap(sizeCap: number): number[];
  getDirectorySize(): number;
}
