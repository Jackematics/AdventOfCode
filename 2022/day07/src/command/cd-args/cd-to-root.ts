import IDirectory from '../../directory/idirectory.interface';
import ICommand from '../icommand.interface';

export default class CdToRoot implements ICommand {
  public execute(sourceDir: IDirectory): IDirectory {
    return this.cdToRootRecursively(sourceDir);
  }

  private cdToRootRecursively(sourceDir: IDirectory): IDirectory {
    if (sourceDir.getName() === '/') {
      return sourceDir;
    }

    return this.cdToRootRecursively(sourceDir.getParent());
  }
}
