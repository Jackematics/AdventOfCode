import Directory from './directory';
import NullDirectory from './null-directory';

describe('Directory', () => {
  describe('cd', () => {
    it('should move in one level to a specified directory', () => {
      const directory = new Directory('/', new NullDirectory());
      directory.addChildDirectory('x');

      const childDir = directory.cd('x');

      expect(childDir.getName()).toBe('x');
    });

    it('should create a new child directory if specified directory does not exist', () => {
      const directory = new Directory('/', new NullDirectory());

      const childDir = directory.cd('x');

      expect(childDir.getName()).toBe('x');
    });

    it('should move out one level', () => {
      const directory = new Directory('/', new NullDirectory());
      const childDir = directory.addChildDirectory('x');

      const parent = childDir.cd('..');

      expect(parent).toBe(directory);
    });

    it('should error when trying to cd into the parent of the root directory', () => {
      const directory = new Directory('/', new NullDirectory());

      const error = () => directory.cd('..');

      expect(error).toThrow('directory has no parent');
    });

    it('should get source directory from source', () => {
      const directory = new Directory('/', new NullDirectory());

      const root = directory.cd('/');

      expect(root).toBe(directory);
    });

    it('should get source directory from child', () => {
      const directory = new Directory('/', new NullDirectory());
      const child = directory.addChildDirectory('a');

      const root = child.cd('/');

      expect(root).toBe(directory);
    });

    it('should get source directory from nest sub-child', () => {
      const directory = new Directory('/', new NullDirectory());
      const child = directory.addChildDirectory('a');
      const subChild = child.addChildDirectory('b');
      const subsubChild = subChild.addChildDirectory('c');

      const root = subsubChild.cd('/');

      expect(root).toBe(directory);
    });
  });

  describe('ls', () => {
    it('should list all files and directories in a directory', () => {
      const directory = new Directory('/', new NullDirectory());
      const childX = directory.addChildDirectory('x');
      const childY = directory.addChildDirectory('y');
      const fileAbc = directory.addFile({ name: 'abc', size: 1000 });
      const fileDef = directory.addFile({ name: 'def', size: 2000 });

      expect(directory.ls()).toStrictEqual({
        directories: [childX, childY],
        files: [fileAbc, fileDef],
      });
    });
  });

  describe('toJson', () => {
    it('should parse directory', () => {
      const directory = new Directory('/', new NullDirectory());

      expect(directory.toJson()).toStrictEqual({ '/': {} });
    });

    it('should parse directory, and a child file', () => {
      const directory = new Directory('/', new NullDirectory());
      directory.addFile({ name: 'abc', size: 1000 });

      expect(directory.toJson()).toStrictEqual({ '/': { abc: 1000 } });
    });

    it('should parse directory, and a child directory', () => {
      const directory = new Directory('/', new NullDirectory());
      directory.addChildDirectory('a');

      expect(directory.toJson()).toStrictEqual({ '/': { a: {} } });
    });

    it('should parse directory, and two child directories', () => {
      const directory = new Directory('/', new NullDirectory());
      directory.addChildDirectory('a');
      directory.addChildDirectory('b');

      expect(directory.toJson()).toStrictEqual({ '/': { a: {}, b: {} } });
    });

    it('should parse directory, two child directories and two files', () => {
      const directory = new Directory('/', new NullDirectory());
      directory.addChildDirectory('a');
      directory.addChildDirectory('b');
      directory.addFile({ name: 'abc', size: 1000 });
      directory.addFile({ name: 'def', size: 500 });

      expect(directory.toJson()).toStrictEqual({
        '/': { a: {}, b: {}, abc: 1000, def: 500 },
      });
    });

    it('should parse directory, two child directories, two sub child directories, and two files', () => {
      const directory = new Directory('/', new NullDirectory());
      const a = directory.addChildDirectory('a');
      directory.addChildDirectory('b');
      a.addChildDirectory('x');
      a.addChildDirectory('y');
      directory.addFile({ name: 'abc', size: 1000 });
      directory.addFile({ name: 'def', size: 500 });

      expect(directory.toJson()).toStrictEqual({
        '/': { a: { x: {}, y: {} }, b: {}, abc: 1000, def: 500 },
      });
    });

    it('should parse directory, two child directories with files, two sub child directories, and two files', () => {
      const directory = new Directory('/', new NullDirectory());
      const a = directory.addChildDirectory('a');
      directory.addChildDirectory('b');
      a.addChildDirectory('x');
      a.addChildDirectory('y');
      a.addFile({ name: '123', size: 12345 });
      a.addFile({ name: '456', size: 45678 });
      directory.addFile({ name: 'abc', size: 1000 });
      directory.addFile({ name: 'def', size: 500 });

      expect(directory.toJson()).toStrictEqual({
        '/': {
          a: { x: {}, y: {}, '123': 12345, '456': 45678 },
          b: {},
          abc: 1000,
          def: 500,
        },
      });
    });
  });

  describe('getDirectoriesBelowSizeCap()', () => {
    it('should add the size of its own files if below size cap', () => {
      const directory = new Directory('/', new NullDirectory());
      directory.addFile({ name: 'abc', size: 1000 });
      directory.addFile({ name: 'def', size: 2000 });

      expect(directory.getDirectoriesBelowSizeCap(100000)).toStrictEqual([
        3000,
      ]);
    });

    it('should not add the size of its own files if above size cap', () => {
      const directory = new Directory('/', new NullDirectory());
      directory.addFile({ name: 'abc', size: 200000 });
      directory.addFile({ name: 'def', size: 300 });

      expect(directory.getDirectoriesBelowSizeCap(100000)).toStrictEqual([]);
    });

    it('should add the size of its own files and child files if above size cap example 1', () => {
      const directory = new Directory('/', new NullDirectory());
      directory.addFile({ name: 'abc', size: 1000 });
      directory.addFile({ name: 'def', size: 2000 });
      const x = directory.addChildDirectory('x');
      x.addFile({ name: 'def', size: 3000 });
      x.addFile({ name: 'ghi', size: 4000 });

      expect(directory.getDirectoriesBelowSizeCap(100000)).toStrictEqual([
        10000, 7000,
      ]);
    });

    it('should add the size of its own files and child files if above size cap example 3', () => {
      const directory = new Directory('/', new NullDirectory());
      directory.addFile({ name: 'abc', size: 1000 });
      directory.addFile({ name: 'def', size: 2000 });

      const childOneDir = directory.addChildDirectory('childOneDir');
      directory.addChildDirectory('subChildDir');
      childOneDir.addFile({ name: 'def', size: 3000 });
      childOneDir.addFile({ name: 'ghi', size: 4000 });

      const subChildDir = childOneDir.addChildDirectory('subSubChildDir');
      subChildDir.addFile({ name: 'abc', size: 100 });

      expect(directory.getDirectoriesBelowSizeCap(100000)).toStrictEqual([
        10100, 7100, 100, 0,
      ]);
    });

    it('should add the size of its own files and child files if above size cap example 3', () => {
      const directory = new Directory('/', new NullDirectory());
      directory.addFile({ name: 'abc', size: 1000 });
      directory.addFile({ name: 'def', size: 2000 });

      const childOneDir = directory.addChildDirectory('childOneDir');
      childOneDir.addFile({ name: 'def', size: 3000 });
      childOneDir.addFile({ name: 'ghi', size: 4000 });
      const childTwoDir = directory.addChildDirectory('childTwoDir');
      childTwoDir.addFile({ name: 'testOne', size: 500 });
      childTwoDir.addFile({ name: 'testTwo', size: 1000 });

      const subChildDir = childOneDir.addChildDirectory('subChildDir');
      subChildDir.addFile({ name: 'abc', size: 100 });

      expect(directory.getDirectoriesBelowSizeCap(100000)).toStrictEqual([
        11600, 7100, 100, 1500,
      ]);
    });
  });
});
