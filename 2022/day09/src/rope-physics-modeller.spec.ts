import RopePhysicsModeller from './rope-physics-modeller';
import Motion from './motion/motion';
import { Direction } from './enum';

describe('RopePhysicsModeller', () => {
  it('should count the number of positions a single tail visited at least once', () => {
    const motions: Motion[] = [
      new Motion(Direction.Right, 4),
      new Motion(Direction.Up, 4),
      new Motion(Direction.Left, 3),
      new Motion(Direction.Down, 1),
      new Motion(Direction.Right, 4),
      new Motion(Direction.Down, 1),
      new Motion(Direction.Left, 5),
      new Motion(Direction.Right, 2),
    ];

    expect(RopePhysicsModeller.calculateVisitedPositions(motions, 2)).toBe(13);
  });

  it('should count the number of positions a single tail visited at least once', () => {
    const motions: Motion[] = [
      new Motion(Direction.Right, 4),
      new Motion(Direction.Up, 4),
      new Motion(Direction.Left, 3),
      new Motion(Direction.Down, 1),
      new Motion(Direction.Right, 4),
      new Motion(Direction.Down, 1),
      new Motion(Direction.Left, 5),
      new Motion(Direction.Right, 2),
    ];

    expect(RopePhysicsModeller.calculateVisitedPositions(motions, 10)).toBe(1);
  });

  it('should count the number of positions the final tail visited at least once', () => {
    const motions: Motion[] = [
      new Motion(Direction.Right, 5),
      new Motion(Direction.Up, 8),
      new Motion(Direction.Left, 8),
      new Motion(Direction.Down, 3),
      new Motion(Direction.Right, 17),
      new Motion(Direction.Down, 10),
      new Motion(Direction.Left, 25),
      new Motion(Direction.Up, 20),
    ];

    expect(RopePhysicsModeller.calculateVisitedPositions(motions, 10)).toBe(36);
  });
});
