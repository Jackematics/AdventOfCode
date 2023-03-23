import * as fs from 'fs';
import Directory from './directory/directory';
import IDirectory from './directory/idirectory.interface';
import NullDirectory from './directory/null-directory';

export default class FileProcessor {
  public static processInput(inputPath: string): IDirectory {
    const rawCommands = fs.readFileSync(inputPath).toString().split('\n');
    const fileSystem = this.mapFileSystemFromCommands(rawCommands);

    return fileSystem;
  }

  private static mapFileSystemFromCommands(
    rawTerminalOutput: string[]
  ): IDirectory {
    let currentDir: IDirectory = new Directory('/', new NullDirectory());
    rawTerminalOutput.map((terminalOutput) => {
      currentDir = this.cdIfCdCommand(terminalOutput, currentDir);

      if (!terminalOutput.includes('$')) {
        this.addDirectoryIfDirectory(terminalOutput, currentDir);
        this.addFileOtherwise(terminalOutput, currentDir);
      }
    });

    return currentDir.cd('/');
  }

  private static cdIfCdCommand(
    command: string,
    currentDir: IDirectory
  ): IDirectory {
    return command.includes('$ cd')
      ? currentDir.cd(command.split(' ')[2])
      : currentDir;
  }

  private static addDirectoryIfDirectory(
    terminalOutput: string,
    currentDir: IDirectory
  ): void {
    if (terminalOutput.includes('dir')) {
      const directoryName = terminalOutput.split(' ')[1];
      if (this.childDirectoryNotInParent(directoryName, currentDir))
        currentDir.addChildDirectory(directoryName);
    }
  }

  private static childDirectoryNotInParent(
    directoryName: string,
    currentDir: IDirectory
  ): boolean {
    return !currentDir
      .ls()
      .directories.find((directory) => directory.getName() === directoryName);
  }

  private static addFileOtherwise(
    terminalOutput: string,
    currentDir: IDirectory
  ): void {
    if (!terminalOutput.includes('dir')) {
      const [size, name] = terminalOutput.split(' ');
      if (this.fileNotInDirectory(name, currentDir))
        currentDir.addFile({ name, size: parseInt(size) });
    }
  }

  private static fileNotInDirectory(
    name: string,
    currentDir: IDirectory
  ): boolean {
    return !currentDir.ls().files.find((file) => file.name === name);
  }
}
