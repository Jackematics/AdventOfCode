import CdArgFactory from '../command/cd-args/cd-arg-factory';
import { Contents, File } from '../types';
import IDirectory from './idirectory.interface';

export default class Directory implements IDirectory {
  private readonly name: string;
  private readonly parentDirectory: IDirectory;
  private readonly childDirectories: Directory[] = [];
  private readonly files: File[] = [];

  constructor(name: string, parentDirectory: IDirectory) {
    this.name = name;
    this.parentDirectory = parentDirectory;
  }

  public getChildDirectories(): Directory[] {
    return this.childDirectories;
  }

  public getName(): string {
    return this.name;
  }

  public getParent(): IDirectory {
    if (this.parentDirectory.getName() === '') {
      throw new Error('directory has no parent');
    }

    return this.parentDirectory;
  }

  public addChildDirectory(targetDirName: string): Directory {
    const childDir = new Directory(targetDirName, this);
    this.childDirectories.push(childDir);

    return childDir;
  }

  public addFile(file: File): File {
    this.files.push(file);

    return file;
  }

  public cd(rawCdArg: string): IDirectory {
    const cdArg = CdArgFactory.createCdArg(rawCdArg);
    return cdArg.execute(this);
  }

  public ls(): Contents {
    return { directories: this.childDirectories, files: this.files };
  }

  public toJson(): any {
    const dirJson = this.childDirectories
      .map((dir) => dir.toJson())
      .reduce(this.dirJsonReducer, {});
    const files = this.files.reduce(this.fileJsonReducer, {});

    return { [this.name]: { ...dirJson, ...files } };
  }

  private dirJsonReducer(acc: any, dir: any) {
    return { ...acc, [Object.keys(dir)[0]]: Object.values(dir)[0] };
  }

  private fileJsonReducer(acc: any, file: File): any {
    return { ...acc, [file.name]: file.size };
  }

  public getDirectoriesBelowSizeCap(sizeCap: number): number[] {
    const currentDirFileSize = this.getDirectorySize();
    const sizes = currentDirFileSize <= sizeCap ? [currentDirFileSize] : [];

    if (this.childDirectories.length === 0) return sizes;

    this.childDirectories.map((dir) => {
      sizes.push(...dir.getDirectoriesBelowSizeCap(sizeCap));
    });

    return sizes;
  }

  public getDirectorySize(): number {
    const currentDirSize = this.files.reduce(
      (acc: number, file: File) => acc + file.size,
      0
    );
    return this.childDirectories.reduce(
      (acc: number, directory: Directory) => acc + directory.getDirectorySize(),
      currentDirSize
    );
  }
}
