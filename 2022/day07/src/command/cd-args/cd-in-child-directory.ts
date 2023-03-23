import IDirectory from '../../directory/idirectory.interface';
import ICommand from '../icommand.interface';

export default class CdInChildDirectory implements ICommand {
  private readonly targetDirName: string;

  constructor(targetDirName: string) {
    this.targetDirName = targetDirName;
  }

  public execute(sourceDir: IDirectory): IDirectory {
    const childDir = sourceDir
      .getChildDirectories()
      .find((dir: IDirectory) => dir.getName() === this.targetDirName);

    if (!childDir) {
      return sourceDir.addChildDirectory(this.targetDirName);
    }

    return childDir;
  }
}
