import Directory from '../directory/directory';
import IDirectory from '../directory/idirectory.interface';

export default interface ICommand {
  execute(sourceDir: Directory): IDirectory;
}
