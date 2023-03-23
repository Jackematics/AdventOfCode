import * as fs from 'fs';
import Motion from './motion/motion';
import MotionFactory from './motion/motion-factory';

export default class FileProcessor {
  public static processInput(inputPath: string): Motion[] {
    return fs
      .readFileSync(inputPath)
      .toString()
      .split('\n')
      .map((rawMotion: string) => {
        const split = rawMotion.split(' ');
        return MotionFactory.createMotion(split[0], split[1]);
      });
  }
}
