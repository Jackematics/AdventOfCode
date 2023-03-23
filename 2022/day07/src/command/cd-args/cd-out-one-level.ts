import IDirectory from '../../directory/idirectory.interface';
import ICommand from '../icommand.interface';

export default class CdOutOneLevel implements ICommand {
  public execute(sourceDir: IDirectory): IDirectory {
    return sourceDir.getParent();
  }
}
