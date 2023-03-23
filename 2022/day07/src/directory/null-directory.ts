import { File, Contents } from '../types';
import IDirectory from './idirectory.interface';

export default class NullDirectory implements IDirectory {
  public getChildDirectories(): IDirectory[] {
    return [];
  }

  getName(): string {
    return '';
  }

  getParent(): IDirectory {
    return {} as IDirectory;
  }

  addChildDirectory(targetDirName: string): IDirectory {
    return {} as IDirectory;
  }

  addFile(file: File): File {
    return {} as File;
  }

  cd(rawCdArg: string): IDirectory {
    return {} as IDirectory;
  }
  ls(): Contents {
    return {} as Contents;
  }

  toJson() {
    return {};
  }

  getDirectoriesBelowSizeCap(sizeCap: number): number[] {
    throw new Error('Method not implemented.');
  }
  getDirectorySize(): number {
    throw new Error('Method not implemented.');
  }
}
