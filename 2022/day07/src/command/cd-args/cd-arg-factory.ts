import ICommand from '../icommand.interface';
import CdInChildDirectory from './cd-in-child-directory';
import CdOutOneLevel from './cd-out-one-level';
import CdToRoot from './cd-to-root';

export default class CdArgFactory {
  public static createCdArg(arg: string): ICommand {
    switch (arg) {
      case '..':
        return new CdOutOneLevel();
      case '/':
        return new CdToRoot();
      default:
        return new CdInChildDirectory(arg);
    }
  }
}
