import { Direction } from '../enum';
import Motion from './motion';

export default class MotionFactory {
  public static createMotion(
    rawDirection: string,
    rawDistance: string
  ): Motion {
    return new Motion(this.getDirection(rawDirection), parseInt(rawDistance));
  }

  private static getDirection(rawDirection: string): Direction {
    switch (rawDirection) {
      case 'U':
        return Direction.Up;
      case 'R':
        return Direction.Right;
      case 'D':
        return Direction.Down;
      case 'L':
        return Direction.Left;
      default:
        throw new Error('Invalid motion');
    }
  }
}
